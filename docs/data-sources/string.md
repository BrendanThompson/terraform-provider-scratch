---
page_title: "string Data source - terraform-provider-scratch"
subcategory: ""
description: |-
  The string data sources allows you to get what you pass in back.
---

# Data Source `scratch_string`

This allows you to pass in a string and have it returned to you.

## Example Usage

```terraform
data "scratch_string" "this" {
  in = "meow"
}
```

## Argument Reference

- `in` - (Required) Input value for the data source.

## Attribute Reference

In addition to all the arguments above, the following attributes are exported.

- `id` - The ID of the resource.
- `out` - The value of `in`.
