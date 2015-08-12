package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/hashcode"
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

			"back_off": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3,
			},

			"failures": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3,
			},

			"http_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},

			"http_body_regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},

			"http_host_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},

			"http_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/",
			},

			"http_status_regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "^[234][0-9][0-9]$",
			},

			"machine": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},

			"note": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},

			"rtsp_body_regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},

			"rtsp_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},

			"rtsp_status_regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "^[234][0-9][0-9]$",
			},

			"script_arguments": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  true,
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
				Set: hashScriptArguments,
			},

			"script_program": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},

			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "pernode",
			},

			"sip_body_regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},

			"sip_status_regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "^[234][0-9][0-9]$",
			},

			"sip_transport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "udp",
			},

			"tcp_close_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},

			"tcp_max_response_len": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2048,
			},

			"tcp_response_regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  ".+",
			},

			"tcp_write_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},

			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3,
			},

			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ping",
			},

			"udp_accept_all": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"use_ssl": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"verbose": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
	c := meta.(*providerConfig).client

	r, resp, err := c.GetMonitor(d.Get("name").(string))
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			// The resource doesn't exist anymore
			d.SetId("")

			return nil
		}

		return fmt.Errorf("Error reading resource: %s", err)
	}

	d.Set("back_off", bool(*r.Basic.BackOff))
	d.Set("delay", int(*r.Basic.Delay))
	d.Set("failures", int(*r.Basic.Failures))
	d.Set("http_authentication", string(*r.HTTP.Authentication))
	d.Set("http_body_regex", string(*r.HTTP.BodyRegex))
	d.Set("http_host_header", string(*r.HTTP.HostHeader))
	d.Set("http_path", string(*r.HTTP.Path))
	d.Set("http_status_regex", string(*r.HTTP.StatusRegex))
	d.Set("machine", string(*r.Basic.Machine))
	d.Set("note", string(*r.Basic.Note))
	d.Set("rtsp_body_regex", string(*r.RTSP.BodyRegex))
	d.Set("rtsp_path", string(*r.RTSP.Path))
	d.Set("rtsp_status_regex", string(*r.RTSP.StatusRegex))
	d.Set("scope", string(*r.Basic.Scope))
	d.Set("script_arguments", flattenScriptArgumentsTable(*r.Script.Arguments))
	d.Set("script_program", string(*r.Script.Program))
	d.Set("sip_body_regex", string(*r.SIP.BodyRegex))
	d.Set("sip_status_regex", string(*r.SIP.StatusRegex))
	d.Set("sip_transport", string(*r.SIP.Transport))
	d.Set("tcp_close_string", string(*r.TCP.CloseString))
	d.Set("tcp_max_response_len", int(*r.TCP.MaxResponseLen))
	d.Set("tcp_response_regex", string(*r.TCP.ResponseRegex))
	d.Set("tcp_write_string", string(*r.TCP.WriteString))
	d.Set("timeout", int(*r.Basic.Timeout))
	d.Set("type", string(*r.Basic.Type))
	d.Set("udp_accept_all", bool(*r.UDP.AcceptAll))
	d.Set("use_ssl", bool(*r.Basic.UseSSL))
	d.Set("verbose", bool(*r.Basic.Verbose))

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
	c := meta.(*providerConfig).client
	r := stingray.NewMonitor(d.Id())

	_, err := c.Delete(r)
	if err != nil {
		return err
	}

	return nil
}

func resourceMonitorSet(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client
	r := stingray.NewMonitor(d.Get("name").(string))

	setBool(&r.Basic.BackOff, d, "back_off")
	setInt(&r.Basic.Delay, d, "delay")
	setInt(&r.Basic.Failures, d, "failures")
	setString(&r.HTTP.Authentication, d, "http_authentication")
	setString(&r.HTTP.BodyRegex, d, "http_body_regex")
	setString(&r.HTTP.HostHeader, d, "http_host_header")
	setString(&r.HTTP.Path, d, "http_path")
	setString(&r.HTTP.StatusRegex, d, "http_status_regex")
	setString(&r.Basic.Machine, d, "machine")
	setString(&r.Basic.Note, d, "note")
	setString(&r.RTSP.BodyRegex, d, "rtsp_body_regex")
	setString(&r.RTSP.Path, d, "rtsp_path")
	setString(&r.RTSP.StatusRegex, d, "rtsp_status_regex")
	setString(&r.Basic.Scope, d, "scope")
	setScriptArgumentsTable(&r.Script.Arguments, d, "script_arguments")
	setString(&r.Script.Program, d, "script_program")
	setString(&r.SIP.BodyRegex, d, "sip_body_regex")
	setString(&r.SIP.StatusRegex, d, "sip_status_regex")
	setString(&r.SIP.Transport, d, "sip_transports")
	setString(&r.TCP.CloseString, d, "tcp_close_string")
	setInt(&r.TCP.MaxResponseLen, d, "tcp_max_response_len")
	setString(&r.TCP.ResponseRegex, d, "tcp_response_regex")
	setString(&r.TCP.WriteString, d, "tcp_write_string")
	setInt(&r.Basic.Timeout, d, "timeout")
	setString(&r.Basic.Type, d, "type")
	setBool(&r.UDP.AcceptAll, d, "udp_accept_all")
	setBool(&r.Basic.UseSSL, d, "use_ssl")
	setBool(&r.Basic.Verbose, d, "verbose")

	_, err := c.Set(r)
	if err != nil {
		return err
	}

	d.SetId(d.Get("name").(string))

	return nil
}

func setScriptArgumentsTable(target **stingray.ScriptArgumentsTable, d *schema.ResourceData, key string) {
	if _, ok := d.GetOk(key); ok {
		table := d.Get(key).(*schema.Set).List()
		*target, _ = expandScriptArgumentsTable(table)
	}
}

func expandScriptArgumentsTable(configured []interface{}) (*stingray.ScriptArgumentsTable, error) {
	table := make(stingray.ScriptArgumentsTable, 0, len(configured))

	for _, raw := range configured {
		data := raw.(map[string]interface{})

		s := stingray.ScriptArgument{
			Description: stingray.String(data["description"].(string)),
			Name:        stingray.String(data["name"].(string)),
			Value:       stingray.String(data["value"].(string)),
		}

		table = append(table, s)
	}

	return &table, nil
}

func flattenScriptArgumentsTable(list stingray.ScriptArgumentsTable) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))
	for _, i := range list {
		s := map[string]interface{}{
			"description": *i.Description,
			"name":        *i.Name,
			"value":       *i.Value,
		}
		result = append(result, s)
	}

	return result
}

func hashScriptArguments(v interface{}) int {
	m := v.(map[string]interface{})
	return hashcode.String(m["name"].(string))
}
