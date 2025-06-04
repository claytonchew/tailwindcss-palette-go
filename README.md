# Tailwind CSS Palette Generator

This command line tool generates a Tailwind CSS-like color palette from a base hex color. Built in pure Go, without external dependencies.

## Features

- Generate a full Tailwind CSS palette from any hex color
- Output in various formats (hex, HSL, RGB)
- Export palette to JSON file

## Installation

### Option 1: Homebrew (macOS)

```bash
brew install claytonchew/tap/tailwindcss-palette
```

### Option 2: Using Go Install

```bash
go install github.com/claytonchew/tailwindcss-palette-go/cmd/tailwindcss-palette@latest
```

### Option 3: Manual build and install

Clone the repository:
```bash
git clone https://github.com/claytonchew/tailwindcss-palette-go.git
cd tailwindcss-palette-go
```

Using Go directly:
```bash
go build -o tailwindcss-palette ./cmd/tailwindcss-palette
```

Using the Makefile:
```bash
# Build for your current platform
make build

# Run tests
make test

# Build for multiple platforms (linux, macOS, Windows)
make cross-platform
```

Built binaries can be found in the `build` directory.

## Usage

```
tailwindcss-palette <hex-color> [-c format] [-o output-file]
```

Where `<hex-color>` can be with or without the `#` prefix.

### Arguments

- `<hex-color>`: The base color in hex format (with or without # prefix)

### Flags

- `-c`: Color format (default: "hex")
  - Available formats: "hex", "hsl", "rgb"
- `-o`: Path to output JSON file (optional)
  - When specified, the palette will be saved as JSON

### Examples

Generate a palette in hex format:
```
tailwindcss-palette #3b82f6
```

Generate a palette in HSL format:
```
tailwindcss-palette 3b82f6 -c hsl
```

Generate a palette in RGB format:
```
tailwindcss-palette 3b82f6 -c rgb
```

Export the palette to a JSON file:
```
tailwindcss-palette 3b82f6 -o palette.json
```

## Example Output

### Hex Format (default)

```
$ tailwindcss-palette #3B82F6

Base color: #3B82F6

Tailwind CSS palette:
---------------------
  50  : #F5F8FE
  100 : #E6EFFD
  200 : #CEDFFC
  300 : #A7C7FA
  400 : #4F8FF6
  500 : #0A5BE0
  600 : #0741A0
  700 : #052F74
  800 : #031B44
  900 : #010D22
  950 : #000713
```

### HSL Format

```
$ tailwindcss-palette #3B82F6 -c hsl

Base color: #3B82F6
HSL: hsl(217, 91%, 60%)

Tailwind CSS palette:
---------------------
  50  : hsl(220, 82%, 98%)
  100 : hsl(217, 85%, 95%)
  200 : hsl(218, 88%, 90%)
  300 : hsl(217, 89%, 82%)
  400 : hsl(217, 90%, 64%)
  500 : hsl(217, 91%, 46%)
  600 : hsl(217, 92%, 33%)
  700 : hsl(217, 92%, 24%)
  800 : hsl(218, 92%, 14%)
  900 : hsl(218, 94%, 7%)
  950 : hsl(218, 100%, 4%)
```

### RGB Format

```
$ tailwindcss-palette #3B82F6 -c rgb

Base color: #3B82F6
RGB: rgb(59, 130, 246)

Tailwind CSS palette:
---------------------
  50  : rgb(245, 248, 254)
  100 : rgb(230, 239, 253)
  200 : rgb(206, 223, 252)
  300 : rgb(167, 199, 250)
  400 : rgb( 79, 143, 246)
  500 : rgb( 10,  91, 224)
  600 : rgb(  7,  65, 160)
  700 : rgb(  5,  47, 116)
  800 : rgb(  3,  27,  68)
  900 : rgb(  1,  13,  34)
  950 : rgb(  0,   7,  19)
```

### JSON Output

When using the `-o` flag, the palette will be saved as a JSON file:

```
$ tailwindcss-palette #3B82F6 -o palette.json
Palette has been written to palette.json
```

The JSON output contains all color formats:

```json
{
  "base": {
    "hex": "#3B82F6",
    "hsl": {
      "h": 217,
      "s": 0.91,
      "l": 0.6
    },
    "rgb": {
      "r": 59,
      "g": 130,
      "b": 246
    }
  },
  "palette": {
    "50": {
      "hex": "#F5F8FE",
      "hsl": {
        "h": 220,
        "s": 0.82,
        "l": 0.98
      },
      "rgb": {
        "r": 245,
        "g": 248,
        "b": 254
      }
    },
    "100": {
      "hex": "#E6EFFD",
      "hsl": {
        "h": 217,
        "s": 0.85,
        "l": 0.95
      },
      "rgb": {
        "r": 230,
        "g": 239,
        "b": 253
      }
    },
    // ...
    "950": {
      "hex": "#000713",
      "hsl": {
        "h": 218,
        "s": 1,
        "l": 0.04
      },
      "rgb": {
        "r": 0,
        "g": 7,
        "b": 19
      }
    }
  }
}
```

## License

[MIT](https://github.com/claytonchew/tailwindcss-palette-go/blob/main/LICENSE)
