# rhymeswith

A fast, standalone command-line tool to find perfect rhymes using the CMU Pronouncing Dictionary.

## Features
- Fully self-contained single binary (dictionary embedded)
- High-quality phonetic rhymes
- Space-separated output for easy copying
- Works on macOS, Linux, Windows

## Quick Start

```bash
make          # downloads dict + builds
./rhymeswith boy
```

## Usage Example

```bash
./rhymeswith beach
Looking for rhymes for: BEACH
Loaded 135166 entries from dictionary
BOY found in dict: true
Rhymes with 'beach' (37):
beech beseech bleach breach breech cheech creach creech dietsch dietsche each impeach inspeech keach keech keetch leach leech leetch leitch long-beach meech peach piech pietsch preach reach reeche screech speech swiech teach veach veatch veech weech wiech 
```
