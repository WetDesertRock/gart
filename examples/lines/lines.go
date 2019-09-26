package main

import (
	"fmt"
	"math"

	"github.com/Knetic/govaluate"
	"github.com/wetdesertrock/gart"
)

type Settings struct {
	Expression  string
	MaxLines    float64
	MarginLines float64
	YScale      float64
	LinePeriod  float64
	Render      struct {
		LineWidth float64
	}
	Variables map[string]MutatingVariableSetting
}

type MutatingVariableSetting struct {
	Type      string
	Value     float64
	Amplitude float64
	Period    float64
	Phase     float64
}

type MutatingVariable struct {
	Value     float64
	offset    float64
	amplitude float64
	phase     float64
	period    float64
}

func NewMutatingVariable(initalValue, amplitude, period, phase float64) *MutatingVariable {
	result := MutatingVariable{}

	result.Value = initalValue
	result.offset = initalValue
	result.amplitude = amplitude
	result.period = period
	result.phase = phase

	return &result
}

func (self *MutatingVariable) Mutate() {
	self.phase += self.period
	self.Value = math.Sin(self.phase)*self.amplitude + self.offset
}

func expressionFromString(expr string) (*govaluate.EvaluableExpression, error) {
	functions := map[string]govaluate.ExpressionFunction{
		"sin": func(args ...interface{}) (interface{}, error) {
			if fval, ok := args[0].(float64); ok {
				val := math.Sin(fval)
				return val, nil
			}
			return nil, fmt.Errorf("Unable to coerce argument to sin() to float")
		},
		"pow": func(args ...interface{}) (interface{}, error) {
			arg1, ok1 := args[0].(float64)
			arg2, ok2 := args[1].(float64)
			if ok1 && ok2 {
				val := math.Pow(arg1, arg2)
				return val, nil
			} else {
				return nil, fmt.Errorf("Unable to coerce argument #1 to pow() to float ")
			}
		},
	}

	return govaluate.NewEvaluableExpressionWithFunctions(expr, functions)
}

func makeMutations(variables map[string]MutatingVariableSetting) map[string]*MutatingVariable {
	mutations := make(map[string]*MutatingVariable)
	for key, config := range variables {
		switch config.Type {
		case "Sin":
			mutation := NewMutatingVariable(config.Value, config.Amplitude, config.Period, config.Phase)
			mutations[key] = mutation
		}
	}

	return mutations
}

func paint(settings Settings, renderer *gart.Renderer, expression *govaluate.EvaluableExpression, mutations map[string]*MutatingVariable) {
	maxLines := settings.MaxLines
	yoff := 1.0 / maxLines
	margin := settings.MarginLines
	yscale := settings.YScale
	lineWidth := settings.Render.LineWidth
	linePeriod := settings.LinePeriod

	parameters := make(map[string]interface{}, 5)

	for line := -margin; line < maxLines+margin; line++ {

		var curLine = make([]gart.Vector2d, 0)

		for x := float64(0); x < 1.1; x += linePeriod {
			parameters["x"] = x

			for key, mutation := range mutations {
				parameters[key] = mutation.Value
			}

			_y, _ := expression.Evaluate(parameters)
			if y, ok := _y.(float64); ok {
				y *= yscale
				y += yoff * line

				curLine = append(curLine, gart.NewVector2d(x, y))
			}
		}

		renderer.SetColor(gart.HSLColor{H: 0, S: 0, L: 0}, 0.25)
		renderer.LineWidth(lineWidth)
		renderer.Line(false, curLine)

		for _, mutation := range mutations {
			mutation.Mutate()
		}
	}
}

func main() {
	painting := gart.InitProgramWithPainting("lines")

	renderer := painting.Renderer

	settings := Settings{}
	painting.Settings.Get("lines", &settings)

	mutations := makeMutations(settings.Variables)

	expressionstr := settings.Expression
	expression, _ := expressionFromString(expressionstr)

	paint(settings, renderer, expression, mutations)

	painting.Save()
}
