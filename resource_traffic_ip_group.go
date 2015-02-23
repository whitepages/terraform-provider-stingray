package main

import (
	"fmt"

	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/hashicorp/terraform/helper/hashcode"
	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/hashicorp/terraform/helper/schema"
	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/whitepages/go-stingray"
)

func resourceTrafficIPGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceTrafficIPGroupCreate,
		Read:   resourceTrafficIPGroupRead,
		Update: resourceTrafficIPGroupUpdate,
		Delete: resourceTrafficIPGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"hash_source_port": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"ipaddresses": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set: func(v interface{}) int {
					return hashcode.String(v.(string))
				},
			},

			"keeptogether": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"location": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"machines": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set: func(v interface{}) int {
					return hashcode.String(v.(string))
				},
			},

			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"multicast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"note": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"slaves": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set: func(v interface{}) int {
					return hashcode.String(v.(string))
				},
			},
		},
	}
}

func resourceTrafficIPGroupCreate(d *schema.ResourceData, meta interface{}) error {
	err := resourceTrafficIPGroupSet(d, meta)
	if err != nil {
		return err
	}

	return resourceTrafficIPGroupRead(d, meta)
}

func resourceTrafficIPGroupRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client

	r, resp, err := c.GetTrafficIPGroup(d.Get("name").(string))
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			// The resource doesn't exist anymore
			d.SetId("")

			return nil
		}

		return fmt.Errorf("Error reading resource: %s", err)
	}

	d.Set("enabled", bool(*r.Basic.Enabled))
	d.Set("hash_source_port", bool(*r.Basic.HashSourcePort))
	d.Set("ipaddresses", []string(*r.Basic.IPAddresses))
	d.Set("keeptogether", bool(*r.Basic.KeepTogether))
	d.Set("location", int(*r.Basic.Location))
	d.Set("machines", []string(*r.Basic.Machines))
	d.Set("mode", string(*r.Basic.Mode))
	d.Set("multicast", string(*r.Basic.Multicast))
	d.Set("note", string(*r.Basic.Note))
	d.Set("slaves", []string(*r.Basic.Machines))

	return nil
}

func resourceTrafficIPGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	err := resourceTrafficIPGroupSet(d, meta)
	if err != nil {
		return err
	}

	return resourceTrafficIPGroupRead(d, meta)
}

func resourceTrafficIPGroupDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client
	r := stingray.NewTrafficIPGroup(d.Id())

	_, err := c.Delete(r)
	if err != nil {
		return err
	}

	return nil
}

func resourceTrafficIPGroupSet(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client
	r := stingray.NewTrafficIPGroup(d.Get("name").(string))

	setBool(&r.Basic.Enabled, d, "enabled")
	setBool(&r.Basic.HashSourcePort, d, "hash_source_port")
	setStringSet(&r.Basic.IPAddresses, d, "ipaddresses")
	setBool(&r.Basic.KeepTogether, d, "keeptogether")
	setInt(&r.Basic.Location, d, "location")
	setStringSet(&r.Basic.Machines, d, "machines")
	setString(&r.Basic.Mode, d, "mode")
	setString(&r.Basic.Multicast, d, "multicast")
	setString(&r.Basic.Note, d, "note")
	setStringSet(&r.Basic.Slaves, d, "slaves")

	_, err := c.Set(r)
	if err != nil {
		return err
	}

	d.SetId(d.Get("name").(string))

	return nil
}
