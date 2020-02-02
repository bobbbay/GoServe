package main

import (
	"fmt"
	"os"
)

func main(){
	os.Remove("goserve.conf")
	os.Create("goserve.conf")
	fmt.Println("GoServe cleaned, run make to start again")
}