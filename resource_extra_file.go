package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/whitepages/go-stingray"
)

func resourceExtraFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtraFileCreate,
		Read:   resourceExtraFileRead,
		Update: resourceExtraFileUpdate,
		Delete: resourceExtraFileDelete,

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

func resourceExtraFileCreate(d *schema.ResourceData, meta interface{}) error {
	err := resourceExtraFileSet(d, meta)
	if err != nil {
		return err
	}

	return resourceExtraFileRead(d, meta)
}

func resourceExtraFileRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client

	r, resp, err := c.GetExtraFile(d.Get("name").(string))
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

func resourceExtraFileUpdate(d *schema.ResourceData, meta interface{}) error {
	err := resourceExtraFileSet(d, meta)
	if err != nil {
		return err
	}

	return resourceExtraFileRead(d, meta)
}

func resourceExtraFileDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client
	r := stingray.NewExtraFile(d.Id())

	_, err := c.Delete(r)
	if err != nil {
		return err
	}

	return nil
}

func resourceExtraFileSet(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client
	r := stingray.NewExtraFile(d.Get("name").(string))

	r.Content = []byte(d.Get("content").(string))

	_, err := c.Set(r)
	if err != nil {
		return err
	}

	d.SetId(d.Get("name").(string))

	return nil
}
