package main

import (
	"fmt"

	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/hashicorp/terraform/helper/hashcode"
	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/hashicorp/terraform/helper/schema"
	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/whitepages/go-stingray"
)

func resourceVirtualServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceVirtualServerCreate,
		Read:   resourceVirtualServerRead,
		Update: resourceVirtualServerUpdate,
		Delete: resourceVirtualServerDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"connection_errors_error_file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"connection_keepalive_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"connection_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"connect_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"gzip_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"gzip_include_mime": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set: func(v interface{}) int {
					return hashcode.String(v.(string))
				},
			},

			"http_location_rewrite": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"listen_on_any": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"listen_on_traffic_ips": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set: func(v interface{}) int {
					return hashcode.String(v.(string))
				},
			},

			"log_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"log_filename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"log_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"log_server_connection_failures": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"note": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"pool": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"recent_connections_save_all": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"request_rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"response_rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"ssl_add_http_headers": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"ssl_decrypt": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"ssl_server_cert_default": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"web_cache_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"web_cache_max_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVirtualServerCreate(d *schema.ResourceData, meta interface{}) error {
	err := resourceVirtualServerSet(d, meta)
	if err != nil {
		return err
	}

	return resourceVirtualServerRead(d, meta)
}

func resourceVirtualServerRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*stingray.Client)

	r, resp, err := c.GetVirtualServer(d.Get("name").(string))
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			// The resource doesn't exist anymore
			d.SetId("")

			return nil
		}

		return fmt.Errorf("Error reading resource: %s", err)
	}

	d.Set("connection_errors_error_file", string(*r.ConnectionErrors.ErrorFile))
	d.Set("connection_keepalive_timeout", int(*r.Connection.KeepaliveTimeout))
	d.Set("connection_timeout", int(*r.Connection.Timeout))
	d.Set("connect_timeout", int(*r.Basic.ConnectTimeout))
	d.Set("enabled", bool(*r.Basic.Enabled))
	d.Set("gzip_enabled", bool(*r.Gzip.Enabled))
	d.Set("gzip_include_mime", []string(*r.Gzip.IncludeMIME))
	d.Set("http_location_rewrite", string(*r.HTTP.LocationRewrite))
	d.Set("listen_on_any", bool(*r.Basic.ListenOnAny))
	d.Set("listen_on_traffic_ips", []string(*r.Basic.ListenOnTrafficIPs))
	d.Set("log_enabled", bool(*r.Log.Enabled))
	d.Set("log_filename", string(*r.Log.Filename))
	d.Set("log_format", string(*r.Log.Format))
	d.Set("log_server_connection_failures", bool(*r.Log.ServerConnectionFailures))
	d.Set("note", string(*r.Basic.Note))
	d.Set("pool", string(*r.Basic.Pool))
	d.Set("port", int(*r.Basic.Port))
	d.Set("protocol", string(*r.Basic.Protocol))
	d.Set("recent_connections_save_all", bool(*r.RecentConnections.SaveAll))
	d.Set("request_rules", []string(*r.Basic.RequestRules))
	d.Set("response_rules", []string(*r.Basic.ResponseRules))
	d.Set("ssl_add_http_headers", bool(*r.SSL.AddHTTPHeaders))
	d.Set("ssl_decrypt", bool(*r.Basic.SSLDecrypt))
	d.Set("ssl_server_cert_default", string(*r.SSL.ServerCertDefault))
	d.Set("web_cache_enabled", bool(*r.WebCache.Enabled))
	d.Set("web_cache_max_time", int(*r.WebCache.MaxTime))

	return nil
}

func resourceVirtualServerUpdate(d *schema.ResourceData, meta interface{}) error {
	err := resourceVirtualServerSet(d, meta)
	if err != nil {
		return err
	}

	return resourceVirtualServerRead(d, meta)
}

func resourceVirtualServerDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*stingray.Client)
	r := stingray.NewVirtualServer(d.Id())

	resp, err := c.Delete(r)
	// Deletes happen twice when a resource is recreated due to
	// ForceNew, so treat a 404 as a success
	// https://github.com/hashicorp/terraform/issues/542
	if err != nil && resp.StatusCode != 404 {
		return err
	}

	return nil
}

func resourceVirtualServerSet(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*stingray.Client)
	r := stingray.NewVirtualServer(d.Get("name").(string))

	setString(&r.ConnectionErrors.ErrorFile, d, "connection_errors_error_file")
	setInt(&r.Connection.KeepaliveTimeout, d, "connection_keepalive_timeout")
	setInt(&r.Connection.Timeout, d, "connection_timeout")
	setInt(&r.Basic.ConnectTimeout, d, "connect_timeout")
	setBool(&r.Basic.Enabled, d, "enabled")
	setStringSet(&r.Gzip.IncludeMIME, d, "gzip_include_mime")
	setBool(&r.Gzip.Enabled, d, "gzip_enabled")
	setString(&r.HTTP.LocationRewrite, d, "http_location_rewrite")
	setBool(&r.Basic.ListenOnAny, d, "listen_on_any")
	setStringSet(&r.Basic.ListenOnTrafficIPs, d, "listen_on_traffic_ips")
	setBool(&r.Log.Enabled, d, "log_enabled")
	setString(&r.Log.Filename, d, "log_filename")
	setString(&r.Log.Format, d, "log_format")
	setBool(&r.Log.ServerConnectionFailures, d, "log_server_connection_failures")
	setString(&r.Basic.Note, d, "note")
	setString(&r.Basic.Pool, d, "pool")
	setInt(&r.Basic.Port, d, "port")
	setString(&r.Basic.Protocol, d, "protocol")
	setBool(&r.RecentConnections.SaveAll, d, "recent_connections_save_all")
	setStringList(&r.Basic.RequestRules, d, "request_rules")
	setStringList(&r.Basic.ResponseRules, d, "response_rules")
	setBool(&r.SSL.AddHTTPHeaders, d, "ssl_add_http_headers")
	setBool(&r.Basic.SSLDecrypt, d, "ssl_decrypt")
	setString(&r.SSL.ServerCertDefault, d, "ssl_server_cert_default")
	setBool(&r.WebCache.Enabled, d, "web_cache_enabled")
	setInt(&r.WebCache.MaxTime, d, "web_cache_max_time")

	_, err := c.Set(r)
	if err != nil {
		return err
	}

	d.SetId(d.Get("name").(string))

	return nil
}
