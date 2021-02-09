#!/bin/bash
git clone git@github.com:bashsiz/nipo.git
cd nipo/nipo
go get gopkg.in/yaml.v2
go build

./nipo --help