package PKCS5_test

import (
	"bytes"
	"fmt"
	"github.com/trylife/go-PKCS5"
	"testing"
)

type PaddingCase struct {
	Original     []byte
	ResultPadded []byte
	BlockSize    int
}

var cases = []PaddingCase{
	0: {
		Original:     []byte("trylife"),
		ResultPadded: []byte{116, 114, 121, 108, 105, 102, 101, 1},
		BlockSize:    8,
	},
	1: {
		Original:     []byte("DDDDDD"),
		ResultPadded: []byte{68, 68, 68, 68, 68, 68, 2, 2},
		BlockSize:    8,
	},
	2: {
		Original:     []byte("DDDDDDDD"),
		ResultPadded: []byte{68, 68, 68, 68, 68, 68, 68, 68, 8, 8, 8, 8, 8, 8, 8, 8},
		BlockSize:    8,
	},
	3: {
		Original:     []byte("DDDDDDDDDD"),
		ResultPadded: []byte{68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 6, 6, 6, 6, 6, 6},
		BlockSize:    8,
	},
	4: {
		Original:     []byte("DDDDDD"),
		ResultPadded: []byte{68, 68, 68, 68, 68, 68, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26,},
		BlockSize:    32,
	},
}

func TestPaddingAndRemove(t *testing.T) {
	for index, c := range cases {
		res := PKCS5.RemovePadding(PKCS5.Padding(c.Original, c.BlockSize))
		if !bytes.Equal(res, c.Original) {
			t.Log("case: ", index)
			t.Log("case.Original: ", c.Original)
			t.Log("case.Original->string: ", string(c.Original))
			t.Log("case.PaddingResult: ", c.ResultPadded)
			t.Log("case.PaddingResult->string: ", string(c.ResultPadded))
			t.Log("case.BlockSize: ", c.BlockSize)
			t.Log("Result: ", res)
			t.Log("Result->string: ", string(res))
			t.Error("origin != RemovePadding(padding(origin))")
		} else {

		}
	}
}

func TestPadding(t *testing.T) {
	for index, c := range cases {
		result := PKCS5.Padding(c.Original, c.BlockSize)
		t.Log("case: ", index)
		t.Log("case.Original: ", c.Original)
		t.Log("case.Original->string: ", string(c.Original))
		t.Log("case.PaddingResult: ", c.ResultPadded)
		t.Log("case.PaddingResult->string: ", string(c.ResultPadded))
		t.Log("case.BlockSize: ", c.BlockSize)
		t.Log("Result: ", result)
		t.Log("Result->string: ", string(result))
		if !bytes.Equal(result, c.ResultPadded) {
			t.Error("result != c.ResultPadding")
		} else {

		}
	}
}

func TestRemovePadding(t *testing.T) {
	for index, c := range cases {
		result := PKCS5.RemovePadding(c.ResultPadded)
		t.Log("case: ", index)
		t.Log("case.Original: ", c.Original)
		t.Log("case.Original->string: ", string(c.Original))
		t.Log("case.PaddingResult: ", c.ResultPadded)
		t.Log("case.PaddingResult->string: ", string(c.ResultPadded))
		t.Log("case.BlockSize: ", c.BlockSize)
		t.Log("Result: ", result)
		t.Log("Result->string: ", string(result))
		if !bytes.Equal(result, c.Original) {
			t.Error("result != c.Original")
		} else {

		}
	}
}


func Example_() {
	padded := PKCS5.Padding([]byte("hi"), 8)
	original := PKCS5.RemovePadding(padded)
	fmt.Println(padded)
	fmt.Println(original)
	// Output:
	// [104 105 6 6 6 6 6 6]
	// [104 105]
}