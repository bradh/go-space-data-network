// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package OEM

import "strconv"

type timeSystem int8

const (
	/// Greenwich Mean Sidereal Time
	timeSystemGMST timeSystem = 0
	/// Global Positioning System
	timeSystemGPS  timeSystem = 1
	/// Mission Elapsed Time
	timeSystemMET  timeSystem = 2
	/// Mission Relative Time
	timeSystemMRT  timeSystem = 3
	/// Spacecraft Clock (receiver) (requires rules for interpretation in ICD)
	timeSystemSCLK timeSystem = 4
	/// International Atomic Time
	timeSystemTAI  timeSystem = 5
	/// Barycentric Coordinate Time
	timeSystemTCB  timeSystem = 6
	/// Barycentric Dynamical Time
	timeSystemTDB  timeSystem = 7
	/// Geocentric Coordinate Time
	timeSystemTCG  timeSystem = 8
	/// Terrestrial Time
	timeSystemTT   timeSystem = 9
	/// Universal Time
	timeSystemUT1  timeSystem = 10
	/// Coordinated Universal Time
	timeSystemUTC  timeSystem = 11
)

var EnumNamestimeSystem = map[timeSystem]string{
	timeSystemGMST: "GMST",
	timeSystemGPS:  "GPS",
	timeSystemMET:  "MET",
	timeSystemMRT:  "MRT",
	timeSystemSCLK: "SCLK",
	timeSystemTAI:  "TAI",
	timeSystemTCB:  "TCB",
	timeSystemTDB:  "TDB",
	timeSystemTCG:  "TCG",
	timeSystemTT:   "TT",
	timeSystemUT1:  "UT1",
	timeSystemUTC:  "UTC",
}

var EnumValuestimeSystem = map[string]timeSystem{
	"GMST": timeSystemGMST,
	"GPS":  timeSystemGPS,
	"MET":  timeSystemMET,
	"MRT":  timeSystemMRT,
	"SCLK": timeSystemSCLK,
	"TAI":  timeSystemTAI,
	"TCB":  timeSystemTCB,
	"TDB":  timeSystemTDB,
	"TCG":  timeSystemTCG,
	"TT":   timeSystemTT,
	"UT1":  timeSystemUT1,
	"UTC":  timeSystemUTC,
}

func (v timeSystem) String() string {
	if s, ok := EnumNamestimeSystem[v]; ok {
		return s
	}
	return "timeSystem(" + strconv.FormatInt(int64(v), 10) + ")"
}
