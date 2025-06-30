package main

import (
	"fmt"
	"gemini-files/cmd/ui-kit/uikit"
	"image"
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type App struct {
	window *app.Window
	kit    *uikit.UIKit

	// Form fields
	nameEditor    widget.Editor
	emailEditor   widget.Editor
	messageEditor widget.Editor
	slider        widget.Float

	// Buttons
	primaryBtn   widget.Clickable
	secondaryBtn widget.Clickable
	outlineBtn   widget.Clickable
	ghostBtn     widget.Clickable
	dangerBtn    widget.Clickable
	successBtn   widget.Clickable

	// Checkboxes
	checkbox1 widget.Bool
	checkbox2 widget.Bool
	checkbox3 widget.Bool

	// Interactive elements
	submitBtn widget.Clickable
	resetBtn  widget.Clickable

	// State
	progress         float32
	notification     string
	notificationType uikit.AlertVariant
	showNotification bool
	formSubmitted    bool
	selectedTab      int

	// Animation
	animationStart time.Time
	lastFrame      time.Time
}

func NewApp() *App {
	app := &App{
		kit:            uikit.NewUIKit(),
		progress:       0.0,
		animationStart: time.Now(),
		lastFrame:      time.Now(),
		slider:         widget.Float{Value: 0.5},
		selectedTab:    0,
	}

	// Set up initial editor content
	app.nameEditor.SetText("John Doe")
	app.emailEditor.SetText("john@example.com")
	app.messageEditor.SetText("This is a sample message to demonstrate the multi-line text editor component.")

	return app
}

func (a *App) handleEvents(gtx layout.Context) {
	// Handle button clicks
	if a.primaryBtn.Clicked(gtx) {
		a.showNotification = true
		a.notification = "Primary button clicked!"
		a.notificationType = uikit.AlertInfo
	}

	if a.secondaryBtn.Clicked(gtx) {
		a.showNotification = true
		a.notification = "Secondary action performed"
		a.notificationType = uikit.AlertSuccess
	}

	if a.outlineBtn.Clicked(gtx) {
		a.showNotification = true
		a.notification = "Outline button pressed"
		a.notificationType = uikit.AlertWarning
	}

	if a.dangerBtn.Clicked(gtx) {
		a.showNotification = true
		a.notification = "Danger! This is a destructive action"
		a.notificationType = uikit.AlertError
	}

	if a.successBtn.Clicked(gtx) {
		a.showNotification = true
		a.notification = "Success! Operation completed"
		a.notificationType = uikit.AlertSuccess
	}

	if a.submitBtn.Clicked(gtx) {
		a.formSubmitted = true
		a.showNotification = true
		a.notification = "Form submitted successfully!"
		a.notificationType = uikit.AlertSuccess
		a.progress = 1.0
	}

	if a.resetBtn.Clicked(gtx) {
		a.nameEditor.SetText("")
		a.emailEditor.SetText("")
		a.messageEditor.SetText("")
		a.formSubmitted = false
		a.progress = 0.0
		a.showNotification = true
		a.notification = "Form reset"
		a.notificationType = uikit.AlertInfo
		a.checkbox1.Value = false
		a.checkbox2.Value = false
		a.checkbox3.Value = false
	}

	// Handle tab clicks
	for i := 0; i < 3; i++ {
		if a.selectedTab != i {
			// Create temporary buttons for tab handling
			btn := widget.Clickable{}
			if btn.Clicked(gtx) {
				a.selectedTab = i
			}
		}
	}

	// Animate progress bar
	now := time.Now()
	if now.Sub(a.lastFrame) > time.Millisecond*50 {
		if !a.formSubmitted && a.progress < 0.8 {
			a.progress += 0.002
		}
		a.lastFrame = now
	}
}

func (a *App) Layout(gtx layout.Context) layout.Dimensions {
	a.handleEvents(gtx)

	// Use all available space
	gtx.Constraints.Min = gtx.Constraints.Max

	// Main layout with header, tabs, and content
	return layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.Spacing(0), // No spacing between header, tabs, and content
	}.Layout(gtx,
		// Header section with fixed height
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints = layout.Exact(image.Pt(gtx.Constraints.Max.X, gtx.Dp(80)))
			return a.renderHeader(gtx)
		}),
		// Tabs section with fixed height
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints = layout.Exact(image.Pt(gtx.Constraints.Max.X, gtx.Dp(48)))
			return a.renderTabs(gtx)
		}),
		// Content area that takes remaining space
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return widget.Border{
				Color:        a.kit.Colors.Border,
				Width:        unit.Dp(1),
				CornerRadius: a.kit.Spacing.Small,
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				// Add padding inside the content area
				gtx.Constraints.Min = gtx.Constraints.Constrain(gtx.Constraints.Min)
				return a.renderCurrentTab(gtx)
			})
		}),
	)
}

func (a *App) renderCurrentTab(gtx layout.Context) layout.Dimensions {
	// Add uniform padding around the content
	return layout.UniformInset(a.kit.Spacing.Medium).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		// Use a list for scrollable content
		list := &widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		}

		return list.Layout(gtx, 1, func(gtx layout.Context, _ int) layout.Dimensions {
			switch a.selectedTab {
			case 0:
				return a.renderComponentsTab(gtx)
			case 1:
				return a.renderFormTab(gtx)
			case 2:
				return a.renderSettingsTab(gtx)
			default:
				return a.renderComponentsTab(gtx)
			}
		})
	})
}

func (a *App) renderHeader(gtx layout.Context) layout.Dimensions {
	return a.kit.Card(gtx, func(gtx layout.Context) layout.Dimensions {
		// Ensure card respects constrained height
		gtx.Constraints.Min = gtx.Constraints.Max

		return layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.Middle,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return a.kit.Text("UI Kit Demo", a.kit.Typography.DisplaySmall, a.kit.Colors.TextPrimary)(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return a.kit.Space(a.kit.Spacing.Small)(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return a.kit.Text("A comprehensive design system",
					a.kit.Typography.BodyMedium, a.kit.Colors.TextSecondary)(gtx)
			}),
		)
	})
}

func (a *App) renderTabs(gtx layout.Context) layout.Dimensions {
	tabs := []string{"Components", "Form", "Settings"}

	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Top: a.kit.Spacing.Medium}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := widget.Clickable{}
						if a.selectedTab == 0 {
							return a.kit.Button(&btn, tabs[0], uikit.ButtonPrimary, uikit.ButtonMedium)(gtx)
						}
						return a.kit.Button(&btn, tabs[0], uikit.ButtonOutline, uikit.ButtonMedium)(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := widget.Clickable{}
						if a.selectedTab == 1 {
							return a.kit.Button(&btn, tabs[1], uikit.ButtonPrimary, uikit.ButtonMedium)(gtx)
						}
						return a.kit.Button(&btn, tabs[1], uikit.ButtonOutline, uikit.ButtonMedium)(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := widget.Clickable{}
						if a.selectedTab == 2 {
							return a.kit.Button(&btn, tabs[2], uikit.ButtonPrimary, uikit.ButtonMedium)(gtx)
						}
						return a.kit.Button(&btn, tabs[2], uikit.ButtonOutline, uikit.ButtonMedium)(gtx)
					}),
				)
			})
		}),
	)
}

func (a *App) renderComponentsTab(gtx layout.Context) layout.Dimensions {
	// Use a flex layout with consistent spacing between sections
	return layout.Flex{
		Axis:      layout.Vertical,
		Spacing:   layout.Spacing(a.kit.Spacing.Medium),
		Alignment: layout.Start,
	}.Layout(gtx,
		// Notification section - only show if there's a notification
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if !a.showNotification {
				return layout.Dimensions{}
			}
			return a.renderNotificationSection(gtx)
		}),
		// Typography section
		layout.Rigid(a.renderTypographySection),
		// Buttons section
		layout.Rigid(a.renderButtonSection),
		// Progress section
		layout.Rigid(a.renderProgressSection),
		// Add some space at the bottom if needed
		layout.Rigid(layout.Spacer{Height: unit.Dp(16)}.Layout),
	)
}

func (a *App) renderNotificationSection(gtx layout.Context) layout.Dimensions {
	if !a.showNotification {
		return layout.Dimensions{}
	}

	return a.kit.Alert("Notification", a.notification, a.notificationType)(gtx)
}

func (a *App) renderTypographySection(gtx layout.Context) layout.Dimensions {
	return a.kit.Card(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return a.kit.Text("Typography", a.kit.Typography.HeadlineSmall, a.kit.Colors.TextPrimary)(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return a.kit.Space(a.kit.Spacing.Large)(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return a.kit.Text("Display Small", a.kit.Typography.DisplaySmall, a.kit.Colors.TextPrimary)(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return a.kit.Space(a.kit.Spacing.Medium)(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return a.kit.Text("Headline Small", a.kit.Typography.HeadlineSmall, a.kit.Colors.TextPrimary)(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return a.kit.Space(a.kit.Spacing.Medium)(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return a.kit.Text("Body Medium - Standard text for most content", a.kit.Typography.BodyMedium, a.kit.Colors.TextPrimary)(gtx)
			}),
		)
	})
}

func (a *App) renderButtonSection(gtx layout.Context) layout.Dimensions {
	return a.kit.Card(gtx, func(gtx layout.Context) layout.Dimensions {
		// Ensure card has minimum height
		gtx.Constraints.Min.Y = gtx.Dp(100)

		return layout.Flex{
			Axis:      layout.Vertical,
			Spacing:   layout.Spacing(a.kit.Spacing.Medium),
			Alignment: layout.Start,
		}.Layout(gtx,
			// Section title
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return a.kit.Text("Buttons", a.kit.Typography.HeadlineSmall, a.kit.Colors.TextPrimary)(gtx)
			}),

			// First row of buttons
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:      layout.Horizontal,
					Spacing:   layout.Spacing(a.kit.Spacing.Medium),
					Alignment:  layout.Start,
				}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return a.kit.Button(&a.primaryBtn, "Primary", uikit.ButtonPrimary, uikit.ButtonMedium)(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return a.kit.Button(&a.secondaryBtn, "Secondary", uikit.ButtonSecondary, uikit.ButtonMedium)(gtx)
					}),
				)
			}),

			// Second row of buttons
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:      layout.Horizontal,
					Spacing:   layout.Spacing(a.kit.Spacing.Medium),
					Alignment:  layout.Start,
				}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return a.kit.Button(&a.outlineBtn, "Outline", uikit.ButtonOutline, uikit.ButtonMedium)(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return a.kit.Button(&a.ghostBtn, "Ghost", uikit.ButtonGhost, uikit.ButtonMedium)(gtx)
					}),
				)
			}),

			// Third row of buttons (status variants)
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:      layout.Horizontal,
					Spacing:   layout.Spacing(a.kit.Spacing.Medium),
					Alignment:  layout.Start,
				}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return a.kit.Button(&a.dangerBtn, "Danger", uikit.ButtonDanger, uikit.ButtonMedium)(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return a.kit.Button(&a.successBtn, "Success", uikit.ButtonSuccess, uikit.ButtonMedium)(gtx)
					}),
				)
			}),
		)
	})
}

func (a *App) renderProgressSection(gtx layout.Context) layout.Dimensions {
	return a.kit.Card(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return a.kit.Text("Progress", a.kit.Typography.HeadlineSmall, a.kit.Colors.TextPrimary)(gtx)
			}),
			layout.Rigid(a.kit.Space(a.kit.Spacing.Medium)),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return a.kit.ProgressBar(a.progress)(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return a.kit.Space(a.kit.Spacing.Small)(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				progressText := fmt.Sprintf("%.0f%%", a.progress*100)
				return a.kit.Text(progressText, a.kit.Typography.BodyMedium, a.kit.Colors.TextSecondary)(gtx)
			}),
		)
	})
}

func (a *App) renderFormTab(gtx layout.Context) layout.Dimensions {
	return a.kit.Card(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return a.kit.Text("Contact Form", a.kit.Typography.HeadlineSmall, a.kit.Colors.TextPrimary)(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return a.kit.Space(a.kit.Spacing.Small)(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return a.kit.Text("Fill out the form below to get in touch", a.kit.Typography.BodyMedium, a.kit.Colors.TextSecondary)(gtx)
					}),
				)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Top: a.kit.Spacing.Medium, Bottom: a.kit.Spacing.Medium}.Layout(gtx, a.kit.Divider())
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceBetween}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return a.kit.Text("Name", a.kit.Typography.LabelMedium, a.kit.Colors.TextPrimary)(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return a.kit.Space(a.kit.Spacing.Tiny)(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return a.kit.Input(&a.nameEditor, "Enter your full name", false)(gtx)
							}),
						)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Top: a.kit.Spacing.Medium}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							hasError := len(a.emailEditor.Text()) > 0 && !contains(a.emailEditor.Text(), "@")
							return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return a.kit.Text("Email", a.kit.Typography.LabelMedium, a.kit.Colors.TextPrimary)(gtx)
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return a.kit.Space(a.kit.Spacing.Tiny)(gtx)
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return a.kit.Input(&a.emailEditor, "your.email@example.com", hasError)(gtx)
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									if hasError {
										return layout.Inset{Top: a.kit.Spacing.Tiny}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
											return a.kit.Text("Please enter a valid email address", a.kit.Typography.LabelSmall, a.kit.Colors.Error)(gtx)
										})
									}
									return layout.Dimensions{}
								}),
							)
						})
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Top: a.kit.Spacing.Medium}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return a.kit.Text("Message", a.kit.Typography.LabelMedium, a.kit.Colors.TextPrimary)(gtx)
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return a.kit.Space(a.kit.Spacing.Tiny)(gtx)
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									// Make message field taller
									gtx.Constraints.Min.Y = gtx.Dp(100)
									gtx.Constraints.Max.Y = gtx.Dp(150)
									return a.kit.Input(&a.messageEditor, "Type your message here...", false)(gtx)
								}),
							)
						})
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Top: a.kit.Spacing.Large}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween, Alignment: layout.Middle}.Layout(gtx,
								layout.Flexed(0.48, a.kit.Button(&a.resetBtn, "Clear Form", uikit.ButtonOutline, uikit.ButtonMedium)),
								layout.Rigid(a.kit.Space(a.kit.Spacing.Medium)),
								layout.Flexed(0.48, a.kit.Button(&a.submitBtn, "Send Message", uikit.ButtonPrimary, uikit.ButtonMedium)),
							)
						})
					}),
				)
			}),
		)
	})
}

func (a *App) renderSettingsTab(gtx layout.Context) layout.Dimensions {
	return a.kit.Card(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return a.kit.Text("Settings", a.kit.Typography.HeadlineSmall, a.kit.Colors.TextPrimary)(gtx)
			}),
			layout.Rigid(a.kit.Space(a.kit.Spacing.Medium)),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return a.renderCheckboxSection(gtx)
			}),
			layout.Rigid(a.kit.Space(a.kit.Spacing.Medium)),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return a.renderSliderSection(gtx)
			}),
		)
	})
}

func (a *App) renderCheckboxSection(gtx layout.Context) layout.Dimensions {
	return a.kit.Card(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return a.kit.Text("Checkbox Options", a.kit.Typography.TitleMedium, a.kit.Colors.TextPrimary)(gtx)
			}),
			layout.Rigid(a.kit.Space(a.kit.Spacing.Small)),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.CheckBox(a.kit.Theme, &a.checkbox1, "Option 1").Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.CheckBox(a.kit.Theme, &a.checkbox2, "Option 2").Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.CheckBox(a.kit.Theme, &a.checkbox3, "Option 3").Layout(gtx)
			}),
		)
	})
}

func (a *App) renderSliderSection(gtx layout.Context) layout.Dimensions {
	return a.kit.Card(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return a.kit.Text("Slider Control", a.kit.Typography.TitleMedium, a.kit.Colors.TextPrimary)(gtx)
			}),
			layout.Rigid(a.kit.Space(a.kit.Spacing.Small)),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.Slider(a.kit.Theme, &a.slider).Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				value := fmt.Sprintf("Value: %.2f", a.slider.Value)
				return a.kit.Text(value, a.kit.Typography.BodyMedium, a.kit.Colors.TextSecondary)(gtx)
			}),
		)
	})
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func loop(w *app.Window, a *App) error {
	var ops op.Ops

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			a.Layout(gtx)
			e.Frame(gtx.Ops)
		}
	}
}

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("UI Kit Demo - Complete Design System"))
		w.Option(app.Size(unit.Dp(900), unit.Dp(700)))
		a := NewApp()
		a.window = w
		if err := loop(w, a); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
