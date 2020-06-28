package crypto

type Crypto struct {
	store Store
}

//设置处理器
func SetStore(store Store) *Crypto {
	return &Crypto{store: store}
}

type Store interface {
	//生成签名
	CreateSign(data, key []byte) (signature []byte)
	//签名验证
	VerySign(data, signature, key []byte) (err error)

	//加密
	Encrypt(data, key []byte) (ciphertext []byte, err error)
	//解密
	Decrypt(ciphertext, key []byte) (data []byte, err error)
}
