---
page_title: "map Resource - terraform-provider-scratch"
subcategory: ""
description: |-
	The map resource allows you to test against map.
---

# Resource `scratch_map`

The map resource allows you to test against maps.

## Example Usage

```terraform
locals {
	map  = {
		cat = {
			name = "Delilah"
			colour = "Ginger"
		}
		dog = {
			name = "Dennis"
			colour = "Brown"
		}
	}
}

resource "scratch_map" "this" {
	for_each = local.map

	in = each.value

	description = format("Key: %s, Value: %s", each.key, each.value.name)
}

```

## Argument Reference

- `in` - (Required) Input value for the resource.
- `description` - (Optional) A description that can be easily consumed.

## Attribute Reference

In addition to all the arguments above, the following attributes are exported.

- `id` - The ID of the resource.
