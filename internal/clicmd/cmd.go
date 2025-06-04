package clicmd

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/claytonchew/tailwindcss-palette-go/internal/color"
	"github.com/claytonchew/tailwindcss-palette-go/internal/generator"
	"github.com/claytonchew/tailwindcss-palette-go/internal/version"
)

type exitCode int

const (
	exitOK    exitCode = 0
	exitError exitCode = 1
)

type ColorFormat string

const (
	HexFormat ColorFormat = "hex"
	HSLFormat ColorFormat = "hsl"
	RGBFormat ColorFormat = "rgb"
)

const (
	colorReset = "\033[0m"
	colorBlock = "    "
)

var (
	ErrorInvalidHexInput = errors.New("invalid hex color: must be in format #RRGGBB or #RGB")
	ErrorInvalidFormat   = errors.New("invalid color format: must be one of 'hex', 'hsl', or 'rgb'")
)

func Main() exitCode {
	flagSet := flag.NewFlagSet("tailwindcss-palette", flag.ExitOnError)
	colorFormat := flagSet.String("c", string(HexFormat), "Color format: hex, hsl, or rgb")
	outputFile := flagSet.String("o", "", "Path to output JSON file (optional)")
	noColorPtr := flagSet.Bool("no-color", false, "Disable colored output")
	_ = flagSet.Bool("v", false, "Print version information and exit")

	flagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: tailwindcss-palette <hex-color> [options]\n\n")
		fmt.Fprintf(os.Stderr, "Arguments:\n")
		fmt.Fprintf(os.Stderr, "  <hex-color>    Hex color code (e.g. #FF5733 or FF5733)\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		fmt.Fprintf(os.Stderr, "  -h, --help     Show this help message\n")
		fmt.Fprintf(os.Stderr, "  -v, --version  Print version information and exit\n")
		flagSet.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  tailwindcss-palette #3B82F6                   # Generate palette in hex format\n")
		fmt.Fprintf(os.Stderr, "  tailwindcss-palette #3B82F6 -c rgb            # Generate palette in RGB format\n")
		fmt.Fprintf(os.Stderr, "  tailwindcss-palette #3B82F6 -o palette.json   # Export to JSON file\n")
	}

	for _, arg := range os.Args[1:] {
		if arg == "-h" || arg == "--help" {
			flagSet.Usage()
			return exitOK
		}
		if arg == "-v" || arg == "--version" {
			fmt.Println("tailwindcss-palette version", version.Info())
			return exitOK
		}
	}

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: Missing hex color argument\n\n")
		flagSet.Usage()
		return exitError
	}

	hexColor := os.Args[1]

	if len(os.Args) > 2 {
		if err := flagSet.Parse(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
			return exitError
		}
	}

	format := ColorFormat(strings.ToLower(*colorFormat))
	if format != HexFormat && format != HSLFormat && format != RGBFormat {
		fmt.Fprintf(os.Stderr, "Error: %v\n", ErrorInvalidFormat)
		return exitError
	}

	if !strings.HasPrefix(hexColor, "#") {
		hexColor = "#" + hexColor
	}

	palette, err := generator.GeneratePaletteFromHex(hexColor, generator.DefaultTailwindOptions())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return exitError
	}

	if *outputFile != "" {
		if err := writeToJSONFile(palette, hexColor, *outputFile); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
			return exitError
		}
		fmt.Printf("Palette has been written to %s\n", *outputFile)
		return exitOK
	}

	useColor := !*noColorPtr && isTerminal()

	baseHex := hexColor
	if useColor {
		colorBlock := getColorBlock(baseHex)
		fmt.Printf("Base color: %s %s\n", baseHex, colorBlock)
	} else {
		fmt.Printf("Base color: %s\n", baseHex)
	}

	h, s, l, _ := color.HexToHSL(baseHex)
	r, g, b, _ := color.HexToRGB(baseHex)

	switch format {
	case HSLFormat:
		fmt.Printf("HSL: hsl(%3.0f, %3.0f%%, %3.0f%%)", h, s*100, l*100)
		if useColor {
			fmt.Printf(" %s\n", getColorBlock(baseHex))
		} else {
			fmt.Println()
		}
	case RGBFormat:
		fmt.Printf("RGB: rgb(%3d, %3d, %3d)", r, g, b)
		if useColor {
			fmt.Printf(" %s\n", getColorBlock(baseHex))
		} else {
			fmt.Println()
		}
	}

	fmt.Println("\nTailwind CSS palette:")
	fmt.Println("---------------------")

	if err := outputPalette(palette, format, useColor); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return exitError
	}

	return exitOK
}

func outputPalette(palette map[string]string, format ColorFormat, useColor bool) error {
	keys := []string{"50", "100", "200", "300", "400", "500", "600", "700", "800", "900", "950"}

	for _, key := range keys {
		hexValue, exists := palette[key]
		if !exists {
			continue
		}

		switch format {
		case HexFormat:
			fmt.Printf("  %-4s: %-9s", key, hexValue)
			if useColor {
				fmt.Printf(" %s\n", getColorBlock(hexValue))
			} else {
				fmt.Println()
			}
		case HSLFormat:
			h, s, l, err := color.HexToHSL(hexValue)
			if err != nil {
				return err
			}
			hslString := fmt.Sprintf("hsl(%3.0f, %3.0f%%, %3.0f%%)", h, s*100, l*100)
			fmt.Printf("  %-4s: %-25s", key, hslString)
			if useColor {
				fmt.Printf(" %s\n", getColorBlock(hexValue))
			} else {
				fmt.Println()
			}
		case RGBFormat:
			r, g, b, err := color.HexToRGB(hexValue)
			if err != nil {
				return err
			}
			rgbString := fmt.Sprintf("rgb(%3d, %3d, %3d)", r, g, b)
			fmt.Printf("  %-4s: %-20s", key, rgbString)
			if useColor {
				fmt.Printf(" %s\n", getColorBlock(hexValue))
			} else {
				fmt.Println()
			}
		}
	}

	return nil
}

func isTerminal() bool {
	fileInfo, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

func getColorBlock(hex string) string {
	r, g, b, err := color.HexToRGB(hex)
	if err != nil {
		return colorBlock + colorReset
	}
	return fmt.Sprintf("\033[48;2;%d;%d;%dm%s%s", r, g, b, colorBlock, colorReset)
}

func writeToJSONFile(palette map[string]string, baseColor string, filePath string) error {
	paletteData := map[string]any{
		"base": map[string]any{
			"hex": baseColor,
		},
		"palette": map[string]map[string]any{},
	}

	h, s, l, err := color.HexToHSL(baseColor)
	if err == nil {
		paletteData["base"].(map[string]any)["hsl"] = map[string]any{
			"h": int(h),
			"s": s,
			"l": l,
		}
	}

	r, g, b, err := color.HexToRGB(baseColor)
	if err == nil {
		paletteData["base"].(map[string]any)["rgb"] = map[string]any{
			"r": r,
			"g": g,
			"b": b,
		}
	}

	for shade, hexValue := range palette {
		shadeData := map[string]any{
			"hex": hexValue,
		}

		if h, s, l, err := color.HexToHSL(hexValue); err == nil {
			shadeData["hsl"] = map[string]any{
				"h": int(h),
				"s": s,
				"l": l,
			}
		}

		if r, g, b, err := color.HexToRGB(hexValue); err == nil {
			shadeData["rgb"] = map[string]any{
				"r": r,
				"g": g,
				"b": b,
			}
		}

		paletteData["palette"].(map[string]map[string]any)[shade] = shadeData
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(paletteData)
}
