package main

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type squareBtn struct {
	theme  *material.Theme
	button *widget.Clickable
	icon   *widget.Icon
}

var (
	yesterdayIcon, tomorrowIcon, refreshIcon *widget.Icon

	yesterdayBtn = new(widget.Clickable)
	nowBtn       = new(widget.Clickable)
	tomorrowBtn  = new(widget.Clickable)
	refreshBtn   = new(widget.Clickable)

	list = &widget.List{
		List: layout.List{
			Axis: layout.Vertical,
		},
	}

	breakfastCaption = material.LabelStyle{TextSize: unit.Dp(18)}
	lunchCaption     = material.LabelStyle{TextSize: unit.Dp(18)}
	dinnerCaption    = material.LabelStyle{TextSize: unit.Dp(18)}
)
