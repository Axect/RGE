package RGE

import (
	"fmt"

	"github.com/Axect/csv"
)

// CalcPotential calculate potential
func CalcPotential(mtint, mtfloat, xi int) {
	dir := fmt.Sprintf("Data/Gauge_%d_%d_%d", mtint, mtfloat, xi)
	Temp := csv.Read(dir)
	Data := RConvert(Temp)
	fmt.Println(Data)
}
