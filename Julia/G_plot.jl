using Winston

println("-----------------------------------")
println("  Welcome to G(t) Plot.jl  ")
println("-----------------------------------")

mt_int = ARGS[1]
mt_float = ARGS[2]
xi = ARGS[3]

Data = readcsv("Data/Gauge_$(mt_int)_$(mt_float)_$(xi).csv")


t = Data[:,1];
# λ = Data[:,2];
# yt = Data[:,3];
# g1 = Data[:,4];
# g2 = Data[:,5];
# g3 = Data[:,6];
G = Data[:,7];

# Gauge Plot
p = FramedPlot(
    title="G(t) Plots",
    xlabel="t",
    ylabel="Gauge");
C = Curve(t, G)
setattr(C, "label", "G(t)")
lgnd = Legend(.9, .9, [C])
add(p, C, lgnd)
savefig(p, "Fig/G_$(mt_int)_$(mt_float)_$(xi).svg", (1000, 600))
run(`inkscape -z Fig/G_$(mt_int)_$(mt_float)_$(xi).svg -e Fig/G_$(mt_int)_$(mt_float)_$(xi).png -d 300 --export-background=WHITE`)
run(`rm Fig/G_$(mt_int)_$(mt_float)_$(xi).svg`)
