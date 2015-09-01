package stingray

import (
	"encoding/json"
	"net/http"
)

// A Pool is a Stingray pool.
type Pool struct {
	jsonResource   `json:"-"`
	PoolProperties `json:"properties"`
}

type PoolProperties struct {
	AutoScaling struct {
		AddNodeDelayTime *int      `json:"addnode_delaytime,omitempty"`
		CloudCredentials *string   `json:"cloud_credentials,omitempty"`
		Cluster          *string   `json:"cluster,omitempty"`
		DataCenter       *string   `json:"data_center,omitempty"`
		DataStore        *string   `json:"data_store,omitempty"`
		Enabled          *bool     `json:"enabled,omitempty"`
		External         *bool     `json:"external,omitempty"`
		Hysteresis       *int      `json:"hysteresis,omitempty"`
		ImageID          *string   `json:"imageid,omitempty"`
		IPsToUse         *string   `json:"ips_to_use,omitempty"`
		LastNodeIdleTime *int      `json:"last_node_idle_time,omitempty"`
		MaxNodes         *int      `json:"max_nodes,omitempty"`
		MinNodes         *int      `json:"min_nodes,omitempty"`
		Name             *string   `json:"name,omitempty"`
		Port             *int      `json:"port,omitempty"`
		Refractory       *int      `json:"refractory,omitempty"`
		ResponseTime     *int      `json:"response_time,omitempty"`
		ScaleDownLevel   *int      `json:"scale_down_level,omitempty"`
		ScaleUpLevel     *int      `json:"scale_up_level,omitempty"`
		SecurityGroupIDs *[]string `json:"securitygroupids,omitempty"`
		SizeID           *string   `json:"size_id,omitempty"`
		SubnetIDs        *[]string `json:"subnetids,omitempty"`
	} `json:"auto_scaling"`
	Basic struct {
		BandwidthClass                *string     `json:"bandwidth_class,omitempty"`
		FailurePool                   *string     `json:"failure_pool,omitempty"`
		MaxConnectionAttempts         *int        `json:"max_connection_attempts,omitempty"`
		MaxIdleConnectionsPerNode     *int        `json:"max_idle_connections_pernode,omitempty"`
		MaxTimedOutConnectionAttempts *int        `json:"max_timed_out_connection_attempts,omitempty"`
		Monitors                      *[]string   `json:"monitors,omitempty"`
		NodeCloseWithRST              *bool       `json:"node_close_with_rst,omitempty"`
		NodeConnectionAttempts        *int        `json:"node_connection_attempts,omitempty"`
		NodeDeleteBehavior            *string     `json:"node_delete_behavior,omitempty"`
		NodeDrainToDeleteTimeout      *int        `json:"node_drain_to_delete_timeout,omitempty"`
		NodesTable                    *NodesTable `json:"nodes_table,omitempty"`
		PassiveMonitoring             *bool       `json:"passive_monitoring,omitempty"`
		PersistenceClass              *string     `json:"persistence_class,omitempty"`
		Note                          *string     `json:"note,omitempty"`
		Transparent                   *bool       `json:"transparent,omitempty"`
	} `json:"basic"`
	Connection struct {
		MaxConnectTime        *int `json:"max_connect_time,omitempty"`
		MaxConnectionsPerNode *int `json:"max_connections_per_node,omitempty"`
		MaxQueueSize          *int `json:"max_queue_size,omitempty"`
		MaxReplyTime          *int `json:"max_reply_time,omitempty"`
		QueueTimeout          *int `json:"queue_timeout,omitempty"`
	} `json:"connection"`
	DNS struct {
		EDNSUDPSize *int `json:"edns_udpsize,omitempty"`
		MaxUDPSize  *int `json:"max_udpsize,omitempty"`
	} `json:"dns"`
	DNSAutoscale struct {
		Enabled   *bool     `json:"enabled,omitempty"`
		Hostnames *[]string `json:"hostnames,omitempty"`
		Port      *int      `json:"port,omitempty"`
	} `json:"dns_autoscale"`
	FTP struct {
		SupportRFC2428 *bool `json:"support_rfc_2428,omitempty"`
	} `json:"ftp"`
	HTTP struct {
		Keepalive              *bool `json:"keepalive,omitempty"`
		KeepaliveNonIdempotent *bool `json:"keepalive_non_idempotent,omitempty"`
	} `json:"http"`
	KerberosProtocolTransition struct {
		Principal *string `json:"principal,omitempty"`
		Target    *string `json:"target,omitempty"`
	} `json:"kerberos_protocol_transition"`
	LoadBalancing struct {
		Algorithm       *string `json:"algorithm,omitempty"`
		PriorityEnabled *bool   `json:"priority_enabled,omitempty"`
		PriorityNodes   *int    `json:"priority_nodes,omitempty"`
	} `json:"load_balancing"`
	Node struct {
		CloseOnDeath  *bool `json:"close_on_death,omitempty"`
		RetryFailTime *int  `json:"retry_fail_time,omitempty"`
	} `json:"node"`
	SMTP struct {
		SendStartTLS *bool `json:"send_starttls,omitempty"`
	} `json:"smtp"`
	SSL struct {
		ClientAuth          *bool     `json:"client_auth,omitempty"`
		CommonNameMatch     *[]string `json:"common_name_match,omitempty"`
		EllipticCurves      *[]string `json:"elliptic_curves,omitempty"`
		Enable              *bool     `json:"enable,omitempty"`
		Enhance             *bool     `json:"enhance,omitempty"`
		SendCloseAlerts     *bool     `json:"send_close_alerts,omitempty"`
		ServerName          *bool     `json:"server_name,omitempty"`
		SignatureAlgorithms *string   `json:"signature_algorithms,omitempty"`
		SSLCiphers          *string   `json:"ssl_ciphers,omitempty"`
		SSLSupportSSL2      *string   `json:"ssl_support_ssl2,omitempty"`
		SSLSupportSSL3      *string   `json:"ssl_support_ssl3,omitempty"`
		SSLSupportTLS1      *string   `json:"ssl_support_tls1,omitempty"`
		SSLSupportTLS11     *string   `json:"ssl_support_tls1_1,omitempty"`
		SSLSupportTLS12     *string   `json:"ssl_support_tls1_2,omitempty"`
		StrictVerify        *bool     `json:"strict_verify,omitempty"`
	} `json:"ssl"`
	TCP struct {
		Nagle *bool `json:"nagle,omitempty"`
	} `json:"tcp"`
	UDP struct {
		AcceptFrom     *string `json:"accept_from,omitempty"`
		AcceptFromMask *string `json:"accept_from_mask,omitempty"`
	} `json:"udp"`
}

type NodesTable []Node

type Node struct {
	Node     *string `json:"node,omitempty"`
	Priority *int    `json:"priority,omitempty"`
	State    *string `json:"state,omitempty"`
	Weight   *int    `json:"weight,omitempty"`
}

func (r *Pool) endpoint() string {
	return "pools"
}

func (r *Pool) String() string {
	s, _ := jsonMarshal(r)
	return string(s)
}

func (r *Pool) decode(data []byte) error {
	return json.Unmarshal(data, &r)
}

func NewPool(name string) *Pool {
	r := new(Pool)
	r.setName(name)
	return r
}

func (c *Client) GetPool(name string) (*Pool, *http.Response, error) {
	r := NewPool(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (c *Client) ListPools() ([]string, *http.Response, error) {
	return c.List(&Pool{})
}
