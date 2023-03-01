---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_antivirus_profile"
description: |-
  Configure fortiadc security antivirus profile info.
---

# fortiadc_security_antivirus_profile
Configure fortiadc security antivirus profile info.

## Example Usage
```hcl
resource "fortiadc_security_antivirus_profile" "av" {
	mkey = "custom_av"
	uncomp_nest_limit = "5"
	scan_bzip2 = "enable"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - antivirus profile name.
* `uncomp_nest_limit` - maximum number of allowed levels of compression nesting that attempt to decompress, range is 2 - 100, default 2. (2,100)
* `streaming_content_bypass` - allow streaming content to be bypassed rather than buffered, default enable. Valid values: enable/disable.
* `emulator` - enable/disable Win32 emulator, default disable to improve throughput. Valid values: enable/disable.
* `scan_bzip2` - enable to scan archives using BZIP2 algorithm, default disable. Valid values: enable/disable.
* `av_block_log` - enable/disable logging files that are blocked by antivirus, default enable. Valid values: enable/disable.
* `comments` - antivirus profile comments. 
* `analytics_db` - use signature database from FortiSandbox to supplement the AV signature databases, default disable. Valid values: enable/disable.
* `oversize_limit` - maximum in-memory file size that will be scanned in kilobytes (KB), range is 1 - 12000000 KB, default 1024 KB. (1,12000000)
* `av_virus_log` - enable/disable logging for antivirus scanning, default enable. Valid values: enable/disable.
* `analytics_max_upload` - maximum upload size to FortiSandbox in kilobytes (KB), range is 1 - 2048 KB, default 1024 KB. (1,2048)
* `uncomp_size_limit` - maximum size in megabytes (MB) of memory buffer that use to temporarily decompress, range is 1 - 2000 MB, default 2 MB. (1,2000)
* `oversize` - bypass means bypass oversize files, log means logging oversize files, block means block oversize files, default bypass. Valid values: 1:log, 0:bypass, 2:block .
* `ftgd_analytics` - control which files are uploaded to FortiSandbox, would affect the destination of quarantine, default disable. Valid values: 1:suspicious, 0:disable, 2:all .
* `options` - avmonitor means scanning and logging, quarantine means scanning, logging and quarantine infected files, default avmonitor. Valid values: 1:quarantine, 0:avmonitor .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Security Antivirus Profile can be imported using any of these accepted formats:
```
$ terraform import fortiadc_security_antivirus_profile.labelname {{mkey}}
```
