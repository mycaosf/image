//16bit channel should big endian
typedef enum {
  imageColorGray8       = 0,
  imageColorGray16,
  imageColorRGB,
  imageColorRGBA,
  imageColorRGBA64,
  imageColorCMYK,
} imageColor;

typedef enum {
  imageTypePNG          = 0,
  imageTypeJPEG,
  imageTypeTIFF,
  imageTypeBMP,
} imageType;

typedef struct {
  const char    *fileName;      //should be UTF8 encode.
  imageType     t;
  imageColor    c;
  size_t        width;
  size_t        height;
  //data size should be:
  //gray8: width * height
  //gray16: width * 2 * height
  //rgba: width * 4 * height
  //rgba64: width * 8 * height
  //cmyk: width * 4 * height
  const void    *data;
} imageEncodeParam;

#ifdef __cplusplus
extern "C" {
#endif
//imageEncode encode data to file.
//Succeeded: return > 0
int     ImageEncode(imageEncodeParam *p);
#ifdef __cplusplus
}
#endif
