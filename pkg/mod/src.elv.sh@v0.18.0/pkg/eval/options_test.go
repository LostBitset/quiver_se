package eval

import (
	"reflect"
	"testing"

	. "src.elv.sh/pkg/tt"
)

type opts struct {
	Foo string
	bar int
}

// Equal is required by cmp.Diff, since opts contains unexported fields.
func (o opts) Equal(p opts) bool { return o == p }

func TestScanOptions(t *testing.T) {
	// A wrapper of ScanOptions, to make it easier to test
	wrapper := func(src RawOptions, dstInit interface{}) (interface{}, error) {
		ptr := reflect.New(reflect.TypeOf(dstInit))
		ptr.Elem().Set(reflect.ValueOf(dstInit))
		err := scanOptions(src, ptr.Interface())
		return ptr.Elem().Interface(), err
	}

	Test(t, Fn("scanOptions", wrapper), Table{
		Args(RawOptions{"foo": "lorem ipsum"}, opts{}).
			Rets(opts{Foo: "lorem ipsum"}, nil),
		Args(RawOptions{"bar": 20}, opts{bar: 10}).
			Rets(opts{bar: 10}, UnknownOption{"bar"}),
	})
}
