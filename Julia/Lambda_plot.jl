using Plots, LaTeXStrings

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

# Background
gr(size=(1000,600), dpi=100)

# Gauge Plot
plot(t, λ, title="Gauge Plots", label=L"$\lambda$", show=false);
xlabel!("t");
ylabel!("gauge");
savefig("Fig/Lambda_$(mt_int)_$(mt_float)_$(xi).svg")
run(`inkscape -z Fig/Lambda_$(mt_int)_$(mt_float)_$(xi).svg -e Fig/Lambda_$(mt_int)_$(mt_float)_$(xi).png -d 600`)
run(`rm Fig/Lambda_$(mt_int)_$(mt_float)_$(xi).svg`)
