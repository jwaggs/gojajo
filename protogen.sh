#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
SWAGGER="${DIR}/docs/swagger"
rm *.pb.go

ls *.proto | while read -r line ; do

  echo "generating go code for $line"

  protoc -I$GOPATH/src \
  --go_out=plugins=grpc:$DIR \
  --grpc-gateway_out=logtostderr=true:$DIR \
  --swagger_out=logtostderr=true:$SWAGGER \
  --proto_path=$DIR $line

done