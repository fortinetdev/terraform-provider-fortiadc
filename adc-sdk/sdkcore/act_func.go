package forticlient

func (c *FortiSDKClient) DeleteSystemGlobal(mkey string, vdom string) (err error) {
	//No unset API for SystemGlobal
	return
}

func (c *FortiSDKClient) UpdateSystemGlobal(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_global"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemGlobal(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_global"

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterAccessList(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/router_access_list"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterAccessList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/router_access_list"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterAccessList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_access_list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterAccessList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_access_list"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceContentRoutingChildMatchCondition(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_content_routing_child_match_condition"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceContentRoutingChildMatchCondition(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_content_routing_child_match_condition"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceContentRoutingChildMatchCondition(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_content_routing_child_match_condition"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceContentRoutingChildMatchCondition(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_content_routing_child_match_condition"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemHealthCheckScript(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_health_check_script"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemHealthCheckScript(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_health_check_script"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemHealthCheckScript(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_health_check_script"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemHealthCheckScript(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_health_check_script"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateLocal(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_certificate_local"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateLocal(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_certificate_local"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateLocal(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_certificate_local"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateLocal(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_certificate_local"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceClonePoolChildPoolMember(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_clone_pool_child_pool_member"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceClonePoolChildPoolMember(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_clone_pool_child_pool_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceClonePoolChildPoolMember(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_clone_pool_child_pool_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceClonePoolChildPoolMember(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_clone_pool_child_pool_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateCaGroupChildGroupMember(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_certificate_ca_group_child_group_member"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateCaGroupChildGroupMember(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_certificate_ca_group_child_group_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateCaGroupChildGroupMember(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_certificate_ca_group_child_group_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateCaGroupChildGroupMember(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_certificate_ca_group_child_group_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemIspAddrChildAddress(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_isp_addr_child_address"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemIspAddrChildAddress(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_isp_addr_child_address"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemIspAddrChildAddress(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_isp_addr_child_address"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemIspAddrChildAddress(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_isp_addr_child_address"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateCaGroup(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_certificate_ca_group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateCaGroup(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_certificate_ca_group"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateCaGroup(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_certificate_ca_group"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateCaGroup(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_certificate_ca_group"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemService(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_service"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemService(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_service"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemService(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_service"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemService(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_service"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceProfileChildServerRequestHeaderErase(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_profile_child_server_request_header_erase"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceProfileChildServerRequestHeaderErase(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_profile_child_server_request_header_erase"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceProfileChildServerRequestHeaderErase(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_profile_child_server_request_header_erase"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceProfileChildServerRequestHeaderErase(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_profile_child_server_request_header_erase"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceCachingChildUriExcludeList(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_caching_child_uri_exclude_list"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceCachingChildUriExcludeList(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_caching_child_uri_exclude_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceCachingChildUriExcludeList(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_caching_child_uri_exclude_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceCachingChildUriExcludeList(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_caching_child_uri_exclude_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceMethod(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_method"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceMethod(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_method"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceMethod(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_method"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceMethod(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_method"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemOverlayTunnelChildRemoteHost(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_overlay_tunnel_child_remote_host"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemOverlayTunnelChildRemoteHost(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_overlay_tunnel_child_remote_host"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemOverlayTunnelChildRemoteHost(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_overlay_tunnel_child_remote_host"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemOverlayTunnelChildRemoteHost(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_overlay_tunnel_child_remote_host"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceProfileChildMssqlUserPassword(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_profile_child_mssql_user_password"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceProfileChildMssqlUserPassword(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_profile_child_mssql_user_password"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceProfileChildMssqlUserPassword(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_profile_child_mssql_user_password"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceProfileChildMssqlUserPassword(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_profile_child_mssql_user_password"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceGeoipList(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_geoip_list"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceGeoipList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_geoip_list"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceGeoipList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_geoip_list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceGeoipList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_geoip_list"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemOverlayTunnel(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_overlay_tunnel"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemOverlayTunnel(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_overlay_tunnel"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemOverlayTunnel(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_overlay_tunnel"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemOverlayTunnel(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_overlay_tunnel"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemAddress(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_address"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemAddress(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_address"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemAddress(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_address"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemAddress(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_address"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceDecompressionChildUriList(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_decompression_child_uri_list"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceDecompressionChildUriList(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_decompression_child_uri_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceDecompressionChildUriList(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_decompression_child_uri_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceDecompressionChildUriList(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_decompression_child_uri_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalancePagespeedChildPageControl(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_pagespeed_child_page_control"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalancePagespeedChildPageControl(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_pagespeed_child_page_control"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalancePagespeedChildPageControl(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_pagespeed_child_page_control"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalancePagespeedChildPageControl(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_pagespeed_child_page_control"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceProfile(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceProfile(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_profile"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceProfile(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_profile"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceProfile(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_profile"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalancePool(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_pool"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalancePool(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_pool"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalancePool(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_pool"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalancePool(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_pool"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemIspAddr(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_isp_addr"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemIspAddr(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_isp_addr"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemIspAddr(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_isp_addr"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemIspAddr(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_isp_addr"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceAuthPolicy(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_auth_policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceAuthPolicy(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_auth_policy"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceAuthPolicy(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_auth_policy"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceAuthPolicy(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_auth_policy"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemSdnConnector(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_sdn_connector"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemSdnConnector(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_sdn_connector"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemSdnConnector(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_sdn_connector"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemSdnConnector(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_sdn_connector"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterOspfChildHaRouterIdList(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/router_ospf_child_ha_router_id_list"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterOspfChildHaRouterIdList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/router_ospf_child_ha_router_id_list"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterOspfChildHaRouterIdList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_ospf_child_ha_router_id_list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterOspfChildHaRouterIdList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_ospf_child_ha_router_id_list"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemHealthCheck(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_health_check"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemHealthCheck(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_health_check"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemHealthCheck(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_health_check"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemHealthCheck(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_health_check"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterAccessListChildRule(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/router_access_list_child_rule"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterAccessListChildRule(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/router_access_list_child_rule"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterAccessListChildRule(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_access_list_child_rule"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterAccessListChildRule(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_access_list_child_rule"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateOcsp(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_certificate_ocsp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateOcsp(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_certificate_ocsp"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateOcsp(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_certificate_ocsp"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateOcsp(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_certificate_ocsp"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemStreamScripting(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_stream_scripting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemStreamScripting(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_stream_scripting"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemStreamScripting(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_stream_scripting"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemStreamScripting(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_stream_scripting"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemInterface(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_interface"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemInterface(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_interface"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemInterface(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_interface"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemInterface(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_interface"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalancePersistenceChildRadiusAttribute(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_persistence_child_radius_attribute"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalancePersistenceChildRadiusAttribute(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_persistence_child_radius_attribute"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalancePersistenceChildRadiusAttribute(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_persistence_child_radius_attribute"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalancePersistenceChildRadiusAttribute(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_persistence_child_radius_attribute"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterPolicy(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/router_policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterPolicy(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/router_policy"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterPolicy(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_policy"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterPolicy(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_policy"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceIppoolChildNodeMember(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_ippool_child_node_member"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceIppoolChildNodeMember(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_ippool_child_node_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceIppoolChildNodeMember(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_ippool_child_node_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceIppoolChildNodeMember(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_ippool_child_node_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceCaptchaProfile(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_captcha_profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceCaptchaProfile(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_captcha_profile"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceCaptchaProfile(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_captcha_profile"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceCaptchaProfile(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_captcha_profile"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemVdom(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_vdom"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemVdom(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_vdom"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemVdom(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_vdom"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemVdom(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_vdom"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemMailserver(mkey string, vdom string) (err error) {
	//No unset API for SystemMailserver
	return
}

func (c *FortiSDKClient) UpdateSystemMailserver(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_mailserver?mkey=-1"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemMailserver(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_mailserver"

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterOspfChildOspfInterface(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/router_ospf_child_ospf_interface"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterOspfChildOspfInterface(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/router_ospf_child_ospf_interface"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterOspfChildOspfInterface(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_ospf_child_ospf_interface"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterOspfChildOspfInterface(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_ospf_child_ospf_interface"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemIspProvince(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_isp_province"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemIspProvince(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_isp_province"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemIspProvince(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_isp_province"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemIspProvince(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_isp_province"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateCertificateVerify(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_certificate_certificate_verify"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateCertificateVerify(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_certificate_certificate_verify"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateCertificateVerify(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_certificate_certificate_verify"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateCertificateVerify(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_certificate_certificate_verify"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceDecompressionChildContentTypes(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_decompression_child_content_types"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceDecompressionChildContentTypes(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_decompression_child_content_types"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceDecompressionChildContentTypes(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_decompression_child_content_types"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceDecompressionChildContentTypes(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_decompression_child_content_types"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemAutoBackup(mkey string, vdom string) (err error) {
	//No unset API for SystemAutoBackup
	return
}

func (c *FortiSDKClient) UpdateSystemAutoBackup(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_auto_backup?mkey=-1"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemAutoBackup(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_auto_backup"

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemInterfaceChildTrustIpList(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_interface_child_trust_ip_list"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemInterfaceChildTrustIpList(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_interface_child_trust_ip_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemInterfaceChildTrustIpList(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_interface_child_trust_ip_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemInterfaceChildTrustIpList(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_interface_child_trust_ip_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterOspfChildArea(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/router_ospf_child_area"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterOspfChildArea(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/router_ospf_child_area"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterOspfChildArea(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_ospf_child_area"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterOspfChildArea(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_ospf_child_area"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceProfileChildMysqlRule(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_profile_child_mysql_rule"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceProfileChildMysqlRule(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_profile_child_mysql_rule"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceProfileChildMysqlRule(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_profile_child_mysql_rule"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceProfileChildMysqlRule(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_profile_child_mysql_rule"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemVdomLink(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_vdom_link"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemVdomLink(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_vdom_link"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemVdomLink(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_vdom_link"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemVdomLink(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_vdom_link"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceRealServerSslProfile(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_real_server_ssl_profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceRealServerSslProfile(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_real_server_ssl_profile"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceRealServerSslProfile(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_real_server_ssl_profile"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceRealServerSslProfile(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_real_server_ssl_profile"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceErrorPage(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_error_page"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceErrorPage(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_error_page"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceErrorPage(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_error_page"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceErrorPage(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_error_page"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceCertificateCaching(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_certificate_caching"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceCertificateCaching(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_certificate_caching"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceCertificateCaching(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_certificate_caching"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceCertificateCaching(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_certificate_caching"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceL2ExceptionList(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_l2_exception_list"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceL2ExceptionList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_l2_exception_list"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceL2ExceptionList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_l2_exception_list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceL2ExceptionList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_l2_exception_list"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterBgpChildNeighbor(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/router_bgp_child_neighbor"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterBgpChildNeighbor(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/router_bgp_child_neighbor"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterBgpChildNeighbor(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_bgp_child_neighbor"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterBgpChildNeighbor(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_bgp_child_neighbor"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateIntermediateCa(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_certificate_intermediate_ca"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateIntermediateCa(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_certificate_intermediate_ca"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateIntermediateCa(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_certificate_intermediate_ca"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateIntermediateCa(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_certificate_intermediate_ca"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemScheduleGroupChildScheduleMember(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_schedule_group_child_schedule_member"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemScheduleGroupChildScheduleMember(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_schedule_group_child_schedule_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemScheduleGroupChildScheduleMember(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_schedule_group_child_schedule_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemScheduleGroupChildScheduleMember(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_schedule_group_child_schedule_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalancePagespeed(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_pagespeed"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalancePagespeed(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_pagespeed"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalancePagespeed(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_pagespeed"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalancePagespeed(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_pagespeed"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceCompressionChildContentTypes(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_compression_child_content_types"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceCompressionChildContentTypes(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_compression_child_content_types"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceCompressionChildContentTypes(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_compression_child_content_types"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceCompressionChildContentTypes(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_compression_child_content_types"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterAccessList6ChildRule(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/router_access_list6_child_rule"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterAccessList6ChildRule(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/router_access_list6_child_rule"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterAccessList6ChildRule(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_access_list6_child_rule"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterAccessList6ChildRule(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_access_list6_child_rule"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateFirewallNatSnat(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/firewall_nat_snat"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteFirewallNatSnat(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/firewall_nat_snat"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateFirewallNatSnat(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/firewall_nat_snat"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadFirewallNatSnat(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/firewall_nat_snat"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceDecompression(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_decompression"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceDecompression(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_decompression"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceDecompression(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_decompression"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceDecompression(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_decompression"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceCompressionChildUriList(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_compression_child_uri_list"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceCompressionChildUriList(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_compression_child_uri_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceCompressionChildUriList(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_compression_child_uri_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceCompressionChildUriList(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_compression_child_uri_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceClientSslProfile(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_client_ssl_profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceClientSslProfile(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_client_ssl_profile"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceClientSslProfile(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_client_ssl_profile"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceClientSslProfile(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_client_ssl_profile"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateLocalCertGroupChildGroupMember(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_certificate_local_cert_group_child_group_member"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateLocalCertGroupChildGroupMember(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_certificate_local_cert_group_child_group_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateLocalCertGroupChildGroupMember(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_certificate_local_cert_group_child_group_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateLocalCertGroupChildGroupMember(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_certificate_local_cert_group_child_group_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterOspf(mkey string, vdom string) (err error) {
	//No unset API for RouterOspf
	return
}

func (c *FortiSDKClient) UpdateRouterOspf(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_ospf?mkey=-1"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterOspf(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_ospf"

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterPrefixList6(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/router_prefix_list6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterPrefixList6(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/router_prefix_list6"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterPrefixList6(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_prefix_list6"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterPrefixList6(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_prefix_list6"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceProfileChildServerResponseHeaderErase(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_profile_child_server_response_header_erase"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceProfileChildServerResponseHeaderErase(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_profile_child_server_response_header_erase"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceProfileChildServerResponseHeaderErase(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_profile_child_server_response_header_erase"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceProfileChildServerResponseHeaderErase(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_profile_child_server_response_header_erase"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSecurityZtnaProfile(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/security_ztna_profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSecurityZtnaProfile(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/security_ztna_profile"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSecurityZtnaProfile(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/security_ztna_profile"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSecurityZtnaProfile(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/security_ztna_profile"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceContentRewritingChildMatchCondition(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_content_rewriting_child_match_condition"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceContentRewritingChildMatchCondition(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_content_rewriting_child_match_condition"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceContentRewritingChildMatchCondition(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_content_rewriting_child_match_condition"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceContentRewritingChildMatchCondition(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_content_rewriting_child_match_condition"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemDns(mkey string, vdom string) (err error) {
	//No unset API for SystemDns
	return
}

func (c *FortiSDKClient) UpdateSystemDns(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_dns"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemDns(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_dns"

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemDnsVdom(mkey string, vdom string) (err error) {
	//No unset API for SystemDnsVdom
	return
}

func (c *FortiSDKClient) UpdateSystemDnsVdom(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_dns_vdom"
	path += "?vdom=" + escapeURLString(vdom)
	path += "&mkey=-1"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemDnsVdom(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_dns_vdom"
	path += "?vdom=" + escapeURLString(vdom)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceProfileChildClientResponseHeaderInsert(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_profile_child_client_response_header_insert"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceProfileChildClientResponseHeaderInsert(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_profile_child_client_response_header_insert"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceProfileChildClientResponseHeaderInsert(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_profile_child_client_response_header_insert"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceProfileChildClientResponseHeaderInsert(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_profile_child_client_response_header_insert"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemServicegrp(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_servicegrp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemServicegrp(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_servicegrp"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemServicegrp(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_servicegrp"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemServicegrp(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_servicegrp"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceWebFilterProfileChildCategoryMembers(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_web_filter_profile_child_category_members"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceWebFilterProfileChildCategoryMembers(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_web_filter_profile_child_category_members"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceWebFilterProfileChildCategoryMembers(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_web_filter_profile_child_category_members"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceWebFilterProfileChildCategoryMembers(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_web_filter_profile_child_category_members"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceCachingChildDynCacheList(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_caching_child_dyn_cache_list"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceCachingChildDynCacheList(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_caching_child_dyn_cache_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceCachingChildDynCacheList(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_caching_child_dyn_cache_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceCachingChildDynCacheList(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_caching_child_dyn_cache_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceProfileChildServerRequestHeaderInsert(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_profile_child_server_request_header_insert"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceProfileChildServerRequestHeaderInsert(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_profile_child_server_request_header_insert"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceProfileChildServerRequestHeaderInsert(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_profile_child_server_request_header_insert"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceProfileChildServerRequestHeaderInsert(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_profile_child_server_request_header_insert"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemAddrgrp(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_addrgrp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemAddrgrp(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_addrgrp"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemAddrgrp(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_addrgrp"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemAddrgrp(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_addrgrp"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterAccessList6(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/router_access_list6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterAccessList6(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/router_access_list6"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterAccessList6(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_access_list6"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterAccessList6(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_access_list6"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceProfileChildClientResponseHeaderErase(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_profile_child_client_response_header_erase"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceProfileChildClientResponseHeaderErase(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_profile_child_client_response_header_erase"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceProfileChildClientResponseHeaderErase(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_profile_child_client_response_header_erase"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceProfileChildClientResponseHeaderErase(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_profile_child_client_response_header_erase"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterPrefixListChildRule(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/router_prefix_list_child_rule"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterPrefixListChildRule(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/router_prefix_list_child_rule"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterPrefixListChildRule(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_prefix_list_child_rule"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterPrefixListChildRule(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_prefix_list_child_rule"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateConfigSyncList(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/config_sync_list"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteConfigSyncList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/config_sync_list"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateConfigSyncList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/config_sync_list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadConfigSyncList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/config_sync_list"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterPrefixList(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/router_prefix_list"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterPrefixList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/router_prefix_list"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterPrefixList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_prefix_list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterPrefixList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_prefix_list"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemAzureLbBackendIp(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_azure_lb_backend_ip"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemAzureLbBackendIp(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_azure_lb_backend_ip"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemAzureLbBackendIp(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_azure_lb_backend_ip"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemAzureLbBackendIp(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_azure_lb_backend_ip"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterBgpChildNetwork(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/router_bgp_child_network"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterBgpChildNetwork(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/router_bgp_child_network"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterBgpChildNetwork(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_bgp_child_network"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterBgpChildNetwork(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_bgp_child_network"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateIntermediateCaGroupChildGroupMember(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_certificate_intermediate_ca_group_child_group_member"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateIntermediateCaGroupChildGroupMember(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_certificate_intermediate_ca_group_child_group_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateIntermediateCaGroupChildGroupMember(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_certificate_intermediate_ca_group_child_group_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateIntermediateCaGroupChildGroupMember(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_certificate_intermediate_ca_group_child_group_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceAllowlist(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_allowlist"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceAllowlist(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_allowlist"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceAllowlist(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_allowlist"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceAllowlist(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_allowlist"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterBgp(mkey string, vdom string) (err error) {
	//No unset API for RouterBgp
	return
}

func (c *FortiSDKClient) UpdateRouterBgp(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_bgp?mkey=-1&"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterBgp(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_bgp"

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterBgpChildHaRouterIdList(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/router_bgp_child_ha_router_id_list"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterBgpChildHaRouterIdList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/router_bgp_child_ha_router_id_list"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterBgpChildHaRouterIdList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_bgp_child_ha_router_id_list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterBgpChildHaRouterIdList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_bgp_child_ha_router_id_list"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterPrefixList6ChildRule(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/router_prefix_list6_child_rule"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterPrefixList6ChildRule(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/router_prefix_list6_child_rule"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterPrefixList6ChildRule(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_prefix_list6_child_rule"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterPrefixList6ChildRule(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_prefix_list6_child_rule"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceSchedulePool(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_schedule_pool"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceSchedulePool(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_schedule_pool"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceSchedulePool(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_schedule_pool"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceSchedulePool(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_schedule_pool"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceWebCategory(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_web_category"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceWebCategory(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_web_category"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceWebCategory(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_web_category"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceWebCategory(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_web_category"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemHa(mkey string, vdom string) (err error) {
	//No unset API for SystemHa
	return
}

func (c *FortiSDKClient) UpdateSystemHa(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_ha"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemHa(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_ha"

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalancePagespeedProfile(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_pagespeed_profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalancePagespeedProfile(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_pagespeed_profile"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalancePagespeedProfile(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_pagespeed_profile"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalancePagespeedProfile(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_pagespeed_profile"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterMd5OspfChildMd5Member(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/router_md5_ospf_child_md5_member"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterMd5OspfChildMd5Member(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/router_md5_ospf_child_md5_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterMd5OspfChildMd5Member(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_md5_ospf_child_md5_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterMd5OspfChildMd5Member(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_md5_ospf_child_md5_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceContentRewriting(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_content_rewriting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceContentRewriting(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_content_rewriting"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceContentRewriting(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_content_rewriting"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceContentRewriting(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_content_rewriting"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateUserAdfsPublish(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/user_adfs_publish"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteUserAdfsPublish(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/user_adfs_publish"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateUserAdfsPublish(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/user_adfs_publish"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadUserAdfsPublish(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/user_adfs_publish"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceVirtualServer(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_virtual_server"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceVirtualServer(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_virtual_server"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceVirtualServer(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_virtual_server"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceVirtualServer(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_virtual_server"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceProfileChildMysqlUserPassword(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_profile_child_mysql_user_password"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceProfileChildMysqlUserPassword(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_profile_child_mysql_user_password"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceProfileChildMysqlUserPassword(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_profile_child_mysql_user_password"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceProfileChildMysqlUserPassword(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_profile_child_mysql_user_password"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceCaching(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_caching"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceCaching(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_caching"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceCaching(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_caching"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceCaching(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_caching"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterMd5Ospf(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/router_md5_ospf"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterMd5Ospf(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/router_md5_ospf"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterMd5Ospf(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_md5_ospf"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterMd5Ospf(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_md5_ospf"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalancePersistence(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_persistence"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalancePersistence(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_persistence"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalancePersistence(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_persistence"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalancePersistence(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_persistence"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalancePersistenceChildIso8583Bitmap(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_persistence_child_iso8583_bitmap"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalancePersistenceChildIso8583Bitmap(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_persistence_child_iso8583_bitmap"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalancePersistenceChildIso8583Bitmap(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_persistence_child_iso8583_bitmap"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalancePersistenceChildIso8583Bitmap(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_persistence_child_iso8583_bitmap"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceHttp2Profile(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_http2_profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceHttp2Profile(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_http2_profile"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceHttp2Profile(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_http2_profile"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceHttp2Profile(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_http2_profile"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemScripting(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_scripting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemScripting(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_scripting"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemScripting(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_scripting"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemScripting(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_scripting"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceRealServer(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_real_server"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceRealServer(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_real_server"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceRealServer(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_real_server"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceRealServer(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_real_server"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemScheduleGroup(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_schedule_group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemScheduleGroup(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_schedule_group"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemScheduleGroup(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_schedule_group"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemScheduleGroup(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_schedule_group"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceProfileChildMysqlSharding(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_profile_child_mysql_sharding"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceProfileChildMysqlSharding(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_profile_child_mysql_sharding"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceProfileChildMysqlSharding(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_profile_child_mysql_sharding"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceProfileChildMysqlSharding(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_profile_child_mysql_sharding"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateCrl(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_certificate_crl"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateCrl(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_certificate_crl"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateCrl(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_certificate_crl"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateCrl(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_certificate_crl"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemAddress6(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_address6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemAddress6(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_address6"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemAddress6(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_address6"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemAddress6(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_address6"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSecurityAntivirusProfile(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/security_antivirus_profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSecurityAntivirusProfile(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/security_antivirus_profile"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSecurityAntivirusProfile(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/security_antivirus_profile"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSecurityAntivirusProfile(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/security_antivirus_profile"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceClonePool(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_clone_pool"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceClonePool(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_clone_pool"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceClonePool(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_clone_pool"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceClonePool(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_clone_pool"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalancePagespeedChildResourceControl(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_pagespeed_child_resource_control"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalancePagespeedChildResourceControl(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_pagespeed_child_resource_control"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalancePagespeedChildResourceControl(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_pagespeed_child_resource_control"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalancePagespeedChildResourceControl(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_pagespeed_child_resource_control"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemTrafficGroup(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_traffic_group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemTrafficGroup(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_traffic_group"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemTrafficGroup(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_traffic_group"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemTrafficGroup(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_traffic_group"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterOspfChildNetwork(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/router_ospf_child_network"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterOspfChildNetwork(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/router_ospf_child_network"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterOspfChildNetwork(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_ospf_child_network"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterOspfChildNetwork(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_ospf_child_network"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterIsp(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/router_isp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterIsp(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/router_isp"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterIsp(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_isp"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterIsp(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_isp"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemInterfaceChildHaNodeIpList(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_interface_child_ha_node_ip_list"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemInterfaceChildHaNodeIpList(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_interface_child_ha_node_ip_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemInterfaceChildHaNodeIpList(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_interface_child_ha_node_ip_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemInterfaceChildHaNodeIpList(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_interface_child_ha_node_ip_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSecurityWafProfile(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/security_waf_profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSecurityWafProfile(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/security_waf_profile"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSecurityWafProfile(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/security_waf_profile"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSecurityWafProfile(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/security_waf_profile"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceProfileChildClientRequestHeaderErase(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_profile_child_client_request_header_erase"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceProfileChildClientRequestHeaderErase(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_profile_child_client_request_header_erase"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceProfileChildClientRequestHeaderErase(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_profile_child_client_request_header_erase"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceProfileChildClientRequestHeaderErase(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_profile_child_client_request_header_erase"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemAddrgrp6(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_addrgrp6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemAddrgrp6(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_addrgrp6"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemAddrgrp6(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_addrgrp6"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemAddrgrp6(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_addrgrp6"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateLocalCertGroup(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_certificate_local_cert_group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateLocalCertGroup(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_certificate_local_cert_group"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateLocalCertGroup(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_certificate_local_cert_group"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateLocalCertGroup(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_certificate_local_cert_group"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemTimeManual(mkey string, vdom string) (err error) {
	//No unset API for SystemTimeManual
	return
}

func (c *FortiSDKClient) UpdateSystemTimeManual(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_time_manual"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemTimeManual(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_time_manual"

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceContentRouting(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_content_routing"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceContentRouting(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_content_routing"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceContentRouting(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_content_routing"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceContentRouting(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_content_routing"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateIntermediateCaGroup(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_certificate_intermediate_ca_group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateIntermediateCaGroup(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_certificate_intermediate_ca_group"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateIntermediateCaGroup(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_certificate_intermediate_ca_group"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateIntermediateCaGroup(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_certificate_intermediate_ca_group"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSecurityDosDosProtectionProfile(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/security_dos_dos_protection_profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSecurityDosDosProtectionProfile(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/security_dos_dos_protection_profile"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSecurityDosDosProtectionProfile(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/security_dos_dos_protection_profile"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSecurityDosDosProtectionProfile(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/security_dos_dos_protection_profile"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceWebSubCategory(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_web_sub_category"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceWebSubCategory(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_web_sub_category"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceWebSubCategory(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_web_sub_category"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceWebSubCategory(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_web_sub_category"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceCompression(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_compression"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceCompression(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_compression"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceCompression(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_compression"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceCompression(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_compression"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemInterfaceChildSecondaryIpList(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_interface_child_secondary_ip_list"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemInterfaceChildSecondaryIpList(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_interface_child_secondary_ip_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemInterfaceChildSecondaryIpList(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_interface_child_secondary_ip_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemInterfaceChildSecondaryIpList(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_interface_child_secondary_ip_list"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateFirewallVip(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/firewall_vip"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteFirewallVip(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/firewall_vip"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateFirewallVip(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/firewall_vip"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadFirewallVip(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/firewall_vip"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceProfileChildServerResponseHeaderInsert(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_profile_child_server_response_header_insert"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceProfileChildServerResponseHeaderInsert(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_profile_child_server_response_header_insert"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceProfileChildServerResponseHeaderInsert(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_profile_child_server_response_header_insert"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceProfileChildServerResponseHeaderInsert(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_profile_child_server_response_header_insert"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceWebFilterProfile(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_web_filter_profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceWebFilterProfile(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_web_filter_profile"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceWebFilterProfile(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_web_filter_profile"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceWebFilterProfile(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_web_filter_profile"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceL2ExceptionListChildMember(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_l2_exception_list_child_member"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceL2ExceptionListChildMember(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_l2_exception_list_child_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceL2ExceptionListChildMember(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_l2_exception_list_child_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceL2ExceptionListChildMember(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_l2_exception_list_child_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterStatic(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/router_static"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterStatic(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/router_static"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterStatic(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/router_static"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterStatic(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/router_static"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSecurityIpsProfile(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/security_ips_profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSecurityIpsProfile(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/security_ips_profile"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSecurityIpsProfile(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/security_ips_profile"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSecurityIpsProfile(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/security_ips_profile"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalancePoolChildPoolMember(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_pool_child_pool_member"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalancePoolChildPoolMember(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_pool_child_pool_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalancePoolChildPoolMember(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_pool_child_pool_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalancePoolChildPoolMember(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_pool_child_pool_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceProfileChildClientRequestHeaderInsert(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_profile_child_client_request_header_insert"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceProfileChildClientRequestHeaderInsert(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_profile_child_client_request_header_insert"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceProfileChildClientRequestHeaderInsert(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_profile_child_client_request_header_insert"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceProfileChildClientRequestHeaderInsert(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_profile_child_client_request_header_insert"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateLoadBalanceIppool(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/load_balance_ippool"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteLoadBalanceIppool(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_ippool"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateLoadBalanceIppool(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/load_balance_ippool"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadLoadBalanceIppool(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_ippool"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateCertificateVerifyChildGroupMember(pkey string, params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_certificate_certificate_verify_child_group_member"
	path += "?pkey=" + escapeURLString(pkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateCertificateVerifyChildGroupMember(pkey, mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_certificate_certificate_verify_child_group_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateCertificateVerifyChildGroupMember(pkey string, params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_certificate_certificate_verify_child_group_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateCertificateVerifyChildGroupMember(pkey, mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_certificate_certificate_verify_child_group_member"
	path += "?pkey=" + escapeURLString(pkey)
	path += "&mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemHaChildRemoteIpMonitorList(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/system_ha_child_remote_ip_monitor_list"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemHaChildRemoteIpMonitorList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_ha_child_remote_ip_monitor_list"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemHaChildRemoteIpMonitorList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_ha_child_remote_ip_monitor_list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemHaChildRemoteIpMonitorList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_ha_child_remote_ip_monitor_list"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemTimeNtp(mkey string, vdom string) (err error) {
	//No unset API for SystemTimeNtp
	return
}

func (c *FortiSDKClient) UpdateSystemTimeNtp(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_time_ntp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemTimeNtp(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_time_ntp"

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) StandardCreate(params *map[string]interface{}, vdom string, path string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) StandardDelete(mkey string, vdom string, path string) (err error) {
	HTTPMethod := "DELETE"

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) StandardUpdate(params *map[string]interface{}, mkey string, vdom string, path string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) StandardRead(mkey string, vdom string, path string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) DeleteGlobalLoadBalanceSetting(mkey string, vdom string) (err error) {
	return
}

func (c *FortiSDKClient) UpdateGlobalLoadBalanceSetting(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/global_load_balance_setting?vdom=test&mkey=-1"
	//path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadGlobalLoadBalanceSetting(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/global_load_balance_setting?vdom=test"

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}
