// Code generated by "stringer -type=RequestReturnValue"; DO NOT EDIT.

package comm

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OK-0]
}

const _RequestReturnValue_name = "OK"

var _RequestReturnValue_index = [...]uint8{0, 2}

func (i RequestReturnValue) String() string {
	if i < 0 || i >= RequestReturnValue(len(_RequestReturnValue_index)-1) {
		return "RequestReturnValue(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RequestReturnValue_name[_RequestReturnValue_index[i]:_RequestReturnValue_index[i+1]]
}
