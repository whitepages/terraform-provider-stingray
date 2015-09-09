package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/whitepages/go-stingray"
)

func resourceServiceLevelMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceLevelMonitorCreate,
		Read:   resourceServiceLevelMonitorRead,
		Update: resourceServiceLevelMonitorUpdate,
		Delete: resourceServiceLevelMonitorDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"note": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},

			"response_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},

			"serious_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},

			"warning_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  50,
			},
		},
	}
}

func resourceServiceLevelMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	err := resourceServiceLevelMonitorSet(d, meta)
	if err != nil {
		return err
	}

	return resourceServiceLevelMonitorRead(d, meta)
}

func resourceServiceLevelMonitorRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client

	r, resp, err := c.GetServiceLevelMonitor(d.Get("name").(string))
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			// The resource doesn't exist anymore
			d.SetId("")

			return nil
		}

		return fmt.Errorf("Error reading resource: %s", err)
	}

	d.Set("note", string(*r.Basic.Note))
	d.Set("response_time", int(*r.Basic.ResponseTime))
	d.Set("serious_threshold", int(*r.Basic.SeriousThreshold))
	d.Set("warning_threshold", int(*r.Basic.WarningThreshold))

	return nil
}

func resourceServiceLevelMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	err := resourceServiceLevelMonitorSet(d, meta)
	if err != nil {
		return err
	}

	return resourceServiceLevelMonitorRead(d, meta)
}

func resourceServiceLevelMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client
	r := stingray.NewServiceLevelMonitor(d.Id())

	_, err := c.Delete(r)
	if err != nil {
		return err
	}

	return nil
}

func resourceServiceLevelMonitorSet(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client
	r := stingray.NewServiceLevelMonitor(d.Get("name").(string))

	setString(&r.Basic.Note, d, "note")
	setInt(&r.Basic.ResponseTime, d, "response_time")
	setInt(&r.Basic.SeriousThreshold, d, "serious_threshold")
	setInt(&r.Basic.WarningThreshold, d, "warning_threshold")

	_, err := c.Set(r)
	if err != nil {
		return err
	}

	d.SetId(d.Get("name").(string))

	return nil
}
