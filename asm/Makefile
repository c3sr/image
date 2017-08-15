
CXX = clang++
CXXFLAGS= -O3 -std=c++11 -S  -mavx -mfma -masm=intel -mno-red-zone -mstackrealign -mllvm -inline-threshold=1000 -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti -fno-jump-tables

SRCDIR=cbits

all: hwc_to_chw.s

hwc_to_chw.cpp.s: $(SRCDIR)/hwc_to_chw.cpp
	$(CXX) $(CXXFLAGS) $^ -o $@

hwc_to_chw.s: hwc_to_chw.cpp.s
	c2goasm -a -s -c -f $^ $@

clean:
	rm -fr *.s
