package svgostuff

import (
	"fmt"
	"os"

	svg "github.com/ajstarks/svgo"
)

// SimpleCircle draws a simple circle to stdout
func SimpleCircle() {
	width := 1500
	height := 2000
	f, _ := os.Create("test1.svg")
	canvas := svg.New(f)
	canvas.Start(width, height)
	canvas.Def()
	canvas.Gid("uni_crema")
	canvas.Image(0, 0, 150, 75, "file:///Users/henrik/Documents/TaPhoenix/Tiles/universal_crema.png", `preserveAspectRatio="none"`)
	canvas.Gend()
	// canvas.ClipPath(`id="cut-off-right"`)
	// canvas.Rect(75, 0, 75, 75)
	// canvas.ClipEnd()
	// TileByCount(canvas, 4, 16, 2, "uni_crema")
	// TileByCountStaggered(canvas, 4, 16, 2, "uni_crema")
	leftArea := TileAreaStaggered("panel3", canvas, Dim{width: 1360, height: 1680}, Alignment{verticalLeft: false, horizontalTop: false}, 2, "uni_crema")
	leftNookArea := TileAreaStaggered("panel4", canvas, Dim{width: 300, height: 1680}, Alignment{verticalLeft: true, horizontalTop: false}, 2, "uni_crema")
	backNookArea := TileAreaStaggered("panel5", canvas, Dim{width: 1100, height: 758}, Alignment{verticalLeft: true, horizontalTop: false}, 2, "uni_crema")
	rightNookArea := TileAreaStaggered("panel6", canvas, Dim{width: 300, height: 1680}, Alignment{verticalLeft: true, horizontalTop: false}, 2, "uni_crema")
	rightArea := TileAreaStaggered("panel7", canvas, Dim{width: 1360, height: 1680}, Alignment{verticalLeft: true, horizontalTop: false}, 2, "uni_crema")
	courtyardArea := TileAreaStaggered("panel8", canvas, Dim{width: 1500, height: 1680}, Alignment{verticalLeft: true, horizontalTop: false}, 2, "uni_crema")
	canvas.DefEnd()

	canvas.Group(`transform="scale(0.25)"`)

	next := 10
	canvas.Use(next, 10, "#panel3")

	next += 40 + leftArea.width
	canvas.Use(next, 10, "#panel4")

	next += 40 + leftNookArea.width
	canvas.Use(next, 10, "#panel5")

	next += 40 + backNookArea.width
	canvas.Use(next, 10, "#panel6")

	next += 40 + rightNookArea.width
	canvas.Use(next, 10, "#panel7")

	next += 40 + rightArea.width
	canvas.Use(next, 10, "#panel8")

	next += 40 + courtyardArea.width

	canvas.Gend()
	// canvas.Use(0, 77, "#uni_crema")
	canvas.End()
}

// Dim is dimension in mm (px)
type Dim struct {
	width  int
	height int
}

// Alignment describes alignment vertically and horizontally.
type Alignment struct {
	verticalLeft  bool
	horizontalTop bool
}

// TileByCount - uniform tiling with given number of tiles
// The panel is filled with given number of tiles starting from the top
//
func TileByCount(canvas *svg.SVG, xcount int, ycount int, gap int, tile string) Dim {
	resultingDim := Dim{width: xcount*(150+gap) - gap, height: ycount*(75+gap) - gap}
	canvas.Gid("panel1")
	// Grouting background panel
	canvas.Rect(0, 0, resultingDim.width, resultingDim.height, "fill:#cccccc")
	var startx, starty int
	for x := 0; x < xcount; x++ {
		startx = (150 + gap) * x
		for y := 0; y < ycount; y++ {
			starty = (75 + gap) * y
			canvas.Use(startx, starty, fmt.Sprintf("#%s", tile))
		}
	}
	canvas.Gend()
	return resultingDim
}

// TileByCountStaggered Defines a staggered panel
func TileByCountStaggered(canvas *svg.SVG, xcount int, ycount int, gap int, tile string) Dim {
	resultingDim := Dim{width: xcount*(150+gap) - gap, height: ycount*(75+gap) - gap}
	canvas.Gid("panel2")
	canvas.ClipPath(`id="panel2-cut"`)
	canvas.Rect(0, 0, resultingDim.width, resultingDim.height)
	canvas.ClipEnd()

	// Grouting background panel
	canvas.Rect(0, 0, resultingDim.width, resultingDim.height, "fill:#cccccc")
	// Clip to the grouted area
	canvas.Group(`clip-path="url(#panel2-cut)"`)
	var startx, starty int
	for y := 0; y < ycount; y++ {
		shifted := false
		extra := 0
		if y%2 != 0 {
			shifted = true
			extra = 1
		}
		starty = (75 + gap) * y
		for x := 0; x < xcount+extra; x++ {
			startx = (150 + gap) * x
			if shifted {
				startx -= 75 + gap
			}
			canvas.Use(startx, starty, fmt.Sprintf("#%s", tile))
		}
	}
	canvas.Gend()
	canvas.Gend()
	return resultingDim
}

// TileAreaStaggered Defines a staggered panel
func TileAreaStaggered(panelName string, canvas *svg.SVG, area Dim, align Alignment, gap int, tile string) Dim {
	th := 75
	tw := 150 // cheat: use Tile object

	xcount := area.width / (tw + gap)
	ycount := area.height / (th + gap)

	// Compute how much to shift when drawing to align at
	// right and bottom instead of top left
	xdiff := tw - (area.width - xcount*(tw+gap))
	ydiff := th - (area.height - ycount*(th+gap))
	if align.verticalLeft {
		xdiff = 0
	}
	if align.horizontalTop {
		ydiff = 0
	}

	canvas.Gid(panelName)
	canvas.ClipPath(fmt.Sprintf(`id="%s-cut"`, panelName))
	canvas.Rect(0, 0, area.width, area.height)
	canvas.ClipEnd()

	// Grouting background panel
	canvas.Rect(0, 0, area.width, area.height, "fill:#cccccc")

	// Clip to the grouted area
	canvas.Group(fmt.Sprintf(`clip-path="url(#%s-cut)"`, panelName))
	var startx, starty int

	for y := 0; y < ycount+1; y++ {
		shifted := false
		extra := 2
		if y%2 != 0 {
			shifted = true
		}
		starty = (th+gap)*y - ydiff
		for x := 0; x < xcount+extra; x++ {
			startx = (tw+gap)*x - xdiff
			if shifted {
				startx -= tw/2 + gap
			}
			canvas.Use(startx, starty, fmt.Sprintf("#%s", tile))
		}
	}
	canvas.Gend()
	canvas.Gend()
	return area
}
