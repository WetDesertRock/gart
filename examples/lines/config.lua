math.randomseed(os.time())
math.random(); math.random(); math.random(); 

function MutatingVariable(initialValue, amplitude, period, phase)
	return {
		Type = "Sin",
		Value = initialValue,
		Amplitude = amplitude,
		Period = period,
		Phase = phase
	}
end

return {
	gart = {
		renderer = {
			Width = 1000*3,
			Height = 1000*3,
			OutPath = "./output",
			BackgroundColor = {H=0, S=0, L=1},
			BackgroundAlpha = 1,
		}
	},
	lines = {
		Expression = "sin(x * m1 + m2) * m3 + sin(x * sin(m4) * m5 + m6) * m7 + sin(x * m8 + m9) * m10 + sin(pow(x+0.3, 3) * 21.5 + 0.2) * 0.1",
		MaxLines = 5000, -- The number of lines that are spread across the height of the painting
		MarginLines = 45, -- How many lines are added before and after the height of the painting
		YScale = 0.15,
		LinePeriod = 0.001, -- How frequent calculations are made
		Render = {
			LineWidth = 0.0007,
		},
		Variables = {
			m1 = MutatingVariable(8, 0.1, 0.02, math.random() * math.pi * 2),
			m2 = MutatingVariable(0.3, 0.025, 0.08, math.random() * math.pi * 2),
			m3 = MutatingVariable(0.3, 0.5, 0.08, math.random() * math.pi * 2),
			m4 = MutatingVariable(3.14, 0.1, 0.1, math.random() * math.pi * 2),
			m5 = MutatingVariable(14.1, 0.5, 0.1, math.random() * math.pi * 2),
			m6 = MutatingVariable(3.14, 0.25, 0.2, math.random() * math.pi * 2),
			m7 = MutatingVariable(0.5, .05, 0.3, math.random() * math.pi * 2),
			m8 = MutatingVariable(5.2, 0.1, -0.0151, math.random() * math.pi * 2),
			m9 = MutatingVariable(0.2, 0.08, -0.0151, math.random() * math.pi * 2),
			m10 = MutatingVariable(0.14, 0.04, -0.02, math.random() * math.pi * 2),
		}
	},
}
