using DataFrames, Gadfly

println("---------------------------")
println("Welcome to RGE Plot.jl")
println("---------------------------")

Data = readcsv("Data/Gauge_170_85_50.csv")


t = Data[:,1];
λ = Data[:,2];
gauge = vcat(Data[:,3], Data[:,4], Data[:,5], Data[:,6]); # yt, g1, g2, g3
G = Data[:,7];

function main()

    println("What do you like draw?")
    println("1.Gauge, 2.G(t), 3.λ(t), 4.Potential")
    choose = readline(STDIN)
    choice = split(choose)
    # Gauge Plot
    if "1" in choice
        df = DataFrame(t=repeat(t, outer=[4]), gauge=gauge, index=repeat(["yt", "g1", "g2", "g3"], inner=[length(t)]));
        pl = plot(df, x=:t, y=:gauge, color=:index, Geom.line, Guide.title("Gauge"))
        draw(SVG("Fig/Gauge_170_85_50.svg", 1000px, 600px), pl)
        println("Complete Gauge")
    end

    if "2" in choice
        dg = DataFrame(t=t, G=G, index=repeat(["G",], inner=[length(t)]));
        pl2 = plot(dg, x=:t, y=:G, color=:index, Geom.line, Guide.title("G(t)"))
        draw(SVG("Fig/G_170_85_50.svg", 1000px, 600px), pl2)
        println("Complete G(t)")
    end

    if "3" in choice
        dh = DataFrame(t=t, λ=λ, index=repeat(["λ",], inner=[length(t)]));
        pl3 = plot(dh, x=:t, y=:λ, color=:index, Geom.line, Guide.title("λ(t)"))
        draw(SVG("Fig/lambda_170_85_50.svg", 1000px, 600px), pl3)
        println("Complete λ(t)")
    end

    println("Complete all process. Do you want to exit julia? ", "y/n")
    a = readline(STDIN)

    if a == "y"
        exit(1)
    end
end
