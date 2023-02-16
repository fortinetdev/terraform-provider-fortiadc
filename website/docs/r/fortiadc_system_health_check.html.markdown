---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_health_check"
description: |-
  Configure fortiadc system health-check info.
---

# fortiadc_system_health_check
Configure fortiadc system health-check info.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - health-check name.
* `counter_value` - snmp counter value: range(0, 2147483647). 
* `service_name` - the service name for connecting oracle db. 
* `basedn` - baseDN. 
* `pwd_type` - RADIUS password type. Valid values: 1:chap-password, 0:user-password .
* `dest_addr6` - dest-addr6. 
* `mssql_row` - row. 
* `disk` - disk threshold for snmp. (1,99)
* `row` - the row for sql searching. 
* `http_version` - http version. Valid values: 1:http_1.0, 2:http_1.1 .
* `passive` - FTP passive enable/disable. Valid values: enable/disable.
* `mssql_column` - column. 
* `status_code` - status-code. 
* `folder` - Mail folder name. 
* `addr_type` - address type. Valid values: 1:ipv4, 2:ipv6 .
* `mssql_receive_string` - mssql receive string. 
* `oracle_receive_string` - the receive-string for connecting oracle db. 
* `host_addr` - host-addr. 
* `password` - password. 
* `sip_request_type` - type of SIP request. Valid values: 1:register, 2:options .
* `remote_password` - remote-password. 
* `timeout` - timeout. (1,3600)
* `acct_appid` - diameter acct application id. 

* `port` - port. (0,65535)
* `allow_ssl_version` - SSL/TLS versions allowed. Valid values: 8:tlsv1.1, 2:sslv3, 4:tlsv1.0, 32:tlsv1.3, 16:tlsv1.2 .
* `remote_port` - remote-port. (0,65535)
* `connect_string` - the connect data for connecting oracle db. 
* `domain_name` - domain-name. 
* `version` - snmp version. Valid values: 1:v1, 2:v2c .
* `product_name` - product name. 
* `nas_ip` - RADIUS NAS IP address. 
* `binddn` - bindDN. 
* `mem` - mem threshold for snmp. (1,99)
* `host_addr6` - host-addr6. 
* `http_extra_string` - additional-string. 
* `column` - the column for sql seaching. 
* `http_connect` - http connect type. Valid values: 1:no_connect, 3:remote_connect, 2:local_connect .
* `auth_appid` - diameter auth application id. 
* `agent_type` - snmp agent type. Valid values: 1:UCD, 2:WIN2000 .
* `filter` - filter. 
* `remote_host` - remote-host. 
* `connect_type` - the type of oracle connect data. Valid values: 1:service_name, 3:connect_string, 2:sid .
* `method_type` - method type. Valid values: 1:http_get, 2:http_head .
* `ssl_ciphers` - ssl ciphers including Perfect Forward Secrecy. Valid values: 32768:, 4194304:Perfect Forward Secrecy cipher, 1099511627776:Perfect Forward Secrecy cipher, 131072:Perfect Forward Secrecy cipher, 262144:Perfect Forward Secrecy cipher, 64:Perfect Forward Secrecy cipher, 65536:, 256:Perfect Forward Secrecy cipher, 33554432:, 8192:Perfect Forward Secrecy cipher, 16777216:, 536870912:Perfect Forward Secrecy cipher, 1048576:Perfect Forward Secrecy cipher, 1:Perfect Forward Secrecy cipher, 2097152:Perfect Forward Secrecy cipher, 2:Perfect Forward Secrecy cipher, 4:Perfect Forward Secrecy cipher, 8:Perfect Forward Secrecy cipher, 16384:, 17179869184:Non Encryption Ciphers, 2199023255552:Perfect Forward Secrecy cipher, 4096:Perfect Forward Secrecy cipher, 134217728:, 137438953472:Perfect Forward Secrecy cipher, 274877906944:Perfect Forward Secrecy cipher, 8388608:, 67108864:Perfect Forward Secrecy cipher, 524288:Perfect Forward Secrecy cipher, 549755813888:Perfect Forward Secrecy cipher, 128:Perfect Forward Secrecy cipher, 2147483648:, 1024:Perfect Forward Secrecy cipher, 268435456:, 34359738368:Perfect Forward Secrecy cipher, 16:Perfect Forward Secrecy cipher, 32:Perfect Forward Secrecy cipher, 2048:Perfect Forward Secrecy cipher, 1073741824:Perfect Forward Secrecy cipher, 512:Perfect Forward Secrecy cipher, 68719476736:Perfect Forward Secrecy cipher .
* `community` - snmp community. 
* `script` - health-check script. 
* `dest_addr` - dest-addr. 
* `type` - health-check type. Valid values: 24:mysql, 25:diameter, 26:oracle, 27:script, 20:sip, 21:sip-tcp, 22:snmp-custom, 23:rtsp, 28:ldap, 29:mssql, 1:icmp, 3:tcp, 2:tcp-echo, 5:https, 4:http, 7:radius, 6:dns, 9:pop3, 8:smtp, 11:radacct, 10:imap4, 12:ftp, 15:tcpssl, 14:tcphalf, 17:ssh, 16:snmp, 19:udp, 18:l2-detection .
* `username` - username. 
* `rtsp_method_type` - type of RTSP method. Valid values: 1:options, 2:describe .
* `mssql_send_string` - mssql send string. 
* `string_value` - snmp string value. 
* `down_retry` - retry. (1,10)
* `origin_realm` - origin realm. 
* `origin_host` - origin host. 
* `secret_key` - RADIUS Secret Key. 
* `compare_type` - snmp compare type. Valid values: 1:equal, 3:greater, 2:less .
* `match_type` - match-type. Valid values: 1:match_string, 3:match_all, 2:match_status .
* `attribute` - attribute. 
* `local_cert` - Local certificate reference. 
* `mysql_server_type` - type of MySQL server. Valid values: 1:primary, 2:secondary .
* `file` - File path on FTP server. 
* `dest_addr_type` - destination address type. Valid values: 1:ipv4, 2:ipv6 .
* `cpu_weight` - snmp cpu weight. (0,200)
* `host_ip_addr` - host ip addr. 
* `host_ip6_addr` - host ip6 addr. 
* `hostname` - hostname. 
* `mem_weight` - snmp mem weight. (0,200)
* `sid` - the sid for connecting oracle db. 

* `disk_weight` - snmp disk weight. (0,200)
* `receive_string` - receive-string. 
* `oid` - snmp oid. 
* `value_type` - snmp value type. Valid values: 2:ASN_INTEGER, 4:ASN_OCTET_STR, 71:ASN_UINTEGER, 65:ASN_COUNTER .
* `radius_reject` - enbale radius reject. Valid values: 1:disable, 2:enable .
* `rtsp_describe_url` - URL for DESCRIBE method. 
* `database` - datebase name of the mssql server. 
* `oracle_send_string` - the send string for connecting oracle db. 
* `send_string` - send-string. 
* `interval` - interval. (1,3600)
* `up_retry` - up-retry. (1,10)
* `remote_username` - remote-username. 
* `cpu` - cpu threshold for snmp. (1,99)
* `vendor_id` - vendor id. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Health Check can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_health_check.labelname {{mkey}}
```
