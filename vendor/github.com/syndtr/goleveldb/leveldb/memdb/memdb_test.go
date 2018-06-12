package memdb

import (
	"bytes"
	"fmt"
	"testing"
)

type cmp struct{}

var c cmp

func (cmp) Compare(a, b []byte) int {
	return bytes.Compare(a, b)
}
func TestMemdb(t *testing.T) {
	db := New(c, 2)
	db.Put([]byte("name"), []byte("lzb"))
	db.Put([]byte{1, 2, 3}, []byte{4, 5, 6})
	v, err := db.Get([]byte("name"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(v)
}
