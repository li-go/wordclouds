package wordclouds

import (
	"image/color"
	"io/fs"
)

type Options struct {
	FontMaxSize     int
	FontMinSize     int
	RandomPlacement bool
	FontFile        fs.File
	Colors          []color.Color
	BackgroundColor color.Color
	Width           int
	Height          int
	Mask            []*Box
	SizeFunction    sizeFunction
	Debug           bool
}

var defaultOptions = Options{
	FontMaxSize:     500,
	FontMinSize:     10,
	RandomPlacement: false,
	FontFile:        nil,
	Colors:          []color.Color{color.RGBA{}},
	BackgroundColor: color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
	Width:           2048,
	Height:          2048,
	Mask:            make([]*Box, 0),
	SizeFunction:    sizeLinear,
	Debug:           false,
}

type Option func(*Options)

// Font file
func FontFile(file fs.File) Option {
	return func(options *Options) {
		options.FontFile = file
	}
}

// Output file background color
func BackgroundColor(color color.Color) Option {
	return func(options *Options) {
		options.BackgroundColor = color
	}
}

// Colors to use for the words
func Colors(colors []color.Color) Option {
	return func(options *Options) {
		options.Colors = colors
	}
}

// Max font size
func FontMaxSize(max int) Option {
	return func(options *Options) {
		options.FontMaxSize = max
	}
}

// Min font size
func FontMinSize(min int) Option {
	return func(options *Options) {
		options.FontMinSize = min
	}
}

// A list of bounding boxes where words can not be placed.
// See Mask
func MaskBoxes(mask []*Box) Option {
	return func(options *Options) {
		options.Mask = mask
	}
}

func Width(w int) Option {
	return func(options *Options) {
		options.Width = w
	}
}

func Height(h int) Option {
	return func(options *Options) {
		options.Height = h
	}
}

// Place words randomly
func RandomPlacement(do bool) Option {
	return func(options *Options) {
		options.RandomPlacement = do
	}
}

// Set word font sizing function
func WordSizeFunction(f string) Option {
	return func(options *Options) {
		switch f {
		case SizeFunctionLinear:
			options.SizeFunction = sizeLinear
		case SizeFunctionSqrt:
			options.SizeFunction = sizeSqrt
		case SizeFunctionSqrtInverse:
			options.SizeFunction = sizeSqrtInverse
		default:
			panic("No such size function " + f)
		}
	}
}

// Draw bounding boxes around words
func Debug() Option {
	return func(options *Options) {
		options.Debug = true
	}
}
