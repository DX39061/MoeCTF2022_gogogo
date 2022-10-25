package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

func check(flag string) bool {
	encFlag := "200c2c3ef00f31999df93d6919aa33e42dde307be02017ebf47067099ed0bddc525d5dba0f83c122159b89ae715907cc"
	key := []byte("---moeCTF2022---")
	iv := []byte("---moeCTF2022---")
	encrypt, err := AesEncrypt([]byte(flag), key, iv)
	if err != nil {
		return false
	}
	if hex.EncodeToString(encrypt) == encFlag {
		return true
	}
	return false
}
func AesEncrypt(origData []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv) //初始向量的长度必须等于块block的长度16字节
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}
func PKCS5Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}
