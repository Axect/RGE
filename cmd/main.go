package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"

	"github.com/Axect/RGE"
	"github.com/Axect/check"
)

const (
	ver         = "0.0.1"
	author      = "Axect"
	page        = "https://github.com/Axect/RGE"
	Julia       = "julia"
	JuliaFolder = "Julia/"
)

var wg sync.WaitGroup

func main() {
	// parameter set (choose can be list)
	// 1.Gauge, 2.G(t), 3.Lambda, 4.Potential
	if err := exec.Command("clear", "").Run(); err != nil {
		log.Fatal("Can't clear")
	}
	mt, xi, choice := Welcome()

	// Running and receive mtint, mtfloat, xi
	fmt.Println("-----------------------------------")
	fmt.Println("  Data Processing...  ")
	fmt.Println("-----------------------------------")
	fmt.Println()
	MX := RGE.Running(mt, xi)
	mtint := MX[0]
	mtfloat := MX[1]
	fmt.Println("Calculation Complete!")
	fmt.Println()

	// Handle Plot with Julia
	fmt.Println("-----------------------------------")
	fmt.Println("  Plotting...  ")
	fmt.Println("-----------------------------------")
	fmt.Println()

	// Cmd Settings
	cmdBody := strings.Fields(fmt.Sprintf("%d %d %d", mtint, mtfloat, int(xi+0.4)))
	var subDir string
	var cmdDir []string

	fmt.Println("Input Parameter: ", cmdBody)

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

	// Potential Plot
	if check.Contains("4", choice) {
		subDir = "Potential_plot.jl"
		cmdDir = append(cmdDir, subDir)
		fmt.Println("Draw Potential Plot...")
	}

	for _, dir := range cmdDir {
		wg.Add(1)
		go Routine(JuliaFolder, dir, cmdBody)
	}
	wg.Wait()

	fmt.Println("All Process Finished")
}

// Routine runs julia for plotting by parallel
func Routine(JuliaFolder, subdir string, cmdBody []string) {
	defer wg.Done()

	cmdArgs := append([]string{JuliaFolder + subdir}, cmdBody...)

	var (
		cmdOut []byte
		err    error
	)

	if cmdOut, err = exec.Command(Julia, cmdArgs...).Output(); err != nil {
		log.Fatal("Can't execute commands")
	}
	comp := string(cmdOut)
	fmt.Println(comp)
	fmt.Println(subdir, " Complete!")
	fmt.Println()
	return
}

// Welcome handle IO, Print
func Welcome() (float64, float64, []string) {
	// Running with Go
	fmt.Println("-----------------------------------")
	fmt.Println("  RGE Solver  ")
	fmt.Printf("  ver %s   \n", ver)
	fmt.Printf("  author %s  \n", author)
	fmt.Printf("  page %s  \n", page)
	fmt.Println("-----------------------------------")
	fmt.Println()
	fmt.Println("__________  ___________________")
	fmt.Println("\\______   \\/  _____/\\_   _____/")
	fmt.Println(" |       _/   \\  ___ |    __)_")
	fmt.Println(" |    |   \\    \\_\\  \\|        \\")
	fmt.Println(" |____|_  /\\______  /_______  /")
	fmt.Println("        \\/        \\/        \\/ ")
	fmt.Println()
	fmt.Println("Input parameters: ")
	fmt.Println("ex) 170.85 50")
	var Mt, Xi float64
	choice := make([]string, 4, 4)
	fmt.Scanln(&Mt, &Xi)
	fmt.Println("Select Plots: 1.Gauge, 2.G(t), 3.Lambda, 4.Potential")
	for i := range choice {
		fmt.Scan(&choice[i])
	}
	fmt.Println()
	return Mt, Xi, choice
}
