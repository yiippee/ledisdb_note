package main

import (
	"fmt"
	lediscfg "github.com/siddontang/ledisdb/config"
	"github.com/siddontang/ledisdb/ledis"
	"time"
)

// Use Ledis's default config
func main() {

	//# Use Ledis's default config
	cfg := lediscfg.NewConfigDefault()
	cfg.DataDir = "C:\\Go\\gopath\\src\\github.com\\siddontang\\ledisdb\\cmd\\packageExample\\var"
	l, _ := ledis.Open(cfg)
	db, _ := l.Select(0)

	db.Set([]byte{1, 2, 3}, []byte{4, 5, 6})
	t1 := time.Now()
	for i := 0; i < 1000000; i ++ {
		db.Get([]byte{1, 2, 3})
	}
	t2 := time.Now()
	d := t2.Sub(t1)

	fmt.Printf("%s", d.String(),)
	//fmt.Println(v)
}
