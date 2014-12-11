package main

import (
	"fmt"

	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/hashicorp/terraform/helper/schema"
	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/whitepages/go-stingray"
)

func resourceMonitorScript() *schema.Resource {
	return &schema.Resource{
		Create: resourceMonitorScriptCreate,
		Read:   resourceMonitorScriptRead,
		Update: resourceMonitorScriptUpdate,
		Delete: resourceMonitorScriptDelete,

		Schema: map[string]*schema.Schema{
			"content": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				StateFunc: hashString,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceMonitorScriptCreate(d *schema.ResourceData, meta interface{}) error {
	err := resourceMonitorScriptSet(d, meta)
	if err != nil {
		return err
	}

	return resourceMonitorScriptRead(d, meta)
}

func resourceMonitorScriptRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*stingray.Client)

	r, resp, err := c.GetMonitorScript(d.Get("name").(string))
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			// The resource doesn't exist anymore
			d.SetId("")

			return nil
		}

		return fmt.Errorf("Error reading resource: %s", err)
	}

	d.Set("content", hashString(string(r.Content)))

	return nil
}

func resourceMonitorScriptUpdate(d *schema.ResourceData, meta interface{}) error {
	err := resourceMonitorScriptSet(d, meta)
	if err != nil {
		return err
	}

	return resourceMonitorScriptRead(d, meta)
}

func resourceMonitorScriptDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*stingray.Client)
	r := stingray.NewMonitorScript(d.Id())

	_, err := c.Delete(r)
	if err != nil {
		return err
	}

	return nil
}

func resourceMonitorScriptSet(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*stingray.Client)
	r := stingray.NewMonitorScript(d.Get("name").(string))

	r.Content = []byte(d.Get("content").(string))

	_, err := c.Set(r)
	if err != nil {
		return err
	}

	d.SetId(d.Get("name").(string))

	return nil
}
