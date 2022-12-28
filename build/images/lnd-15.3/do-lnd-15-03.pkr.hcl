packer {
  required_plugins {
    digitalocean = {
      version = ">= 1.0.4"
      source  = "github.com/digitalocean/digitalocean"
    }
  }
}

source "digitalocean" "lnd-15-03" {
  image                = "ubuntu-22-04-x64"
  region               = "nyc3"
  size                 = "s-1vcpu-1gb"
  ssh_username         = "root"
  snapshot_name        = "lnd-15-05"
  ssh_key_id           = 36505685
  ssh_private_key_file = "/Users/ramin/.ssh/id_ed25519"
}

build {
  sources = ["source.digitalocean.lnd-15-03"]
}
