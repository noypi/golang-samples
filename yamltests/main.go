package main

import (
	"fmt"
	"gopkg.in/yaml.v1"
	"io/ioutil"
)

type T struct {
	C int
	D []string ",flow"
}

func main() {

	mm := map[string]*T{}

	bb, err := ioutil.ReadFile("./t.yaml")
	if nil != err {
		fmt.Println("err:", err)
		return
	}
	if err = yaml.Unmarshal(bb, mm); nil != err {
		fmt.Println("err:", err)
		return
	}
	for k, v := range mm {
		fmt.Printf("%v: %v\n", k, v)
	}

}
