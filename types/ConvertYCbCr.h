
#include <stdint.h>

#define UINT8 uint8_t
#define INT16 int16_t

void image_types_ImagingConvertYCbCr2RGB(UINT8 *out, const UINT8 *Y, const UINT8 *Cb,
                                         const UINT8 *Cr, int yStride, int cStride,
                                         int width, int height);
void image_types_ImagingConvertYCbCr2BGR(UINT8 *out, const UINT8 *Y, const UINT8 *Cb,
                                         const UINT8 *Cr, int yStride, int cStride,
                                         int width, int height);
