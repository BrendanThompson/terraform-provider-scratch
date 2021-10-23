---
page_title: "list Resource - terraform-provider-scratch"
subcategory: ""
description: |-
	The list resource allows you to test against list.
---

# Resource `scratch_list`

The list resource allows you to test against lists.

## Example Usage

```terraform
locals {
	list  = {
		cat = ["ginger", "black"]
		dog = ["brown", "white"]
	}
}

resource "scratch_list" "this" {
	for_each = local.list

	in = each.value

	description = format("Key: %s, Value: %s", each.key, each.value[0])
}

```

## Argument Reference

- `in` - (Required) Input value for the resource.
- `description` - (Optional) A description that can be easily consumed.

## Attribute Reference

In addition to all the arguments above, the following attributes are exported.

- `id` - The ID of the resource.
