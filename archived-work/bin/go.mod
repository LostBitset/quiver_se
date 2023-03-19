module LostBitset/quiver_se/bin

go 1.19

replace LostBitset/quiver_se/lib => ../lib

require (
	LostBitset/quiver_se/EIDIN/proto_lib v0.0.0-00010101000000-000000000000 // indirect
	LostBitset/quiver_se/lib v0.0.0-00010101000000-000000000000 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	src.elv.sh v0.18.0 // indirect
)

replace LostBitset/quiver_se/EIDIN/proto_lib => ../js_concolic/EIDIN/proto_lib