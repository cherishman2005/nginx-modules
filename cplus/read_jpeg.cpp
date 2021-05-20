// https://blog.csdn.net/Cashey1991/article/details/6769038

#include <stdio.h>
#include <string>

void readfile(const char* uid)
{
  char buf[300*1024];
  uint32_t size = 300*1024;

  FILE* f = fopen((std::string(uid) + ".jpeg").c_str(), "r");
  if (f == nullptr) {
    printf("scy open file fail %s\n", uid);
    return;
  }

  //fwrite(buf, 1, size, f);
  int n = fread(buf, sizeof(uint8_t), size, f);
  fclose(f);
  
  printf("n=%u\n", n);
}

int main() {
  readfile("2314776959");
  return 0;
}
