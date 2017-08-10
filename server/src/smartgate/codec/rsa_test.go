package codec_test

import (
	"testing"
	"github.com/carsonsx/log4g"
	"smartgate/codec"
)

func TestRSA(t *testing.T)  {
	text := "helloworld"
	encrypt_text, _ := codec.PrivateEncrypt(text)
	log4g.Debug(encrypt_text)
	//log4g.Debug(codec.PublicVerfiy(text, encrypt_text))
	ntext, err := codec.PublicDecrypt(encrypt_text)
	log4g.Error(err)
	log4g.Debug(ntext)
}