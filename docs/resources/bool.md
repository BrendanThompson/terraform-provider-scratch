---
page_title: "bool Resource - terraform-provider-scratch"
subcategory: ""
description: |-
	The bool resource allows you to test against booleans.
---

# Resource `scratch_bool`

The bool resource allows you to test against bools.

## Example Usage

```terraform
locals {
	bool  = {
		cat = true
		dog = true
		cow = false
	}
}

resource "scratch_bool" "this" {
	for_each = local.bool

	in = each.value

	description = format("Key: %s, Value: %s", each.key, each.value)
}

```

## Argument Reference

- `in` - (Required) Input value for the resource.
- `description` - (Optional) A description that can be easily consumed.

## Attribute Reference

In addition to all the arguments above, the following attributes are exported.

- `id` - The ID of the resource.
