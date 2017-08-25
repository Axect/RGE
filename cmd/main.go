package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/Axect/RGE"
	"github.com/Axect/check"
)

const (
	cmdName     = "julia"
	JuliaFolder = "Julia/"
)

func main() {
	// parameter set (choose can be list)
	// 1.Gauge, 2.G(t), 3.Lambda, 4.Potential
	Mt, Xi, choice := os.Args[1], os.Args[2], os.Args[3:]
	mt, err1 := strconv.ParseFloat(Mt, 64)
	xi, err2 := strconv.ParseFloat(Xi, 64)

	if err1 != nil || err2 != nil {
		log.Fatal("Can't convert string to float64. Plz input proper value")
	}
	Welcome()
	// Running and receive mtint, mtfloat, xi
	fmt.Println("Data Processing...")
	MX := RGE.RGERunning(mt, xi)
	mtint := MX[0]
	mtfloat := MX[1]

	// Handle Plot with Julia
	fmt.Println("-------------------------")
	fmt.Println("Welcome to RGE Plot")
	fmt.Println("-------------------------")
	fmt.Println()

	// Cmd Settings
	cmdBody := strings.Fields(fmt.Sprintf("%d %d %d", mtint, mtfloat, int(xi+0.4)))
	var subDir string
	var cmdDir []string

	fmt.Println("Input Parameter: ", cmdBody)

	// Parallel Setting
	PlotList := make(chan string)

	// Gauge Plot
	if check.Contains("1", choice) {
		subDir = "Gauge_plot.jl"
		cmdDir = append(cmdDir, subDir)
		fmt.Println("Draw Gauge Plot...")
	}

	// G(t) Plot
	if check.Contains("2", choice) {
		subDir = "G_plot.jl"
		cmdDir = append(cmdDir, subDir)
		fmt.Println("Draw G(t) Plot...")
	}

	// Lambda Plot
	if check.Contains("3", choice) {
		subDir = "Lambda_plot.jl"
		cmdDir = append(cmdDir, subDir)
		fmt.Println("Draw Lambda Plot...")
	}

	for _, dir := range cmdDir {
		go Routine(JuliaFolder, dir, cmdBody, PlotList)
	}

	for plot := range PlotList {
		fmt.Println(plot)
	}
}

func Routine(JuliaFolder, subdir string, cmdBody []string, c chan string) {
	cmdArgs := append([]string{JuliaFolder + subdir}, cmdBody...)

	var (
		cmdOut []byte
		err    error
	)

	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		log.Fatal("Can't execute commands")
	}
	comp := string(cmdOut)
	fmt.Println(comp)
	Assign(c, (subdir + "complete!"))
}

func Assign(c chan string, text string) {
	c <- text
}

func Welcome() {
	// Running with Go
	fmt.Println("--------------------------------")
	fmt.Println("Welcome to RGE.go")
	fmt.Println("--------------------------------")
	fmt.Println()
}
