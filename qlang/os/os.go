package os

import (
	"os"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "os",

	"DevNull":           os.DevNull,
	"ModeAppend":        os.ModeAppend,
	"ModeCharDevice":    os.ModeCharDevice,
	"ModeDevice":        os.ModeDevice,
	"ModeDir":           uint64(os.ModeDir),
	"ModeExclusive":     os.ModeExclusive,
	"ModeNamedPipe":     os.ModeNamedPipe,
	"ModePerm":          os.ModePerm,
	"ModeSetgid":        os.ModeSetgid,
	"ModeSetuid":        os.ModeSetuid,
	"ModeSocket":        os.ModeSocket,
	"ModeSticky":        os.ModeSticky,
	"ModeSymlink":       os.ModeSymlink,
	"ModeTemporary":     os.ModeTemporary,
	"ModeType":          uint64(os.ModeType),
	"O_APPEND":          os.O_APPEND,
	"O_CREATE":          os.O_CREATE,
	"O_EXCL":            os.O_EXCL,
	"O_RDONLY":          os.O_RDONLY,
	"O_RDWR":            os.O_RDWR,
	"O_SYNC":            os.O_SYNC,
	"O_TRUNC":           os.O_TRUNC,
	"O_WRONLY":          os.O_WRONLY,
	"PathListSeparator": os.PathListSeparator,
	"PathSeparator":     os.PathSeparator,
	"SEEK_CUR":          os.SEEK_CUR,
	"SEEK_END":          os.SEEK_END,
	"SEEK_SET":          os.SEEK_SET,

	"Args":          os.Args,
	"ErrExist":      os.ErrExist,
	"ErrInvalid":    os.ErrInvalid,
	"ErrNotExist":   os.ErrNotExist,
	"ErrPermission": os.ErrPermission,
	"Interrupt":     os.Interrupt,
	"Kill":          os.Kill,
	"Stderr":        os.Stderr,
	"Stdin":         os.Stdin,
	"Stdout":        os.Stdout,

	"chdir":           os.Chdir,
	"chmod":           os.Chmod,
	"chown":           os.Chown,
	"chtimes":         os.Chtimes,
	"clearenv":        os.Clearenv,
	"environ":         os.Environ,
	"exit":            os.Exit,
	"expand":          os.Expand,
	"expandEnv":       os.ExpandEnv,
	"getegid":         os.Getegid,
	"getenv":          os.Getenv,
	"geteuid":         os.Geteuid,
	"getgid":          os.Getgid,
	"getgroups":       os.Getgroups,
	"getpagesize":     os.Getpagesize,
	"getpid":          os.Getpid,
	"getppid":         os.Getppid,
	"getuid":          os.Getuid,
	"getwd":           os.Getwd,
	"hostname":        os.Hostname,
	"isExist":         os.IsExist,
	"isNotExist":      os.IsNotExist,
	"isPathSeparator": os.IsPathSeparator,
	"isPermission":    os.IsPermission,
	"lchown":          os.Lchown,
	"link":            os.Link,
	"mkdir":           os.Mkdir,
	"mkdirAll":        os.MkdirAll,
	"newSyscallError": os.NewSyscallError,
	"readlink":        os.Readlink,
	"remove":          os.Remove,
	"removeAll":       os.RemoveAll,
	"rename":          os.Rename,
	"sameFile":        os.SameFile,
	"setenv":          os.Setenv,
	"symlink":         os.Symlink,
	"tempDir":         os.TempDir,
	"truncate":        os.Truncate,
	"unsetenv":        os.Unsetenv,

	"lstat": os.Lstat,
	"stat":  os.Stat,

	"File":         qlang.StructOf((*os.File)(nil)),
	"file":         os.NewFile,
	"create":       os.Create,
	"open":         os.Open,
	"openFile":     os.OpenFile,
	"pipe":         os.Pipe,
	"ProcAttr":     qlang.StructOf((*os.ProcAttr)(nil)),
	"Process":      qlang.StructOf((*os.Process)(nil)),
	"findProcess":  os.FindProcess,
	"startProcess": os.StartProcess,
	"ProcessState": qlang.StructOf((*os.ProcessState)(nil)),
}