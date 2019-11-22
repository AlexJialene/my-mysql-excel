package main_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestMain(t *testing.M) {
	t.Run()
}

func TestIn(t *testing.T) {
}

func Test1(t *testing.T) {

}

func Test2(t *testing.T) {
	i := 2
	var str string
	for {
		if i <= 0 {
			break
		}
		i--
		i2 := i % 26
		str += string(i2 + 97)
		i = (i-i2 )/26

	}

	var s string
	for i:=len(str); i > 0 ; i--  {
		s += string(str[i-1])

	}
	fmt.Println(strings.ToUpper(s ))

}

func test(v interface{}) {
	of := reflect.TypeOf(v)
	fmt.Println(of.String())
}
