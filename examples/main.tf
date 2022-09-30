terraform {
  required_providers {
    scratch = {
      source = "BrendanThompson/scratch"
    }
  }
}

provider "scratch" {
  # use = true
}

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
  in = "hello, moto!"
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
  in          = ["cat", "dog", "mouse", "elephant"]
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
  in = { name = "Brendan Thompson", Age = "32" }
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

resource "scratch_block" "single" {
  description = "example"
  in {
    string = "one"
    number = 666777
    bool   = false
  }
}

output "block_single" {
  value = scratch_block.single
}

resource "scratch_block" "multi" {
  dynamic "in" {
    for_each = local.map

    content {
      string = in.value.name

    }
  }
}

output "block_multi" {
  value = scratch_block.multi
}

data "scratch_string" "this" {
  in = "meow"
}

output "data_string" {
  value = data.scratch_string.this
}
