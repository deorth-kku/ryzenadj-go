// Code generated by "stringer -type=RyzenAdjErr"; DO NOT EDIT.

package lib

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ADJ_ERR_FAM_UNSUPPORTED - -1]
	_ = x[ADJ_ERR_SMU_TIMEOUT - -2]
	_ = x[ADJ_ERR_SMU_UNSUPPORTED - -3]
	_ = x[ADJ_ERR_SMU_REJECTED - -4]
	_ = x[ADJ_ERR_MEMORY_ACCESS - -5]
}

const _RyzenAdjErr_name = "ADJ_ERR_MEMORY_ACCESSADJ_ERR_SMU_REJECTEDADJ_ERR_SMU_UNSUPPORTEDADJ_ERR_SMU_TIMEOUTADJ_ERR_FAM_UNSUPPORTED"

var _RyzenAdjErr_index = [...]uint8{0, 21, 41, 64, 83, 106}

func (i RyzenAdjErr) String() string {
	i -= -5
	if i < 0 || i >= RyzenAdjErr(len(_RyzenAdjErr_index)-1) {
		return "RyzenAdjErr(" + strconv.FormatInt(int64(i+-5), 10) + ")"
	}
	return _RyzenAdjErr_name[_RyzenAdjErr_index[i]:_RyzenAdjErr_index[i+1]]
}
