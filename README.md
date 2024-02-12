# GRPC Demo Project

A sample project to demonstrate working of gRPC using GO Lang.

## Pre-requisite

 - [Protocol Buffer Compiler](https://grpc.io/docs/protoc-installation/)

## Installation

Install protocol buffer

CMD:

```bash
set PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO %PB_REL%/download/v3.15.8/protoc-3.15.8-win64.zip
```
PowerShell:

```bash
$PB_REL = "https://github.com/protocolbuffers/protobuf/releases"
Invoke-WebRequest -Uri "$PB_REL/download/v3.15.8/protoc-3.15.8-win64.zip" -OutFile protoc-3.15.8-win64.zip
```

Extract the downloaded zip file

CMD:

```bash
mkdir %USERPROFILE%\AppData\Local\protoc
tar -xf protoc-3.15.8-win64.zip -C %USERPROFILE%\AppData\Local\protoc
```
PowerShell:

```bash
Expand-Archive -Path protoc-3.15.8-win64.zip -DestinationPath ~\AppData\Local
```

Add ..\protoc\bin folder to PATH Environment Variable, run
```bash
protoc --version
```
to verify Installation.

#### Install GO plugins for protocol buffer compiler
```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
