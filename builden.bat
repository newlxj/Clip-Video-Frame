@echo off
rename main.go main.go.cn
rename main.go.en main.go
SET CGO_ENABLED=0
SET GOARCH=amd64

SET GOOS=linux
go build -ldflags "-w -s" -o build-en\ClipVideoFrame-linux-x64

SET GOOS=windows
go build  -ldflags "-w -s" -o build-en\ClipVideoFrame.exe

SET GOOS=darwin
go build -ldflags "-w -s" -o build-en\ClipVideoFrame-macos-x64

SET GOARCH=386
SET GOOS=windows
go build  -ldflags "-w -s" -o build-en\ClipVideoFrame-x86.exe


SET GOOS=linux
go build -ldflags "-w -s" -o build-en\ClipVideoFrame-linux86

SET GOOS=linux
SET GOARCH=arm
go build -ldflags "-w -s" -o build-en\ClipVideoFrame-linux-arm


SET GOOS=darwin
go build -ldflags "-w -s" -o build-en\ClipVideoFrame-macos-arm


rename main.go main.go.en
rename main.go.cn main.go