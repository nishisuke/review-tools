package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	bar string

}
func main() {
	dat := Foo{"ssssss"}
	byt,_ := json.Marshal(dat)
	fmt.Println(byt)
}
