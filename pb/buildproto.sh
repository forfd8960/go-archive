#!/usr/bin/env bash

echo $GOPATH
protoc  \
	--proto_path=$GOPATH/src  \
	--proto_path=.  \
	--go_out=plugins=grpc:$GOPATH/src/github.com/forfd8960/go-archive/pb \
	archivepb.proto