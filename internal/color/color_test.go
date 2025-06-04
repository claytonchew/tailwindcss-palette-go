package color

import (
	"testing"
)

func TestHexToRGB(t *testing.T) {
	tests := []struct {
		name     string
		hex      string
		wantR    uint8
		wantG    uint8
		wantB    uint8
		wantErr  bool
		errValue error
	}{
		{
			name:    "Valid 6 character hex",
			hex:     "#FF5733",
			wantR:   255,
			wantG:   87,
			wantB:   51,
			wantErr: false,
		},
		{
			name:    "Valid 6 character hex but with lowercase",
			hex:     "#ff5733",
			wantR:   255,
			wantG:   87,
			wantB:   51,
			wantErr: false,
		},
		{
			name:    "Valid 6 character hex without #",
			hex:     "FF5733",
			wantR:   255,
			wantG:   87,
			wantB:   51,
			wantErr: false,
		},
		{
			name:    "Valid 3 character hex",
			hex:     "#F53",
			wantR:   255,
			wantG:   85,
			wantB:   51,
			wantErr: false,
		},
		{
			name:    "Valid 3 character hex but with lowercase",
			hex:     "#f53",
			wantR:   255,
			wantG:   85,
			wantB:   51,
			wantErr: false,
		},
		{
			name:    "Valid 3 character hex without #",
			hex:     "F53",
			wantR:   255,
			wantG:   85,
			wantB:   51,
			wantErr: false,
		},
		{
			name:     "Invalid hex length",
			hex:      "FF57",
			wantErr:  true,
			errValue: ErrorInvalidHexFormat,
		},
		{
			name:    "Invalid hex characters",
			hex:     "FFZZ33",
			wantErr: true,
		},
		{
			name:    "Black color",
			hex:     "000000",
			wantR:   0,
			wantG:   0,
			wantB:   0,
			wantErr: false,
		},
		{
			name:    "White color",
			hex:     "FFFFFF",
			wantR:   255,
			wantG:   255,
			wantB:   255,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, g, b, err := HexToRGB(tt.hex)

			if (err != nil) != tt.wantErr {
				t.Errorf("HexToRGB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && tt.errValue != nil && err != tt.errValue {
				t.Errorf("HexToRGB() error = %v, wantErr %v", err, tt.errValue)
				return
			}

			if !tt.wantErr && (r != tt.wantR || g != tt.wantG || b != tt.wantB) {
				t.Errorf("HexToRGB() = (%v, %v, %v), want (%v, %v, %v)",
					r, g, b, tt.wantR, tt.wantG, tt.wantB)
			}
		})
	}
}

func TestHSLToHex(t *testing.T) {
	tests := []struct {
		name    string
		h       float64
		s       float64
		l       float64
		want    string
		wantErr bool
	}{
		{
			name:    "Red",
			h:       0,
			s:       1,
			l:       0.5,
			want:    "#FF0000",
			wantErr: false,
		},
		{
			name:    "Green",
			h:       120,
			s:       1,
			l:       0.5,
			want:    "#00FF00",
			wantErr: false,
		},
		{
			name:    "Blue",
			h:       240,
			s:       1,
			l:       0.5,
			want:    "#0000FF",
			wantErr: false,
		},
		{
			name:    "Black",
			h:       0,
			s:       0,
			l:       0,
			want:    "#000000",
			wantErr: false,
		},
		{
			name:    "White",
			h:       0,
			s:       0,
			l:       1,
			want:    "#FFFFFF",
			wantErr: false,
		},
		{
			name:    "Gray",
			h:       0,
			s:       0,
			l:       0.5,
			want:    "#7F7F7F",
			wantErr: false,
		},
		{
			name:    "Invalid saturation",
			h:       120,
			s:       1.5,
			l:       0.5,
			wantErr: true,
		},
		{
			name:    "Invalid lightness",
			h:       120,
			s:       0.5,
			l:       1.5,
			wantErr: true,
		},
		{
			name:    "Invalid hue",
			h:       400,
			s:       0.5,
			l:       0.5,
			wantErr: true,
		},
		{
			name:    "Negative hue",
			h:       -10,
			s:       0.5,
			l:       0.5,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HSLToHex(tt.h, tt.s, tt.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("HSLToHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("HSLToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexToHSL(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		wantH   float64
		wantS   float64
		wantL   float64
		wantErr bool
	}{
		{
			name:    "Red",
			hex:     "#FF0000",
			wantH:   0,
			wantS:   1,
			wantL:   0.5,
			wantErr: false,
		},
		{
			name:    "Green",
			hex:     "#00FF00",
			wantH:   120,
			wantS:   1,
			wantL:   0.5,
			wantErr: false,
		},
		{
			name:    "Blue",
			hex:     "#0000FF",
			wantH:   240,
			wantS:   1,
			wantL:   0.5,
			wantErr: false,
		},
		{
			name:    "Black",
			hex:     "#000000",
			wantH:   0,
			wantS:   0,
			wantL:   0,
			wantErr: false,
		},
		{
			name:    "White",
			hex:     "#FFFFFF",
			wantH:   0,
			wantS:   0,
			wantL:   1,
			wantErr: false,
		},
		{
			name:    "Gray",
			hex:     "#808080",
			wantH:   0,
			wantS:   0,
			wantL:   0.5,
			wantErr: false,
		},
		{
			name:    "Invalid hex",
			hex:     "#ZZ00FF",
			wantErr: true,
		},
		{
			name:    "Invalid length",
			hex:     "#1234",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, s, l, err := HexToHSL(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToHSL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if h != tt.wantH || s != tt.wantS || l != tt.wantL {
					t.Errorf("HexToHSL() = (%v, %v, %v), want (%v, %v, %v)",
						h, s, l, tt.wantH, tt.wantS, tt.wantL)
				}
			}
		})
	}
}

func TestRoundtrip(t *testing.T) {
	tests := []struct {
		name string
		hex  string
	}{
		{
			name: "Red",
			hex:  "#FF0000",
		},
		{
			name: "Green",
			hex:  "#00FF00",
		},
		{
			name: "Blue",
			hex:  "#0000FF",
		},
		{
			name: "Black",
			hex:  "#000000",
		},
		{
			name: "White",
			hex:  "#FFFFFF",
		},
		{
			name: "Custom color",
			hex:  "#1A2B3C",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, s, l, err := HexToHSL(tt.hex)
			if err != nil {
				t.Errorf("HexToHSL() error = %v", err)
				return
			}

			hex, err := HSLToHex(h, s, l)
			if err != nil {
				t.Errorf("HSLToHex() error = %v", err)
				return
			}

			r1, g1, b1, _ := HexToRGB(tt.hex)
			r2, g2, b2, _ := HexToRGB(hex)

			if abs(int(r1)-int(r2)) > 1 || abs(int(g1)-int(g2)) > 1 || abs(int(b1)-int(b2)) > 1 {
				t.Errorf("Roundtrip conversion failed: original = %v (%v,%v,%v), result = %v (%v,%v,%v)",
					tt.hex, r1, g1, b1, hex, r2, g2, b2)
			}
		})
	}
}

func TestHex2uint8(t *testing.T) {
	tests := []struct {
		name    string
		hexStr  string
		want    uint8
		wantErr bool
	}{
		{
			name:   "Valid hex 00",
			hexStr: "00",
			want:   0,
		},
		{
			name:   "Valid hex FF",
			hexStr: "FF",
			want:   255,
		},
		{
			name:   "Valid hex 80",
			hexStr: "80",
			want:   128,
		},
		{
			name:    "Invalid hex",
			hexStr:  "ZZ",
			wantErr: true,
		},
		{
			name:    "Too large hex",
			hexStr:  "FFF",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hex2uint8(tt.hexStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("hex2uint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("hex2uint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
