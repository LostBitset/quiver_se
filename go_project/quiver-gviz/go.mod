module github.com/LostBitset/quiver_se/quiver_gviz

go 1.20

replace github.com/LostBitset/quiver_se/lib => ../lib

replace github.com/LostBitset/quiver_se/lib_synthetic => ../lib_synthetic

require (
	github.com/LostBitset/quiver_se/lib v0.0.0-20230421205926-194398f471d9
	github.com/LostBitset/quiver_se/lib_synthetic v0.0.0-00010101000000-000000000000
)

require (
	github.com/emicklei/dot v1.4.2
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	src.elv.sh v0.18.0 // indirect
)
