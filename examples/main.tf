terraform {
  required_providers {
    scratch = {
      source = "registry.terraform.io/brendanthompson/scratch"
    }
  }
}

provider "scratch" {}

locals {
  string = {
    cat  = "meow"
    dog  = "woof"
    pig  = "oink"
    lion = "roar"
  }

  number = {
    cat = 10
    dog = 8
    pig = 6
  }

  bool = {
    cat = true
    dog = true
    pig = false
  }

  list = {
    cat = ["ginger", "black", "white"]
    dog = ["brown", "black", "golden"]
  }

  map = {
    cat = {
      name = "Delilah"
    }
    dog = {
      name = "Gracie"
    }
  }
}

resource "scratch_string" "single" {
  in = "hello"
}

output "string_single" {
  value = scratch_string.single.in
}

resource "scratch_string" "multi" {
  for_each = local.string

  description = format("The animal '%s' goes '%s'", each.key, each.value)
  in          = each.value
}

output "string_multi" {
  value = scratch_string.multi[*]
}

resource "scratch_number" "single" {
  in = 8
}

output "number_single" {
  value = scratch_number.single.in
}

resource "scratch_number" "multi" {
  for_each = local.number

  description = format("%s - %s", each.key, each.value)
  in          = each.value
}

output "number_multi" {
  value = scratch_number.multi[*]
}

resource "scratch_bool" "single" {
  in = true
}

output "bool_single" {
  value = scratch_bool.single.in
}

resource "scratch_bool" "multi" {
  for_each = local.bool

  description = format("%s - %s", each.key, each.value)
  in          = each.value
}

resource "scratch_list" "single" {
  description = "exmaple"
  in          = ["cat", "dog", "mouse"]
}

output "list_single" {
  value = scratch_list.single.in
}

resource "scratch_list" "multi" {
  for_each = local.list

  description = format("%s - %s", each.key, each.value[0])
  in          = each.value
}

output "list_multi" {
  value = scratch_list.multi[*]
}

resource "scratch_map" "single" {
  in = { name = "Brendan", Age = "31" }
}

output "map_single" {
  value = scratch_map.single.in
}

resource "scratch_map" "multi" {
  for_each = local.map

  description = format("%s - %s", each.key, each.value.name)
  in          = each.value
}

output "map_multi" {
  value = scratch_map.multi[*]
}

resource "scratch_dynamic" "single" {
  description = "example"
  in {
    first = "one"
    # second = "two"
    third = "three"
  }
}

output "dynamic_single" {
  value = scratch_dynamic.single
}

resource "scratch_dynamic" "multi" {
  dynamic "in" {
    for_each = local.map

    content {
      first = in.value.name
    }
  }
}

output "dynamic_multi" {
  value = scratch_dynamic.multi
}
