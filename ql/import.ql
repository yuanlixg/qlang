#!/usr/bin/env qlang

importBase = "github.com/yuanlixg/qlang/qlang"

// =====================================================================
print_head = fn() {
	print(`package main

// Auto generation, edit and go fmt it

import (
`)
}

// =====================================================================
print_middle = fn() {
	print(`
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

`)
}

// =====================================================================
print_tail = fn() {
	print(`
}

// ---------------------------------------------------------------------------
`)
}

// =====================================================================
mModule = make(map[string]bool)

walkFunc = fn(path, info, err) {
	if err != nil {
		return filepath.SkipDir
	}

	if !info.IsDir() && filepath.ext(path) == ".go" {
		dir, _ = filepath.split(path)
		dir = dir[len("qlang/"):len(dir)-1]
		mModule[dir] = false
	}

	return nil
}

// =====================================================================
main {
	filepath.walk("qlang", walkFunc)
	sModule = make([]string, 0, len(mModule))
	for k, _ = range mModule {
		sModule = append(sModule, k)
	}
	sort.strings(sModule)

	print_head()

	for i = 0; i < len(sModule); i++ {
		_, name = filepath.split(sModule[i])
		switch sModule[i] {
		case "crypto/rand":
			name = "crand"
			print(" ", name)
		case "regexp/syntax":
			name = "rsyntax"
			print(" ", name)
		case "runtime/pprof":
			name = "rpprof"
			print(" ", name)
		}
		printf(` "%s/%s"`, importBase, sModule[i])
		println()
		sModule[i] = name
	}

	print_middle()

	for i = 0; i < len(sModule); i++ {
		printf(` qlang.Import("%s", %s.Exports)`, sModule[i], sModule[i])
		println()
	}

	print_tail()
}
