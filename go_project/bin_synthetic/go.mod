module github.com/LostBitset/quiver_se/bin_synthetic

go 1.20

replace github.com/LostBitset/quiver_se/lib => ../lib

replace github.com/LostBitset/quiver_se/lib_synthetic => ../lib_synthetic

require github.com/LostBitset/quiver_se/lib_synthetic v0.0.0-20230423010004-bcfc48842235

require (
	github.com/LostBitset/quiver_se/lib v0.0.0-20230421205926-194398f471d9 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	src.elv.sh v0.18.0 // indirect
)
