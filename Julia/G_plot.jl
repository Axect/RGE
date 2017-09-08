using Plots

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

# Background
gr(size=(1000,600), dpi=100)

# Gauge Plot
plot(t, G, title="Gauge Plots", label="G", show=false);
xlabel!("t");
ylabel!("gauge");
savefig("Fig/G_$(mt_int)_$(mt_float)_$(xi).svg")
run(`inkscape -z Fig/G_$(mt_int)_$(mt_float)_$(xi).svg -e Fig/G_$(mt_int)_$(mt_float)_$(xi).png -d 600`)
run(`rm Fig/G_$(mt_int)_$(mt_float)_$(xi).svg`)
