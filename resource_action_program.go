package main

import (
	"fmt"

	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/hashicorp/terraform/helper/schema"
	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/whitepages/go-stingray"
)

func resourceActionProgram() *schema.Resource {
	return &schema.Resource{
		Create: resourceActionProgramCreate,
		Read:   resourceActionProgramRead,
		Update: resourceActionProgramUpdate,
		Delete: resourceActionProgramDelete,

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

func resourceActionProgramCreate(d *schema.ResourceData, meta interface{}) error {
	err := resourceActionProgramSet(d, meta)
	if err != nil {
		return err
	}

	return resourceActionProgramRead(d, meta)
}

func resourceActionProgramRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client

	r, resp, err := c.GetActionProgram(d.Get("name").(string))
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

func resourceActionProgramUpdate(d *schema.ResourceData, meta interface{}) error {
	err := resourceActionProgramSet(d, meta)
	if err != nil {
		return err
	}

	return resourceActionProgramRead(d, meta)
}

func resourceActionProgramDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client
	r := stingray.NewActionProgram(d.Id())

	_, err := c.Delete(r)
	if err != nil {
		return err
	}

	return nil
}

func resourceActionProgramSet(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client
	r := stingray.NewActionProgram(d.Get("name").(string))

	r.Content = []byte(d.Get("content").(string))

	_, err := c.Set(r)
	if err != nil {
		return err
	}

	d.SetId(d.Get("name").(string))

	return nil
}
