---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_persistence_child_iso8583_bitmap"
description: |-
  Configure fortiadc load-balance persistence info.
---

# fortiadc_load_balance_persistence_child_iso8583_bitmap
Configure fortiadc load-balance persistence info.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - iso8583 bitmap id.
* `type` - iso8583 bitmap type. Valid values: 67:67-ext-payment-code, 68:68-recv-inst-cy-code, 69:69-settle-inst-cy-code, 80:80-inquiries-num, 24:24-function-code, 20:20-PAN-ext-cy-code, 21:21-fwd-inst-cy-code, 22:22-POS-data, 23:23-card-seq-num, 42:42-card-accept-id-code, 40:40-serv-restrict-code, 41:41-card-accept-term-id, 3:3-process-code, 2:2-primary-acct-num, 7:7-date-time-trans, 77:77-debits-reversal-num, 76:76-debits-num, 75:75-credits-reversal-num, 74:74-credits-num, 73:73-date-action, 79:79-transfer-reversal-num, 78:78-tranfer-num, 11:11-sys-trace-audit-num, 12:12-date-tm-loc-trans, 19:19-acq-inst-cy-code, 18:18-merchant-type, 37:37-retrieval-ref-num, 52:52-PIN-data, 33:33-fwd-inst-id-code, 32:32-acq-inst-id-code .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Persistence Child Iso8583 Bitmap can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_persistence_child_iso8583_bitmap.labelname {{mkey}}
```
