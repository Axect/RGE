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
	cmdName = "julia"
	cmdDir1 = "Julia/"
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
	fmt.Println(mtint, mtfloat, xi)

	// Handle Plot with Julia
	fmt.Println("-------------------------")
	fmt.Println("Welcome to RGE Plot")
	fmt.Println("-------------------------")
	fmt.Println()

	// Cmd Settings
	cmdBody := strings.Fields(fmt.Sprintf("%d %d %d", mtint, mtfloat, int(xi+0.4)))
	var cmdDir2 string
	var cmdDir []string

	// Parallel Setting
	done := make(chan bool)

	fmt.Println("Input Parameter: ", cmdBody)
	// Gauge Plot
	if check.Contains("1", choice) {
		cmdDir2 = "Gauge_plot.jl"
		cmdDir = append(cmdDir, cmdDir2)
		fmt.Println("Draw Gauge Plot...")
	}

	// G(t) Plot
	if check.Contains("2", choice) {
		cmdDir2 = "G_plot.jl"
		cmdDir = append(cmdDir, cmdDir2)
		fmt.Println("Draw G(t) Plot...")
	}

	// Lambda Plot
	if check.Contains("3", choice) {
		cmdDir2 = "Lambda_plot.jl"
		cmdDir = append(cmdDir, cmdDir2)
		fmt.Println("Draw Lambda Plot...")
	}

	for _, dir := range cmdDir {
		go Routine(cmdDir1, dir, cmdBody, done)
	}
	<-done
}

func Routine(cmdDir1, dir string, cmdBody []string, c chan bool) {
	defer Assign(c)
	cmdArgs := append([]string{cmdDir1 + dir}, cmdBody...)

	var (
		cmdOut []byte
		err    error
	)

	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		log.Fatal("Can't execute commands")
	}
	comp := string(cmdOut)
	fmt.Println(comp)
	fmt.Println("Complete!")
}

func Assign(c chan bool) {
	c <- true
}

func Welcome() {
	// Running with Go
	fmt.Println("--------------------------------")
	fmt.Println("Welcome to RGE.go")
	fmt.Println("--------------------------------")
	fmt.Println()
}
