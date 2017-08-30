using DataFrames, Gadfly

println("---------------------------")
println("Welcome to Gauge Plot.jl")
println("---------------------------")

mt_int = ARGS[1]
mt_float = ARGS[2]
xi = ARGS[3]

Data = readcsv("Data/Gauge_$(mt_int)_$(mt_float)_$(xi).csv")


t = Data[:,1];
# λ = Data[:,2];
gauge = vcat(Data[:,3], Data[:,4], Data[:,5], Data[:,6]); # yt, g1, g2, g3
# G = Data[:,7];

# Gauge Plot
df = DataFrame(t=repeat(t, outer=[4]), gauge=gauge, index=repeat(["yt", "g1", "g2", "g3"], inner=[length(t)]));
pl = plot(df, x=:t, y=:gauge, color=:index, Geom.line, Guide.title("Gauge"), Theme(background_color=color("white")))
draw(SVG("Fig/Gauge_$(mt_int)_$(mt_float)_$(xi).svg", 1000px, 600px), pl)

