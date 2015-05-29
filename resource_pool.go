package main

import (
	"fmt"

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

			"bandwidth_class": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"connection_max_connect_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},

			"connection_max_connections_per_node": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},

			"connection_max_queue_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},

			"connection_max_reply_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},

			"connection_queue_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},

			"failure_pool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"load_balancing_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "round_robin",
			},

			"load_balancing_priority_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"max_connection_attempts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"max_idle_connections_pernode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"max_timed_out_connection_attempts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"monitors": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"node_close_with_rst": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"node_connection_attempts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"nodes": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"note": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},

			"passive_monitoring": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"persistence_class": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"tcp_nagle": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"transparent": &schema.Schema{
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
	c := meta.(*providerConfig).client

	r, resp, err := c.GetPool(d.Get("name").(string))
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			// The resource doesn't exist anymore
			d.SetId("")

			return nil
		}

		return fmt.Errorf("Error reading resource: %s", err)
	}

	d.Set("bandwidth_class", string(*r.Basic.BandwidthClass))
	d.Set("connection_max_connect_time", int(*r.Connection.MaxConnectTime))
	d.Set("connection_max_connections_per_node", int(*r.Connection.MaxConnectionsPerNode))
	d.Set("connection_max_queue_size", int(*r.Connection.MaxQueueSize))
	d.Set("connection_max_reply_time", int(*r.Connection.MaxReplyTime))
	d.Set("connection_queue_timeout", int(*r.Connection.QueueTimeout))
	d.Set("failure_pool", string(*r.Basic.FailurePool))
	d.Set("load_balancing_algorithm", string(*r.LoadBalancing.Algorithm))
	d.Set("load_balancing_priority_enabled", bool(*r.LoadBalancing.PriorityEnabled))
	d.Set("max_connection_attempts", int(*r.Basic.MaxConnectionAttempts))
	d.Set("max_idle_connections_pernode", int(*r.Basic.MaxIdleConnectionsPerNode))
	d.Set("max_timed_out_connection_attempts", int(*r.Basic.MaxTimedOutConnectionAttempts))
	d.Set("monitors", []string(*r.Basic.Monitors))
	d.Set("node_close_with_rst", bool(*r.Basic.NodeCloseWithRST))
	d.Set("node_connection_attempts", int(*r.Basic.NodeConnectionAttempts))
	d.Set("nodes", nodesTableToNodes(*r.Basic.NodesTable))
	d.Set("note", string(*r.Basic.Note))
	d.Set("passive_monitoring", bool(*r.Basic.PassiveMonitoring))
	d.Set("persistence_class", string(*r.Basic.PersistenceClass))
	d.Set("tcp_nagle", bool(*r.TCP.Nagle))
	d.Set("transparent", bool(*r.Basic.Transparent))

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
	c := meta.(*providerConfig).client
	r := stingray.NewPool(d.Id())

	_, err := c.Delete(r)
	if err != nil {
		return err
	}

	return nil
}

func resourcePoolSet(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client
	r := stingray.NewPool(d.Get("name").(string))

	setString(&r.Basic.BandwidthClass, d, "bandwidth_class")
	setInt(&r.Connection.MaxConnectTime, d, "connection_max_connect_time")
	setInt(&r.Connection.MaxConnectionsPerNode, d, "connection_max_connections_per_node")
	setInt(&r.Connection.MaxQueueSize, d, "connection_max_queue_size")
	setInt(&r.Connection.MaxReplyTime, d, "connection_max_reply_time")
	setInt(&r.Connection.QueueTimeout, d, "connection_queue_timeout")
	setString(&r.Basic.FailurePool, d, "failure_pool")
	setString(&r.LoadBalancing.Algorithm, d, "load_balancing_algorithm")
	setBool(&r.LoadBalancing.PriorityEnabled, d, "load_balancing_priority_enabled")
	setInt(&r.Basic.MaxConnectionAttempts, d, "max_connection_attempts")
	setInt(&r.Basic.MaxIdleConnectionsPerNode, d, "max_idle_connections_pernode")
	setInt(&r.Basic.MaxTimedOutConnectionAttempts, d, "max_timed_out_connection_attempts")
	setStringSet(&r.Basic.Monitors, d, "monitors")
	setBool(&r.Basic.NodeCloseWithRST, d, "node_close_with_rst")
	setInt(&r.Basic.NodeConnectionAttempts, d, "node_connection_attempts")
	setNodesTable(&r.Basic.NodesTable, d, "nodes")
	setString(&r.Basic.Note, d, "note")
	setBool(&r.Basic.PassiveMonitoring, d, "passive_monitoring")
	setString(&r.Basic.PersistenceClass, d, "persistence_class")
	setBool(&r.TCP.Nagle, d, "tcp_nagle")
	setBool(&r.Basic.Transparent, d, "transparent")

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
