#include "CryptoManager.h"


CryptoManager * CryptoManager::mng=nullptr;
mutex CryptoManager::mtx;
CryptoManager::CryptoManager()
{
    const char keyPub[]= "-----BEGIN PUBLIC KEY-----\n"\
              "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv\n"\
              "ppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd\n"\
              "wYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL\n"\
              "AUeJ6PeW+DAkmJWF6QIDAQAB\n"\
              "-----END PUBLIC KEY-----";

      const char keyPri[]="-----BEGIN RSA PRIVATE KEY-----\n"\
              "MIICXQIBAAKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y\n"\
              "7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7\n"\
              "Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB\n"\
              "AoGBAJlNxenTQj6OfCl9FMR2jlMJjtMrtQT9InQEE7m3m7bLHeC+MCJOhmNVBjaM\n"\
              "ZpthDORdxIZ6oCuOf6Z2+Dl35lntGFh5J7S34UP2BWzF1IyyQfySCNexGNHKT1G1\n"\
              "XKQtHmtc2gWWthEg+S6ciIyw2IGrrP2Rke81vYHExPrexf0hAkEA9Izb0MiYsMCB\n"\
              "/jemLJB0Lb3Y/B8xjGjQFFBQT7bmwBVjvZWZVpnMnXi9sWGdgUpxsCuAIROXjZ40\n"\
              "IRZ2C9EouwJBAOPjPvV8Sgw4vaseOqlJvSq/C/pIFx6RVznDGlc8bRg7SgTPpjHG\n"\
              "4G+M3mVgpCX1a/EU1mB+fhiJ2LAZ/pTtY6sCQGaW9NwIWu3DRIVGCSMm0mYh/3X9\n"\
              "DAcwLSJoctiODQ1Fq9rreDE5QfpJnaJdJfsIJNtX1F+L3YceeBXtW0Ynz2MCQBI8\n"\
              "9KP274Is5FkWkUFNKnuKUK4WKOuEXEO+LpR+vIhs7k6WQ8nGDd4/mujoJBr5mkrw\n"\
              "DPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO\n"\
              "AQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O\n"\
              "-----END RSA PRIVATE KEY-----";

      this->rsa1024=new RSA1024(keyPub,keyPri);
      this->aes128=new AES128("5454395434473454","6916665466156476");

}

CryptoManager *CryptoManager::getInstance()
{
    unique_lock<mutex> lock(mtx);
    if(mng==nullptr)
    {
        mng=new CryptoManager();
    }
    return mng;
}

void CryptoManager::changeRSAPubKey(string key)
{
    if(this->rsa1024!=nullptr) delete this->rsa1024;

    this->rsa1024=new RSA1024(key.c_str(),nullptr);
}


int CryptoManager::aesDecrypt(uint8_t* enc,int length,uint8_t*raw)
{
    if(this->aes128==nullptr) return -1;

    return this->aes128->decrypto(enc,length,raw);
}

int CryptoManager::rsaDecrypt(uint8_t* rsa,uint8_t*raw)
{
    if(this->rsa1024==nullptr) return -1;

    return this->rsa1024->decrypto(rsa,128,raw);
}
