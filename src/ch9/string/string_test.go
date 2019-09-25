package string_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s, len(s)) //初始化为默认零值

	//s[1]=3 // string 为不可变的byte slice
	s = "Hello"
	t.Log(s, len(s))

	s = "\xE4\xB8\xA5" //可以存储任意二进制数据
	t.Log(s, len(s))

	s = "中"
	t.Log(len(s)) //是byte数

	c := []rune(s)
	t.Log(len(c))
	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 utf8 %x", s)

	// Unicode ：code point
	// utf8
}

func TestStrings(t *testing.T) {
	s1 := "helloworld"
	s2 := strings.Replace(s1, "l", "@", -1)
	fmt.Println(s2)
}

func TestStrconv(t *testing.T) {
	s1 := "true"
	b1, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T, %t\n", b1, b1)

	ss1 := strconv.FormatBool(b1)
	fmt.Printf("%T, %s\n", ss1, ss1)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
	}
}
