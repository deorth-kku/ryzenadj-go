//go:generate stringer -type=RyzenAdjErr
package lib

import "C"
import (
	"errors"
	"fmt"
)

type RyzenAdjErr int32

const (
	ADJ_ERR_FAM_UNSUPPORTED RyzenAdjErr = -1
	ADJ_ERR_SMU_TIMEOUT     RyzenAdjErr = -2
	ADJ_ERR_SMU_UNSUPPORTED RyzenAdjErr = -3
	ADJ_ERR_SMU_REJECTED    RyzenAdjErr = -4
	ADJ_ERR_MEMORY_ACCESS   RyzenAdjErr = -5
)

var ErrInitFailed = errors.New("unable to init ryzenadj, check stdout for more information")

func (rae RyzenAdjErr) Error() (str string) {
	return fmt.Sprintf("ryzenadj return with error %s", rae.String())
}

func NewRyzenAdjErr(errorno C.int) error {
	if errorno == 0 {
		return nil
	}
	return RyzenAdjErr(errorno)
}
