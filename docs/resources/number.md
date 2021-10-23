---
page_title: "number Resource - terraform-provider-scratch"
subcategory: ""
description: |-
	The number resource allows you to test against number.
---

# Resource `scratch_number`

The number resource allows you to test against numbers.

## Example Usage

```terraform
locals {
	number  = {
		cat = 1
		dog = 2
		cow = 3
	}
}

resource "scratch_number" "this" {
	for_each = local.number

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
