@echo off

SET CGO_ENABLED=0
SET GOARCH=amd64

SET GOOS=linux
go build -ldflags "-w -s" -o build\ClipVideoFrame-linux-x64

SET GOOS=windows
go build  -ldflags "-w -s" -o build\ClipVideoFrame.exe


SET GOOS=darwin
go build -ldflags "-w -s" -o build\ClipVideoFrame-macos-x64

SET GOARCH=386
SET GOOS=windows
go build  -ldflags "-w -s" -o build\ClipVideoFrame-x86.exe


SET GOOS=linux
go build -ldflags "-w -s" -o build\ClipVideoFrame-linux86

SET GOOS=linux
SET GOARCH=arm
go build -ldflags "-w -s" -o build\ClipVideoFrame-linux-arm


SET GOOS=darwin
go build -ldflags "-w -s" -o build\ClipVideoFrame-macos-arm