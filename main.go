package main

import (
	"fmt"

	"github.com/deorth-kku/ryzenadj-go/lib"
)

func main() {
	ry, err := lib.NewRyzenAccess()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ry.Cleanup()
	err = ry.Refresh()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("CPU Family: %s\n", ry.GetCpuFamily())
	fmt.Printf("SMU BIOS Interface Version: %d\n", ry.GetBiosIfVer())
}
