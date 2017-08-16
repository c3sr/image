
CXX = clang++
AMD64CXXFLAGS= -O3 -std=c++1z -S -mavx2 -mfma -masm=intel -mno-red-zone -mstackrealign -mllvm -inline-threshold=1000 -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti -fno-jump-tables

SRCDIR=cbits
TMPDIR=tmp

all: hwc_to_chw_amd64.s

hwc_to_chw.cpp.s: $(SRCDIR)/hwc_to_chw.cpp
	mkdir -p tmp
	$(CXX) $(AMD64CXXFLAGS) $^ -o tmp/$@

hwc_to_chw_amd64.s: hwc_to_chw.cpp.s
	c2goasm -a -c -s -f tmp/$^ $@

clean:
	rm -fr tmp *.s
