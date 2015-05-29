* pool: Set default value for passive_monitoring

## 2.2.2

* virtual_server: Set connection timeout default to documented value
  (300)

## 2.2.1

* Update godeps for terraform v0.5.2

## 2.2.0

* virtual_server: Set a default value for connection_timeout

## 2.1.1

* Update godeps for terraform v0.4.2

## 2.1.0

* Update godeps for terraform v0.4.0

## 2.0.0

* TypeList parameters (monitor/script_arguments and
  virtual_server/request_rules, response_rules) are no longer computed

## 1.5.2

* Update godeps for go-stingray v1.1.0

## 1.5.1

* Update godeps for go-stingray v1.0.2

## 1.5.0

* virtual_server: Support ssl_server_cert_host_mapping
* traffic_ip_group: Only parse IP addresses if `valid_networks` is set

## 1.4.2

* Fix problem where default value of stingray_virtual_server
  listen_any was not being set

## 1.4.1

* Update godeps for terraform v0.3.7

## 1.4.0

* Add valid_networks option to provider

## 1.3.0

* Support provider configuration via environment variables

## 1.2.0

* pool: Support load_balancing_algorithm

## 1.1.2

* Remove workaround for hashicorp/terraform#452 (fixed upstream)

## 1.1.1

* Update godeps for terraform v0.3.5

## 1.1.0

* virtual_server: Support syslog_format

## 1.0.0

 * Initial release
