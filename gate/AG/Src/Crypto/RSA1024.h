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
     RSA* rsaKey;
public:
    RSA1024(string & fileName);
    int decrypto(uint8_t* encrypto,uint8_t* raw);
};

#endif // RSA1024_H
