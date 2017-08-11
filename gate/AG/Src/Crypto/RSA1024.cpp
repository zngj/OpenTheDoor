#include "RSA1024.h"

#include <stdio.h>
#include <string.h>

RSA1024::RSA1024(string &fileName)
{
    this->rsaPubKey=NULL;
    FILE* fileKey = fopen(fileName.c_str(), "rb");
    if(fileKey==NULL) return;

    this->rsaPubKey = PEM_read_RSA_PUBKEY(fileKey, NULL, NULL, NULL);


    fclose(fileKey);

}

RSA1024::RSA1024(const char *mem)
{
   BIO * bio=BIO_new_mem_buf(mem,-1);
   if(bio==NULL) return;
   this->rsaPubKey=PEM_read_bio_RSA_PUBKEY(bio,NULL,NULL,NULL);

   BIO_set_close(bio, BIO_CLOSE);
   BIO_free(bio);

}

RSA1024::RSA1024(const char *memPub, const char *memPri)
{
    if(memPub==nullptr) return ;
    BIO * bioPub=BIO_new_mem_buf(memPub,-1);
    if(bioPub==NULL) return;

    this->rsaPubKey=PEM_read_bio_RSA_PUBKEY(bioPub,NULL,NULL,NULL);

    BIO_set_close(bioPub, BIO_CLOSE);
    BIO_free(bioPub);


    if(memPri==nullptr) return;
    BIO * bioPri=BIO_new_mem_buf(memPri,-1);
    if(bioPri==NULL) return;
    this->rsaPriKey=PEM_read_bio_RSAPrivateKey(bioPri,NULL,NULL,NULL);

    BIO_set_close(bioPri, BIO_CLOSE);
    BIO_free(bioPri);


}


int RSA1024::encrypto(uint8_t* raw,int dataLen,uint8_t* enc)
{
    if(this->rsaPriKey==NULL) return -1;



    return RSA_private_encrypt(dataLen,(unsigned char *)raw,(unsigned char*)enc,this->rsaPriKey,RSA_PKCS1_PADDING);

}

int RSA1024::decrypto(uint8_t *encrypto,int dataLen, uint8_t *raw)
{
    if(this->rsaPubKey==NULL) return -1;

    int ret = RSA_public_decrypt(dataLen,encrypto, raw, this->rsaPubKey, RSA_PKCS1_PADDING);

    return ret;
}
