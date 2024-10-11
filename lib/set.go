package lib

/*
#include <stdint.h>
#include <ryzenadj.h>
*/
import "C"

func (ry RyzenAccess) SetStapmLimit(value uint32) (err error) {
	return RyzenAdjErr(C.set_stapm_limit(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetFastLimit(value uint32) (err error) {
	return RyzenAdjErr(C.set_fast_limit(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetSlowLimit(value uint32) (err error) {
	return RyzenAdjErr(C.set_slow_limit(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetSlowTime(value uint32) (err error) {
	return RyzenAdjErr(C.set_slow_time(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetStapmTime(value uint32) (err error) {
	return RyzenAdjErr(C.set_stapm_time(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetTctlTemp(value uint32) (err error) {
	return RyzenAdjErr(C.set_tctl_temp(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetVrmCurrent(value uint32) (err error) {
	return RyzenAdjErr(C.set_vrm_current(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetVrmsocCurrent(value uint32) (err error) {
	return RyzenAdjErr(C.set_vrmsoc_current(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetVrmgfxCurrent(value uint32) (err error) {
	return RyzenAdjErr(C.set_vrmgfx_current(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetVrmcvipCurrent(value uint32) (err error) {
	return RyzenAdjErr(C.set_vrmcvip_current(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetVrmmaxCurrent(value uint32) (err error) {
	return RyzenAdjErr(C.set_vrmmax_current(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetVrmgfxmaxCurrent(value uint32) (err error) {
	return RyzenAdjErr(C.set_vrmgfxmax_current(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetVrmsocmaxCurrent(value uint32) (err error) {
	return RyzenAdjErr(C.set_vrmsocmax_current(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetPsi0Current(value uint32) (err error) {
	return RyzenAdjErr(C.set_psi0_current(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetPsi3cpuCurrent(value uint32) (err error) {
	return RyzenAdjErr(C.set_psi3cpu_current(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetPsi0socCurrent(value uint32) (err error) {
	return RyzenAdjErr(C.set_psi0soc_current(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetPsi3gfxCurrent(value uint32) (err error) {
	return RyzenAdjErr(C.set_psi3gfx_current(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetMaxGfxclkFreq(value uint32) (err error) {
	return RyzenAdjErr(C.set_max_gfxclk_freq(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetMinGfxclkFreq(value uint32) (err error) {
	return RyzenAdjErr(C.set_min_gfxclk_freq(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetMaxSocclkFreq(value uint32) (err error) {
	return RyzenAdjErr(C.set_max_socclk_freq(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetMinSocclkFreq(value uint32) (err error) {
	return RyzenAdjErr(C.set_min_socclk_freq(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetMaxFclkFreq(value uint32) (err error) {
	return RyzenAdjErr(C.set_max_fclk_freq(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetMinFclkFreq(value uint32) (err error) {
	return RyzenAdjErr(C.set_min_fclk_freq(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetMaxVcn(value uint32) (err error) {
	return RyzenAdjErr(C.set_max_vcn(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetMinVcn(value uint32) (err error) {
	return RyzenAdjErr(C.set_min_vcn(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetMaxLclk(value uint32) (err error) {
	return RyzenAdjErr(C.set_max_lclk(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetMinLclk(value uint32) (err error) {
	return RyzenAdjErr(C.set_min_lclk(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetProchotDeassertionRamp(value uint32) (err error) {
	return RyzenAdjErr(C.set_prochot_deassertion_ramp(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetApuSkinTempLimit(value uint32) (err error) {
	return RyzenAdjErr(C.set_apu_skin_temp_limit(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetDgpuSkinTempLimit(value uint32) (err error) {
	return RyzenAdjErr(C.set_dgpu_skin_temp_limit(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetApuSlowLimit(value uint32) (err error) {
	return RyzenAdjErr(C.set_apu_slow_limit(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetSkinTempPowerLimit(value uint32) (err error) {
	return RyzenAdjErr(C.set_skin_temp_power_limit(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetGfxClk(value uint32) (err error) {
	return RyzenAdjErr(C.set_gfx_clk(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetOcClk(value uint32) (err error) {
	return RyzenAdjErr(C.set_oc_clk(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetPerCoreOcClk(value uint32) (err error) {
	return RyzenAdjErr(C.set_per_core_oc_clk(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetOcVolt(value uint32) (err error) {
	return RyzenAdjErr(C.set_oc_volt(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetCoall(value uint32) (err error) {
	return RyzenAdjErr(C.set_coall(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetCoper(value uint32) (err error) {
	return RyzenAdjErr(C.set_coper(ry.access, C.uint(value)))
}
func (ry RyzenAccess) SetCogfx(value uint32) (err error) {
	return RyzenAdjErr(C.set_cogfx(ry.access, C.uint(value)))
}

func (ry RyzenAccess) SetDisableOc() (err error) {
	return RyzenAdjErr(C.set_disable_oc(ry.access))
}
func (ry RyzenAccess) SetEnableOc() (err error) {
	return RyzenAdjErr(C.set_enable_oc(ry.access))
}
func (ry RyzenAccess) SetPowerSaving() (err error) {
	return RyzenAdjErr(C.set_power_saving(ry.access))
}
func (ry RyzenAccess) SetMaxPerformance() (err error) {
	return RyzenAdjErr(C.set_max_performance(ry.access))
}
