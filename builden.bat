@echo off
rename main.go main.go.cn
rename main.go.en main.go
SET CGO_ENABLED=0
SET GOARCH=amd64

SET GOOS=linux
go build -ldflags "-w -s" -o build-en\ClipVideo-linux-x64

SET GOOS=windows
go build  -ldflags "-w -s" -o build-en\ClipVideo.exe

SET GOOS=darwin
go build -ldflags "-w -s" -o build-en\ClipVideo-macos-x64

SET GOARCH=386
SET GOOS=windows
go build  -ldflags "-w -s" -o build-en\ClipVideo-x86.exe


SET GOOS=linux
go build -ldflags "-w -s" -o build-en\ClipVideo-linux86

SET GOOS=linux
SET GOARCH=arm
go build -ldflags "-w -s" -o build-en\ClipVideo-linux-arm


SET GOOS=darwin
go build -ldflags "-w -s" -o build-en\ClipVideo-macos-arm


rename main.go main.go.en
rename main.go.cn main.go