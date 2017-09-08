using Plots

println("-----------------------------------")
println("  Welcome to Potential Plot.jl")
println("-----------------------------------")

mt_int = ARGS[1]
mt_float = ARGS[2]
xi = ARGS[3]

Data = readcsv("Data/Cosmo_$(mt_int)_$(mt_float)_$(xi).csv")


t = Data[:,1];
V = Data[:,2];

# Background
gr(size=(1000,600), dpi=100)

# Gauge Plot
plot(t, V, xlims=(0,5), ylims=(0,0.00000006), title="Potential Plots", label="V", show=false);
xlabel!("t");
ylabel!("V");
savefig("Fig/Potential_$(mt_int)_$(mt_float)_$(xi).svg")
run(`inkscape -z Fig/Potential_$(mt_int)_$(mt_float)_$(xi).svg -e Fig/Potential_$(mt_int)_$(mt_float)_$(xi).png -d 600`)
run(`rm Fig/Potential_$(mt_int)_$(mt_float)_$(xi).svg`)
