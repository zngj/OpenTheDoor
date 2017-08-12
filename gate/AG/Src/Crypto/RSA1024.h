#ifndef RSA1024_H
#define RSA1024_H

#include <openssl/rsa.h>
#include <openssl/err.h>
#include <openssl/pem.h>
#include <string>
using namespace std;

class RSA1024
{
private:
     RSA* rsaPubKey;
     RSA* rsaPriKey;
public:
    RSA1024(string & fileName);
    RSA1024(const char* memPub);
    RSA1024(const char *memPub,const char*memPri);
    virtual ~RSA1024();
    int decrypto(uint8_t *encrypto, int dataLen, uint8_t *raw);
    int encrypto(uint8_t *raw, int dataLen, uint8_t *enc);
};

#endif // RSA1024_H
