#ifndef CRYPTOMANAGER_H
#define CRYPTOMANAGER_H

#include <string>
#include <mutex>

#include "Crypto/AES128.h"
#include "Crypto/RSA1024.h"

using namespace std;
class CryptoManager
{
private:
    static CryptoManager* mng;
    static mutex mtx;
private:
    CryptoManager();
    RSA1024 *rsa1024;
    AES128 * aes128;
public:
    static CryptoManager* getInstance();
    void changeRSAPubKey(string key);
    int aesDecrypt(string base64, uint8_t *raw);
    int rsaDecrypt(uint8_t *rsa, uint8_t *raw);
};

#endif // CRYPTOMANAGER_H
