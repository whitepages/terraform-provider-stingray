package main

import (
	"fmt"

	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/hashicorp/terraform/helper/schema"
	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/whitepages/go-stingray"
)

func resourceMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceMonitorCreate,
		Read:   resourceMonitorRead,
		Update: resourceMonitorUpdate,
		Delete: resourceMonitorDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"failures": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"http_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"http_body_regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"http_host_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"http_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"http_status_regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"note": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// TODO: TypeSet might be better for script_arguments
			"script_arguments": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"value": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"script_program": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"tcp_close_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"tcp_max_response_len": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"tcp_response_regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"tcp_write_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"udp_accept_all": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"use_ssl": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	err := resourceMonitorSet(d, meta)
	if err != nil {
		return err
	}

	return resourceMonitorRead(d, meta)
}

func resourceMonitorRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*stingray.Client)

	r, resp, err := c.GetMonitor(d.Get("name").(string))
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			// The resource doesn't exist anymore
			d.SetId("")

			return nil
		}

		return fmt.Errorf("Error reading resource: %s", err)
	}

	d.Set("delay", int(*r.Basic.Delay))
	d.Set("failures", int(*r.Basic.Failures))
	d.Set("http_authentication", string(*r.HTTP.Authentication))
	d.Set("http_body_regex", string(*r.HTTP.BodyRegex))
	d.Set("http_host_header", string(*r.HTTP.HostHeader))
	d.Set("http_path", string(*r.HTTP.Path))
	d.Set("http_status_regex", string(*r.HTTP.StatusRegex))
	d.Set("note", string(*r.Basic.Note))
	for i, arg := range *r.Script.Arguments {
		prefix := fmt.Sprintf("script_arguments.%d", i)
		d.Set(prefix+".description", arg.Description)
		d.Set(prefix+".name", arg.Value)
		d.Set(prefix+".value", arg.Value)
	}
	d.Set("script_program", string(*r.Script.Program))
	d.Set("tcp_close_string", string(*r.TCP.CloseString))
	d.Set("tcp_max_response_len", int(*r.TCP.MaxResponseLen))
	d.Set("tcp_response_regex", string(*r.TCP.ResponseRegex))
	d.Set("tcp_write_string", string(*r.TCP.WriteString))
	d.Set("timeout", int(*r.Basic.Timeout))
	d.Set("type", string(*r.Basic.Type))
	d.Set("udp_accept_all", bool(*r.UDP.AcceptAll))
	d.Set("use_ssl", bool(*r.Basic.UseSSL))

	return nil
}

func resourceMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	err := resourceMonitorSet(d, meta)
	if err != nil {
		return err
	}

	return resourceMonitorRead(d, meta)
}

func resourceMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*stingray.Client)
	r := stingray.NewMonitor(d.Id())

	resp, err := c.Delete(r)
	// Deletes happen twice when a resource is recreated due to
	// ForceNew, so treat a 404 as a success
	// https://github.com/hashicorp/terraform/issues/542
	if err != nil && resp.StatusCode != 404 {
		return err
	}

	return nil
}

func resourceMonitorSet(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*stingray.Client)
	r := stingray.NewMonitor(d.Get("name").(string))

	setInt(&r.Basic.Delay, d, "delay")
	setInt(&r.Basic.Failures, d, "failures")
	setString(&r.HTTP.Authentication, d, "http_authentication")
	setString(&r.HTTP.BodyRegex, d, "http_body_regex")
	setString(&r.HTTP.HostHeader, d, "http_host_header")
	setString(&r.HTTP.Path, d, "http_path")
	setString(&r.HTTP.StatusRegex, d, "http_status_regex")
	setString(&r.Basic.Note, d, "note")
	setScriptArgumentsTable(&r.Script.Arguments, d, "script_arguments")
	setString(&r.Script.Program, d, "script_program")
	setString(&r.TCP.CloseString, d, "tcp_close_string")
	setInt(&r.TCP.MaxResponseLen, d, "tcp_max_response_len")
	setString(&r.TCP.ResponseRegex, d, "tcp_response_regex")
	setString(&r.TCP.WriteString, d, "tcp_write_string")
	setInt(&r.Basic.Timeout, d, "timeout")
	setString(&r.Basic.Type, d, "type")
	setBool(&r.UDP.AcceptAll, d, "udp_accept_all")
	setBool(&r.Basic.UseSSL, d, "use_ssl")

	_, err := c.Set(r)
	if err != nil {
		return err
	}

	d.SetId(d.Get("name").(string))

	return nil
}

func setScriptArgumentsTable(target **stingray.ScriptArgumentsTable, d *schema.ResourceData, key string) {
	if _, ok := d.GetOk(key); ok {
		t := stingray.ScriptArgumentsTable{}
		count := d.Get(key + ".#").(int)
		for i := 0; i < count; i++ {
			a := stingray.ScriptArgument{}
			prefix := fmt.Sprintf("%s.%d", key, i)
			setString(&a.Description, d, prefix+".description")
			a.Name = stingray.String(d.Get(prefix + ".name").(string))
			a.Value = stingray.String(d.Get(prefix + ".value").(string))
			t = append(t, a)
		}
		*target = &t
	}
}
