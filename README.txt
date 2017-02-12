
modules custom tools for Q Language ( https://github.com/qiniu/qlang )

Install:
  go get -d -v -x github.com/yuanlixg/qlang

Custom:
  cd $GOPATH/src/github.com/yuanlixg/qlang

add your module export folder to "qlang/" folder.
or remove modules export under "qlang/" folder.

Build:
  sh Scripts/import.sh > import.go
  go build -v github.com/yuanlixg/qlang

Test:
  ./qlang ql/test.ql


Windows only (self serve), first:
  cd %GOPATH%\src\github.com\yuanlixg\qlang
  go build -v github.com/yuanlixg/qlang

after custom your modules:
  qlang.exe ql\import.ql > import.go

  cd %GOBIN%
  go build -v github.com/yuanlixg/qlang

