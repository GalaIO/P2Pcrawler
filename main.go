package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"reflect"
	"time"
)

func panicTest() (ret string) {

	defer func() {
		if err := recover(); err != nil {
			ret = err.(string)
		}
	}()

	panic("123")
	return "456"
}

func main() {
	fmt.Println("hello world")

	var l []interface{}
	l = nil

	for e := range l {
		fmt.Println(e)
	}

	var m map[string]interface{}

	m = nil
	for k, v := range m {
		fmt.Println(k, v)
	}

	s1 := sha1.New()
	s1.Write([]byte("test"))
	sha1sum := s1.Sum(nil)
	fmt.Println(hex.EncodeToString(sha1sum[:]))

	type Dict map[string]interface{}

	d := Dict{"name": "xiaoguo"}
	m1 := map[string]interface{}{"name": "xiaoguo"}

	t := reflect.TypeOf(d)
	mt := reflect.TypeOf(m1)
	fmt.Printf("%s\n", t.Name())
	fmt.Printf("%s\n", t.Kind())
	fmt.Printf("%s\n", mt.Name())
	fmt.Printf("%s\n", mt.Kind())

	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))

	fmt.Println(panicTest())
}
