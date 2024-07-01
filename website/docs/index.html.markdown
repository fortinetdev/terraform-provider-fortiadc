---
layout: "fortiadc"
page_title: "Provider: FortiADC"
sidebar_current: "docs-fortiadc-index"
description: |-
  The FortiADC provider interacts with FortiADC.
---

# FortiADC Provider

The FortiADC provider is used to interact with the resources supported by FortiADC. We need to configure the provider with the proper credentials before it can be used. Please use the navigation on the left to read more details about the available resources.

## Example Usage

```hcl
# Configure the FortiADC Provider for FortiADC
provider "fortiadc" {
  hostname     = "192.168.52.177"
  token        = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
  insecure     = "false"
  cabundlefile = "/path/yourCA.crt"
}

# Create a Static Route Item
resource "fortiadc_router_static" "test1" {
  mkey    = "1"
  dst     = "110.2.2.122/32"
  gateway = "2.2.2.2"
  # ...
}
```

If it is used for testing, you can set `insecure` to "true" and unset `cabundlefile` to quickly set the provider up, for example:

```hcl
provider "fortiadc" {
  hostname = "192.168.52.177"
  token    = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
  insecure = "true"
}
```

Please refer to the Argument Reference below for more help on `insecure` and `cabundlefile`.

## Authentication

The FortiADC provider offers a means of providing credentials for authentication. The following methods are supported:

- Static credentials
- Environment variables

### Static credentials

Static credentials can be provided by adding a `token` key in-line in the FortiADC provider block.

Usage:

```hcl
provider "fortiadc" {
  hostname     = "192.168.52.177"
  token        = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
  insecure     = "false"
  cabundlefile = "/path/yourCA.crt"
}
```

### Generate an API token for FortiADC

See the left navigation: `Guides` -> `Generate an API token for FortiADC`.

### Environment variables

You can provide your credentials via the `FORTIADC_ACCESS_HOSTNAME`, `FORTIADC_ACCESS_TOKEN`, `FORTIADC_INSECURE` and `FORTIADC_CA_CABUNDLE` environment variables. Note that setting your FortiADC credentials using static credentials variables will override the environment variables.

Usage:

```shell
$ export "FORTIADC_ACCESS_HOSTNAME"="192.168.52.177"
$ export "FORTIADC_INSECURE"="false"
$ export "FORTIADC_ACCESS_TOKEN"="09m441wrwc10yGyrtQ4nk6mjbqcfz9"
$ export "FORTIADC_CA_CABUNDLE"="/path/yourCA.crt"
```

Then configure the FortiADC Provider as following:

```hcl
provider "fortiadc" {}

# Create a Static Route Item
resource "fortiadc_router_static" "test1" {
  dst       = "110.2.2.122/32"
  gateway   = "2.2.2.2"
  mkey      = "3"
  # …
}
```

## VDOM

If the FortiADC unit is running in VDOM mode, the `vdom` configuration needs to be added.

Usage:

```hcl
provider "fortiadc" {
  hostname     = "192.168.52.177"
  token        = "q3Hs49jxts195gkd9Hjsxnjtmr6k39"
  insecure     = "false"
  cabundlefile = "/path/yourCA.crt"
  vdom         = "vdomtest"
}

# Create a Static Route Item
resource "fortiadc_router_static" "test1" {
  dst       = "110.2.2.122/32"
  gateway   = "2.2.2.2"
  mkey      = "3"
  # …
}
```

By default, each resource inherits the provider's global vdom settings, but it can also set its own vdom through the `vdom` of each resource. See the `vdom` argument of each resource for details.



## Release
Check out the FortiADC provider release notes and additional information from: [the FortiADC provider releases](https://github.com/fortinetdev/terraform-provider-fortiadc/releases).

## Versioning

The provider can cover FortiADC 7.6 versions, the configuration of all parameters should be based on the relevant FortiADC version manual.
