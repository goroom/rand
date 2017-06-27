
````go
package main

import (
	"fmt"

	"github.com/goroom/rand"
)

func main() {
	fmt.Println(rand.Int(), rand.Intn(10))
	fmt.Println(rand.Int32(), rand.Int32n(10))
	fmt.Println(rand.Int64(), rand.Int64n(10))
	fmt.Println(rand.Float32(), rand.Float64())
	fmt.Println(rand.Perm(10))
	fmt.Println(rand.RandRange(1, 3),
		rand.RandRangeInt32(1291212412, 2091212412),
		rand.RandRangeInt64(129121241212912124, 229121241212912124))
	fmt.Println(rand.Bool())

	fmt.Println(rand.String(32, rand.RST_LOWER))
	fmt.Println(rand.String(32, rand.RST_UPPER))
	fmt.Println(rand.String(32, rand.RST_NUMBER))
	fmt.Println(rand.String(32, rand.RST_SYMBOL))
	fmt.Println(rand.String(32, 0))

	fmt.Println(rand.StringArray(&[]string{"张三", "李四", "王五", "赵六"}))
	fmt.Println(rand.CarPlate())
	fmt.Println(rand.ChineseName())

	fmt.Println(rand.RandPhone())
}
````

结果
````go
4156960477472916945 8
1591767834 8
1849074015167768215 6
0.3724794 0.24850495941917342
[9 3 1 0 5 4 6 7 8 2]
3 2058348008 173108310343143870
false
embvdawhehdptibctdarfwfrolunwecf
NRSLUFQAEHYUOZMYKHNSEGEWSRIEYMOZ
09052341237931033459193029727749
\+&/(@`~!,+],>(+!<,#+$_==}^%'+>\
r3GFZ@0FxTsN9#Il]Q0z6mJj_,i=dHVZ
李四
浙H3CVXL
花力启
15303145889
````
