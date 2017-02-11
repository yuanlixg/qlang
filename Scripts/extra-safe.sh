#!/bin/sh

tar -zcf extra-safe.tar.gz \
  `find qlang -type f -name extra.go` \
  `find qlang -type f -name safe.go`
