Release Notes
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

