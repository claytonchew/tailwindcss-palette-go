package generator

import (
	"errors"

	"github.com/claytonchew/tailwindcss-palette-go/internal/color"
)

type Shade struct {
	name      string
	lightness uint8
}

func NewShade(name string, lightness uint8) Shade {
	return Shade{
		name:      name,
		lightness: lightness,
	}
}

type Options struct {
	shades []Shade
}

func NewOptions(shades []Shade) Options {
	return Options{
		shades: shades,
	}
}

func DefaultTailwindOptions() Options {
	return Options{
		shades: []Shade{
			NewShade("50", 98),
			NewShade("100", 95),
			NewShade("200", 90),
			NewShade("300", 82),
			NewShade("400", 64),
			NewShade("500", 46),
			NewShade("600", 33),
			NewShade("700", 24),
			NewShade("800", 14),
			NewShade("900", 7),
			NewShade("950", 4),
		},
	}
}

var (
	ErrorInvalidLightness = errors.New("lightness must be between 0 and 100")
)

func GeneratePaletteFromHex(hex string, opts Options) (palette map[string]string, err error) {
	h, s, _, err := color.HexToHSL(hex)
	if err != nil {
		return nil, err
	}
	palette = make(map[string]string)

	for _, shade := range opts.shades {
		l := float64(shade.lightness) / 100.0
		if l < 0 || l > 1 {
			return nil, ErrorInvalidLightness
		}

		palette[shade.name], err = color.HSLToHex(h, s, l)
		if err != nil {
			return nil, err
		}

	}

	return palette, nil
}
