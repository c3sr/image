#define UNROLL_FACTOR 8

typedef float float3 __attribute__((ext_vector_type(3)));

extern "C" int hwc_to_chw(float *__restrict__ output,
                          const float *__restrict__ input0, const int height,
                          const int width) {
  const int channels = 3;
  float *const __restrict__ firstOutputPlane = &output[0 * width * height];
  float *const __restrict__ secondOutputPlane = &output[1 * width * height];
  float *const __restrict__ thirdOutputPlane = &output[2 * width * height];
  const float3 *__restrict__ input = (float3 *)input0;
  int offset = 0;
  for (int yy = 0; yy < height; yy++) {
    int xx = 0;
    for (xx = 0; xx < width / UNROLL_FACTOR; xx++) {
      __builtin_prefetch(&input[offset * UNROLL_FACTOR], 0, 1);
#pragma unroll
      for (int ii = 0; ii < UNROLL_FACTOR; ii++) {
        const float3 rgb = input[offset + ii];
        firstOutputPlane[offset + ii] = rgb.x;
        secondOutputPlane[offset + ii] = rgb.y;
        thirdOutputPlane[offset + ii] = rgb.x;
      }
      offset += UNROLL_FACTOR;
    }

    for (; xx < width; xx++) {
      const float3 rgb = input[offset];
      firstOutputPlane[offset] = rgb.x;
      secondOutputPlane[offset] = rgb.y;
      thirdOutputPlane[offset] = rgb.x;
      offset++;
    }
  }
  return 0;
}
