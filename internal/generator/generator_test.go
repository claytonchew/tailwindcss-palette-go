package generator

import (
	"testing"
)

func TestGeneratePaletteFromHex(t *testing.T) {
	tests := map[string]struct {
		hex     string
		opts    Options
		want    map[string]string
		wantErr bool
	}{
		"Red with basic shades": {
			hex: "#FF0000",
			opts: Options{
				shades: []Shade{
					{name: "100", lightness: 90},
					{name: "500", lightness: 50},
					{name: "900", lightness: 10},
				},
			},
			want: map[string]string{
				"100": "#FECCCC",
				"500": "#FF0000",
				"900": "#330000",
			},
		},
		"Blue with named shades": {
			hex: "#0000FF",
			opts: Options{
				shades: []Shade{
					{name: "light", lightness: 80},
					{name: "medium", lightness: 50},
					{name: "dark", lightness: 20},
				},
			},
			want: map[string]string{
				"light":  "#9999FF",
				"medium": "#0000FF",
				"dark":   "#000066",
			},
		},
		"Default Tailwind scale": {
			hex:  "#0066FF",
			opts: DefaultTailwindOptions(),
			want: map[string]string{
				"50":  "#F4F8FF",
				"100": "#E5EFFF",
				"200": "#CCE0FE",
				"300": "#A3C7FE",
				"400": "#4790FF",
				"500": "#005DEA",
				"600": "#0043A8",
				"700": "#00307A",
				"800": "#001C47",
				"900": "#000E23",
				"950": "#000814",
			},
		},
		"Gray scale": {
			hex: "#808080",
			opts: NewOptions([]Shade{
				NewShade("100", 90),
				NewShade("200", 80),
				NewShade("300", 70),
				NewShade("400", 60),
				NewShade("500", 50),
				NewShade("600", 40),
				NewShade("700", 30),
				NewShade("800", 20),
				NewShade("900", 10),
			}),
			want: map[string]string{
				"100": "#E5E5E5",
				"200": "#CCCCCC",
				"300": "#B2B2B2",
				"400": "#999999",
				"500": "#7F7F7F",
				"600": "#666666",
				"700": "#4C4C4C",
				"800": "#333333",
				"900": "#191919",
			},
		},
		"Invalid hex color": {
			hex: "NOTAHEX",
			opts: NewOptions([]Shade{
				NewShade("100", 90),
			}),
			wantErr: true,
		},
		"Invalid lightness value": {
			hex: "#FF0000",
			opts: NewOptions([]Shade{
				NewShade("invalid", 101),
			}),
			wantErr: true,
		},
		"Empty shades": {
			hex:  "#FF0000",
			opts: NewOptions([]Shade{}),
			want: map[string]string{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GeneratePaletteFromHex(tt.hex, tt.opts)

			hasErr := err != nil
			if hasErr != tt.wantErr {
				t.Errorf("error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(got) != len(tt.want) {
					t.Errorf("got %d shades, want %d", len(got), len(tt.want))
					return
				}

				for name, hexValue := range tt.want {
					if gotHex, exists := got[name]; !exists || gotHex != hexValue {
						t.Errorf("shade %s: got %q, want %q", name, gotHex, hexValue)
					}
				}
			}
		})
	}
}

func TestNewShade(t *testing.T) {
	tests := map[string]struct {
		name      string
		lightness uint8
		want      Shade
	}{
		"Standard shade": {
			name:      "500",
			lightness: 50,
			want:      Shade{name: "500", lightness: 50},
		},
		"Light shade": {
			name:      "100",
			lightness: 90,
			want:      Shade{name: "100", lightness: 90},
		},
		"Dark shade": {
			name:      "900",
			lightness: 10,
			want:      Shade{name: "900", lightness: 10},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := NewShade(tt.name, tt.lightness)
			if got.name != tt.want.name || got.lightness != tt.want.lightness {
				t.Errorf("got {%q, %d}, want {%q, %d}",
					got.name, got.lightness, tt.want.name, tt.want.lightness)
			}
		})
	}
}

func TestDefaultTailwindOptions(t *testing.T) {
	expected := map[string]uint8{
		"50": 98, "100": 95, "200": 90, "300": 82, "400": 64,
		"500": 46, "600": 33, "700": 24, "800": 14, "900": 7, "950": 4,
	}

	opts := DefaultTailwindOptions()

	if len(opts.shades) != len(expected) {
		t.Errorf("got %d shades, want %d", len(opts.shades), len(expected))
	}

	for _, shade := range opts.shades {
		if lightness, exists := expected[shade.name]; !exists || lightness != shade.lightness {
			expectedVal := "not found"
			if exists {
				expectedVal = string(lightness)
			}
			t.Errorf("shade %s: got lightness %d, want %s", shade.name, shade.lightness, expectedVal)
		}
	}
}
