package toolkit

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"
)

// GenUID 随机生成UUID
func GenUID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	b64 := base64.URLEncoding.EncodeToString(b)
	h := md5.New()
	h.Write([]byte(b64))
	return hex.EncodeToString(h.Sum(nil))
}

// Base64Encode base64编码
func Base64Encode(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

// Base64Decode base64解码
func Base64Decode(content string) string {
	b, err := base64.URLEncoding.DecodeString(content)
	if err != nil {
		return ""
	}
	return string(b)
}

// GetTime 获取当前时间
func GetTime() time.Time {
	return time.Now()
}

// Md5 获取MD5值
func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// SHA256 获取SHA256
func SHA256(buf []byte) string {
	hash := sha256.New()
	hash.Write(buf)
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}

// RawURLEncode 获取转义后的URL
func RawURLEncode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}


// AesEncrypt AES算法加密字符串
func AesEncrypt(mess []byte, key []byte) ([]byte, error) {
    var iv = key[:aes.BlockSize]
    encrypted := make([]byte, len(mess))
    aesBlockEncrypter, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
    aesEncrypter.XORKeyStream(encrypted, mess)
    return encrypted, nil
}

// AesDecrypt AES算法解密字符串
func AesDecrypt(src []byte, key []byte) (mess []byte, err error) {
    defer func() {
        //错误处理
        if e := recover(); e != nil {
            err = e.(error)
        }
    }()
    var iv = key[:aes.BlockSize]
    decrypted := make([]byte, len(src))
    var aesBlockDecrypter cipher.Block
    aesBlockDecrypter, err = aes.NewCipher([]byte(key))
    if err != nil {
        return nil, err
    }
	
    aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
    aesDecrypter.XORKeyStream(decrypted, src)
    return decrypted, nil
}