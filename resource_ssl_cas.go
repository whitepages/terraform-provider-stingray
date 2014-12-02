package main

import (
	"fmt"

	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/hashicorp/terraform/helper/schema"
	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/whitepages/go-stingray"
)

func resourceSSLCAs() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSLCAsCreate,
		Read:   resourceSSLCAsRead,
		Update: resourceSSLCAsUpdate,
		Delete: resourceSSLCAsDelete,

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

func resourceSSLCAsCreate(d *schema.ResourceData, meta interface{}) error {
	err := resourceSSLCAsSet(d, meta)
	if err != nil {
		return err
	}

	return resourceSSLCAsRead(d, meta)
}

func resourceSSLCAsRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*stingray.Client)

	r, resp, err := c.GetSSLCAs(d.Get("name").(string))
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

func resourceSSLCAsUpdate(d *schema.ResourceData, meta interface{}) error {
	err := resourceSSLCAsSet(d, meta)
	if err != nil {
		return err
	}

	return resourceSSLCAsRead(d, meta)
}

func resourceSSLCAsDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*stingray.Client)
	r := stingray.NewSSLCAs(d.Id())

	resp, err := c.Delete(r)
	// Deletes happen twice when a resource is recreated due to
	// ForceNew, so treat a 404 as a success
	// https://github.com/hashicorp/terraform/issues/542
	if err != nil && resp.StatusCode != 404 {
		return err
	}

	return nil
}

func resourceSSLCAsSet(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*stingray.Client)
	r := stingray.NewSSLCAs(d.Get("name").(string))

	r.Content = []byte(d.Get("content").(string))

	_, err := c.Set(r)
	if err != nil {
		return err
	}

	d.SetId(d.Get("name").(string))

	return nil
}
