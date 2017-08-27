


#define ROUND_DOWN(x, s) ((x) & ~((s)-1))

extern void hwc2cwh_(uniform int len,
                     uniform int lda,
                     uniform float mean[3],
                     uniform float input[50176], uniform float result[50176]) {
    uniform int stride = ROUND_DOWN(len, 4);
    for (uniform int i = 0; i < len ; i += stride) {
        float x1, y1, z1;
        float x2, y2, z2;
        float x3, y3, z3;
        float x4, y4, z4;


        aos_to_soa3(&input[i + 0], &x1, &y1, &z1);
        aos_to_soa3(&input[i + 1], &x2, &y2, &z2);
        aos_to_soa3(&input[i + 2], &x3, &y3, &z3);
        aos_to_soa3(&input[i + 3], &x4, &y4, &z4);

        float<4> xs = {x1, x2, x3, x4};
        float<4> ys = {y1, y2, y3, y4};
        float<4> zs = {z1, z2, z3, z4};

        xs -= mean[0];
        ys -= mean[1];
        zs -= mean[2];

        soa_to_aos4(xs.x, xs.y, xs.z, xs.w, &result[0*lda + i]);
        soa_to_aos4(y1, y2, y3, y4, &result[1*lda + i]);
        soa_to_aos4(z1, z2, z3, z4, &result[2*lda + i]);
    }
}

/*
static inline void hwc2cwh_(uniform int width, uniform int height,
                    uniform float mean[3],
                    uniform float input[], uniform float result[]) {
    uniform int stride = ROUND_DOWN(width, 4);
    for (uniform int j = 0; j < height; ++j) {
        uniform int offset = j * height;
        for (uniform int i = 0; i < width ; i += stride) {
            float x1, y1, z1;
            float x2, y2, z2;
            float x3, y3, z3;
            float x4, y4, z4;


            aos_to_soa3(&input[offset + i + 0], &x1, &y1, &z1);
            aos_to_soa3(&input[offset + i + 1], &x2, &y2, &z2);
            aos_to_soa3(&input[offset + i + 2], &x3, &y3, &z3);
            aos_to_soa3(&input[offset + i + 3], &x4, &y4, &z4);

            float<4> xs = {x1, x2, x3, x4};
            float<4> ys = {y1, y2, y3, y4};
            float<4> zs = {z1, z2, z3, z4};

            xs -= mean[0];
            ys -= mean[1];
            zs -= mean[2];

            soa_to_aos4(xs.x, xs.y, xs.z, xs.w, &result[0*(width * height) + offset + i]);
            soa_to_aos4(y1, y2, y3, y4, &result[1*(width * height) + offset + i]);
            soa_to_aos4(z1, z2, z3, z4, &result[2*(width * height) + offset + i]);
        }
    }
}

export void hwc2cwh(uniform int width, uniform int height,
                    uniform float mean[3],
                    uniform float input[], uniform float result[]) {
    hwc2cwh_(width, height, mean, input, result);
}
export void hwc2cwh2(uniform int width, uniform int height,
                    uniform float input[], uniform float result[]) {
    uniform float mean[3] = {0, 0, 0};
    hwc2cwh_(width, height, mean, input, result);
}
*/
