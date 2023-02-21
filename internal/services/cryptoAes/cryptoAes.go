package cryptoAes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"redisadmin/internal/configs"
	"redisadmin/internal/consts"
)

// 密钥
var cryptoKey string

func init() {
	cryptoKey = configs.GetEnvVal(consts.ENV_CONF_CRYPTOKEY)
}

// Key 初始16位字符串，24位字符串或32位字符串
func Encrypt(plaintext string, keyStr string) (string, error) {
	if keyStr == "" {
		keyStr = cryptoKey
	}

	key := []byte(keyStr)

	block, err := aes.NewCipher(key) // 返回一个使用AES算法的cipher.Block接口
	if err != nil {
		return "", err
	}

	plaintextBytes := []byte(plaintext)
	plaintextPad := pkcs7Pad(plaintextBytes, block.BlockSize()) // 填充

	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()]) // 返回一个CBC加密模式的cipher.BlockMode接口
	cryptograph := make([]byte, len(plaintextPad))                      // 存放密文的数组
	blockMode.CryptBlocks(cryptograph, plaintextPad)                    // 加密

	return hex.EncodeToString(cryptograph), nil
}

// 填充函数
func pkcs7Pad(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding) // 生成填充的文本
	return append(ciphertext, padText...)
}

//Key 初始16位字符串，24位字符串或32位字符串
func Decrypt(ciphertextHex string, keyStr string) (plaintext string, err error) {
	defer func() {
		if err1 := recover(); err1 != nil {
			plaintext = ""
			err = errors.New("cryptoAes解密错误" + fmt.Sprintf("%v", err1))
		}
	}()
	if keyStr == "" {
		keyStr = cryptoKey
	}

	key := []byte(keyStr)

	// 转为16进制字符串
	cipherBytes, _ := hex.DecodeString(ciphertextHex)

	block, err := aes.NewCipher(key) // 返回一个使用AES算法的cipher.Block接口
	if err != nil {
		return
	}

	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()]) // 返回一个CBC解密模式的cipher.BlockMode接口
	plaintextPad := make([]byte, len(cipherBytes))                      // 存放明文的数组
	blockMode.CryptBlocks(plaintextPad, cipherBytes)                    // 解密

	plaintextBytes, err := pkcs7Unpad(plaintextPad) // 去除填充
	if err != nil {
		return
	}

	plaintext = string(plaintextBytes)

	return
}

//Padding去除函数
func pkcs7Unpad(plaintext []byte) ([]byte, error) {
	length := len(plaintext)
	unpadNum := int(plaintext[length-1])
	index := length - unpadNum
	if index < 0 {
		return nil, errors.New("长度错误")
	}
	return plaintext[:index], nil
}
