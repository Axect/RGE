package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/Axect/RGE"
	"github.com/Axect/check"
)

func main() {
	// Input mt, xi what you want to draw
	mt, xi := 170.85, 50.

	Welcome()
	DrawPlot(mt, xi)
}

func DrawPlot(mt, xi float64) {
	// Running and receive mtint, mtfloat, xi
	MX := RGE.RGERunning(mt, xi)
	mtint := MX[0]
	mtfloat := MX[1]

	// Handle Plot with Julia
	fmt.Println("-------------------------")
	fmt.Println("Welcome to RGE Plot")
	fmt.Println("-------------------------")
	fmt.Println()
	fmt.Println("Choose what you want to draw")
	fmt.Println("1.Gauge, 2.G(t), 3.Lambda, 4.Potential")

	var choose string
	_, err := fmt.Scanf("Number: %s", &choose)

	if err != nil {
		log.Fatal("Can't Scan")
	}
	choice := strings.Fields(choose)

	if check.Contains("1", choice) {
		command := fmt.Sprintf("julia Julia/Gauge_plot.jl %d %d %d", mtint, mtfloat, xi)
		fmt.Println("Draw Gauge Plot...")
		exec.Command("sh", "-c", command).Run()
		fmt.Println("Complete!")
	}
	if check.Contains("2", choice) {
		command := fmt.Sprintf("julia Julia/G_plot.jl %d %d %d", mtint, mtfloat, xi)
		fmt.Println("Draw G(t) Plot...")
		exec.Command("sh", "-c", command).Run()
		fmt.Println("Complete!")
	}
	if check.Contains("3", choice) {
		command := fmt.Sprintf("julia Julia/Lambda_plot.jl %d %d %d", mtint, mtfloat, xi)
		fmt.Println("Draw Lambda Plot...")
		exec.Command("sh", "-c", command).Run()
		fmt.Println("Complete!")
	}

}

func Welcome() {
	// Running with Go
	fmt.Println("--------------------------------")
	fmt.Println("Welcome to RGE.go")
	fmt.Println("--------------------------------")
	fmt.Println()
}
