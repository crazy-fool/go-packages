package main

import (
	"fmt"
	"go-packages/packages/encrypt"
)

func main() {
	str := "123456"
	encodeStr := encrypt.Base64Encode(str)
	fmt.Println(encodeStr) // MTIzNDU2
	decodeStr := encrypt.Base64Decode(encodeStr)
	fmt.Println(decodeStr)
	fmt.Println(decodeStr == str)

	key := "1234567890123456"
	iv := "1234567890123456"
	encryptedStr := encrypt.NewCbcPkcs7Aes(key, iv).AesEncrypt(str)
	fmt.Println(encrypt.Base64Encode(encryptedStr))
	fmt.Println(encrypt.NewCbcPkcs7Aes(key, iv).AesDecrypt(encryptedStr))
}
