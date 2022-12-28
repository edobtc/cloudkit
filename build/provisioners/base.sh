#!/usr/bin/env bash

# base setup
# await unattended upgrades on boot before proceeding to run update and upgrade
while sudo lsof /var/lib/dpkg/lock-frontend ; do sleep 10; done;

# proceed
export DEBIAN_FRONTEND=noninteractive
sudo apt-get update
sudo apt-get upgrade -y
sudo apt install -y apt-transport-https ca-certificates curl software-properties-common gnupg lsb-release

# install docker
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
sudo echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update -y
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin

# pull images for faster setup once running
docker pull lightninglabs/lnd:v0.15.5-beta
docker network create bslightning

# set up copy directory
sudo mkdir -p /var/app/lnd
