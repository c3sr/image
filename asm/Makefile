
CXX = clang++
AMD64CXXFLAGS= -O3 -Dsincosf=__sincosf -Dsincos=__sincos -std=c++1z -S \
						   -mavx2 -mfma -masm=intel -ffast-math -mno-red-zone -mstackrealign \
							 -mllvm -inline-threshold=1000 -fno-asynchronous-unwind-tables \
							 -fno-exceptions -fno-rtti -fno-jump-tables

SRCDIR=cbits
TMPDIR=tmp

all: hwc_to_chw_amd64.s

hwc_to_chw.cpp: $(SRCDIR)/hwc_to_chw.ispc
	ispc --target=generic-4 --emit-c++ --c++-include-file=intrinsics/sse4.h --opt=fast-math $^ -o cbits/$@

hwc_to_chw.cpp.s: # hwc_to_chw.cpp
	mkdir -p tmp
	$(CXX) $(AMD64CXXFLAGS) -I $(SRCDIR) $(SRCDIR)/hwc_to_chw.cpp -o tmp/$@

hwc_to_chw_amd64.s: hwc_to_chw.cpp.s
	c2goasm -a -c -s -f tmp/$^ $@


clean:
	rm -fr tmp *.s
