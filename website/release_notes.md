Release Notes
==============================

Release Version 1.3.2
--------------------

Release Targets
---------------

FortiADC Terraform 1.3.2

Features
---------------

- Add new modules:

  - fortiadc_security_waf
  - fortiadc_security_waf_adaptive_learning.markdown
  - fortiadc_security_waf_adaptive_learning_child_url_list
  - fortiadc_security_waf_advanced_protection
  - fortiadc_security_waf_advanced_protection_child_advanced_protection_rule
  - fortiadc_security_waf_allowed_origin
  - fortiadc_security_waf_allowed_origin_child_allowed_origin_list
  - fortiadc_security_waf_api_discovery
  - fortiadc_security_waf_api_discovery_child_api_security_rule
  - fortiadc_security_waf_cors_headers
  - fortiadc_security_waf_cors_headers_child_cors_headers_list
  - fortiadc_security_waf_cors_protection
  - fortiadc_security_waf_cors_protection_child_cors_rule_list


- FortiADC version: v7.6 and v8.0

Notice
---------------

- For detailed configurations of the new modules, please refer to the internal examples within each module.

==============================

Release Version 1.3.1
--------------------

Release Targets
---------------

FortiADC Terraform 1.3.1

Features
---------------

- Add new modules:

  - fortiadc_global_load_balance_virtual_server_content_routing_attach

- Updated modules:

  - fortiadc_load_balance_http2_profile


- FortiADC version: v7.6 and v8.0

Notice
---------------

- For detailed configurations of the new modules, please refer to the internal examples within each module.

==============================

Release Version 1.3.0
--------------------

Release Targets
---------------

FortiADC Terraform 1.3.0

Features
---------------

- Add new modules:

  - fortiadc_firewall_policy
  - fortiadc_firewall_policy_child_rule
  - fortiadc_global_dns_server_address_group
  - fortiadc_global_dns_server_address_group_child_member
  - fortiadc_global_dns_server_general
  - fortiadc_global_dns_server_policy
  - fortiadc_global_dns_server_remote_dns_server
  - fortiadc_global_dns_server_remote_dns_server_child_member
  - fortiadc_global_dns_server_response_rate_limit
  - fortiadc_global_dns_server_zone
  - fortiadc_global_load_balance_data_center
  - fortiadc_global_load_balance_host
  - fortiadc_global_load_balance_host_child_virtual_server_pool_list
  - fortiadc_global_load_balance_servers
  - fortiadc_global_load_balance_servers_child_virtual_server_list
  - fortiadc_global_load_balance_setting
  - fortiadc_global_load_balance_virtual_server_pool
  - fortiadc_global_load_balance_virtual_server_pool_child_member
  - fortiadc_security_waf_exception
  - fortiadc_security_waf_exception_child_exception_rule
  - fortiadc_security_waf_sig_profile_category_id_list_group
  - fortiadc_security_waf_sig_profile_sub_category_id_list_group
  - fortiadc_security_waf_threshold_based_detection
  - fortiadc_security_waf_url_protection
  - fortiadc_security_waf_url_protection_child_file_extension_rule
  - fortiadc_security_waf_url_protection_child_url_access_rule
  - fortiadc_security_waf_web_attack_signature
  - fortiadc_security_waf_web_attack_signature_child_category
  - fortiadc_security_waf_web_attack_signature_child_signature
  - fortiadc_security_waf_web_attack_signature_child_sub_category


- FortiADC version: v7.4 and v7.6

Notice
---------------

- For detailed configurations of the new modules, please refer to the internal examples within each module.

==============================

Release Version 1.2.0
--------------------

Release Targets
---------------

FortiADC Terraform 1.2.0

Features
---------------

- Add support for the proxy environment variable no_proxy.
  - If HTTP_PROXY, HTTPS_PROXY, and NO_PROXY are all present, the values of the uppercase parameters take precedence.
  - A comma-separated list of host names that should not go through any proxy is set in no_proxy (only an asterisk, '*', matches all hosts).  

- Fix bugs of fadcos_system_stream_scripting and fadcons_system_scripting

- Revise the error examples for fortiadc_certificate_intmed_caupload and fortiadc_certificate_caupload

- For changes in DNS settings starting from FortADC 7.4:  
  - for the root vdom, please use fadcos_system_global  
  - for non-root vdoms, please use fadcos_system_dns_vdom  

- Add new modules:

  - fadcos_system_dns_vdom

- FortiADC version: v7.4 and v7.6

Notice
---------------

- For detailed configurations of fadcos_system_stream_scripting, fadcos_system_scritping and fadcos_system_dns_vdom, please refer to the internal example within each module.

