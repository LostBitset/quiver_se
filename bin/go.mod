module LostBitset/quiver_se/bin

go 1.19

replace LostBitset/quiver_se/lib => ../lib

require (
	LostBitset/quiver_se/lib v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.0
	github.com/stretchr/testify v1.8.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	src.elv.sh v0.18.0 // indirect
)
