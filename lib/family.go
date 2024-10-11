//go:generate stringer -type=RyzenFamily
package lib

type RyzenFamily int32

const (
	WAIT_FOR_LOAD RyzenFamily = iota - 2
	UNKNOWN
	RAVEN
	PICASSO
	RENOIR
	CEZANNE
	DALI
	LUCIENNE
	VANGOGH
	REMBRANDT
	MENDOCINO
	PHOENIX
	HAWKPOINT
	STRIXPOINT
	END
)
