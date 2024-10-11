package lib

/*
#include <stdint.h>
#include <ryzenadj.h>
*/
import "C"

func (ry RyzenAccess) GetCpuFamily() RyzenFamily {
	return RyzenFamily(C.get_cpu_family(ry.access))
}
func (ry RyzenAccess) GetBiosIfVer() int32 {
	return int32(C.get_bios_if_ver(ry.access))
}

func (ry RyzenAccess) GetStapmLimit() float32 {
	return float32(C.get_stapm_limit(ry.access))
}
func (ry RyzenAccess) GetStapmValue() float32 {
	return float32(C.get_stapm_value(ry.access))
}
func (ry RyzenAccess) GetFastLimit() float32 {
	return float32(C.get_fast_limit(ry.access))
}
func (ry RyzenAccess) GetFastValue() float32 {
	return float32(C.get_fast_value(ry.access))
}
func (ry RyzenAccess) GetSlowLimit() float32 {
	return float32(C.get_slow_limit(ry.access))
}
func (ry RyzenAccess) GetSlowValue() float32 {
	return float32(C.get_slow_value(ry.access))
}
func (ry RyzenAccess) GetApuSlowLimit() float32 {
	return float32(C.get_apu_slow_limit(ry.access))
}
func (ry RyzenAccess) GetApuSlowValue() float32 {
	return float32(C.get_apu_slow_value(ry.access))
}
func (ry RyzenAccess) GetVrmCurrent() float32 {
	return float32(C.get_vrm_current(ry.access))
}
func (ry RyzenAccess) GetVrmCurrentValue() float32 {
	return float32(C.get_vrm_current_value(ry.access))
}
func (ry RyzenAccess) GetVrmsocCurrent() float32 {
	return float32(C.get_vrmsoc_current(ry.access))
}
func (ry RyzenAccess) GetVrmsocCurrentValue() float32 {
	return float32(C.get_vrmsoc_current_value(ry.access))
}
func (ry RyzenAccess) GetVrmmaxCurrent() float32 {
	return float32(C.get_vrmmax_current(ry.access))
}
func (ry RyzenAccess) GetVrmmaxCurrentValue() float32 {
	return float32(C.get_vrmmax_current_value(ry.access))
}
func (ry RyzenAccess) GetVrmsocmaxCurrent() float32 {
	return float32(C.get_vrmsocmax_current(ry.access))
}
func (ry RyzenAccess) GetVrmsocmaxCurrentValue() float32 {
	return float32(C.get_vrmsocmax_current_value(ry.access))
}
func (ry RyzenAccess) GetTctlTemp() float32 {
	return float32(C.get_tctl_temp(ry.access))
}
func (ry RyzenAccess) GetTctlTempValue() float32 {
	return float32(C.get_tctl_temp_value(ry.access))
}
func (ry RyzenAccess) GetApuSkinTempLimit() float32 {
	return float32(C.get_apu_skin_temp_limit(ry.access))
}
func (ry RyzenAccess) GetApuSkinTempValue() float32 {
	return float32(C.get_apu_skin_temp_value(ry.access))
}
func (ry RyzenAccess) GetDgpuSkinTempLimit() float32 {
	return float32(C.get_dgpu_skin_temp_limit(ry.access))
}
func (ry RyzenAccess) GetDgpuSkinTempValue() float32 {
	return float32(C.get_dgpu_skin_temp_value(ry.access))
}
func (ry RyzenAccess) GetPsi0Current() float32 {
	return float32(C.get_psi0_current(ry.access))
}
func (ry RyzenAccess) GetPsi0socCurrent() float32 {
	return float32(C.get_psi0soc_current(ry.access))
}
func (ry RyzenAccess) GetStapmTime() float32 {
	return float32(C.get_stapm_time(ry.access))
}
func (ry RyzenAccess) GetSlowTime() float32 {
	return float32(C.get_slow_time(ry.access))
}
func (ry RyzenAccess) GetCclkSetpoint() float32 {
	return float32(C.get_cclk_setpoint(ry.access))
}
func (ry RyzenAccess) GetCclkBusyValue() float32 {
	return float32(C.get_cclk_busy_value(ry.access))
}
func (ry RyzenAccess) GetL3Clk() float32 {
	return float32(C.get_l3_clk(ry.access))
}
func (ry RyzenAccess) GetL3Logic() float32 {
	return float32(C.get_l3_logic(ry.access))
}
func (ry RyzenAccess) GetL3Vddm() float32 {
	return float32(C.get_l3_vddm(ry.access))
}
func (ry RyzenAccess) GetL3Temp() float32 {
	return float32(C.get_l3_temp(ry.access))
}
func (ry RyzenAccess) GetGfxClk() float32 {
	return float32(C.get_gfx_clk(ry.access))
}
func (ry RyzenAccess) GetGfxTemp() float32 {
	return float32(C.get_gfx_temp(ry.access))
}
func (ry RyzenAccess) GetGfxVolt() float32 {
	return float32(C.get_gfx_volt(ry.access))
}
func (ry RyzenAccess) GetMemClk() float32 {
	return float32(C.get_mem_clk(ry.access))
}
func (ry RyzenAccess) GetFclk() float32 {
	return float32(C.get_fclk(ry.access))
}
func (ry RyzenAccess) GetSocPower() float32 {
	return float32(C.get_soc_power(ry.access))
}
func (ry RyzenAccess) GetSocVolt() float32 {
	return float32(C.get_soc_volt(ry.access))
}
func (ry RyzenAccess) GetSocketPower() float32 {
	return float32(C.get_socket_power(ry.access))
}

func (ry RyzenAccess) GetCoreClk(core uint32) float32 {
	return float32(C.get_core_clk(ry.access, C.uint(core)))
}
func (ry RyzenAccess) GetCoreVolt(core uint32) float32 {
	return float32(C.get_core_volt(ry.access, C.uint(core)))
}
func (ry RyzenAccess) GetCorePower(core uint32) float32 {
	return float32(C.get_core_power(ry.access, C.uint(core)))
}
func (ry RyzenAccess) GetCoreTemp(core uint32) float32 {
	return float32(C.get_core_temp(ry.access, C.uint(core)))
}
