package main

import (
	"fmt"
	"w7_widgets/utils"

	"github.com/dgraph-io/ristretto"
)

func main() {

	fmt.Println("kkkddddd")

	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}

	// set a value with a cost of 1
	cache.Set("key", "value", 1)
	// wait for value to pass through buffers
	cache.Wait()
	// get value from cache
	value, found := cache.Get("key")
	if !found {
		panic("missing value")
	}
	fmt.Println(value)

	// del value from cache
	cache.Del("key")

	var text []string
	filter := utils.NewNodeFilter([]string{"希特勒", "法西斯"})
	result, err := filter.FilterResult("希特勒")
	if nil != err {
		return
	}
	for k := range result {
		text = append(text, k)
	}
	fmt.Println(text)

	cache.Set("38", filter, 1)
	cache.Wait()

	var v2 interface{}
	v2, f2 := cache.Get("38")
	if !f2 {
		fmt.Println("get dirty filter from cache success!")
	}
	var v3 utils.DirtyFilter
	v3 = v2.(utils.DirtyFilter)
	result2, _ := v3.FilterResult("希特勒")
	for k := range result2 {
		text = append(text, k)
	}
	fmt.Println(text)

	//var v3 utils.DirtyFilter
	//v3 = 123
	//fmt.Println(v3)
}
