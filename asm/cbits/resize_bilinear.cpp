#include <cstdint>
#include <math.h>
#include <stdio.h>
#include <stdlib.h>

#define UNROLL_FACTOR 8

typedef float float3 __attribute__((ext_vector_type(3)));
#define ROUND_DOWN(x, s) ((x) & ~((s)-1))

#if 1
extern "C" void resize_hori(uint8_t *dst, uint8_t *src, uint64_t h,
                            uint64_t dst_w, uint64_t src_w) {
  float real_scale = ((float)src_w - 1) / ((float)dst_w - 1);
  for (int i = 0; i < h; i++) {
    for (int j = 0; j < dst_w; j++) {
      int x = floor((float)j * real_scale);
      float dx = j * real_scale - x;
      if (x >= (src_w - 1)) {
        dst[i * dst_w + j] = src[i * src_w + src_w - 1];
      } else {
        dst[i * dst_w + j] = (uint8_t)((1.0f - dx) * src[i * src_w + x] +
                                       (dx)*src[i * src_w + x + 1]);
      }
    }
  }
}

extern "C" void resize_vert(uint8_t *dst, uint8_t *src, uint64_t w,
                            uint64_t dst_h, uint64_t src_h) {
  float real_scale = ((float)src_h - 1) / ((float)dst_h - 1);
  for (int i = 0; i < w; i++) {
    for (int j = 0; j < dst_h; j++) {
      int y = floor((float)j * real_scale);
      float dy = j * real_scale - y;
      if (y >= (src_h - 1)) {
        dst[j * w + i] = src[i + (src_h - 1) * w];
      } else {
        dst[j * w + i] =
            (uint8_t)((1.0f - dy) * src[i + y * w] + dy * src[i + (y + 1) * w]);
      }
    }
  }
}

extern "C" void resize_bilinear(uint8_t *dst, uint8_t *src, uint64_t dst_h,
                                uint64_t dst_w, uint64_t src_h,
                                uint64_t src_w) {
  const int channels = 1;
  uint8_t *tmp = (uint8_t *)malloc(sizeof(uint8_t) * dst_w * src_h);
  resize_hori(tmp, src, src_h, dst_w, src_w);
  resize_vert(dst, tmp, dst_w, dst_h, src_h);
  free(tmp);
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
//   uint8_t a[16] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15};
//   int m = 2, n = 2;
//   uint8_t *b = (uint8_t *)malloc(sizeof(uint8_t) * m * n);
//   resize_bilinear(b, a, m, n, 4, 4);
//   for (int i = 0; i < m; i++) {
//     for (int j = 0; j < n; j++) {
//       printf("%d ", b[i * n + j]);
//     }
//     printf("\n");
//   }
//   return 0;
// }
