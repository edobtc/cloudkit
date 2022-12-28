packer {
  required_plugins {
    digitalocean = {
      version = ">= 1.0.4"
      source  = "github.com/digitalocean/digitalocean"
    }
  }
}

source "digitalocean" "lnd-15-5" {
  image         = "ubuntu-22-04-x64"
  region        = "nyc3"
  size          = "s-2vcpu-2gb"
  ssh_username  = "root"
  snapshot_name = "lnd-15.5"
}

build {
  sources = ["source.digitalocean.lnd-15-5"]

  # base install
  provisioner "shell" {
    script = "../../provisioners/base.sh"
  }

  # systemd to start docker-compose/lnd on start
  provisioner "file" {
    source      = "../../provisioners/systemd/lnd.service"
    destination = "/etc/systemd/system/lnd.service"
  }

  # add the lnd config files
  provisioner "file" {
    source      = "../../config/docker"
    destination = "/var/app/lnd/docker"
  }

  # create lnd.conf
  provisioner "file" {
    source      = "../../config/lnd.conf"
    destination = "/var/app/lnd/docker/lnd/lnd.conf"
  }

  # ensure entrypoint script and systemd stuff are enabled
  provisioner "shell" {
    inline = [
      "chmod +x /var/app/lnd/docker/start-up.sh"
      "systemctl enable docker-compose-app"
    ]
  }

  post-processor "manifest" {}
}
