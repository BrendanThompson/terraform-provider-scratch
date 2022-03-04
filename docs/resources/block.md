---
page_title: "block Resource - terraform-provider-scratch"
subcategory: ""
description: |-
	The block resource allows you to test dynamic blocks.
---

# Resource `scratch_block`

The block resource allows you to test dynamic blocks.

## Example Usage

```terraform
locals {
	map  = {
		cat = {
			name = "Delilah"
		}
		dog = {
			name = "Dennis"
		}
	}
}

resource "scratch_block" "multi" {
  block "in" {
    for_each = local.map

    content {
      first = in.value.name
    }
  }
}

```

## Argument Reference

- `in` - (Optional) Input value block for the resource.
- `description` - (Optional) A description that can be easily consumed.

---

An `in` block represents a block that can be used for `block`:

- `string` - (Optional) String argument.
- `number` - (Optional) Number argument.
- `bool` - (Optional) Bool argument.

## Attribute Reference

In addition to all the arguments above, the following attributes are exported.

- `id` - The ID of the resource.
