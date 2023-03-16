package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//
	type demo struct {
		Id  string
		Pwd string
	}
	s := "{\"Id\":\"123\",\"Pwd\":\"123456\"}"
	var d demo
	err := json.Unmarshal([]byte(s), &d)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("id = %v, pwd = %v", d.Id, d.Pwd)

}
