#include "RSA1024.h"

#include <stdio.h>
#include <string.h>

RSA1024::RSA1024(string &fileName)
{
    this->rsaKey=NULL;
    FILE* fileKey = fopen(fileName.c_str(), "rb");
    if(fileKey==NULL) return;

    this->rsaKey = PEM_read_RSA_PUBKEY(fileKey, NULL, NULL, NULL);


    fclose(fileKey);

}

int RSA1024::decrypto(uint8_t *encrypto, uint8_t *raw)
{
    if(this->rsaKey==NULL) return -1;

    int rsaLen=RSA_size(this->rsaKey);
    memset(raw,0,rsaLen+1);
    int ret = RSA_public_decrypt(rsaLen,encrypto, raw, this->rsaKey, RSA_NO_PADDING);

    return ret;
}
