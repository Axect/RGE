package RGE

import (
	"fmt"
	"log"
	"strconv"
)

// =============================================================================
// You can define simple method in this file.
// =============================================================================

// Copy protects RGE value which is in Container
func (R RGE) Copy() RGE {
	var NR RGE

	NR.t = R.t
	NR.lH = R.lH
	NR.yt = R.yt
	NR.g1 = R.g1
	NR.g2 = R.g2
	NR.g3 = R.g3
	NR.phi = R.phi
	NR.G = R.G

	return NR
}

// Convert supports csv.Write
func Convert(List []float64) []string {
	Temp := make([]string, len(List), len(List))
	for i := range List {
		Temp[i] = fmt.Sprintf("%v", List[i])
	}
	return Temp
}

// RConvert convert reading csv file to []RGE
func RConvert(Temp [][]string) []RGE {
	Array := make([]RGE, len(Temp), len(Temp))
	for i := range Array {
		var err error
		Array[i].t, err = strconv.ParseFloat(Temp[i][0], 64)
		Array[i].lH, err = strconv.ParseFloat(Temp[i][1], 64)
		Array[i].yt, err = strconv.ParseFloat(Temp[i][2], 64)
		Array[i].g1, err = strconv.ParseFloat(Temp[i][3], 64)
		Array[i].g2, err = strconv.ParseFloat(Temp[i][4], 64)
		Array[i].g3, err = strconv.ParseFloat(Temp[i][5], 64)
		Array[i].G, err = strconv.ParseFloat(Temp[i][6], 64)
		if err != nil {
			log.Fatal(err)
		}
	}
	return Array
}
