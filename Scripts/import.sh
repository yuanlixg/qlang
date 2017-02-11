#!/bin/sh

cat <<EOF
package main

// Auto generation, edit and go fmt it

import (
EOF

# Import base path: github.com/yuanlixg/qlang
#
# import "qlang/crypto/rand" as "crand" if match
# import "qlang/regexp/syntax" as "rsyntax" if match
# import "qlang/runtime/pprof" as "rpprof" if match
find qlang -type f -name '*.go' -exec dirname {} \; | sort -u | \
  sed -e 's@^qlang/@ "github.com/yuanlixg/qlang/qlang/@' -e 's/$/"/' \
      -e '/\/crypto\/rand"$/s/^/ crand/' \
      -e '/\/regexp\/syntax"$/s/^/ rsyntax/' \
      -e '/\/runtime\/pprof"$/s/^/ rpprof/'

cat <<EOF

	"qlang.io/qlang/meta"
	qlang "qlang.io/spec"

	// qlang builtin modules
	_ "qlang.io/qlang/builtin"
	_ "qlang.io/qlang/chan"
)

// InitSafe inits qlang and imports modules.
//
func InitSafe(safeMode bool) {

	qlang.SafeMode = safeMode

	qlang.Import("meta", meta.Exports) // import meta package

EOF

# replace to "crand" if match "/crypto/rand"
# replace to "rsyntax" if match "/regexp/syntax"
# replace to "rpprof" if match "/runtime/pprof"
find qlang -type f -name '*.go' -exec dirname {} \; | sort -u | \
  sed -e 's@/crypto/rand$@/crypto/crand@' \
      -e 's@/regexp/syntax$@/regexp/rsyntax@' \
      -e 's@/runtime/pprof$@/regexp/rpprof@' \
      -e 's@^.*/@@' \
      -e 's/^.*$/ qlang.Import("&", &.Exports)/'

cat <<EOF

}

// ---------------------------------------------------------------------------
EOF
