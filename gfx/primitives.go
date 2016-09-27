package gfx

import (
	"fmt"
	"math"

	"github.com/go-gl/gl/v2.1/gl"
)

var lineWidth float32 = 1
var lineR float64 = 1
var lineG float64 = 1
var lineB float64 = 1
var lineA float64 = 1
var smoothLines bool

// SetLineWidth sets the width of the line for all primitives.
func SetLineWidth(w float64) {
	lineWidth = float32(w)
}

// SetLineColor sets the color of the line for all primitives.
func SetLineColor(r, g, b, a float64) {
	lineR, lineG, lineB, lineA = r, g, b, a
}

// SetSmoothLines sets if lines should be drawn with jagged edges or anti-aliasing.
func SetSmoothLines(b bool) {
	smoothLines = b
}

// RenderLines renders multiple lines for the given coords, coords must be an aven number of floats with alternating x and y coordinates.
func RenderLines(coords ...float64) error {
	if len(coords)%4 != 0 {
		return fmt.Errorf("Can only render an even number of x, y coords")
	}
	renderPoints(gl.LINES, coords...)
	return nil
}

// RenderPolygon renders a ploygon from the given coords. When closed is true, the last point will be connected to the first.
func RenderPolygon(closed bool, coords ...float64) error {
	if len(coords)%2 != 0 {
		return fmt.Errorf("Can only render an even number of x, y coords")
	}
	var mode uint32 = gl.LINE_STRIP
	if closed {
		mode = gl.LINE_LOOP
	}
	renderPoints(mode, coords...)
	return nil
}

// RenderCircle renders a cirlc at the given coordinates with the given radius and segments.
func RenderCircle(x, y, radius float64, segments int) {
	coords := make([]float64, segments*2)
	diff := 2 * math.Pi / float64(segments)
	angle := 0.0
	for i := 0; i < segments*2; i += 2 {
		coords[i] = math.Sin(angle)*radius + x
		coords[i+1] = math.Cos(angle)*radius + y
		angle += diff
	}

	RenderPolygon(true, coords...)
}

// RenderRectangle redners a rectangle for the given upper left and lower right corner.
func RenderRectangle(x1, y1, x2, y2 float64) {
	RenderPolygon(true, x1, y1, x1, y2, x2, y2, x2, y1)
}

func renderPoints(mode uint32, coords ...float64) {
	gl.LoadIdentity()
	gl.LineWidth(lineWidth)
	if smoothLines {
		gl.Enable(gl.POINT_SMOOTH)
		gl.Enable(gl.LINE_SMOOTH)
		gl.Enable(gl.POLYGON_SMOOTH)
	} else {
		gl.Disable(gl.POINT_SMOOTH)
		gl.Disable(gl.LINE_SMOOTH)
		gl.Disable(gl.POLYGON_SMOOTH)
	}

	gl.Begin(mode)
	for i := 0; i < len(coords); i += 2 {
		gl.Color4d(lineR, lineG, lineB, lineA)
		gl.Vertex3d(coords[i], -coords[i+1], 0)
	}
	gl.End()
}
