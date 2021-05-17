package controldisplay

import (
	"fmt"
	"log"

	"github.com/turbot/go-kit/helpers"
)

type GroupRenderer struct {
	title string

	failedControls    int
	totalControls     int
	maxFailedControls int
	maxTotalControls  int
	// screen width
	width int
}

func NewGroupRenderer(title string, failed, total, maxFailed, maxTotal, width int) *GroupRenderer {
	return &GroupRenderer{
		title:             title,
		failedControls:    failed,
		totalControls:     total,
		maxFailedControls: maxFailed,
		maxTotalControls:  maxTotal,
		width:             width,
	}
}

func (r GroupRenderer) Render() string {
	log.Println("[TRACE] begin group render")
	defer log.Println("[TRACE] end group render")

	if r.width <= 0 {
		log.Printf("[WARN] group renderer has width of %d\n", r.width)
		return ""
	}

	counterString := NewCounterRenderer(r.failedControls, r.totalControls, r.maxFailedControls, r.maxTotalControls).Render()
	counterWidth := helpers.PrintableLength(counterString)

	graphString := NewCounterGraphRenderer(r.failedControls, r.totalControls, r.maxTotalControls).Render()
	graphWidth := helpers.PrintableLength(graphString)

	// figure out how much width we have available for the title
	availableWidth := r.width - counterWidth - graphWidth

	// now availableWidth is all we have - if it is not enough we need to truncate the title
	titleString := NewGroupTitleRenderer(r.title, availableWidth).Render()
	titleWidth := helpers.PrintableLength(titleString)

	// is there any room for a spacer
	spacerWidth := availableWidth - titleWidth
	var spacerString string
	if spacerWidth > 0 {
		spacerString = NewSpacerRenderer(spacerWidth).Render()
	}

	// now put these all together
	str := fmt.Sprintf("%s%s%s%s", titleString, spacerString, counterString, graphString)
	return str
}
