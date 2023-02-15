#!/bin/bash

echo "[compilation] [.../EIDIN/compile_proto.sh] Compiling (Golang) Protobuf definition of EIDIN..."

echo "[compilation] [.../EIDIN/compile_proto.sh] Setting up module directory..."

rm -r proto_lib
mkdir proto_lib

echo "[compilation] [.../EIDIN/compile_proto.sh] Invoking protobuf compiler (with \"--go_out\")..."

protoc -I=proto --go_out=proto_lib proto/eidin.proto

echo "[compilation] [.../EIDIN/compile_proto.sh] Protobuf compiler has finished, cleaning up..."

cp -r proto_lib/LostBitset/quiver_se/EIDIN/proto_lib .
rm -r proto_lib/LostBitset

echo "[compilation] [.../EIDIN/compile_proto.sh] Creating go module..."

cd proto_lib
go mod init LostBitset/quiver_se/EIDIN/proto_lib 2>/dev/null
go mod tidy
cd ..

echo "[compilation] [.../EIDIN/compile_proto.sh] Module generation finished. All Done."
