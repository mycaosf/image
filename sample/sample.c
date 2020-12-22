#include <stdio.h>
#include "../image.h"

int main()
{
  unsigned char data[64 * 64];
  imageEncodeParam      param = {
    u8"./test.png", imageTypePNG, imageColorGray8, 64, 64, data,
  };
  memset(data, 0xff, sizeof(data));
  memset(data + 64 * 10, 0, 64);
  for(int i = 0; i < 64; i++)
    data[i * 64 + 32] = 0x80;

  printf("ImageEncode return code: %d\n", ImageEncode(&param));

  return 0;
}
