package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"qlang.io/cl/qlang"

	qipt "qlang.io/cl/interpreter"
)

const (
	scriptCode = `x = 1 + 2`
)

var (
	osArg0 string
	absDir string

	stderr io.Writer
)

func main() {
	os.Exit(Qmain(os.Args))
}

func mainInit(args []string) int {
	if len(args) < 1 {
		return 1
	}
	osArg0 = args[0]
	p, _ := exec.LookPath(osArg0)
	p, _ = filepath.Abs(p)
	absDir, _ = filepath.Split(p)

	stderr = os.Stderr
	return 0
}

func Qmain(args []string) int {
	if ret := mainInit(args); ret != 0 {
		return ret
	}

	InitSafe(false)
	qlang.Import("", qipt.Exports)
	qlang.Import("qlang", qlang.Exports)
	qlang.SetDumpCode(os.Getenv("QLANG_DUMPCODE"))

	libs := os.Getenv("QLANG_PATH")
	if libs == "" {
		libs = os.Getenv("HOME")
		if libs == "" {
			libs = os.Getenv("HOMEPATH")
		}
		libs = absDir[len(filepath.VolumeName(absDir)):] + "qlang:" + libs + "/qlang"
	}

	lang := qlang.New()
	lang.SetLibs(libs)

	// exec source
	if len(args) > 1 {
		fname := os.Args[1]
		b, err := ioutil.ReadFile(fname)
		if err != nil {
			fmt.Fprintln(stderr, err)
			return 2
		}
		err = lang.SafeExec(b, fname)
		if err != nil {
			fmt.Fprintln(stderr, err)
			return 3
		}
		return 0
	}

	fmt.Fprintf(stderr, "Usage: %v <script.ql> [arg ...]\n"+
		"\tnow test engine: `%v`\n", osArg0, scriptCode)

	err := lang.SafeExec([]byte(scriptCode), "")
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 3
	}

	v, _ := lang.GetVar("x")
	fmt.Fprintln(stderr, "x:", v)
	return 0
}
