#include <cstdint>
#include <iostream>
#include <math.h>
#include <stdio.h>
#include <stdlib.h>

#if 1
extern "C" void resize_bilinear(float *dst, float *src, uint64_t dst_h,
                                uint64_t dst_w, uint64_t src_h,
                                uint64_t src_w) {
  float scale_y = (float)src_h / (float)dst_h;
  float scale_x = (float)src_w / (float)dst_w;

  // printf("real_scale_x =%f\n", scale_x);
  // printf("real_scale_y =%f\n", scale_y);

  for (int i = 0; i < dst_h; i++) {
    int y = i * scale_y;

    if (y > src_h) {
      y = src_h - 1;
    }
    for (int j = 0; j < dst_w; j++) {

      int x = j * scale_x;

      if (x > src_w) {
        x = src_w - 1;
      }
      // printf("j = %d  x =%d\n", j, x);
      // printf("i = %d  y =%d\n", i, y);
      #pragma unroll
      for (int k = 0; k < 3; k++) {
        dst[3 * (i * dst_w + j) + k] = src[3 * (y * src_w + x) + k];
      }
    }
  }
}
#else
extern "C" void resize(float *__restrict__ output,
                       const float *const __restrict__ input0,
                       const uint64_t height, const uint64_t width) {
  const int channels = 3;
  const float3 *__restrict__ input = (float3 *)input0;
  float *const __restrict__ firstOutputPlane = &output[0 * width * height];
  float *const __restrict__ secondOutputPlane = &output[1 * width * height];
  float *const __restrict__ thirdOutputPlane = &output[2 * width * height];
  for (uint64_t yy = 0; yy < height; yy++) {
    uint64_t xx = 0;
    for (; xx < ROUND_DOWN(width, UNROLL_FACTOR); xx += UNROLL_FACTOR) {
      const uint64_t offset = yy * width + xx;
      __builtin_prefetch(&input[(offset + 1) * UNROLL_FACTOR], 0, 1);
#pragma unroll
      for (int ii = 0; ii < UNROLL_FACTOR; ii++) {
        const float3 rgb = input[offset + ii];
        firstOutputPlane[offset + ii] = rgb.x;
        secondOutputPlane[offset + ii] = rgb.y;
        thirdOutputPlane[offset + ii] = rgb.z;
      }
    }

    for (; xx < width; xx++) {
      const uint64_t offset = yy * width + xx;
      const float3 rgb = input[offset];
      firstOutputPlane[offset] = rgb.x;
      secondOutputPlane[offset] = rgb.y;
      thirdOutputPlane[offset] = rgb.z;
    }
  }
}
#endif

// int main() {
//   float a[][12] = {{0, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 3},
//                    {0, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 3},
//                    {0, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 3},
//                    {0, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 3}};
//   int m = 4, n = 6;
//   float *b = (float *)malloc(sizeof(float) * m * n * 3);
//   resize_bilinear(b, (float *)a, m, n, 4, 4);
//   for (int i = 0; i < m; i++) {
//     for (int j = 0; j < n; j++) {
//       for (int k = 0; k < 3; k++) {
//         printf("%f ", b[3 * (i * n + j) + k]);
//       }
//     }
//     printf("\n");
//   }
//   return 0;
// }
