packer {
  required_plugins {
    amazon = {
      version = ">= 0.0.2"
      source  = "github.com/hashicorp/amazon"
    }
  }
}

variable "region" {
  type    = string
  default = "us-west-2"
}

variable "source_ami" {
  type    = string
  default = "ami-017fecd1353bcc96e"
}

variable "instance-size" {
  type    = string
  default = "t2.micro"
}

source "amazon-ebs" "bitcoin-24-0" {
  ami_name      = "edo-bitcoin-24.0"
  instance_type = var.instance-size
  region        = var.region
  source_ami    = var.source_ami
  ssh_username  = "ubuntu"
}

build {
  name = "edo-bitcoin-24.0"

  sources = [
    "source.amazon-ebs.bitcoin-24-0"
  ]
}
