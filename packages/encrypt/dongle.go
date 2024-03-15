package encrypt

import "github.com/golang-module/dongle"

// Base64Encode base 64 encode
func Base64Encode(str string) string {
	return dongle.Encode.FromString(str).ByBase64().ToString()
}

// Base64Decode base 64 decode
func Base64Decode(str string) string {
	return dongle.Decode.FromString(str).ByBase64().ToString()
}

type DongleAes struct {
	cipher *dongle.Cipher
}

func NewCbcPkcs7Aes(key string, iv string) *DongleAes {
	cipher := dongle.NewCipher()
	cipher.SetMode(dongle.CBC)      // CBC、CFB、OFB、CTR、ECB
	cipher.SetPadding(dongle.PKCS7) // No、Empty、Zero、PKCS5、PKCS7、AnsiX923、ISO97971
	cipher.SetKey(key)              // key 长度必须是 16、24 或 32 字节
	cipher.SetIV(iv)                // iv 长度必须是 16 字节，ECB 模式不需要设置 iv
	return &DongleAes{
		cipher: cipher,
	}
}

func (c *DongleAes) AesEncrypt(str string) string {
	return dongle.Encrypt.FromString(str).ByAes(c.cipher).ToRawString()
}

func (c *DongleAes) AesDecrypt(encryptedStr string) string {
	return dongle.Decrypt.FromRawString(encryptedStr).ByAes(c.cipher).ToString()
}
