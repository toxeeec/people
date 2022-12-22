// Code generated by "stringer -type=ErrorKind"; DO NOT EDIT.

package people

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ValidationError-0]
	_ = x[AuthError-1]
	_ = x[NotFoundError-2]
	_ = x[ConflictError-3]
	_ = x[ResourceError-4]
}

const _ErrorKind_name = "ValidationErrorAuthErrorNotFoundErrorConflictErrorResourceError"

var _ErrorKind_index = [...]uint8{0, 15, 24, 37, 50, 63}

func (i ErrorKind) String() string {
	if i >= ErrorKind(len(_ErrorKind_index)-1) {
		return "ErrorKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrorKind_name[_ErrorKind_index[i]:_ErrorKind_index[i+1]]
}
