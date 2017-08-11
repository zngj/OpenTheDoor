#include "AES128.h"

#include <string.h>
#include "Lib/Base64/Base64.h"

AES128::AES128(string key, string iv)
{
    this->key=key;
    this->iv=iv;
}


//最大支持512字节
int AES128::encrypto(uint8_t *raw, int length, uint8_t *encode)
{
    uint8_t key[16+1];
    uint8_t ivec[AES_BLOCK_SIZE];
    AES_KEY AesKey;

    uint8_t input[512];
    memset(input,0,512);
    memset(key,0,sizeof(key));
    memcpy(key,this->key.c_str(),16);
    memcpy(ivec,this->iv.c_str(),AES_BLOCK_SIZE);

    memset(&AesKey, 0x00, sizeof(AES_KEY));
    if(AES_set_encrypt_key(key, 128, &AesKey) < 0) return -1;

    int encLen=length+AES_BLOCK_SIZE-(length%AES_BLOCK_SIZE);


    memcpy(input, raw, length);

    memset(encode,0,encLen);
    AES_cbc_encrypt(input, encode,encLen, &AesKey, ivec, AES_ENCRYPT);

    return encLen;


}

int AES128::decrypto(uint8_t* encode,int length,uint8_t *raw)
{
    uint8_t key[16+1];
    uint8_t ivec[AES_BLOCK_SIZE];
    AES_KEY AesKey;
    memcpy(key,this->key.c_str(),16);
    key[16]=0;
    memcpy(ivec,this->iv.c_str(),AES_BLOCK_SIZE);

    memset(&AesKey, 0x00, sizeof(AES_KEY));
    if(AES_set_decrypt_key(key, 128, &AesKey) < 0) return -1;


    memset(raw,0,length);

    AES_cbc_encrypt(encode, raw,length, &AesKey, ivec, AES_DECRYPT);

    return length;
}

int AES128::decrypto(string base64, uint8_t *raw)
{

    char dec64[256];
    int num= base64_decode(base64,dec64);

    return decrypto((uint8_t*)dec64,num,raw);




}

