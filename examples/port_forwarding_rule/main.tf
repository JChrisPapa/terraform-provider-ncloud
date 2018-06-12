provider "ncloud" {
  access_key = "${var.access_key}"
  secret_key = "${var.secret_key}"
  region = "${var.region}"
}

data "ncloud_port_forwarding_rules" "rules" {}

resource "ncloud_instance" "instance" {
  "server_name" = "${var.server_name}"
  "server_image_product_code" = "${var.server_image_product_code}"
  "server_product_code" = "${var.server_product_code}"
}

resource "ncloud_port_forwarding_rule" "test" {
  "port_forwarding_configuration_no" = "${data.ncloud_port_forwarding_rules.rules.id}"
  "server_instance_no" = "${ncloud_instance.instance.id}"
  "port_forwarding_external_port" = "6022"
  "port_forwarding_internal_port" = "22"
}