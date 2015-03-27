package stingray

import (
	"encoding/json"
	"net/http"
)

// A VirtualServer is a Stingray virtual server.
type VirtualServer struct {
	jsonResource            `json:"-"`
	VirtualServerProperties `json:"properties"`
}

type VirtualServerProperties struct {
	Aptimizer struct {
		Enabled *bool         `json:"enabled,omitempty"`
		Profile *ProfileTable `json:"profile,omitempty"`
	} `json:"aptimizer"`
	Basic struct {
		AddClusterIP       *bool   `json:"add_cluster_ip,omitempty"`
		AddXForwardedFor   *bool   `json:"add_x_forwarded_for,omitempty"`
		AddXForwardedProto *bool   `json:"add_x_forwarded_proto,omitempty"`
		BandwidthClass     *string `json:"bandwidth_class,omitempty"`
		CloseWithRST       *bool   `json:"close_with_rst,omitempty"`
		// To disable a rule, add a / before the name (e.g.
		// "example" -> "/example")
		CompletionRules      *[]string `json:"completionrules,omitempty"`
		ConnectTimeout       *int      `json:"connect_timeout,omitempty"`
		Enabled              *bool     `json:"enabled,omitempty"`
		FTPForceServerSecure *bool     `json:"ftp_force_server_secure,omitempty"`
		GLBServices          *[]string `json:"glb_services,omitempty"`
		ListenOnAny          *bool     `json:"listen_on_any,omitempty"`
		ListenOnHosts        *[]string `json:"listen_on_hosts,omitempty"`
		ListenOnTrafficIPs   *[]string `json:"listen_on_traffic_ips,omitempty"`
		Note                 *string   `json:"note,omitempty"`
		Pool                 *string   `json:"pool,omitempty"` // mandatory
		Port                 *int      `json:"port,omitempty"` // mandatory
		ProtectionClass      *string   `json:"protection_class,omitempty"`
		Protocol             *string   `json:"protocol,omitempty"`
		RequestRules         *[]string `json:"request_rules,omitempty"`
		ResponseRules        *[]string `json:"response_rules,omitempty"`
		SMLClass             *string   `json:"slm_class,omitempty"`
		SoNagle              *bool     `json:"so_nagle,omitempty"`
		SSLClientCertHeaders *string   `json:"ssl_client_cert_headers,omitempty"`
		SSLDecrypt           *bool     `json:"ssl_decrypt,omitempty"`
	} `json:"basic"`
	Connection struct {
		Keepalive              *bool   `json:"keepalive,omitempty"`
		KeepaliveTimeout       *int    `json:"keepalive_timeout,omitempty"`
		MaxClientBuffer        *int    `json:"max_client_buffer,omitempty"`
		MaxServerBuffer        *int    `json:"max_server_buffer,omitempty"`
		MaxTransactionDuration *int    `json:"max_transaction_duration,omitempty"`
		ServerFirstBanner      *string `json:"server_first_banner,omitempty"`
		Timeout                *int    `json:"timeout,omitempty"`
	} `json:"connection"`
	ConnectionErrors struct {
		ErrorFile *string `json:"error_file,omitempty"`
	} `json:"connection_errors"`
	Cookie struct {
		Domain      *string `json:"domain,omitempty"`
		NewDomain   *string `json:"new_domain,omitempty"`
		PathRegex   *string `json:"path_regex,omitempty"`
		PathReplace *string `json:"path_replace,omitempty"`
		Secure      *string `json:"secure,omitempty"`
	} `json:"cookie"`
	FTP struct {
		DataSourcePort    *int  `json:"data_source_port,omitempty"`
		ForceClientSecure *bool `json:"force_client_secure,omitempty"`
		PortRangeHigh     *int  `json:"port_range_high,omitempty"`
		PortRangeLow      *int  `json:"port_range_low,omitempty"`
		SSLData           *bool `json:"ssl_data,omitempty"`
	} `json:"ftp"`
	Gzip struct {
		CompressLevel *int      `json:"compress_level,omitempty"`
		Enabled       *bool     `json:"enabled,omitempty"`
		IncludeMIME   *[]string `json:"include_mime,omitempty"`
		MaxSize       *int      `json:"max_size,omitempty"`
		MinSize       *int      `json:"min_size,omitempty"`
		NoSize        *bool     `json:"no_size,omitempty"`
	} `json:"gzip"`
	HTTP struct {
		ChunkOverheadForwarding *string `json:"chunk_overhead_forwarding,omitempty"`
		LocationRegex           *string `json:"location_regex,omitempty"`
		LocationReplace         *string `json:"location_replace,omitempty"`
		LocationRewrite         *string `json:"location_rewrite,omitempty"`
		MIMEDefault             *string `json:"mime_default,omitempty"`
		MIMEDetect              *bool   `json:"mime_detect,omitempty"`
	} `json:"http"`
	KerberosProtocolTransition struct {
		Enabled   *bool   `json:"enabled,omitempty"`
		Principal *string `json:"principal,omitempty"`
		Target    *string `json:"target,omitempty"`
	} `json:"kerberos_protocol_transition"`
	Log struct {
		ClientConnectionFailures  *bool   `json:"client_connection_failures,omitempty"`
		Enabled                   *bool   `json:"enabled,omitempty"`
		Filename                  *string `json:"filename,omitempty"`
		Format                    *string `json:"format,omitempty"`
		Save                      *bool   `json:"save_all,omitempty"`
		ServerConnectionFailures  *bool   `json:"server_connection_failures,omitempty"`
		SessionPersistenceVerbose *bool   `json:"session_persistence_verbose,omitempty"`
		SSLFailures               *bool   `json:"ssl_failures,omitempty"`
	} `json:"log"`
	RecentConnections struct {
		Enabled *bool `json:"enabled,omitempty"`
		SaveAll *bool `json:"save_all,omitempty"`
	} `json:"recent_connections"`
	RequestTracing struct {
		Enabled *bool `json:"enabled,omitempty"`
		TraceIO *bool `json:"trace_io,omitempty"`
	} `json:"request_tracing"`
	RTSP struct {
		StreamingPortRangeHigh *int `json:"streaming_port_range_high,omitempty"`
		StreamingPortRangeLow  *int `json:"streaming_port_range_low,omitempty"`
		StreamingTimeout       *int `json:"streaming_timeout,omitempty"`
	} `json:"rtsp"`
	SIP struct {
		DangerousRequests      *string `json:"dangerous_requests,omitempty"`
		FollowRoute            *bool   `json:"follow_route,omitempty"`
		MaxConnectionMem       *int    `json:"max_connection_mem,omitempty"`
		Mode                   *string `json:"mode,omitempty"`
		RewriteURI             *bool   `json:"rewrite_uri,omitempty"`
		StreamingPortRangeHigh *int    `json:"streaming_port_range_high,omitempty"`
		StreamingPortRangeLow  *int    `json:"streaming_port_range_low,omitempty"`
		StreamingTimeout       *int    `json:"streaming_timeout,omitempty"`
		TimeoutMessages        *bool   `json:"timeout_messages,omitempty"`
		TransactionTimeout     *int    `json:"transaction_timeout,omitempty"`
	} `json:"sip"`
	SMTP struct {
		ExpectStartTLS *bool `json:"expect_starttls,omitempty"`
	} `json:"smtp"`
	SSL struct {
		AddHTTPHeaders         *bool                       `json:"add_http_headers,omitempty"`
		ClientCertCAs          *[]string                   `json:"client_cert_cas,omitempty"`
		IssuedCertsNeverExpire *[]string                   `json:"issued_certs_never_expire,omitempty"`
		OCSPEnable             *bool                       `json:"ocsp_enable,omitempty"`
		OCSPIssuers            *OCSPIssuersTable           `json:"ocsp_issuers,omitempty"`
		OCSPMaxResponseAge     *int                        `json:"ocsp_max_response_age,omitempty"`
		OCSPStapling           *bool                       `json:"ocsp_stapling,omitempty"`
		OCSPTimeTolerance      *int                        `json:"ocsp_time_tolerance,omitempty"`
		OCSPTimeout            *int                        `json:"ocsp_timeout,omitempty"`
		PreferSSLv3            *bool                       `json:"prefer_sslv3,omitempty"`
		RequestClientCert      *string                     `json:"request_client_cert,omitempty"`
		SendCloseAlerts        *bool                       `json:"send_close_alerts,omitempty"`
		ServerCertDefault      *string                     `json:"server_cert_default,omitempty"`
		ServerCertHostMapping  *ServerCertHostMappingTable `json:"server_cert_host_mapping,omitempty"`
		SignatureAlgorithms    *string                     `json:"signature_algorithms,omitempty"`
		SSLCiphers             *string                     `json:"ssl_ciphers,omitempty"`
		SSLSupportSSL2         *string                     `json:"ssl_support_ssl2,omitempty"`
		SSLSupportSSL3         *string                     `json:"ssl_support_ssl3,omitempty"`
		SSLSupportTLS1         *string                     `json:"ssl_support_tls1,omitempty"`
		SSLSupportTLS11        *string                     `json:"ssl_support_tls1_1,omitempty"`
		SSLSupportTLS12        *string                     `json:"ssl_support_tls1_2,omitempty"`
		TrustMagic             *bool                       `json:"trust_magic,omitempty"`
	} `json:"ssl"`
	Syslog struct {
		Enabled     *bool   `json:"enabled,omitempty"`
		Format      *string `json:"format,omitempty"`
		IPEndPoint  *string `json:"ip_end_point,omitempty"`
		MsgLenLimit *int    `json:"msg_len_limit,omitempty"`
	} `json:"syslog"`
	TCP struct {
		ProxyClose *bool `json:"proxy_close,omitempty"`
	} `json:"tcp"`
	UDP struct {
		EndPointPersistence       *bool `json:"end_point_persistence,omitempty"`
		PortSMP                   *bool `json:"port_smp,omitempty"`
		ResponseDatagramsExpected *int  `json:"response_datagrams_expected,omitempty"`
		Timeout                   *int  `json:"timeout,omitempty"`
	} `json:"udp"`
	WebCache struct {
		ControlOut    *string `json:"control_out,omitempty"`
		Enabled       *bool   `json:"enabled,omitempty"`
		ErrorPageTime *int    `json:"error_page_time,omitempty"`
		MaxTime       *int    `json:"max_time,omitempty"`
		RefreshTime   *int    `json:"refresh_time,omitempty"`
	} `json:"web_cache"`
}

type ProfileTable []Profile

type Profile struct {
	Name *string   `json:"name,omitempty"`
	URLs *[]string `json:"urls,omitempty"`
}

type OCSPIssuersTable []OCSPIssuer

type OCSPIssuer struct {
	AIA           *bool   `json:"aia,omitempty"`
	Issuer        *string `json:"issuer,omitempty"`
	Nonce         *string `json:"nonce,omitempty"`
	Required      *string `json:"required,omitempty"`
	ResponderCert *string `json:"responder_cert,omitempty"`
	Signer        *string `json:"signer,omitempty"`
	URL           *string `json:"url,omitempty"`
}

type ServerCertHostMappingTable []ServerCertHostMapping

type ServerCertHostMapping struct {
	Certificate *string `json:"certificate,omitempty"`
	Host        *string `json:"host,omitempty"`
}

func (r *VirtualServer) endpoint() string {
	return "virtual_servers"
}

func (r *VirtualServer) String() string {
	s, _ := jsonMarshal(r)
	return string(s)
}

func (r *VirtualServer) decode(data []byte) error {
	return json.Unmarshal(data, &r)
}

func NewVirtualServer(name string) *VirtualServer {
	r := new(VirtualServer)
	r.setName(name)
	return r
}

func (c *Client) GetVirtualServer(name string) (*VirtualServer, *http.Response, error) {
	r := NewVirtualServer(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (c *Client) ListVirtualServers() ([]string, *http.Response, error) {
	return c.List(&VirtualServer{})
}
