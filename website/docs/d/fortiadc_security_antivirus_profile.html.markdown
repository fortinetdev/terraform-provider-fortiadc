---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_antivirus_profile"
description: |-
  Get information on an fortiadc security antivirus profile
---

# Data Source: fortiadc_security_antivirus_profile
Use this data source to get information on an fortiadc security antivirus profile

## Example Usage

```hcl
 data "fortiadc_security_antivirus_profile" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_security_antivirus_profile.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  security antivirus profile.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - antivirus profile name.
* `uncomp_nest_limit` - maximum number of allowed levels of compression nesting that attempt to decompress, range is 2 - 100, default 2. (2,100)
* `streaming_content_bypass` - allow streaming content to be bypassed rather than buffered, default enable. 
* `emulator` - enable/disable Win32 emulator, default disable to improve throughput. 
* `scan_bzip2` - enable to scan archives using BZIP2 algorithm, default disable. 
* `av_block_log` - enable/disable logging files that are blocked by antivirus, default enable. 
* `comments` - antivirus profile comments. 
* `analytics_db` - use signature database from FortiSandbox to supplement the AV signature databases, default disable. 
* `oversize_limit` - maximum in-memory file size that will be scanned in kilobytes (KB), range is 1 - 12000000 KB, default 1024 KB. (1,12000000)
* `av_virus_log` - enable/disable logging for antivirus scanning, default enable. 
* `analytics_max_upload` - maximum upload size to FortiSandbox in kilobytes (KB), range is 1 - 2048 KB, default 1024 KB. (1,2048)
* `uncomp_size_limit` - maximum size in megabytes (MB) of memory buffer that use to temporarily decompress, range is 1 - 2000 MB, default 2 MB. (1,2000)
* `oversize` - bypass means bypass oversize files, log means logging oversize files, block means block oversize files, default bypass. 
* `ftgd_analytics` - control which files are uploaded to FortiSandbox, would affect the destination of quarantine, default disable. 
* `options` - avmonitor means scanning and logging, quarantine means scanning, logging and quarantine infected files, default avmonitor. 

