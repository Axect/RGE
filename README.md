# RGE

<font size=2> : Golang Package which solves Renormalization Group Equation

## Installation

```bash
go get -u github.com/Axect/RGE
```

## Prerequisites

* Golang (ver >= 1.8.3)
* Julia (ver >= 0.6.0)

## How to use

1. Set global constants in ```const.go```  
	
	```Go
	// -----------------------
	//	Declare Constants
	// -----------------------

	const (
		Mp       = 1.221 * 1E+19 // Plank Mass
		MpR      = 2.4 * 1E+18   // Reduced Planck Mass
		MW       = 80.385        // Mass of W
		MZ       = 91.1876       // Mass of Z
		MH       = 125.09        // Mass of Higgs
		alphasMZ = 0.1182        // alphas(MZ)

		// Running Constants (Precision & Number of Lists)
		h    = 1E-04      // precision
		Step = 1E+04 * 44 // Number of lists
	)
	```

2. Set RGE variables in ```var.go```

	```Go
	//----------------------------
	// Declare your rge variables
	//----------------------------

	type RGE struct {
		t   float64
		lH  float64
		yt  float64
		g1  float64
		g2  float64
		g3  float64
		phi float64
		G   float64
	}
	```
	You can change elements name (t->x, lH->omega etc..). But do not change type name.
	
3. Enter your initial value of RGE variables to ```init.go```

	```Go
	import "math"

	// Initialize RGE variables by some constants
	func Initialize(mt float64) RGE {
		var R RGE
		R.t = 0
		R.lH = 0.12604 + 0.00206*(MH-125.15) - 0.00004*(mt-173.34)
		R.yt = (0.93690 + 0.00556*(mt-173.34) - 0.00003*(MH-125.15) - 0.00042*(alphasMZ-0.1184)/0.0007)
		R.g1 = 0.35830 + 0.00011*(mt-173.34) - 0.00020*(MW-80.384)/0.014
		R.g2 = 0.64779 + 0.00004*(mt-173.34) + 0.00011*(MW-80.384)/0.014
		R.g3 = 1.1666 + 0.00314*(alphasMZ-0.1184)/0.007 - 0.00046*(mt-173.34)
		R.phi = math.Sqrt(2.) / R.yt * mt * math.Exp(R.t)
		R.G = 1

		return R
	}
	```
	You can customize input variable - It will be declared in ```main.go```.
	And using your declared constants in ```const.go```.
	
4. Enter beta function formula to ```beta.go```
	1. Declare all beta function (1-loop, 2-loop and total(1-loop + 2-loop))
	
		```Go
		//-------------------------
		// Declare Beta Functions
		//-------------------------
		
		// Beta Function: You can declare all of beta function list here
		type Beta struct {
			// 1-loop order
			B1lH   float64
			B1yt   float64
			B1g1   float64
			B1g2   float64
			B1g3   float64
			gamma1 float64
		
			// 2-loop order
			B2lH   float64
			B2yt   float64
			B2g1   float64
			B2g2   float64
			B2g3   float64
			gamma2 float64
		
			// Total
			BlH   float64
			Byt   float64
			Bg1   float64
			Bg2   float64
			Bg3   float64
			gamma float64
		}
		```
	2. Enter the formulas
	
		```Go
		// InputFormula to Beta functions
		func (B *Beta) InputFormula(R RGE, mt, xi float64) {
			// Necessary Variables
			hg := math.Sqrt(2.) / R.yt * mt * math.Exp(R.t)
			sh := (1. + xi*math.Pow(hg, 2)/math.Pow(MpR, 2)) / (1. + (1.+6.*xi)*xi*math.Pow(hg, 2)/math.Pow(MpR, 2))
	
			// 1-loop Beta Function
			B.B1lH = 6.*(1.+3.*math.Pow(sh, 2))*math.Pow(R.lH, 2) + 12.*R.lH*math.Pow(R.yt, 2) - 6.*math.Pow(R.yt, 4) - 3.*R.lH*(3.*math.Pow(R.g2, 2)+math.Pow(R.g1, 2)) + 3./8*(2.*math.Pow(R.g2, 4)	+math.Pow((math.Pow(R.g1, 2)+math.Pow(R.g2, 2)), 2))
			B.B1yt = R.yt * ((23./6+2./3*sh)*math.Pow(R.yt, 2) -(8.*math.Pow(R.g3, 2) + 9./4*math.Pow(R.g2, 2) + 17./12*math.Pow(R.g1, 2)))
			B.B1g1 = (81. + sh) / 12 * math.Pow(R.g1, 3)
			B.B1g2 = (sh - 39.) / 12 * math.Pow(R.g2, 3)
			B.B1g3 = -7. * math.Pow(R.g3, 3)
			B.gamma1 = -1. / (16. * math.Pow(math.Pi, 2))(9./4*math.Pow(R.g2,2) + 3./4*math.Pow(R.g1, 2) - 3.*math.Pow(R.yt, 2))

			// 2-loop Beta function
			B.B2lH = 1./48*((912.+3.*sh)*math.Pow(R.g2, 6)-(290.-sh)*math.Pow(R.g1, 2)*math.Pow(R.g2, 4)-(560.-sh)*math.Pow(R.g1, 4)*math.Pow(R.g2, 2)-(380.-sh)*math.Pow(R.g1, 6)) + (38.-8*sh)*math.Pow(R.yt, 6) - math.Pow(R.yt, 4)*(8./3*math.Pow(R.g1, 2)+32.*math.Pow(R.g3, 2)+(12.-117.*sh+108.*math.Pow(sh, 2))*R.lH) + R.lH*(-1./8*(181.+54.*sh-162.*math.Pow(sh, 2))*math.Pow(R.g2, 4)+1./4*(3.-18.*sh+54.*math.Pow(sh, 2))*math.Pow(R.g1, 2)*math.Pow(R.g2, 2)+1./24*(90.+377.*sh+162.*math.Pow(sh, 2))*math.Pow(R.g1, 4)+(27.+54.*sh+27.*math.Pow(sh, 2))*math.Pow(R.g2, 2)*R.lH+(9.+18.*sh+9*math.Pow(sh, 2))*math.Pow(R.g1, 2)*R.lH-(48.+288.*sh-324.*math.Pow(sh, 2)+624.*math.Pow(sh, 3)-324.*math.Pow(sh, 4))*math.Pow(R.lH, 2)) + math.Pow(R.yt, 2)*(-9./4*math.Pow(R.g2, 4)+21./2*math.Pow(R.g1, 2)*math.Pow(R.g2, 2)-19./4*math.Pow(R.g1, 4)+R.lH*(45./2*math.Pow(R.g2, 2)+85./6*math.Pow(R.g1, 2)+80.*math.Pow(R.g3, 2)-(36.+108.*math.Pow(sh, 2))*R.lH))
			B.B2yt = R.yt * (-23./4*math.Pow(R.g2, 4) - 3./4*math.Pow(R.g1, 2)*math.Pow(R.g2, 2) + 1187./216*math.Pow(R.g1, 4) + 9.*math.Pow(R.g2, 2)*math.Pow(R.g3, 2) + 19./9*math.Pow(R.g1, 2)*math.Pow(R.g3, 2) - 108.*math.Pow(R.g3, 4) + (225./16*math.Pow(R.g2, 2)+131./16*math.Pow(R.g1, 2)+36.*math.Pow(R.g3, 2))*sh*math.Pow(R.yt, 2) + 6.*(-2.*math.Pow(sh, 2)*math.Pow(R.yt, 4)-2.*math.Pow(sh, 3)*math.Pow(R.yt, 2)*R.lH+math.Pow(sh, 2)*math.Pow(R.lH, 2)))
			B.B2g1 = 199./18*math.Pow(R.g1, 5) + 9./2*math.Pow(R.g1, 3)*math.Pow(R.g2, 2) + 44./3*math.Pow(R.g1, 3)*math.Pow(R.g3, 2) - 17./6*sh*math.Pow(R.g1, 3)*math.Pow(R.yt, 2)
			B.B2g2 = 3./2*math.Pow(R.g1, 2)*math.Pow(R.g2, 3) + 35./6*math.Pow(R.g2, 5) + 12.*math.Pow(R.g2, 3)*math.Pow(R.g3, 2) - 3./2*sh*math.Pow(R.g2, 3)*math.Pow(R.yt, 2)
			B.B2g3 = 11./6*math.Pow(R.g1, 2)*math.Pow(R.g3, 3) + 9./2*math.Pow(R.g2, 2)*math.Pow(R.g3, 3) - 26.*math.Pow(R.g3, 5) - 2.*sh*math.Pow(R.g3, 3)*math.Pow(R.yt, 2)
			B.gamma2 = -1. / (math.Pow(16*math.Pow(math.Pi, 2), 2)) * (271./32*math.Pow(R.g2, 4) - 9./16*math.Pow(R.g1, 2)*math.Pow(R.g2, 2) - 431./96*sh*math.Pow(R.g1, 4) - 5./2*(9./4*math.Pow(R.g2, 2)+17./12*math.Pow(R.g1, 2)+8*math.Pow(R.g3, 2))*math.Pow(R.yt, 2) + 27./4*sh*math.Pow(R.yt, 4) - 6*math.Pow(sh, 3)*math.Pow(R.lH, 2))
		}
		```
	3. Calculate total Beta function (1-loop + 2-loop)
		
		```Go
		// Calculate Total Beta Function
		func (B *Beta) BetaFunction() {
			B.gamma = 1./(16*math.Pow(math.Pi, 2))*B.gamma1 + 1./math.Pow(16*math.Pow(math.Pi, 2), 2)*B.gamma2
			g := MakeBeta(B.gamma)
			
			// Use : total = g(1-loop, 2-loop)
			B.BlH = g(B.B1lH, B.B2lH)
			B.Byt = g(B.B1yt, B.B2yt)
			B.Bg1 = g(B.B1g1, B.B2g1)
			B.Bg2 = g(B.B1g2, B.B2g2)
			B.Bg3 = g(B.B1g3, B.B2g3)
		}
		```

5. Declare method to solve RGE in `rge.go`

	```Go
	// Single Running - You can change Numerical Integration method here (Default: Euler)
	func (R *RGE) Run(mt, xi float64) {
		var B Beta
	
		B.InputFormula(*R, mt, xi) // reverse pointer
		B.BetaFunction()
	
		// Real Running
		R.lH += h * B.BlH
		R.yt += h * B.Byt
		R.g1 += h * B.Bg1
		R.g2 += h * B.Bg2
		R.g3 += h * B.Bg3
		R.t += h
		R.phi = math.Sqrt(2.) / R.yt * mt * math.Exp(R.t)
		R.G -= h * B.gamma / (1 + B.gamma)
	}
	
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
	
	// SolveRGE running RGE for Step
	func (C *Container) SolveRGE(mt, xi float64) {
		R := Initialize(mt)
		C[0] = R.Copy()
	
		for i := range C {
			R.Run(mt, xi)
			C[i] = R.Copy()
		}
	}
	
	// RGERunning is main tool
	func RGERunning(mt, xi float64) []int {
		var C Container
		C.SolveRGE(mt, xi)
	
		W := make([][]string, len(C), len(C))
	
		for i, elem := range C {
			W[i] = Convert([]float64{elem.t, elem.lH, elem.yt, elem.g1, elem.g2, elem.g3, elem.G})
		}
		mtint := int(mt)
		mtfloat := int((mt-float64(mtint))*100 + 0.49)
		xiint := int(xi)
		title := fmt.Sprintf("Data/Gauge_%d_%d_%d.csv", mtint, mtfloat, xiint)
		csv.Write(W, title)
		return []int{mtint, mtfloat, xiint}
	}
	
	// Convert supports csv.Write
	func Convert(List []float64) []string {
		Temp := make([]string, len(List), len(List))
		for i := range List {
			Temp[i] = fmt.Sprintf("%v", List[i])
		}
		return Temp
	}
	```
	* RGERunning : Transfer mtint, mtfloat, xi for file name \& Main Action - solve RGE
	* Convert : `[]float64` to `[]string`

6. Handle `cmd/main.go` : From Go to Julia

	```Go
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
	// ...
	```

7. Handle Julia Files in `Julia/`

	```Julia
	using Winston

	println("-----------------------------------")
	println("  Welcome to Gauge Plot.jl")
	println("-----------------------------------")

	mt_int = ARGS[1]
	mt_float = ARGS[2]
	xi = ARGS[3]

	Data = readcsv("Data/Gauge_$(mt_int)_$(mt_float)_$(xi).csv")


	t = Data[:,1];
	# λ = Data[:,2];
	yt = Data[:,3];
	g1 = Data[:,4];
	g2 = Data[:,5];
	g3 = Data[:,6];
	# G = Data[:,7];

	# Gauge Plot
	p = FramedPlot(
		title="Gauge Plots",
		xlabel="t",
		ylabel="Gauge");
	C0 = Curve(t, yt, color="purple")
	C1 = Curve(t, g1, color="red")
	C2 = Curve(t, g2, color="blue")
	C3 = Curve(t, g3, color="green")
	setattr(C0, "label", "yt")
	setattr(C1, "label", "g1")
	setattr(C2, "label", "g2")
	setattr(C3, "label", "g3")
	lgnd = Legend(.9, .9, [C0, C1, C2, C3]);
	add(p, C0, C1, C2, C3, lgnd)
	savefig(p, "Fig/Gauge_$(mt_int)_$(mt_float)_$(xi).svg", (1000, 600))
	run(`inkscape -z Fig/Gauge_$(mt_int)_$(mt_float)_$(xi).svg -e Fig/Gauge_$(mt_int)_$(mt_float)_$(xi).png -d 300 --export-background=WHITE`)
	run(`rm Fig/Gauge_$(mt_int)_$(mt_float)_$(xi).svg`)

	```
	* Handle Julia is so easy.
	* Requirements:
		* Julia Winston Package
		* Inkscape
	
8. Make

	```bash
	make
	```
	
9. Run

	```bash
	./main
	```