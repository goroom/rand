package main

import (
	"fmt"

	rd "github.com/goroom/rand"
)

func main() {
	rand := rd.GetRand()
	fmt.Println(rand.Int(), rand.Intn(10))
	fmt.Println(rand.Int31(), rand.Int31n(10))
	fmt.Println(rand.Int63(), rand.Int63n(10))
	fmt.Println(rand.Float32(), rand.Float64())
	fmt.Println(rand.Perm(10))
	fmt.Println(rand.RandRange(1, 3),
		rand.RandRangeInt32(1291212412, 2091212412),
		rand.RandRangeInt64(129121241212912124, 229121241212912124))
	fmt.Println(rand.Bool())

	fmt.Println(rand.String(32, rd.RST_LOWER))
	fmt.Println(rand.String(32, rd.RST_UPPER))
	fmt.Println(rand.String(32, rd.RST_NUMBER))
	fmt.Println(rand.String(32, rd.RST_SYMBOL))
	fmt.Println(rand.String(32, 0))

	fmt.Println(rand.StringArray(&[]string{"张三", "李四", "王五", "赵六"}))
	fmt.Println(rand.CarPlate())
	fmt.Println(rand.ChineseName())

	fmt.Println(rand.Phone())
}
