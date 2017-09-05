#include "stdint.h"
#include "string.h"
typedef float float3 __attribute__((ext_vector_type(3)));

template <typename T> static const T min(const T &x, const T &y) {
  return x < y ? x : y;
}

extern "C" void resize_bilinear(float *dst,const  float *src, uint64_t dst_h,
                                uint64_t dst_w, uint64_t src_h,
                                uint64_t src_w) {
  float scale_y = (float)src_h / (float)dst_h;
  float scale_x = (float)src_w / (float)dst_w;

  // printf("real_scale_x =%f\n", scale_x);
  // printf("real_scale_y =%f\n", scale_y);

  for (int i = 0; i < dst_h; i++) {
    const int y = min((uint64_t)(i * scale_y), src_h - 1);
    for (int j = 0; j < dst_w; j++) {

      const int x = min((uint64_t)(j * scale_x), src_w - 1);
      // printf("j = %d  x =%d\n", j, x);
      // printf("i = %d  y =%d\n", i, y);
      const float *input = &src[3 * (y * src_w + x)];
      float *output = &dst[3 * (i * dst_w + j)];
#pragma unroll
      for (int k = 0; k < 3; k++) {
        *output++ = *input++;
      }
    }
  }
}
