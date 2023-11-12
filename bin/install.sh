#!/bin/bash

docker pull iimeta/iim-client:1.1.0

mkdir -p /data/iim-client/manifest/config
mkdir -p /data/iim-client/resource/public
mkdir -p /data/iim-client/resource/private

wget -P /data/iim-client/manifest/config https://github.com/iimeta/iim-client/raw/docker/manifest/config/config.yaml
wget https://github.com/iimeta/iim-client/raw/docker/bin/start.sh
