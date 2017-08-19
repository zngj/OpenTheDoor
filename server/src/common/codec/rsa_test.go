package codec_test

import (
	"common/codec"
	"github.com/carsonsx/log4g"
	"testing"
)

func TestRSA(t *testing.T) {
	text := "851e58b7d54d43ab8148789b839d2fd21502356760"
	encrypt_text, _ := codec.PrivateEncrypt(text)
	log4g.Debug(encrypt_text)
	//log4g.Debug(codec.PublicVerfiy(text, encrypt_text))
	ntext, err := codec.PublicDecrypt(encrypt_text)
	log4g.Error(err)
	log4g.Debug(ntext)
}
