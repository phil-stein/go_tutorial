package main

import "fmt"
import "time"
import "net/url"
import "math/rand"
import "image/color"
import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/storage"
    "fyne.io/fyne/v2/container"
)


type Right_Panel_Type int
const (
  Right_Panel_Widgets = iota
  Right_Panel_Draw
  Right_Panel_Max
)

type data_t struct{
  right_panel_type Right_Panel_Type
  window fyne.Window
  // app    app.fyneApp
}


func main() {

  var data data_t
  data.right_panel_type = Right_Panel_Widgets

  // create app and window
  app         := app.New()
  data.window = app.NewWindow("Hello")

  // resize the window
  data.window.Resize( fyne.NewSize( 1000, 550 ) )
  

  // cont_00 := container.NewVBox( widget.NewLabel( "cock" ) )
  // fmt.Printf( "cont_00 T: %T\n", cont_00 )
  // fmt.Printf( "a T: %T\n", a )
  // fmt.Printf( "w T: %T\n", w )


  var right_panel *container.Scroll
  switch data.right_panel_type {
    case Right_Panel_Widgets: right_panel = right_panel_widgets( &data )
    case Right_Panel_Draw: right_panel = right_panel_draw( &data )
    default: fmt.Println( "[ERROR] data.right_panel_type is wrong:", data.right_panel_type )
             return
  }
  data.window.SetContent( container.NewHBox( left_panel( &data ), right_panel ) )

  data.window.Show()

  app.Run()
}

func left_panel( data *data_t ) *fyne.Container {
  // create label widget
  label_00 := widget.NewLabel( "category:" )

  // create new button widget
  btn_00 := widget.NewButton(
      "widgets",
      func() { 
        (*data).right_panel_type = Right_Panel_Widgets 
        var right_panel *container.Scroll
        switch data.right_panel_type {
          case Right_Panel_Widgets: right_panel = right_panel_widgets( data )
          case Right_Panel_Draw: right_panel = right_panel_draw( data )
          default: fmt.Println( "[ERROR] data.right_panel_type is wrong:", data.right_panel_type )
                   return
        }
        (*data).window.SetContent( container.NewHBox( left_panel( data ), right_panel ) )
      },
  )
  btn_01 := widget.NewButton(
      "draw",
      func() { 
        (*data).right_panel_type = Right_Panel_Draw 
        var right_panel *container.Scroll
        switch data.right_panel_type {
          case Right_Panel_Widgets: right_panel = right_panel_widgets( data )
          case Right_Panel_Draw: right_panel = right_panel_draw( data )
          default: fmt.Println( "[ERROR] data.right_panel_type is wrong:", data.right_panel_type )
                   return
        }
        (*data).window.SetContent( container.NewHBox( left_panel( data ), right_panel ) )
      },
  )

  return container.NewVBox( label_00, btn_00, btn_01 )
}

func right_panel_widgets( data *data_t ) *container.Scroll {
  // create label widget
  label_00 := widget.NewLabel( "label text" )

  // create new button widget
  btn_00 := widget.NewButton(
      "Click me",
      func() { println("Hello!") },
  )
  // create new checkbox widget
  checkbox_00 := widget.NewCheck( "checkbox", func(b bool) {
    fmt.Println( "checkbox:", b )
  })
 
  // --- start of entry_box ---
  entry_00 := widget.NewEntry()
  entry_00.SetPlaceHolder("Enter text...")

  // create new button widget
  btn_01 := widget.NewButton(
      "print",
      func() { fmt.Println( "entry_00.Text:", entry_00.Text ) },
  )

  entry_box := container.NewHBox( entry_00, btn_01 )
  // --- end of entry_box ---
	
  radio_00 := widget.NewRadioGroup([]string{"Option 1", "Option 2"}, func(value string) {
		fmt.Println("Radio set to", value)
	})
	combo_00 := widget.NewSelect([]string{"Option 1", "Option 2"}, func(value string) {
		fmt.Println("Select set to", value)
	})

  // --- start of form ---
	entry := widget.NewEntry()
	textArea := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Entry", Widget: entry}},
		OnSubmit: func() { // optional, handle form submission
			fmt.Println("Form submitted:", entry.Text)
			fmt.Println("multiline:", textArea.Text)
			// myWindow.Close()
		},
	}
	// we can also append items
	form.Append("Text", textArea)
  // --- end of form ---

  // --- start progress bar ---
	progress_00 := widget.NewProgressBar()
	infinite_00 := widget.NewProgressBarInfinite()

	// go func() {
	// 	for i := 0.0; i <= 1.0; i += 0.1 {
	// 		time.Sleep(time.Millisecond * 250)
	// 		progress_00.SetValue(i)
	// 	}
	// }()
  // --- end progress bar ---

  // --- start toolbar ---
	toolbar_00 := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			fmt.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			fmt.Println("Display help")
		}),
	)
  // --- end toolbar ---

  // icons
	icon_home_00 := widget.NewIcon(theme.HomeIcon())
	icon_info_00 := widget.NewIcon(theme.InfoIcon())

  // hyperlinks / urls
	link, err := url.Parse("https://fyne.io/")
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}
  hyperlink_00 := widget.NewHyperlink("Fyne website", link)


  // --- start rich text ---
  text_00 := widget.NewRichTextFromMarkdown(`
# Rich Text

This paragraph has some *emphasis* and some **strong** words,
as well as a [link](https://fyne.io/).

* it can show lists
* and ` + "`code`" + ` as well`)
	text_00.Wrapping = fyne.TextWrapWord
  // --- end rich text ---

  // --- start grid text ---
	grid_00 := widget.NewTextGridFromString("Hello there\nThis is a text grid")
	grid_00.ShowLineNumbers = true

	highlight_00 := &widget.CustomTextGridStyle{FGColor: color.NRGBA{R: 0xcc, A: 0xff}}
	grid_00.SetStyleRange(0, 0, 0, 4, highlight_00)
  // --- end grid text ---

  // card 
  content_00 := widget.NewLabel("Card content can be any widget or container")
	card_00    := widget.NewCard("Card title", "with a subtitle", content_00)

  // --- start accordion ---
	accordion_00 := widget.NewAccordion(
		widget.NewAccordionItem("Details", widget.NewLabel("Information about this item")),
		widget.NewAccordionItem("Settings", widget.NewCheck("Enabled", func(bool) {})),
	)
	accordion_00.Open(0)
  // --- end accordion ---
	
  // seperator
  seperator_00 := container.NewVBox(
		widget.NewLabel("Above the line"),
		widget.NewSeparator(),
		widget.NewLabel("Below the line"),
	)

  // slider
	slider_00 := widget.NewSlider(0, 100)
	slider_00.Step = 5
	slider_00.OnChangeEnded = func(value float64) {
		fmt.Println("Slider set to", value)
	}

  // --- start select entry ---
	combo_01 := widget.NewSelectEntry([]string{"Red", "Green", "Blue"})
	combo_01.PlaceHolder = "Choose or type a colour"
	combo_01.OnChanged = func(value string) {
		fmt.Println("Colour is", value)
	}
  // --- end select entry ---

  // --- start file icons ---
	text_01 := widget.NewFileIcon(storage.NewFileURI("./readme.txt"))
	image_00 := widget.NewFileIcon(storage.NewFileURI("./photo.jpg"))
  file_icons_00 := container.NewHBox(text_01, image_00)
  // --- end file icons ---

  // --- start activity ---
	activity_00 := widget.NewActivity()
	activity_00.Start()

	go func() {
		time.Sleep(time.Second * 60) // pretend to do some work
		activity_00.Stop()
		activity_00.Hide()
	}()
  activity_cont_00 := container.NewCenter(activity_00)
  // --- end activity ---

  // --- start popup ---
	var popUp *widget.PopUp
	content_01 := container.NewVBox(
		widget.NewLabel("Content of the pop up"),
		widget.NewButton("Close", func() {
			popUp.Hide()
		}),
	)
	popUp = widget.NewModalPopUp(content_01, (*data).window.Canvas())
	
	open_btn_00 := widget.NewButton("Show pop up", func() {
		popUp.Show()
	})

  popup_cont_00 := container.NewCenter( open_btn_00 )
  // --- end popup ---

  container_00 := container.NewVBox( label_00, btn_00, checkbox_00, entry_box, radio_00, combo_00, form, progress_00, infinite_00, toolbar_00, icon_home_00, icon_info_00, hyperlink_00, container.NewScroll(text_00), grid_00, card_00, accordion_00, seperator_00, slider_00, combo_01, file_icons_00, activity_cont_00, popup_cont_00 )
  return container.NewVScroll( container_00 )
}

func right_panel_draw( data *data_t ) *container.Scroll {
  // create label widget
  label_00 := widget.NewLabel( "draw widgets" )

  // rectangle
  rect_00 := canvas.NewRectangle( color.White )
  rect_00.Resize( fyne.NewSize( 200, 200 ) )
  rect_00.SetMinSize( fyne.NewSize( 200, 200 ) )

  // text
	text_00 := canvas.NewText("AlignLeftItalic", color.White)
	text_00.Alignment = fyne.TextAlignTrailing
	text_00.TextStyle = fyne.TextStyle{Italic: true}

  // line
	line_00 := canvas.NewLine( color.White )
	line_00.StrokeWidth = 5

  // circle
	circle_00 := canvas.NewCircle( color.White )
	circle_00.StrokeColor = color.Gray{Y: 0x99}
	circle_00.StrokeWidth = 5
  circle_00.Resize( fyne.NewSize( 200, 200 ) )
  circle_00.Refresh()
  // @NOTE: tried to resize but doesnt work idk
  // circle_00.Position2 = fyne.NewPos( circle_00.Position1.X +200, circle_00.Position1.Y +200 )
  // fmt.Println( "circle_00.MinSize():", circle_00.MinSize() )
  // circle_00.SetMinSize( fyne.NewSize( 200, 200 ) )

  // image
	image_00 := canvas.NewImageFromResource(theme.FyneLogo())
	// image_00 := canvas.NewImageFromURI(uri)
	// image_00 := canvas.NewImageFromImage(src)
	// image_00 := canvas.NewImageFromReader(reader, name)
	// image_00 := canvas.NewImageFromFile(fileName)
	image_00.FillMode = canvas.ImageFillOriginal

  // raster
	raster_00 := canvas.NewRasterWithPixels(
		func(_, _, w, h int) color.Color {
			return color.RGBA{R: uint8(rand.Intn(255)),
				G: uint8(rand.Intn(255)),
				B: uint8(rand.Intn(255)), A: 0xff}
		})
  raster_00.SetMinSize( fyne.NewSize( 200, 200 ) )
  
  // gradient
	gradient_00 := canvas.NewHorizontalGradient(color.White, color.Transparent)
  gradient_01 := canvas.NewRadialGradient(color.White, color.Transparent)
  gradient_00.SetMinSize( fyne.NewSize( 200, 200 ) )
  gradient_01.SetMinSize( fyne.NewSize( 200, 200 ) )

  // animation
  // @NOTE: idk doesnt work
	anim_obj_00 := canvas.NewRectangle(color.Black)
	anim_obj_00.Resize(fyne.NewSize( 200, 200))
  anim_obj_00.SetMinSize( fyne.NewSize( 200, 200 ) )
  anim_cont_00 := container.NewWithoutLayout( anim_obj_00 )

	red_00  := color.NRGBA{R:0xff, A:0xff}
	blue_00 := color.NRGBA{B:0xff, A:0xff}
	canvas.NewColorRGBAAnimation(red_00, blue_00, time.Second*2, func(c color.Color) {
    anim_obj_00.FillColor = c
		canvas.Refresh(anim_obj_00)
	}).Start()

  container_00 := container.NewVBox( label_00, rect_00, text_00, line_00, circle_00, image_00, raster_00, gradient_00, gradient_01, anim_cont_00 )
  return container.NewVScroll( container_00 )
}
