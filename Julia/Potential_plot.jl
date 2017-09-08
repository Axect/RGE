using Winston

println("-----------------------------------")
println("  Welcome to Potential Plot.jl")
println("-----------------------------------")

mt_int = ARGS[1]
mt_float = ARGS[2]
xi = ARGS[3]

Data = readcsv("Data/Cosmo_$(mt_int)_$(mt_float)_$(xi).csv")


phi = Data[:,1];
V = Data[:,2];

# Potential Plot
p = FramedPlot(
    title="Potential Plots",
    xlabel="\\phi",
    ylabel="V",
    xrange=(0,5),
    yrange=(0,6e-08));
C = Curve(phi, V)
setattr(C, "label", "V")
lgnd = Legend(.9, .9, [C])
add(p, C, lgnd)
savefig(p, "Fig/Potential_$(mt_int)_$(mt_float)_$(xi).svg", (1000, 600))
run(`inkscape -z Fig/Potential_$(mt_int)_$(mt_float)_$(xi).svg -e Fig/Potential_$(mt_int)_$(mt_float)_$(xi).png -d 300 --export-background=WHITE`)
run(`rm Fig/Potential_$(mt_int)_$(mt_float)_$(xi).svg`)
