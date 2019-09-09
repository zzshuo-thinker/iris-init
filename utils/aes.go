package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"strings"
	"os"
)

type AesCryptor struct {
	key []byte
	iv  []byte
}

var (
	aesCryptorI = AesCryptor{
		key: []byte(os.Getenv("aesKey")),
		iv: []byte(os.Getenv("aesIv")),
	}
)

func GetAesCryptor() *AesCryptor {
	return &aesCryptorI
}

//加密数据
func (a *AesCryptor) Encrypt(data string) (string) {
	aesBlockEncrypter, err := aes.NewCipher(a.key)
	content := PKCS5Padding([]byte(data), aesBlockEncrypter.BlockSize())
	encrypted := make([]byte, len(content))
	if err != nil {
		fmt.Println("ase util err", err.Error())
		return ""
	}
	aesEncrypter := cipher.NewCBCEncrypter(aesBlockEncrypter, a.iv)
	aesEncrypter.CryptBlocks(encrypted, content)
	str := hex.EncodeToString(encrypted)
	return strings.ToUpper(str)
}

//解密数据
func (a *AesCryptor) Decrypt(src string) (string, error) {
	aesBlockDecrypter, err := aes.NewCipher(a.key)
	if err != nil {
		fmt.Println("ase util decrypt aesBlockDecrypter err", err.Error())
		return "", err
	}
	var encryptByt []byte
	encryptByt, err = hex.DecodeString(src)
	if err != nil {
		fmt.Println("ase util decrypt encryptByt err", err.Error())
		return "", err
	}
	decrypted := make([]byte, len(encryptByt))
	aesDecrypter := cipher.NewCBCDecrypter(aesBlockDecrypter, a.iv)
	aesDecrypter.CryptBlocks(decrypted, encryptByt)
	return string(PKCS5Trimming(decrypted)), nil
}

/**
PKCS5包装
*/
func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

/*
解包装
*/
func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
