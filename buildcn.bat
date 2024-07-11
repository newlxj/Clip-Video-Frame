@echo off

SET CGO_ENABLED=0
SET GOARCH=amd64

SET GOOS=linux
go build -ldflags "-w -s" -o build\ClipVideo-linux-x64

SET GOOS=windows
go build  -ldflags "-w -s" -o build\ClipVideo.exe


SET GOOS=darwin
go build -ldflags "-w -s" -o build\ClipVideo-macos-x64

SET GOARCH=386
SET GOOS=windows
go build  -ldflags "-w -s" -o build\ClipVideo-x86.exe


SET GOOS=linux
go build -ldflags "-w -s" -o build\ClipVideo-linux86

SET GOOS=linux
SET GOARCH=arm
go build -ldflags "-w -s" -o build\ClipVideo-linux-arm


SET GOOS=darwin
go build -ldflags "-w -s" -o build\ClipVideo-macos-arm