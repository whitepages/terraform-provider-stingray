package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/whitepages/go-stingray"
)

func resourceLicenseKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceLicenseKeyCreate,
		Read:   resourceLicenseKeyRead,
		Update: resourceLicenseKeyUpdate,
		Delete: resourceLicenseKeyDelete,

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

func resourceLicenseKeyCreate(d *schema.ResourceData, meta interface{}) error {
	err := resourceLicenseKeySet(d, meta)
	if err != nil {
		return err
	}

	return resourceLicenseKeyRead(d, meta)
}

func resourceLicenseKeyRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client

	r, resp, err := c.GetLicenseKey(d.Get("name").(string))
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

func resourceLicenseKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	err := resourceLicenseKeySet(d, meta)
	if err != nil {
		return err
	}

	return resourceLicenseKeyRead(d, meta)
}

func resourceLicenseKeyDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client
	r := stingray.NewLicenseKey(d.Id())

	_, err := c.Delete(r)
	if err != nil {
		return err
	}

	return nil
}

func resourceLicenseKeySet(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client
	r := stingray.NewLicenseKey(d.Get("name").(string))

	r.Content = []byte(d.Get("content").(string))

	_, err := c.Set(r)
	if err != nil {
		return err
	}

	d.SetId(d.Get("name").(string))

	return nil
}
