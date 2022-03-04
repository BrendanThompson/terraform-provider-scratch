---
page_title: "dynamic Resource - terraform-provider-scratch"
subcategory: ""
description: |-
	The dynamic resource allows you to test dynamic Blocks.
---

# Resource `scratch_dynamic`

The dynamic resource allows you to test dynamic blocks.

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

resource "scratch_dynamic" "multi" {
  dynamic "in" {
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

An `in` block represents a block that can be used for `dynamic`:

- `first` - (Optional) First argument.
- `second` - (Optional) Second argument.
- `third` - (Optional) Third argument.

## Attribute Reference

In addition to all the arguments above, the following attributes are exported.

- `id` - The ID of the resource.
