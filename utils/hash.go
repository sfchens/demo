package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
)

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Md5 计算字符串的 MD5 散列值
// 如果可选的 raw_output 被设置为 true，那么 md5 摘要将以 16 字符长度的原始二进制格式返回。
func Md5(s string, rawOutput ...bool) string {
	h := md5.New()
	h.Write([]byte(s))

	if len(rawOutput) > 0 && rawOutput[0] == true {
		return string(h.Sum(nil))
	}
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1 计算字符串的 sha1 散列值
// 如果可选的 raw_output 参数被设置为 true， 那么 sha1 摘要将以 20 字符长度的原始二进制格式返回， 否则返回值为 40 字符长度的十六进制数
func Sha1(s string, rawOutput ...bool) string {
	o := sha1.New()
	o.Write([]byte(s))

	if len(rawOutput) > 0 && rawOutput[0] == true {
		return string(o.Sum(nil))
	}
	return hex.EncodeToString(o.Sum(nil))
}

func AESKey(key string, keyLen int) []byte {
	bytes := make([]byte, keyLen) //generate a random 32 byte key for AES-256
	for i, ch := range key {
		if i >= keyLen {
			break
		}
		bytes[i] = byte(ch)
	}
	return bytes
}

func AESEncrypt(stringToEncrypt string, keyString string) (encryptedString string, err error) {
	//Since the key is in string, we need to convert decode it to bytes
	key := AESKey(keyString, 32)
	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

func AESDecrypt(encryptedString string, keyString string) (decryptedString string, err error) {
	key := AESKey(keyString, 32)
	enc, err := hex.DecodeString(encryptedString)

	if err != nil {
		return
	} else if len(enc) == 0 {
		return
	}

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	if len(enc) < nonceSize {
		return "", errors.New("the length of encrypted data is less than the nonce size")
	}

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return
	}

	return fmt.Sprintf("%s", plaintext), nil
}
