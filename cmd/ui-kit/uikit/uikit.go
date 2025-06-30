package uikit

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// Design System Configuration
const (
	// Base unit for consistent spacing
	BaseUnit = unit.Dp(4)

	// Corner radius scale
	RadiusSmall  = unit.Dp(4)
	RadiusMedium = unit.Dp(8)
	RadiusLarge  = unit.Dp(12)
	RadiusXL     = unit.Dp(16)

	// Shadow blur amounts
	ShadowSmall  = unit.Dp(2)
	ShadowMedium = unit.Dp(4)
	ShadowLarge  = unit.Dp(8)
)

// Comprehensive Color Palette - Modern Blue-Gray Theme
type ColorPalette struct {
	// Primary Colors
	Primary50  color.NRGBA
	Primary100 color.NRGBA
	Primary200 color.NRGBA
	Primary300 color.NRGBA
	Primary400 color.NRGBA
	Primary500 color.NRGBA // Main primary
	Primary600 color.NRGBA
	Primary700 color.NRGBA
	Primary800 color.NRGBA
	Primary900 color.NRGBA

	// Neutral/Gray Colors
	Gray50  color.NRGBA
	Gray100 color.NRGBA
	Gray200 color.NRGBA
	Gray300 color.NRGBA
	Gray400 color.NRGBA
	Gray500 color.NRGBA
	Gray600 color.NRGBA
	Gray700 color.NRGBA
	Gray800 color.NRGBA
	Gray900 color.NRGBA

	// Semantic Colors
	Success      color.NRGBA
	SuccessLight color.NRGBA
	Warning      color.NRGBA
	WarningLight color.NRGBA
	Error        color.NRGBA
	ErrorLight   color.NRGBA
	Info         color.NRGBA
	InfoLight    color.NRGBA

	// Surface Colors
	White           color.NRGBA
	Background      color.NRGBA
	Surface         color.NRGBA
	SurfaceElevated color.NRGBA

	// Text Colors
	TextPrimary   color.NRGBA
	TextSecondary color.NRGBA
	TextDisabled  color.NRGBA
	TextInverse   color.NRGBA

	// Border Colors
	Border      color.NRGBA
	BorderLight color.NRGBA
	BorderHover color.NRGBA

	// Special Colors
	Shadow  color.NRGBA
	Overlay color.NRGBA
	Focus   color.NRGBA

	OnBackground      color.NRGBA
	OnSurface         color.NRGBA
	OnPrimary         color.NRGBA
	OnSecondary       color.NRGBA
	OnError           color.NRGBA
	OnSuccess         color.NRGBA
	OnWarning         color.NRGBA
	OnInfo            color.NRGBA
	OnSurfaceElevated color.NRGBA
	OnSurfaceDisabled color.NRGBA
	OnSurfaceVariant  color.NRGBA
}

// NewColorPalette creates a modern, accessible color palette
func NewColorPalette() ColorPalette {
	return ColorPalette{
		// Primary Blue-Gray Scale
		Primary50:  color.NRGBA{R: 0xF8, G: 0xFA, B: 0xFC, A: 0xFF},
		Primary100: color.NRGBA{R: 0xF1, G: 0xF5, B: 0xF9, A: 0xFF},
		Primary200: color.NRGBA{R: 0xE2, G: 0xE8, B: 0xF0, A: 0xFF},
		Primary300: color.NRGBA{R: 0xCB, G: 0xD5, B: 0xE1, A: 0xFF},
		Primary400: color.NRGBA{R: 0x94, G: 0xA3, B: 0xB8, A: 0xFF},
		Primary500: color.NRGBA{R: 0x64, G: 0x74, B: 0x8B, A: 0xFF}, // Main primary
		Primary600: color.NRGBA{R: 0x47, G: 0x56, B: 0x69, A: 0xFF},
		Primary700: color.NRGBA{R: 0x33, G: 0x4E, B: 0x68, A: 0xFF},
		Primary800: color.NRGBA{R: 0x2D, G: 0x3F, B: 0x58, A: 0xFF},
		Primary900: color.NRGBA{R: 0x1E, G: 0x29, B: 0x3B, A: 0xFF},

		// Neutral Gray Scale
		Gray50:  color.NRGBA{R: 0xF9, G: 0xFA, B: 0xFB, A: 0xFF},
		Gray100: color.NRGBA{R: 0xF3, G: 0xF4, B: 0xF6, A: 0xFF},
		Gray200: color.NRGBA{R: 0xE5, G: 0xE7, B: 0xEB, A: 0xFF},
		Gray300: color.NRGBA{R: 0xD1, G: 0xD5, B: 0xDB, A: 0xFF},
		Gray400: color.NRGBA{R: 0x9C, G: 0xA3, B: 0xAF, A: 0xFF},
		Gray500: color.NRGBA{R: 0x6B, G: 0x72, B: 0x80, A: 0xFF},
		Gray600: color.NRGBA{R: 0x4B, G: 0x55, B: 0x63, A: 0xFF},
		Gray700: color.NRGBA{R: 0x37, G: 0x41, B: 0x51, A: 0xFF},
		Gray800: color.NRGBA{R: 0x1F, G: 0x29, B: 0x37, A: 0xFF},
		Gray900: color.NRGBA{R: 0x11, G: 0x18, B: 0x27, A: 0xFF},

		// Semantic Colors
		Success:      color.NRGBA{R: 0x10, G: 0xB9, B: 0x81, A: 0xFF},
		SuccessLight: color.NRGBA{R: 0xD1, G: 0xFA, B: 0xE5, A: 0xFF},
		Warning:      color.NRGBA{R: 0xF5, G: 0x9E, B: 0x0B, A: 0xFF},
		WarningLight: color.NRGBA{R: 0xFE, G: 0xF3, B: 0xC7, A: 0xFF},
		Error:        color.NRGBA{R: 0xEF, G: 0x44, B: 0x44, A: 0xFF},
		ErrorLight:   color.NRGBA{R: 0xFE, G: 0xE2, B: 0xE2, A: 0xFF},
		Info:         color.NRGBA{R: 0x38, G: 0x94, B: 0xF6, A: 0xFF},
		InfoLight:    color.NRGBA{R: 0xDB, G: 0xEA, B: 0xFE, A: 0xFF},

		// Surface Colors
		White:           color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
		Background:      color.NRGBA{R: 0xF9, G: 0xFA, B: 0xFB, A: 0xFF},
		Surface:         color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
		SurfaceElevated: color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},

		// Text Colors
		TextPrimary:   color.NRGBA{R: 0x11, G: 0x18, B: 0x27, A: 0xFF},
		TextSecondary: color.NRGBA{R: 0x6B, G: 0x72, B: 0x80, A: 0xFF},
		TextDisabled:  color.NRGBA{R: 0x9C, G: 0xA3, B: 0xAF, A: 0xFF},
		TextInverse:   color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},

		// Border Colors
		Border:      color.NRGBA{R: 0xE5, G: 0xE7, B: 0xEB, A: 0xFF},
		BorderLight: color.NRGBA{R: 0xF3, G: 0xF4, B: 0xF6, A: 0xFF},
		BorderHover: color.NRGBA{R: 0xD1, G: 0xD5, B: 0xDB, A: 0xFF},

		// Special Colors
		Shadow:  color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x0F},
		Overlay: color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x40},
		Focus:   color.NRGBA{R: 0x38, G: 0x94, B: 0xF6, A: 0x60},

		// On Colors - Fixed for proper contrast
		OnBackground:      color.NRGBA{R: 0x11, G: 0x18, B: 0x27, A: 0xFF}, // Dark text on light background
		OnSurface:         color.NRGBA{R: 0x11, G: 0x18, B: 0x27, A: 0xFF}, // Dark text on white surface
		OnPrimary:         color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}, // White text on primary color
		OnSecondary:       color.NRGBA{R: 0x11, G: 0x18, B: 0x27, A: 0xFF}, // Dark text on light secondary
		OnError:           color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}, // White text on error color
		OnSuccess:         color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}, // White text on success color
		OnWarning:         color.NRGBA{R: 0x11, G: 0x18, B: 0x27, A: 0xFF}, // Dark text on warning color (better contrast)
		OnInfo:            color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}, // White text on info color
		OnSurfaceElevated: color.NRGBA{R: 0x11, G: 0x18, B: 0x27, A: 0xFF}, // Dark text on elevated surface
		OnSurfaceDisabled: color.NRGBA{R: 0x9C, G: 0xA3, B: 0xAF, A: 0xFF}, // Muted text on disabled surface
		OnSurfaceVariant:  color.NRGBA{R: 0x47, G: 0x56, B: 0x69, A: 0xFF}, // Medium contrast text on variant surface
	}
}

// Spacing system based on 4px base unit
type Spacing struct {
	None     unit.Dp // 0px
	Tiny     unit.Dp // 4px
	Small    unit.Dp // 8px
	Medium   unit.Dp // 16px
	Large    unit.Dp // 24px
	XLarge   unit.Dp // 32px
	XXLarge  unit.Dp // 48px
	XXXLarge unit.Dp // 64px
}

func NewSpacing() Spacing {
	return Spacing{
		None:     unit.Dp(0),
		Tiny:     BaseUnit,
		Small:    BaseUnit * 2,
		Medium:   BaseUnit * 4,
		Large:    BaseUnit * 6,
		XLarge:   BaseUnit * 8,
		XXLarge:  BaseUnit * 12,
		XXXLarge: BaseUnit * 16,
	}
}

// Typography system with consistent hierarchy
type Typography struct {
	DisplayLarge   TypographyStyle
	DisplayMedium  TypographyStyle
	DisplaySmall   TypographyStyle
	HeadlineLarge  TypographyStyle
	HeadlineMedium TypographyStyle
	HeadlineSmall  TypographyStyle
	TitleLarge     TypographyStyle
	TitleMedium    TypographyStyle
	TitleSmall     TypographyStyle
	BodyLarge      TypographyStyle
	BodyMedium     TypographyStyle
	BodySmall      TypographyStyle
	LabelLarge     TypographyStyle
	LabelMedium    TypographyStyle
	LabelSmall     TypographyStyle
}

type TypographyStyle struct {
	Size       unit.Sp
	LineHeight float32
	Weight     string
}

func NewTypography() Typography {
	return Typography{
		DisplayLarge:   TypographyStyle{Size: unit.Sp(57), LineHeight: 64, Weight: "400"},
		DisplayMedium:  TypographyStyle{Size: unit.Sp(45), LineHeight: 52, Weight: "400"},
		DisplaySmall:   TypographyStyle{Size: unit.Sp(36), LineHeight: 44, Weight: "400"},
		HeadlineLarge:  TypographyStyle{Size: unit.Sp(32), LineHeight: 40, Weight: "400"},
		HeadlineMedium: TypographyStyle{Size: unit.Sp(28), LineHeight: 36, Weight: "400"},
		HeadlineSmall:  TypographyStyle{Size: unit.Sp(24), LineHeight: 32, Weight: "400"},
		TitleLarge:     TypographyStyle{Size: unit.Sp(22), LineHeight: 28, Weight: "500"},
		TitleMedium:    TypographyStyle{Size: unit.Sp(16), LineHeight: 24, Weight: "500"},
		TitleSmall:     TypographyStyle{Size: unit.Sp(14), LineHeight: 20, Weight: "500"},
		BodyLarge:      TypographyStyle{Size: unit.Sp(16), LineHeight: 24, Weight: "400"},
		BodyMedium:     TypographyStyle{Size: unit.Sp(14), LineHeight: 20, Weight: "400"},
		BodySmall:      TypographyStyle{Size: unit.Sp(12), LineHeight: 16, Weight: "400"},
		LabelLarge:     TypographyStyle{Size: unit.Sp(14), LineHeight: 20, Weight: "500"},
		LabelMedium:    TypographyStyle{Size: unit.Sp(12), LineHeight: 16, Weight: "500"},
		LabelSmall:     TypographyStyle{Size: unit.Sp(11), LineHeight: 16, Weight: "500"},
	}
}

// Main UI Kit struct
type UIKit struct {
	Colors     ColorPalette
	Spacing    Spacing
	Typography Typography
	Theme      *material.Theme
}

// NewUIKit creates a new UI kit instance
func NewUIKit() *UIKit {
	kit := &UIKit{
		Colors:     NewColorPalette(),
		Spacing:    NewSpacing(),
		Typography: NewTypography(),
		Theme:      material.NewTheme(),
	}

	// Configure theme with our colors
	kit.Theme.Palette.Bg = kit.Colors.Background
	kit.Theme.Palette.Fg = kit.Colors.TextPrimary
	kit.Theme.Palette.ContrastBg = kit.Colors.Primary500
	kit.Theme.Palette.ContrastFg = kit.Colors.TextInverse

	return kit
}

// Button Variants
type ButtonVariant int

const (
	ButtonPrimary ButtonVariant = iota
	ButtonSecondary
	ButtonOutline
	ButtonGhost
	ButtonDanger
	ButtonSuccess
)

type ButtonSize int

const (
	ButtonSmall ButtonSize = iota
	ButtonMedium
	ButtonLarge
)

// Button creates a styled button with consistent design
func (kit *UIKit) Button(btn *widget.Clickable, text string, variant ButtonVariant, size ButtonSize) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		var bg, fg color.NRGBA
		var inset unit.Dp
		var fontSize unit.Sp
		var borderWidth unit.Dp = 0
		var borderColor color.NRGBA

		// Size configuration
		switch size {
		case ButtonSmall:
			inset = kit.Spacing.Small
			fontSize = kit.Typography.LabelSmall.Size
		case ButtonMedium:
			inset = kit.Spacing.Medium
			fontSize = kit.Typography.LabelMedium.Size
		case ButtonLarge:
			inset = kit.Spacing.Large
			fontSize = kit.Typography.LabelLarge.Size
		}

		// Variant configuration
		switch variant {
		case ButtonPrimary:
			bg = kit.Colors.Primary500
			fg = kit.Colors.OnPrimary
		case ButtonSecondary:
			bg = kit.Colors.Gray100
			fg = kit.Colors.OnSecondary
		case ButtonOutline:
			bg = color.NRGBA{A: 0}
			fg = kit.Colors.Primary500
			borderWidth = unit.Dp(1)
			borderColor = kit.Colors.Primary500
		case ButtonGhost:
			bg = color.NRGBA{A: 0}
			fg = kit.Colors.Primary500
		case ButtonDanger:
			bg = kit.Colors.Error
			fg = kit.Colors.OnError
		case ButtonSuccess:
			bg = kit.Colors.Success
			fg = kit.Colors.OnSuccess
		}

		// Handle hover state (simplified)
		if btn.Hovered() {
			if variant == ButtonPrimary {
				bg = kit.Colors.Primary600
			} else if variant == ButtonSecondary {
				bg = kit.Colors.Gray200
			}
		}

		return widget.Border{
			Color:        borderColor,
			CornerRadius: RadiusMedium,
			Width:        borderWidth,
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return material.ButtonLayoutStyle{
				Background:   bg,
				CornerRadius: RadiusMedium,
				Button:       btn,
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(inset).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					label := material.Label(kit.Theme, fontSize, text)
					label.Color = fg
					return label.Layout(gtx)
				})
			})
		})
	}
}

// Input field with consistent styling
func (kit *UIKit) Input(editor *widget.Editor, hint string, hasError bool) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		borderColor := kit.Colors.Border
		bg := kit.Colors.Surface

		if hasError {
			borderColor = kit.Colors.Error
		}

		if gtx.Focused(editor) {
			borderColor = kit.Colors.Primary500
		}

		return widget.Border{
			Color:        borderColor,
			CornerRadius: RadiusMedium,
			Width:        unit.Dp(1),
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			defer clip.RRect{
				Rect: image.Rectangle{Max: gtx.Constraints.Max},
				NW:   int(RadiusMedium), NE: int(RadiusMedium),
				SE: int(RadiusMedium), SW: int(RadiusMedium),
			}.Push(gtx.Ops).Pop()

			paint.Fill(gtx.Ops, bg)

			return layout.UniformInset(kit.Spacing.Medium).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				ed := material.Editor(kit.Theme, editor, hint)
				ed.Color = kit.Colors.OnSurface
				ed.HintColor = kit.Colors.TextSecondary
				return ed.Layout(gtx)
			})
		})
	}
}

// Card component with shadow and consistent styling
func (kit *UIKit) Card(gtx layout.Context, content layout.Widget) layout.Dimensions {
	// Draw shadow
	shadowRect := image.Rectangle{
		Max: image.Point{
			X: gtx.Constraints.Max.X,
			Y: gtx.Constraints.Max.Y + int(ShadowSmall),
		},
	}

	return layout.Stack{}.Layout(gtx,
		// Shadow layer
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			defer clip.RRect{
				Rect: shadowRect,
				NW:   int(RadiusLarge), NE: int(RadiusLarge),
				SE: int(RadiusLarge), SW: int(RadiusLarge),
			}.Push(gtx.Ops).Pop()

			paint.Fill(gtx.Ops, kit.Colors.Shadow)
			return layout.Dimensions{Size: shadowRect.Max}
		}),
		// Main card
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return widget.Border{
				Color:        kit.Colors.BorderLight,
				CornerRadius: RadiusLarge,
				Width:        unit.Dp(1),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				defer clip.RRect{
					Rect: image.Rectangle{Max: gtx.Constraints.Max},
					NW:   int(RadiusLarge), NE: int(RadiusLarge),
					SE: int(RadiusLarge), SW: int(RadiusLarge),
				}.Push(gtx.Ops).Pop()

				paint.Fill(gtx.Ops, kit.Colors.Surface)
				return layout.UniformInset(kit.Spacing.Large).Layout(gtx, content)
			})
		}),
	)
}

// Badge component for status indicators
type BadgeVariant int

const (
	BadgeDefault BadgeVariant = iota
	BadgeSuccess
	BadgeWarning
	BadgeError
	BadgeInfo
)

func (kit *UIKit) Badge(text string, variant BadgeVariant) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		var bg, fg color.NRGBA

		switch variant {
		case BadgeDefault:
			bg = kit.Colors.Gray200
			fg = kit.Colors.TextPrimary
		case BadgeSuccess:
			bg = kit.Colors.SuccessLight
			fg = kit.Colors.Success
		case BadgeWarning:
			bg = kit.Colors.WarningLight
			fg = kit.Colors.OnWarning
		case BadgeError:
			bg = kit.Colors.ErrorLight
			fg = kit.Colors.Error
		case BadgeInfo:
			bg = kit.Colors.InfoLight
			fg = kit.Colors.Info
		}

		return widget.Border{
			Color:        bg,
			CornerRadius: unit.Dp(12),
			Width:        unit.Dp(0),
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			defer clip.RRect{
				Rect: image.Rectangle{Max: gtx.Constraints.Max},
				NW:   12, NE: 12, SE: 12, SW: 12,
			}.Push(gtx.Ops).Pop()

			paint.Fill(gtx.Ops, bg)

			return layout.Inset{
				Top: kit.Spacing.Tiny, Bottom: kit.Spacing.Tiny,
				Left: kit.Spacing.Small, Right: kit.Spacing.Small,
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				label := material.Label(kit.Theme, kit.Typography.LabelSmall.Size, text)
				label.Color = fg
				return label.Layout(gtx)
			})
		})
	}
}

// Divider component
func (kit *UIKit) Divider() layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		paint.FillShape(gtx.Ops, kit.Colors.Border, clip.Rect{
			Max: image.Point{X: gtx.Constraints.Max.X, Y: int(unit.Dp(1))},
		}.Op())
		return layout.Dimensions{
			Size: image.Point{X: gtx.Constraints.Max.X, Y: int(unit.Dp(1))},
		}
	}
}

// Alert component for notifications
type AlertVariant int

const (
	AlertInfo AlertVariant = iota
	AlertSuccess
	AlertWarning
	AlertError
)

func (kit *UIKit) Alert(title, message string, variant AlertVariant) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		var bg, borderColor, iconColor color.NRGBA
		var icon string

		switch variant {
		case AlertInfo:
			bg = kit.Colors.InfoLight
			borderColor = kit.Colors.Info
			iconColor = kit.Colors.Info
			icon = "ℹ"
		case AlertSuccess:
			bg = kit.Colors.SuccessLight
			borderColor = kit.Colors.Success
			iconColor = kit.Colors.Success
			icon = "✓"
		case AlertWarning:
			bg = kit.Colors.WarningLight
			borderColor = kit.Colors.Warning
			iconColor = kit.Colors.Warning
			icon = "⚠"
		case AlertError:
			bg = kit.Colors.ErrorLight
			borderColor = kit.Colors.Error
			iconColor = kit.Colors.Error
			icon = "✗"
		}

		return widget.Border{
			Color:        borderColor,
			CornerRadius: RadiusMedium,
			Width:        unit.Dp(1),
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			defer clip.RRect{
				Rect: image.Rectangle{Max: gtx.Constraints.Max},
				NW:   int(RadiusMedium), NE: int(RadiusMedium),
				SE: int(RadiusMedium), SW: int(RadiusMedium),
			}.Push(gtx.Ops).Pop()

			paint.Fill(gtx.Ops, bg)

			return layout.UniformInset(kit.Spacing.Medium).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Alignment: layout.Start}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						iconLabel := material.Label(kit.Theme, unit.Sp(20), icon)
						iconLabel.Color = iconColor
						return iconLabel.Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Spacer{Width: kit.Spacing.Medium}.Layout(gtx)
					}),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								if title == "" {
									return layout.Dimensions{}
								}
								titleLabel := material.Label(kit.Theme, kit.Typography.LabelMedium.Size, title)
								titleLabel.Color = kit.Colors.OnSurface
								return titleLabel.Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								if title != "" {
									return layout.Spacer{Height: kit.Spacing.Tiny}.Layout(gtx)
								}
								return layout.Dimensions{}
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								messageLabel := material.Label(kit.Theme, kit.Typography.BodyMedium.Size, message)
								messageLabel.Color = kit.Colors.OnSurface
								return messageLabel.Layout(gtx)
							}),
						)
					}),
				)
			})
		})
	}
}

// Progress bar component
func (kit *UIKit) ProgressBar(progress float32) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		// Clamp progress between 0 and 1
		if progress < 0 {
			progress = 0
		}
		if progress > 1 {
			progress = 1
		}

		return widget.Border{
			Color:        kit.Colors.BorderLight,
			CornerRadius: unit.Dp(4),
			Width:        unit.Dp(0),
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			// Background
			defer clip.RRect{
				Rect: image.Rectangle{Max: gtx.Constraints.Max},
				NW:   4, NE: 4, SE: 4, SW: 4,
			}.Push(gtx.Ops).Pop()

			paint.Fill(gtx.Ops, kit.Colors.Gray200)

			// Progress fill
			progressWidth := int(float32(gtx.Constraints.Max.X) * progress)
			if progressWidth > 0 {
				defer clip.RRect{
					Rect: image.Rectangle{Max: image.Point{X: progressWidth, Y: gtx.Constraints.Max.Y}},
					NW:   4, NE: 4, SE: 4, SW: 4,
				}.Push(gtx.Ops).Pop()

				paint.Fill(gtx.Ops, kit.Colors.Primary500)
			}

			return layout.Dimensions{
				Size: image.Point{X: gtx.Constraints.Max.X, Y: int(kit.Spacing.Small)},
			}
		})
	}
}

// Helper function to create consistent spacing
func (kit *UIKit) Space(size unit.Dp) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Spacer{Width: size, Height: size}.Layout(gtx)
	}
}

// Helper function for consistent text styles
func (kit *UIKit) Text(text string, style TypographyStyle, color color.NRGBA) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		label := material.Label(kit.Theme, style.Size, text)
		label.Color = color
		return label.Layout(gtx)
	}
}
