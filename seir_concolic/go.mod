module github.com/LostBitset/quiver_se/seir_concolic

go 1.20

replace github.com/LostBitset/quiver_se/lib => ../go_project/lib

require github.com/sirupsen/logrus v1.9.0

require (
	github.com/LostBitset/quiver_se/lib v0.0.0-20230421205926-194398f471d9
	github.com/google/go-cmp v0.5.9 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	src.elv.sh v0.18.0 // indirect
)
