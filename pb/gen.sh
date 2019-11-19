#!/bin/bash

PIGGY_PROTO_DEF="$( cd "$( dirname "${BASH_SOURCE[0]}" )/def" >/dev/null && pwd )"
PIGGY_GO_OUT="$( dirname "${BASH_SOURCE[0]}" )"
PIGGY_SWAGGER_OUT="$( dirname "${BASH_SOURCE[0]}" )/swagger"

ls $PIGGY_PROTO_DEF/*.proto | while read -r line ; do
    protoc -I$GOPATH/src \
    --go_out=plugins=grpc:$PIGGY_GO_OUT \
    --grpc-gateway_out=logtostderr=true:$PIGGY_GO_OUT \
    --swagger_out=logtostderr=true:$PIGGY_SWAGGER_OUT \
    --proto_path=$PIGGY_PROTO_DEF $line
done
