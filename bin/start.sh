#!/bin/bash

docker run -d \
  --network host \
  --restart=always \
  -p 8000:8000 \
  -v /etc/localtime:/etc/localtime:ro \
  -v /data/iim-client/manifest/config/config.yaml:/app/manifest/config/config.yaml \
  -v /data/iim-client/resource/public:/app/resource/public \
  -v /data/iim-client/resource/private:/app/resource/private \
  --name iim-client \
  iimeta/iim-client:1.1.0
