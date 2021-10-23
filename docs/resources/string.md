---
page_title: "string Resource - terraform-provider-scratch"
subcategory: ""
description: |-
	The string resource allows you to test against string.
---

# Resource `scratch_string`

The string resource allows you to test against strings.

## Example Usage

```terraform
locals {
	string  = {
		cat = "meow"
		dog = "woof"
		cow = "moo"
	}
}

resource "scratch_string" "this" {
	for_each = local.string

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
