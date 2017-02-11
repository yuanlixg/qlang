#!/bin/sh

if [ -z "$1" ] ; then
  echo "Usage: $0 <package>"
  exit 1
fi

go list -json "$1" | grep -v internal | grep -v "vendor/" | \
  sed -n -e '/"Deps": \[/,/\],/s/",$/"/p'
