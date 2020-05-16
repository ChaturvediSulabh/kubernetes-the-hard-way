#!/bin/sh
# Configure nfs server and setup /shared dir
sudo apt-get update
sudo apt install nfs-kernel-server -y
sudo mkdir -p /shared
sudo useradd -r postgres
sudo chown nobody:nogroup /shared
sudo chmod 777 /shared
sudo mkdir -p /shared/mysql
sudo useradd -r mysql
sudo chown mysql:mysql -R /shared/mysql
echo '/shared 192.168.5.0/24(rw,sync,no_subtree_check)' > /etc/exports
sudo exportfs -a

# Restart nfs server
sudo systemctl restart nfs-kernel-server
sudo systemctl enable nfs-kernel-server

