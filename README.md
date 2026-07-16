# rhymeswith

A fast, standalone command-line tool to find perfect rhymes using the CMU Pronouncing Dictionary.

## Features
1. One self-contained binary (dictionary is embedded).
2. High-quality phonetic rhymes (from CMU Pronouncing Dictionary).
3. Space-separated output for easy reading and copying.
4. Works on macOS (tested).
5. Probably works on Linux, Windows (with compilation).

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
