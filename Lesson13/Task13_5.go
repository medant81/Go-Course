package main

import (
	"fmt"
	"github.com/medant81/myformat"
	myformatv2 "github.com/medant81/myformat/v2"
)

func main() {

	err := myformat.Do("file_in", "file_out")
	if err != nil {
		fmt.Println(err)
	}

	err = myformatv2.Do("file_in", "file_out_v2")
	if err != nil {
		fmt.Println(err)
	}

}
