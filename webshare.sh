#!/bin/sh
set -e

cd "$SNAP_DATA"
mkdir -p shared
webshare -dir shared -port 8888 &
