# thumb

A simple command-line tool written in Go to generate WebP thumbnails for images in a directory. It efficiently processes JPEG, PNG, and WebP files, creating optimized thumbnails based on specified dimensions and quality settings.

## Features

*   Reads images (JPEG, PNG, WebP) from an input directory recursively.
*   Generates thumbnails constrained by a maximum width.
*   Outputs thumbnails to a specified directory in WebP format.
*   Allows customization of output quality (1-100).
*   Provides sensible defaults:
    *   Maximum width: 512px (if not specified or invalid).
    *   Quality: Automatically determined based on width (e.g., 85 for <=512px, 80 for <=1024px, etc.) if not specified or invalid.
*   Displays progress and reports total file size savings upon completion.
*   Uses efficient image resizing (Lanczos3) and WebP encoding.

## Installation

```bash
go install github.com/coalaura/thumb@latest
```

Alternatively, clone the repository and build it manually:

```bash
git clone https://github.com/coalaura/thumb.git
cd thumb
go build
```

This will create a `thumb` executable in the project directory.

## Usage

Run the tool from your terminal:

```bash
thumb -i <input_directory> -o <output_directory> [options]
```

**Required Flags:**

*   `-i`, `--input`: Path to the directory containing the original images.
*   `-o`, `--output`: Path to the directory where thumbnails will be saved (will be created if it doesn't exist).

**Optional Flags:**

*   `-w`, `--width`: Maximum width for the thumbnails (default: 512). Values outside 1-4096 will use the default.
*   `-q`, `--quality`: Quality of the output WebP thumbnails (1-100). If omitted or invalid, a sensible quality is derived from the width.
*   `-h`, `--help`: Show the help message listing all flags.

**Examples:**

1.  **Basic Usage (Defaults):**
    Generate thumbnails (max 512px width, auto quality) from `./images` and save them to `./thumbnails`:
    ```bash
    thumb -i ./images -o ./thumbnails
    ```

2.  **Custom Width and Quality:**
    Generate thumbnails with a maximum width of 1024px and a fixed quality of 90:
    ```bash
    thumb -i ./images -o ./thumbnails -w 1024 -q 90
    ```
