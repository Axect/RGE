using Winston

println("-----------------------------------")
println("  Welcome to Lambda Plot.jl")
println("-----------------------------------")

mt_int = ARGS[1]
mt_float = ARGS[2]
xi = ARGS[3]

Data = readcsv("Data/Gauge_$(mt_int)_$(mt_float)_$(xi).csv")


t = Data[:,1];
λ = Data[:,2];
# gauge = vcat(Data[:,3], Data[:,4], Data[:,5], Data[:,6]); # yt, g1, g2, g3
# G = Data[:,7];

# Gauge Plot
p = FramedPlot(
    title="Lambda Plots",
    xlabel="t",
    ylabel="\\lambda");
C = Curve(t, λ)
setattr(C, "label", "\\lambda")
lgnd = Legend(.9, .9, [C])
add(p, C, lgnd)
savefig(p, "Fig/Lambda_$(mt_int)_$(mt_float)_$(xi).svg", (1000, 600))
run(`inkscape -z Fig/Lambda_$(mt_int)_$(mt_float)_$(xi).svg -e Fig/Lambda_$(mt_int)_$(mt_float)_$(xi).png -d 300 --export-background=WHITE`)
run(`rm Fig/Lambda_$(mt_int)_$(mt_float)_$(xi).svg`)
