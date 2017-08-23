package codec

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"github.com/carsonsx/log4g"
	"math/big"
)

var Private_Key = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y
7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7
Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB
AoGBAJlNxenTQj6OfCl9FMR2jlMJjtMrtQT9InQEE7m3m7bLHeC+MCJOhmNVBjaM
ZpthDORdxIZ6oCuOf6Z2+Dl35lntGFh5J7S34UP2BWzF1IyyQfySCNexGNHKT1G1
XKQtHmtc2gWWthEg+S6ciIyw2IGrrP2Rke81vYHExPrexf0hAkEA9Izb0MiYsMCB
/jemLJB0Lb3Y/B8xjGjQFFBQT7bmwBVjvZWZVpnMnXi9sWGdgUpxsCuAIROXjZ40
IRZ2C9EouwJBAOPjPvV8Sgw4vaseOqlJvSq/C/pIFx6RVznDGlc8bRg7SgTPpjHG
4G+M3mVgpCX1a/EU1mB+fhiJ2LAZ/pTtY6sCQGaW9NwIWu3DRIVGCSMm0mYh/3X9
DAcwLSJoctiODQ1Fq9rreDE5QfpJnaJdJfsIJNtX1F+L3YceeBXtW0Ynz2MCQBI8
9KP274Is5FkWkUFNKnuKUK4WKOuEXEO+LpR+vIhs7k6WQ8nGDd4/mujoJBr5mkrw
DPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO
AQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O
-----END RSA PRIVATE KEY-----
`)

var Public_Key = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv
ppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd
wYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL
AUeJ6PeW+DAkmJWF6QIDAQAB
-----END PUBLIC KEY-----
`)

// 从RSA的原理来看，公钥加密私钥解密和私钥加密公钥解密应该是等价的，
// 在某些情况下，比如共享软件加密，我们需要用私钥加密注册码或注册文件，
// 发给用户，用户用公钥解密注册码或注册文件进行合法性验证。

func hash(text string) []byte {
	hash := sha512.New()
	hash.Write([]byte(text))
	return hash.Sum(nil)
}

//私钥加密
func PrivateEncrypt(text string) (string, error) {
	block, _ := pem.Decode(Private_Key)
	if block == nil {
		text := "private key error"
		log4g.Error(text)
		return "", errors.New(text)
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log4g.Error(err)
		return "", err
	}
	data, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.Hash(0), []byte(text))
	if err != nil {
		log4g.Error(err)
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

//公钥解密
func PublicVerfiy(text, ciphertext string) (bool, error) {
	cipherbyte, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		log4g.Error(err)
		return false, err
	}
	block, _ := pem.Decode(Public_Key)
	if block == nil {
		text := "public key error"
		log4g.Error(text)
		return false, errors.New(text)
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log4g.Error(err)
		return false, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA512, hash(text), cipherbyte)
	if err != nil {
		log4g.Error(err)
		return false, err
	}
	return true, nil
}

var (
	ErrDataToLarge     = errors.New("message too long for RSA public key size")
	ErrDataLen         = errors.New("data length error")
	ErrDataBroken      = errors.New("data broken, first byte is not zero")
	ErrKeyPairDismatch = errors.New("data is not encrypted by the private key")
	ErrDecryption      = errors.New("decryption error")
	ErrPublicKey       = errors.New("get public key error")
	ErrPrivateKey      = errors.New("get private key error")
)

/*公钥解密*/
func PublicDecrypt(ciphertext string) (string, error) {

	cipherbyte, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		log4g.Error(err)
		return "", err
	}
	block, _ := pem.Decode(Public_Key)
	if block == nil {
		text := "public key error"
		log4g.Error(text)
		return "", errors.New(text)
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log4g.Error(err)
		return "", err
	}
	pub := pubInterface.(*rsa.PublicKey)

	k := (pub.N.BitLen() + 7) / 8
	if k != len(cipherbyte) {
		return "", ErrDataLen
	}
	m := new(big.Int).SetBytes(cipherbyte)
	if m.Cmp(pub.N) > 0 {
		return "", ErrDataToLarge
	}
	m.Exp(m, big.NewInt(int64(pub.E)), pub.N)
	d := leftPad(m.Bytes(), k)
	if d[0] != 0 {
		return "", ErrDataBroken
	}
	if d[1] != 0 && d[1] != 1 {
		return "", ErrKeyPairDismatch
	}
	var i = 2
	for ; i < len(d); i++ {
		if d[i] == 0 {
			break
		}
	}
	i++
	if i == len(d) {
		return "", nil
	}
	return string(d[i:]), nil
}

/*从crypto/rsa复制 */
func leftPad(input []byte, size int) (out []byte) {
	n := len(input)
	if n > size {
		n = size
	}
	out = make([]byte, size)
	copy(out[len(out)-n:], input)
	return
}

// 公钥加密
func PublicEncrypt(text string) (string, error) {
	block, _ := pem.Decode(Public_Key)
	if block == nil {
		text := "public key error"
		log4g.Error(text)
		return "", errors.New(text)
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log4g.Error(err)
		return "", err
	}
	pub := pubInterface.(*rsa.PublicKey)
	data, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(text))
	if err != nil {
		log4g.Error(err)
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// 私钥解密
func PrivateDecrypt(ciphertext string) (string, error) {
	cipherbyte, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		log4g.Error(err)
		return "", err
	}
	block, _ := pem.Decode(Private_Key)
	if block == nil {
		text := "private key error"
		log4g.Error(text)
		return "", errors.New(text)
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log4g.Error(err)
		return "", err
	}
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, cipherbyte)
	if err != nil {
		log4g.Error(err)
		return "", err
	}
	return string(data), nil
}
