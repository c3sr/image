#include <immintrin.h>

typedef float float3 __attribute__((ext_vector_type(3)));

#define UNROLL_FACTOR 4
int hwc_to_chw(float *__restrict__ output, const float *__restrict__ input,
               const int height, const int width) {
  const int channels = 3;
  float *const __restrict__ rOutputPlane = &output[0 * width * height];
  float *const __restrict__ bOutputPlane = &output[1 * width * height];
  float *const __restrict__ gOutputPlane = &output[2 * width * height];
  int inputOffset = 0;
  int outputOffset = 0;
  for (int yy = 0; yy < height; yy++) {
    for (int xx = 0; xx < width / UNROLL_FACTOR; xx++) {
      __builtin_prefetch(&input[inputOffset + channels * UNROLL_FACTOR], 0, 1);
#pragma unroll
      for (int ii = 0; ii < UNROLL_FACTOR; ii++) {
        const float3 rgb{input[inputOffset], input[inputOffset + 1],
                         input[inputOffset + 2]};
        rOutputPlane[outputOffset] = rgb.x;
        bOutputPlane[outputOffset] = rgb.y;
        gOutputPlane[outputOffset] = rgb.z;
        inputOffset += channels;
        outputOffset++;
      }
    }
  }
  return 0;
}
