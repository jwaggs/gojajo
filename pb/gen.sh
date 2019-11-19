#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/def" >/dev/null && pwd )"
PIGGY_PROTO_PKG="/Users/jon/Documents/code/piggy/pb"
PIGGY_SWAGGER_PKG="/Users/jon/Documents/code/piggy/pb"

ls $DIR/*.proto | while read -r line ; do
    protoc -I$GOPATH/src \
    --go_out=plugins=grpc:$PIGGY_PROTO_PKG \
    --grpc-gateway_out=logtostderr=true:$PIGGY_PROTO_PKG \
    --swagger_out=logtostderr=true:$PIGGY_SWAGGER_PKG \
    --proto_path=$DIR $line
done
