package main

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/whitepages/go-stingray"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VTM_URL", nil),
			},

			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VTM_USERNAME", nil),
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VTM_PASSWORD", nil),
			},

			"valid_networks": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VTM_VALID_NETWORKS", ""),
			},

			"verify_ssl": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VTM_VERIFY_SSL", true),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"vtm_action_program":        resourceActionProgram(),
			"vtm_extra_file":            resourceExtraFile(),
			"vtm_license_key":           resourceLicenseKey(),
			"vtm_monitor_script":        resourceMonitorScript(),
			"vtm_monitor":               resourceMonitor(),
			"vtm_pool":                  resourcePool(),
			"vtm_rate":                  resourceRate(),
			"vtm_rule":                  resourceRule(),
			"vtm_service_level_monitor": resourceServiceLevelMonitor(),
			"vtm_ssl_cas":               resourceSSLCAs(),
			"vtm_ssl_server_key":        resourceSSLServerKey(),
			"vtm_traffic_ip_group":      resourceTrafficIPGroup(),
			"vtm_virtual_server":        resourceVirtualServer(),
		},

		ConfigureFunc: providerConfigure,
	}
}

type providerConfig struct {
	client        *stingray.Client
	validNetworks netList
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		URL:       d.Get("url").(string),
		Username:  d.Get("username").(string),
		Password:  d.Get("password").(string),
		VerifySSL: d.Get("verify_ssl").(bool),
	}
	client, err := config.Client()
	if err != nil {
		return nil, err
	}

	validNetworks := d.Get("valid_networks").(string)
	ns := netList{}

	if len(validNetworks) > 0 {
		cidrList := strings.Split(validNetworks, ",")
		ns, err = parseCIDRList(cidrList)
		if err != nil {
			return nil, err
		}
	}

	return &providerConfig{client: client, validNetworks: ns}, nil
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

func setBool(target **bool, d *schema.ResourceData, key string) {
	*target = stingray.Bool(d.Get(key).(bool))
}

func setInt(target **int, d *schema.ResourceData, key string) {
	*target = stingray.Int(d.Get(key).(int))
}

func setString(target **string, d *schema.ResourceData, key string) {
	*target = stingray.String(d.Get(key).(string))
}

func setStringList(target **[]string, d *schema.ResourceData, key string) {
	list := expandStringList(d.Get(key).([]interface{}))
	*target = &list
}

func setStringSet(target **[]string, d *schema.ResourceData, key string) {
	list := expandStringList(d.Get(key).(*schema.Set).List())
	*target = &list
}
