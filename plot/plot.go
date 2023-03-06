package plot

import (
	"github.com/GuiiFerrari/go_training/algelin"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func NewPlot(xtitle string, ytitle string, title string) *plot.Plot {
	canvas := plot.New()
	canvas.X.Label.Text = xtitle
	canvas.Y.Label.Text = ytitle
	canvas.Title.Text = title
	//  Increase the default font size.
	canvas.Title.TextStyle.Font.Size = 20
	canvas.X.Label.TextStyle.Font.Size = 20
	canvas.Y.Label.TextStyle.Font.Size = 20
	return canvas
}

func Plot_test_1() {
	point_a := algelin.NewPoint2D(0.0, 0.0)
	versor := algelin.NewPoint2D(0.0, 0.45)
	canvas := NewPlot("x", "y", "Test")
	// Plot the points
	pts := plotter.XYs{{X: point_a.X, Y: point_a.Y}}
	plotutil.AddLinePoints(canvas, "Point A", pts)
	// Plot the vector
	vector := algelin.NewVector2D(point_a, versor)
	// vector_points := []float64{vector.Origin.X, vector.Origin.Y, vector.PointAt(1.0).X, vector.PointAt(1.0).Y}
	pts = plotter.XYs{{X: vector.Origin.X, Y: vector.Origin.Y}, {X: vector.PointAt(1.0).X, Y: vector.PointAt(1.0).Y}}
	plotutil.AddLinePoints(canvas, "Vector", pts)
	// Save the plot to a PNG file.
	if err := canvas.Save(6*vg.Inch, 6*vg.Inch, "test.png"); err != nil {
		panic(err)
	}
}
