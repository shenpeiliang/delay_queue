package rsa2

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
)

type Rsa2 struct {
}

//创建实例
func NewStore() *Rsa2 {
	return &Rsa2{}
}

//数据签名 - 私钥
func (r Rsa2) CreateSign(data, key []byte) (signature []byte) {
	h := sha256.New()
	h.Write(data)
	//数据的sha256校验和
	hashed := h.Sum(nil)

	//解析pem格式的私钥
	block, _ := pem.Decode(key)

	if block == nil {
		return
	}

	//解析私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}

	//使用私钥生成签名
	signature, err = rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)

	return
}

//签名验证 - 公钥
func (r Rsa2) VerySign(data, signature, key []byte) (err error) {
	//解析pem格式的私钥
	block, _ := pem.Decode(key)

	if block == nil {
		return
	}

	//解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}

	hashed := sha256.Sum256(data)

	err = rsa.VerifyPKCS1v15(pubInterface.(*rsa.PublicKey), crypto.SHA256, hashed[:], signature)

	return
}

//使用公钥加密
func (r Rsa2) Encrypt(data, key []byte) (ciphertext []byte, err error) {
	//解析pem格式的公钥
	block, _ := pem.Decode(key)

	if block == nil {
		return
	}

	//解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}

	//类型断言
	publicKey := pubInterface.(*rsa.PublicKey)

	//获取加密密文
	ciphertext, err = rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)

	return
}

//使用私钥解密
func (r Rsa2) Decrypt(ciphertext, key []byte) (data []byte, err error) {
	//解析pem格式的公钥
	block, _ := pem.Decode(key)

	if block == nil {
		return
	}

	//解析私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}

	//解密
	data, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)

	return
}
