#include <string>  
  
std::string base64_encode(unsigned char const* , unsigned int len);  
int base64_decode(std::string const& s,char * raw);
