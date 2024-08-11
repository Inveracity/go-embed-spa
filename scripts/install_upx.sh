#!/bin/env bash

set -eo pipefail

package=upx-4.2.4-amd64_linux
curl -sL https://github.com/upx/upx/releases/download/v4.2.4/${package}.tar.xz -o /tmp/$package.tar.xz
tar -xf /tmp/$package.tar.xz -C /tmp/
sudo install /tmp/$package/upx /usr/local/bin/upx
rm -rf /tmp/$package
rm -f /tmp/$package.tar.xz
