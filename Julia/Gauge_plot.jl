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
