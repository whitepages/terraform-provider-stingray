package main

import (
	"fmt"

	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/hashicorp/terraform/helper/hashcode"
	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/hashicorp/terraform/helper/schema"
	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/whitepages/go-stingray"
)

func resourcePool() *schema.Resource {
	return &schema.Resource{
		Create: resourcePoolCreate,
		Read:   resourcePoolRead,
		Update: resourcePoolUpdate,
		Delete: resourcePoolDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"connection_max_connect_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"connection_max_connections_per_node": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"connection_max_queue_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"connection_max_reply_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"connection_queue_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"load_balancing_priority_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"monitors": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set: func(v interface{}) int {
					return hashcode.String(v.(string))
				},
			},

			"nodes": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set: func(v interface{}) int {
					return hashcode.String(v.(string))
				},
			},

			"note": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"passive_monitoring": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

      "tcp_nagle": &schema.Schema{
        Type:     schema.TypeBool,
        Optional: true,
        Computed: true,
      },
		},
	}
}

func resourcePoolCreate(d *schema.ResourceData, meta interface{}) error {
	err := resourcePoolSet(d, meta)
	if err != nil {
		return err
	}

	return resourcePoolRead(d, meta)
}

func resourcePoolRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*stingray.Client)

	r, resp, err := c.GetPool(d.Get("name").(string))
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			// The resource doesn't exist anymore
			d.SetId("")

			return nil
		}

		return fmt.Errorf("Error reading resource: %s", err)
	}

	d.Set("connection_max_connect_time", int(*r.Connection.MaxConnectTime))
	d.Set("connection_max_connections_per_node", int(*r.Connection.MaxConnectionsPerNode))
	d.Set("connection_max_queue_size", int(*r.Connection.MaxQueueSize))
	d.Set("connection_max_reply_time", int(*r.Connection.MaxReplyTime))
	d.Set("connection_queue_timeout", int(*r.Connection.QueueTimeout))
	d.Set("load_balancing_priority_enabled", bool(*r.LoadBalancing.PriorityEnabled))
	d.Set("monitors", []string(*r.Basic.Monitors))
	d.Set("nodes", nodesTableToNodes(*r.Basic.NodesTable))
	d.Set("note", string(*r.Basic.Note))
	d.Set("passive_monitoring", bool(*r.Basic.PassiveMonitoring))
  d.Set("tcp_nagle", bool(*r.Basic.TcpNagle))

	return nil
}

func resourcePoolUpdate(d *schema.ResourceData, meta interface{}) error {
	err := resourcePoolSet(d, meta)
	if err != nil {
		return err
	}

	return resourcePoolRead(d, meta)
}

func resourcePoolDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*stingray.Client)
	r := stingray.NewPool(d.Id())

	resp, err := c.Delete(r)
	// Deletes happen twice when a resource is recreated due to
	// ForceNew, so treat a 404 as a success
	// https://github.com/hashicorp/terraform/issues/542
	if err != nil && resp.StatusCode != 404 {
		return err
	}

	return nil
}

func resourcePoolSet(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*stingray.Client)
	r := stingray.NewPool(d.Get("name").(string))

	setInt(&r.Connection.MaxConnectTime, d, "connection_max_connect_time")
	setInt(&r.Connection.MaxConnectionsPerNode, d, "connection_max_connections_per_node")
	setInt(&r.Connection.MaxQueueSize, d, "connection_max_queue_size")
	setInt(&r.Connection.MaxReplyTime, d, "connection_max_reply_time")
	setInt(&r.Connection.QueueTimeout, d, "connection_queue_timeout")
	setBool(&r.LoadBalancing.PriorityEnabled, d, "load_balancing_priority_enabled")
	setStringSet(&r.Basic.Monitors, d, "monitors")
	setNodesTable(&r.Basic.NodesTable, d, "nodes")
	setString(&r.Basic.Note, d, "note")
	setBool(&r.Basic.PassiveMonitoring, d, "passive_monitoring")
  setBool(&r.Basic.TcpNagle, d, "tcp_nagle")

	_, err := c.Set(r)
	if err != nil {
		return err
	}

	d.SetId(d.Get("name").(string))

	return nil
}

func setNodesTable(target **stingray.NodesTable, d *schema.ResourceData, key string) {
	if _, ok := d.GetOk(key); ok {
		var nodes []string
		if v := d.Get(key).(*schema.Set); v.Len() > 0 {
			nodes = make([]string, v.Len())
			for i, v := range v.List() {
				nodes[i] = v.(string)
			}
		}
		nodesTable := nodesToNodesTable(nodes)
		*target = &nodesTable
	}
}

func nodesToNodesTable(nodes []string) stingray.NodesTable {
	t := []stingray.Node{}

	for _, v := range nodes {
		t = append(t, stingray.Node{Node: stingray.String(v)})
	}

	return t
}

func nodesTableToNodes(t []stingray.Node) []string {
	nodes := []string{}

	for _, v := range t {
		// A node deleted from the web UI will still exist in
		// the nodes_table, but state and weight will not
		// exist
		if v.State != nil {
			nodes = append(nodes, *v.Node)
		}
	}

	return nodes
}
