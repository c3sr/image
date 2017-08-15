//go:noescape
func _hwc_to_chw(output, input unsafe.Pointer, height, width int) int

func hwc_to_chw(output , input float32[], width, height int ) {

	_hwc_to_chw(output, input, width, height)
}
