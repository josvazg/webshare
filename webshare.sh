#!/bin/sh
set -e

cd "$HOME"
mkdir -p shared
webshare -dir shared -port 8888 &
