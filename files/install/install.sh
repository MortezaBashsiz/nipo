#!/bin/bash

apt install golang
cd /opt
git clone git@github.com:bashsiz/nipo.git
cp -r /opt/nipo/nipolib/go/nipo ~/go/src/
go get gopkg.in/yaml.v2
cd /opt/nipo/nipo
go build
cp nipo /usr/local/bin/
mkdir /etc/nipo
cp /opt/nipo/files/config/nipo-cfg.yaml /etc/nipo/
cp /opt/nipo/files/config/nipo.service /lib/systemd/system/
mkdir /var/log/nipo