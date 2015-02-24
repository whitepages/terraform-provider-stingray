package main

import (
	"fmt"

	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/hashicorp/terraform/helper/schema"
	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/whitepages/go-stingray"
)

func resourceRate() *schema.Resource {
	return &schema.Resource{
		Create: resourceRateCreate,
		Read:   resourceRateRead,
		Update: resourceRateUpdate,
		Delete: resourceRateDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"max_rate_per_minute": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"max_rate_per_second": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"note": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRateCreate(d *schema.ResourceData, meta interface{}) error {
	err := resourceRateSet(d, meta)
	if err != nil {
		return err
	}

	return resourceRateRead(d, meta)
}

func resourceRateRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client

	r, resp, err := c.GetRate(d.Get("name").(string))
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			// The resource doesn't exist anymore
			d.SetId("")

			return nil
		}

		return fmt.Errorf("Error reading resource: %s", err)
	}

	d.Set("max_rate_per_minute", int(*r.Basic.MaxRatePerMinute))
	d.Set("max_rate_per_second", int(*r.Basic.MaxRatePerSecond))
	d.Set("note", string(*r.Basic.Note))

	return nil
}

func resourceRateUpdate(d *schema.ResourceData, meta interface{}) error {
	err := resourceRateSet(d, meta)
	if err != nil {
		return err
	}

	return resourceRateRead(d, meta)
}

func resourceRateDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client
	r := stingray.NewRate(d.Id())

	_, err := c.Delete(r)
	if err != nil {
		return err
	}

	return nil
}

func resourceRateSet(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client
	r := stingray.NewRate(d.Get("name").(string))

	setInt(&r.Basic.MaxRatePerMinute, d, "max_rate_per_minute")
	setInt(&r.Basic.MaxRatePerSecond, d, "max_rate_per_second")
	setString(&r.Basic.Note, d, "note")

	_, err := c.Set(r)
	if err != nil {
		return err
	}

	d.SetId(d.Get("name").(string))

	return nil
}
