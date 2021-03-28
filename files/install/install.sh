#!/bin/bash

declare -A osInfo;
osInfo[/etc/debian_version]="apt-get install -y"
osInfo[/etc/alpine-release]="apk --update add"
osInfo[/etc/centos-release]="yum install -y"
osInfo[/etc/fedora-release]="dnf install -y"

for f in ${!osInfo[@]}
do
    if [[ -f $f ]];then
        package_manager=${osInfo[$f]}
    fi
done

package="golang"

${package_manager} ${package}


cd /opt
git clone https://github.com/bashsiz/nipo.git
cp -r /opt/nipo/nipolib/go/nipo ~/go/src/
go get gopkg.in/yaml.v2
cd /opt/nipo/nipo
go build
cp nipo /usr/local/bin/
mkdir /etc/nipo
cp /opt/nipo/files/config/nipo-cfg.yaml /etc/nipo/
cp /opt/nipo/files/config/nipo.service /lib/systemd/system/
mkdir /var/log/nipo
rm -fr /opt/nipo