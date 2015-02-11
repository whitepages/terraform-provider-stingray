package main

import (
	"crypto/sha1"
	"encoding/hex"
	"os"

	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/hashicorp/terraform/helper/schema"
	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/hashicorp/terraform/terraform"
	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/whitepages/go-stingray"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"verify_ssl": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"stingray_action_program":        resourceActionProgram(),
			"stingray_extra_file":            resourceExtraFile(),
			"stingray_license_key":           resourceLicenseKey(),
			"stingray_monitor_script":        resourceMonitorScript(),
			"stingray_monitor":               resourceMonitor(),
			"stingray_pool":                  resourcePool(),
			"stingray_rate":                  resourceRate(),
			"stingray_rule":                  resourceRule(),
			"stingray_service_level_monitor": resourceServiceLevelMonitor(),
			"stingray_ssl_cas":               resourceSSLCAs(),
			"stingray_ssl_server_key":        resourceSSLServerKey(),
			"stingray_traffic_ip_group":      resourceTrafficIPGroup(),
			"stingray_virtual_server":        resourceVirtualServer(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		URL:       d.Get("url").(string),
		Username:  d.Get("username").(string),
		Password:  d.Get("password").(string),
		VerifySSL: d.Get("verify_ssl").(bool),
	}

	return config.Client()
}

// Takes the result of flatmap.Expand for an array of strings
// and returns a []string
func expandStringList(configured []interface{}) []string {
	vs := make([]string, 0, len(configured))
	for _, v := range configured {
		vs = append(vs, v.(string))
	}
	return vs
}

// hashString returns a hash of the input for use as a StateFunc
func hashString(v interface{}) string {
	switch v.(type) {
	case string:
		hash := sha1.Sum([]byte(v.(string)))
		return hex.EncodeToString(hash[:])
	default:
		return ""
	}
}

// setBool sets the target if the key is set in the schema config
func setBool(target **bool, d *schema.ResourceData, key string) {
	if v, ok := d.GetOk(key); ok {
		*target = stingray.Bool(v.(bool))
	}
}

// setInt sets the target if the key is set in the schema config
func setInt(target **int, d *schema.ResourceData, key string) {
	if v, ok := d.GetOk(key); ok {
		*target = stingray.Int(v.(int))
	}
}

// setString sets the target if the key is set in the schema config
func setString(target **string, d *schema.ResourceData, key string) {
	if v, ok := d.GetOk(key); ok {
		*target = stingray.String(v.(string))
	}
}

// setStringList sets the target if the key is set in the schema config
func setStringList(target **[]string, d *schema.ResourceData, key string) {
	if v, ok := d.GetOk(key); ok {
		list := expandStringList(v.([]interface{}))
		*target = &list
	}
}

// setStringSet sets the target if the key is set in the schema config
func setStringSet(target **[]string, d *schema.ResourceData, key string) {
	if _, ok := d.GetOk(key); ok {
		list := expandStringList(d.Get(key).(*schema.Set).List())
		*target = &list
	}
}

func envDefaultFunc(k string, alt interface{}) schema.SchemaDefaultFunc {
	return func() (interface{}, error) {
		if v := os.Getenv(k); v != "" {
			return v, nil
		}

		return alt, nil
	}
}
