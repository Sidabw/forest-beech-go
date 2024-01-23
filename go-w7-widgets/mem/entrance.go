package mem

import (
	"fmt"
	"github.com/dgraph-io/ristretto"
)

func DoTest() {
	a1 := 1e7
	a2 := 1 << 30
	fmt.Println(a1)
	fmt.Println(a2 / 1024 / 1024)

	a3 := 1 << 29
	fmt.Println(a3 / 1024 / 1024)
	serTest()

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
	//time.Sleep(time.Second * 5)
	//cache.Wait()
	// get value from cache
	value, _ := cache.Get("key")
	fmt.Println(value)

	// del value from cache
	cache.Del("key")

	var text []string
	filter := NewNodeFilter([]string{"希特勒", "法西斯"})
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
		fmt.Println("get dirty filter from cache fail!")
	}
	var v3 DirtyFilter
	v3 = v2.(DirtyFilter)
	result2, _ := v3.FilterResult("希特勒")
	for k := range result2 {
		text = append(text, k)
	}
	fmt.Println(text)

	//var v3 mem.DirtyFilter
	//v3 = 123
	//fmt.Println(v3)

}

func serTest() {
	filter := NewNodeFilter([]string{"希特勒", "法西斯"})
	res, err := filter.ToJson()
	if err != nil {
		return
	}
	fmt.Println(res)
}
