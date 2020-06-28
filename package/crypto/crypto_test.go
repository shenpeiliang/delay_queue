package crypto

import (
	"encoding/hex"
	"fmt"
	"gin/package/crypto/rsa2"
	"testing"
)

func TestCrypt_Encrypt(t *testing.T) {
	privateKey := `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCgPwlGJrWqTYaoMkI8jXkEI8ewQ7E57G2Fi91WTXMMK7X6GsT9
VmnRcq++Rk/VS+4IPBlfWyVRg0NfQDyuKjed21fUPa9AIbpYWHgP/tojyeYC1+Ra
Xncrt9kLp7nW4FZMJmzwU9hfxIB0nhDQqhJenjdBZuYZfkICfMqyqbVkAwIDAQAB
AoGAJRcSDXOuPrHdBhdD74ILTaL+eFTis3Z+zxdVbsFUbK+9WhtSFxUmPv1dohvi
JIuDl9JZSRHurFRGhsh2gxVwc7JXwWfD0DmD8dvdzr8q85Jml9YVZ7uhHFqSO4cY
I7dlBOd7Uwjnc39E/d+1E/kWVNfKt7opPHgt02zOHLSxkbECQQDS7H3myu3oLOi0
Slpd1MmmHVOo2cqJ1b3H6E8JtEjmHGswWTYvQNAe4yZ+Kffsp5LUYujedncPKvEj
4G+iz44bAkEAwn4Bx30FKTri/tybgSnCWKwTGSX479829Xucrm5pYU/3D5/PeJQL
Ra4YSyg2/hU3ZBrue6CdzYJgGXNGEWhAOQJBALMlOB4A96X+FruidzRA2fBj8j10
lakSSHl1H0RfwpbnRkcvTm0+AEZrqbL4lGGFRplrVNw2BBN25o8RPeArp0cCQEhu
kw0PI1fqhVUzJXqh6a4KT4aDHMWAlMAxi/VuSzKhjDo2Yxbd06DcqFF9JZXUou9W
FFDYTUyW7GEuC/85mwkCQCOEjUQX0C3JCSr6fyZIjpEr+znyc9eFHyBp+533Ur4g
eFu2ewJ3ufJiUBmEj1rEQku8W7h9DS2rXl10IiSwUAA=
-----END RSA PRIVATE KEY-----`

	publicKey := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCgPwlGJrWqTYaoMkI8jXkEI8ew
Q7E57G2Fi91WTXMMK7X6GsT9VmnRcq++Rk/VS+4IPBlfWyVRg0NfQDyuKjed21fU
Pa9AIbpYWHgP/tojyeYC1+RaXncrt9kLp7nW4FZMJmzwU9hfxIB0nhDQqhJenjdB
ZuYZfkICfMqyqbVkAwIDAQAB
-----END PUBLIC KEY-----`

	store := rsa2.NewStore()
	crypto := SetStore(store)

	data := "hello go"
	fmt.Println("对消息进行签名操作...")
	signData := crypto.store.CreateSign([]byte(data), []byte(privateKey))
	fmt.Println("消息的签名信息： ", hex.EncodeToString(signData))
	fmt.Println("\n对签名信息进行验证...")
	if err := crypto.store.VerySign([]byte(data), signData, []byte(publicKey)); err == nil {
		fmt.Println("签名信息验证成功，确定是正确私钥签名！！")
	}

	fmt.Println("-------------------------------进行加密解密操作-----------------------------------------")
	ciphertext, _ := crypto.store.Encrypt([]byte(data), []byte(publicKey))
	fmt.Println("公钥加密后的数据：", hex.EncodeToString(ciphertext))
	sourceData, _ := crypto.store.Decrypt(ciphertext, []byte(privateKey))
	fmt.Println("私钥解密后的数据：", string(sourceData))

}
