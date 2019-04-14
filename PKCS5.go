// PKCS 5
// Public Key Cryptography Standards #5
//
//     padded := PKCS5.Padding([]byte("hi"), 8)
//     original := PKCS5.RemovePadding(padded)
//     fmt.Println(padded)
//     fmt.Println(original)
//     // Output:
//     [104 105 6 6 6 6 6 6]
//     [104 105]
//
package PKCS5

import (
	"bytes"
)

// padding
// 填充
func Padding(original []byte, blockSize int) []byte {
	n := blockSize - len(original)%blockSize
	padded := bytes.Repeat([]byte{byte(n)}, n)
	return append(original, padded...)
}

// remove padding
// 拆除填充返回原
func RemovePadding(padded []byte) []byte {
	length := len(padded)
	// 通过最后一位找到填充了几，填充了几个
	padding := int(padded[length-1])
	return padded[:(length - padding)]
}
