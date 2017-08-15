
static inline int hwc_to_chw(float *output, float *input, int height,
                             int width) {
  float *rOutputPlane = &output[0 * width * height];
  float *bOutputPlane = &output[1 * width * height];
  float *gOutputPlane = &output[2 * width * height];
  for (int yy = 0; yy < height; yy++) {
    for (int xx = 0; xx < width; xx++) {
      *rOutputPlane = *input++;
      *bOutputPlane = *input++;
      *gOutputPlane = *input++;
    }
  }
  return 0;
}
