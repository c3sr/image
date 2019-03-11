
#include <stdint.h>

#define UINT8 uint8_t
#define INT16 int16_t

void ImagingConvertRGB2YCbCr(UINT8 *out, const UINT8 *in, int pixels);
void ImagingConvertBGR2YCbCr(UINT8 *out, const UINT8 *in, int pixels);
void ImagingConvertYCbCr2RGB(UINT8 *out, const UINT8 *Y, const UINT8 *Cb,
                             const UINT8 *Cr, int yStride, int cStride,
                             int width, int height);
void ImagingConvertYCbCr2BGR(UINT8 *out, const UINT8 *Y, const UINT8 *Cb,
                             const UINT8 *Cr, int yStride, int cStride,
                             int width, int height);
