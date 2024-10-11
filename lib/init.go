package lib

/*
#cgo LDFLAGS: -lryzenadj
#include <stdint.h>
#include <ryzenadj.h>
*/
import "C"

type RyzenAccess struct {
	access C.ryzen_access
}

func NewRyzenAccess() (acc RyzenAccess, err error) {
	acc.access = C.init_ryzenadj()
	if acc.access == nil {
		return acc, ErrInitFailed
	}
	errorno := C.init_table(acc.access)
	if errorno != 0 {
		err = RyzenAdjErr(errorno)
	}
	return
}

func (acc RyzenAccess) Cleanup() {
	C.cleanup_ryzenadj(acc.access)
}

func (acc RyzenAccess) Refresh() error {
	return RyzenAdjErr(C.refresh_table(acc.access))
}
