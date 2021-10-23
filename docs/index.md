---
page_title: "Provider: Scratch"
subcategory: ""
description: |-
	Terraform provider to assist with the development of complex for_each loops.
---

# Scratch Provider

The aim of this provider is to assist with the development phase of complex `for_each` loops on
resources and data sources. By utilising this providers resources you will easily be able
to see the output of any `for_each` loops in an easy to consume way.

## Example Usage

```terraform
provider "scratch" {}
```

## Schema

### Optional

- **use** (Bool, Optional) If the provider should be used (defaults to `true`)
