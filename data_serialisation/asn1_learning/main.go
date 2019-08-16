package main

import (
	"encoding/asn1"
	"fmt"
	"os"
)

type SerialClass struct {
	TestInt    int
	TestString string
}

func main() {
	testInt(13)
	testString("hello")
	testObj(SerialClass{TestInt: 13, TestString: "hello"})
}

func testInt(data int) {
	fmt.Println("=========Star int test============")
	mdata, err := asn1.Marshal(data)
	checkError(err)

	fmt.Printf("Data after marshal: %v\n", mdata)
	var n int
	_, err1 := asn1.Unmarshal(mdata, &n)
	checkError(err1)
	fmt.Println("Data after unmarshal: ", n)
}

func testString(data string) {
	fmt.Println("=========Star string test============")
	mdata, err := asn1.Marshal(data)
	checkError(err)

	fmt.Printf("Data after marshal: %v\n", mdata)
	var n string
	_, err1 := asn1.Unmarshal(mdata, &n)
	checkError(err1)
	fmt.Println("Data after unmarshal: ", n)
}

func testObj(data SerialClass) {
	fmt.Println("=========Star object test============")
	mdata, err := asn1.Marshal(data)
	checkError(err)

	fmt.Printf("Data after marshal: %v\n", mdata)
	n := SerialClass{}
	_, err1 := asn1.Unmarshal(mdata, &n)
	checkError(err1)
	fmt.Println("Data after unmarshal: ", n)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
