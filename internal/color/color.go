package color

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	ErrorInvalidHexFormat = errors.New("invalid hex color format, must be 3 or 6 characters long without prefix")
	ErrorInvalidHSLValues = errors.New("HSL values must be in the range: 0 <= H < 360, 0 <= S <= 1, 0 <= L <= 1")
)

func HSLToHex(h, s, l float64) (string, error) {
	if s < 0 || s > 1 || l < 0 || l > 1 || h < 0 || h >= 360 {
		return "", ErrorInvalidHSLValues
	}

	var r, g, b float64

	if s == 0 {
		r = l * 255
		g = l * 255
		b = l * 255
	} else {
		var q, p float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - (l * s)
		}
		p = 2*l - q

		h /= 360
		r = calculateRGBComponent(p, q, h+1.0/3.0) * 255
		g = calculateRGBComponent(p, q, h) * 255
		b = calculateRGBComponent(p, q, h-1.0/3.0) * 255
	}

	return "#" + strings.ToUpper(fmt.Sprintf("%02X%02X%02X", uint(r), uint(g), uint(b))), nil
}

func HexToHSL(hex string) (h, s, l float64, err error) {
	r, g, b, err := HexToRGB(hex)
	if err != nil {
		return 0, 0, 0, err
	}

	Rnot := float64(r) / 255.0
	Gnot := float64(g) / 255.0
	Bnot := float64(b) / 255.0
	max, min := max(Rnot, Gnot, Bnot), min(Rnot, Gnot, Bnot)
	alpha := max - min
	l = (max + min) / 2.0
	if alpha == 0 {
		h = 0
		s = 0
	} else {
		switch max {
		case Rnot:
			h = 60 * (math.Mod((Gnot-Bnot)/alpha, 6))
		case Gnot:
			h = 60 * ((Bnot-Rnot)/alpha + 2)
		case Bnot:
			h = 60 * ((Rnot-Gnot)/alpha + 4)
		}
		if h < 0 {
			h += 360
		}

		s = alpha / (1 - math.Abs(2*l-1))
	}

	return h, round(s), round(l), nil
}

func HexToRGB(hex string) (r, g, b uint8, err error) {
	hex = strings.Replace(hex, "0x", "", -1)
	hex = strings.Replace(hex, "#", "", -1)
	if (utf8.RuneCountInString(hex) != 6) && (utf8.RuneCountInString(hex) != 3) {
		return 0, 0, 0, ErrorInvalidHexFormat
	}

	if utf8.RuneCountInString(hex) == 3 {
		hex = string([]rune{rune(hex[0]), rune(hex[0]), rune(hex[1]), rune(hex[1]), rune(hex[2]), rune(hex[2])})
	}

	r, err = hex2uint8(hex[0:2])
	if err != nil {
		return 0, 0, 0, err
	}
	g, err = hex2uint8(hex[2:4])
	if err != nil {
		return 0, 0, 0, err
	}
	b, err = hex2uint8(hex[4:6])
	if err != nil {
		return 0, 0, 0, err
	}
	return r, g, b, nil
}

func calculateRGBComponent(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6
	}
	return p
}

func hex2uint8(hexStr string) (uint8, error) {
	result, err := strconv.ParseUint(hexStr, 16, 8)
	if err != nil {
		return 0, err
	}
	return uint8(result), nil
}

func max(a, b, c float64) float64 {
	return math.Max(math.Max(a, b), c)
}

func min(a, b, c float64) float64 {
	return math.Min(math.Min(a, b), c)
}

func round(x float64) float64 {
	return math.Round(x*100) / 100
}

func RGBToHex(r, g, b uint8) (string, error) {
	return "#" + strings.ToUpper(fmt.Sprintf("%02X%02X%02X", r, g, b)), nil
}
