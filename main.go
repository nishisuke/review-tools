package main

import (
	"encoding/json"
	"fmt"

	"github.com/nishisuke/try-review-tools/internal/foo"
)

type Foo struct {
	bar string

}
func main() {
	dat := Foo{"ssssss"}
	byt,_ := json.Marshal(dat)
	fmt.Println(byt)
	foo.Bar()
}
