//go:generate stringer -type=RyzenFamily
package lib

type RyzenFamily int32

const (
	WAIT_FOR_LOAD RyzenFamily = iota - 2
	FAM_UNKNOWN
	FAM_RAVEN
	FAM_PICASSO
	FAM_RENOIR
	FAM_CEZANNE
	FAM_DALI
	FAM_LUCIENNE
	FAM_VANGOGH
	FAM_REMBRANDT
	FAM_MENDOCINO
	FAM_PHOENIX
	FAM_HAWKPOINT
	FAM_STRIXPOINT
	FAM_END
)