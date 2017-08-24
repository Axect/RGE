using DataFrames, Gadfly

println("---------------------------")
println("Welcome to Lambda Plot.jl")
println("---------------------------")

mt_int = ARGS[1]
mt_float = ARGS[2]
xi = ARGS[3]

Data = readcsv("../Data/Lambda_$(mt_int)_$(mt_float)_$(xi).csv")


t = Data[:,1];
λ = Data[:,2];
# gauge = vcat(Data[:,3], Data[:,4], Data[:,5], Data[:,6]); # yt, g1, g2, g3
# G = Data[:,7];

function main()
    # Gauge Plot
    dh = DataFrame(t=t, λ=λ, index=repeat(["λ",], inner=[length(t)]));
    pl3 = plot(dh, x=:t, y=:λ, color=:index, Geom.line, Guide.title("λ(t)"), Theme(background_color=color("white")))
    draw(SVG("Fig/lambda_$(mt_int)_$(mt_xi)_$(xi).svg", 1000px, 600px), pl3)
    println("Complete λ(t)")
end

