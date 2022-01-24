package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"golang.org/x/exp/shiny/materialdesign/icons"
	"image/color"
	"log"
	"os"
)

func main() {
	yesterdayIcon, _ = widget.NewIcon(icons.NavigationArrowBack)
	tomorrowIcon, _ = widget.NewIcon(icons.NavigationArrowForward)
	refreshIcon, _ = widget.NewIcon(icons.NavigationRefresh)

	go func() {
		w := app.NewWindow(
			app.Title("Test App"),
			app.Size(unit.Dp(400), unit.Dp(700)),
		)

		if err := loop(w); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()

	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())

	var ops op.Ops
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)

				top(gtx, th)
				e.Frame(gtx.Ops)
			}
		}
	}
}

func (b squareBtn) Layout(gtx layout.Context) layout.Dimensions {
	return material.ButtonLayout(b.theme, b.button).Layout(gtx, func(gtx C) D {
		return layout.UniformInset(unit.Dp(7)).Layout(gtx, func(gtx C) D {
			iconAndLabel := layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}

			layIcon := layout.Rigid(func(gtx C) D {
				return layout.Inset{}.Layout(gtx, func(gtx C) D {
					return b.icon.Layout(gtx, b.theme.ContrastFg)
				})
			})

			return iconAndLabel.Layout(gtx, layIcon)
		})
	})
}

func top(gtx layout.Context, th *material.Theme) layout.Dimensions {
	widgets := []layout.Widget{
		func(gtx C) D {
			in := layout.UniformInset(unit.Dp(7))

			return layout.Flex{
				Alignment: layout.Middle,
				Spacing:   layout.SpaceEvenly,
			}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, squareBtn{theme: th, icon: yesterdayIcon, button: yesterdayBtn}.Layout)
				}),

				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						for nowBtn.Clicked() {

						}

						dims := material.Button(th, nowBtn, "2022-01-24").Layout(gtx)

						return dims
					})
				}),

				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, squareBtn{theme: th, icon: tomorrowIcon, button: tomorrowBtn}.Layout)
				}),
			)
		},

		func(gtx C) D { // 아침
			breakfastCaption = material.Caption(th, "")

			border := widget.Border{Color: color.NRGBA{A: 0xff}, CornerRadius: unit.Dp(8), Width: unit.Px(2)}

			return border.Layout(gtx, func(gtx C) D {
				return layout.UniformInset(unit.Dp(8)).Layout(gtx, breakfastCaption.Layout)
			})
		},

		func(gtx C) D { // 점심
			lunchCaption = material.Caption(th, "")

			border := widget.Border{Color: color.NRGBA{A: 0xff}, CornerRadius: unit.Dp(8), Width: unit.Px(2)}

			return border.Layout(gtx, func(gtx C) D {
				return layout.UniformInset(unit.Dp(8)).Layout(gtx, lunchCaption.Layout)
			})
		},

		func(gtx C) D { // 저녁
			dinnerCaption = material.Caption(th, "")

			border := widget.Border{Color: color.NRGBA{A: 0xff}, CornerRadius: unit.Dp(8), Width: unit.Px(2)}

			return border.Layout(gtx, func(gtx C) D {
				return layout.UniformInset(unit.Dp(8)).Layout(gtx, dinnerCaption.Layout)
			})
		},

		func(gtx C) D {
			in := layout.UniformInset(unit.Dp(10))
			in.Left = unit.Dp(300)
			in.Top = unit.Dp(330)

			return layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, material.IconButton(th, refreshBtn, refreshIcon, "Refresh").Layout)
				}),
			)
		},
	}

	return material.List(th, list).Layout(gtx, len(widgets), func(gtx C, i int) D {
		return layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
	})
}
