#ifndef AES128_H
#define AES128_H

#include <string>
#include <iostream>

#include <openssl/aes.h>
using namespace std;

class AES128
{
private:
    string key;
    string iv;
public:
    AES128(string key,string iv);

    int encrypto(uint8_t *raw,int length,uint8_t*encode);

    int decrypto(uint8_t* encode,int length,uint8_t *raw);

    int decrypto(string base64,uint8_t *raw);

};

#endif // AES128_H
