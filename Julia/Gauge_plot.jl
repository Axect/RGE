using Plots

println("---------------------------")
println("Welcome to Gauge Plot.jl")
println("---------------------------")

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

# Background
gr(size=(1000,600), dpi=600)

# Gauge Plot
plot(t, yt, title="Gauge Plots", label="yt", show=false);
plot!(t, g1, label="g1");
plot!(t, g2, label="g2");
plot!(t, g3, label="g3");
xlabel!("t");
ylabel!("gauge");
savefig("Fig/Gauge_$(mt_int)_$(mt_float)_$(xi).svg")
