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

# Set Figure Axis
function KeyGenerator(phi)
    for (i, elem) in enumerate(phi)
        if abs(elem-5) < 1e-03
            key_index = i
            return key_index
        end
    end
end
key_index = KeyGenerator(phi)
y_key = V[key_index];

# Potential Plot
p = FramedPlot(
    title="Potential Plots",
    xlabel="\\phi",
    ylabel="V",
    xrange=(0,5),
    yrange=(0,y_key+0.1y_key));
C = Curve(phi, V)
setattr(C, "label", "V")
lgnd = Legend(.9, .9, [C])
add(p, C, lgnd)
savefig(p, "Fig/Potential_$(mt_int)_$(mt_float)_$(xi).svg", (1000, 600))
run(`inkscape -z Fig/Potential_$(mt_int)_$(mt_float)_$(xi).svg -e Fig/Potential_$(mt_int)_$(mt_float)_$(xi).png -d 300 --export-background=WHITE`)
run(`rm Fig/Potential_$(mt_int)_$(mt_float)_$(xi).svg`)
