package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type SerialClass struct {
	TestInt    int    `json:test_int`
	TestString string `json:test_string`
}

func main() {
	testObj(SerialClass{TestInt: 13, TestString: "hello"})
}

func testObj(data SerialClass) {
	fmt.Println("=========Star object test============")
	mdata, err := json.Marshal(data)
	checkError(err)

	fmt.Printf("Data after marshal: %v\n", mdata)
	n := SerialClass{}
	err1 := json.Unmarshal(mdata, &n)
	checkError(err1)
	fmt.Println("Data after unmarshal: ", n)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
