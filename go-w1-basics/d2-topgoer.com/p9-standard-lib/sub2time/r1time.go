package main

import (
	"fmt"
	"time"
)

func main() {
	//part-1, 初识
	now := time.Now()
	fmt.Printf("current time: %v \n", now)
	fmt.Println("current time: : ", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	//%02d 表示留两位，不足补零（往前补）
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	//part-2， 获取时间戳
	now2 := time.Now()
	//时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数。它也被称为Unix时间戳（UnixTimestamp）
	//这个不对吧？？
	//now2.Unix()得到的是秒，UnixMilli得到的才是毫秒
	fmt.Printf("current timestamp, default timestamp[UnixTimeStamp]: %v \n", now2.Unix())
	fmt.Printf("UnixMilli: %v, UnixMicro: %v, UnixNano: %v \n", now2.UnixMilli(), now2.UnixMicro(), now2.UnixNano())
	tmp1 := time.Now().UnixMilli()
	time.Sleep(time.Millisecond * 233)
	tmp2 := time.Now().UnixMilli()
	fmt.Println(tmp2 - tmp1)

	//part-3，时间戳转时间
	now3 := time.Unix(now2.Unix(), 0)
	fmt.Println(now3)

	//其他time.Duration 时间操作（Add、Sub、Equal、Before、After...） 定时器、时间格式化，先不看了...
	//我要去复习Java的LocalDateTime了...
}
