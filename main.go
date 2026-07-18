package main

import (
	"bufio"
	"bytes"
	"embed"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

//go:embed cmudict.txt
var embeddedDict embed.FS

type Entry struct {
	Word   string
	Phones []string
}

func main() {
	dictPath := flag.String("dict", "cmudict.txt", "Path to cmudict.txt")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Usage: go run main.go <word>")
		return
	}

	word := strings.ToUpper(flag.Arg(0))

	entries, err := loadDict(*dictPath)
	if err != nil {
		fmt.Printf("Load error: %v\n", err)
		return
	}

	rhymes := findRhymes(word, entries)
	if len(rhymes) == 0 {
		fmt.Printf("No rhymes found for '%s' (or word not in dict).\n", flag.Arg(0))
	} else {
		sort.Strings(rhymes)
		printWrapped(rhymes, terminalWidth())
	}
}

func loadDict(path string) ([]Entry, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return parseDict(bytes.NewReader(data))
}

func parseDict(r io.Reader) ([]Entry, error) {
	var entries []Entry
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		word := strings.ToUpper(fields[0])  // Force upper
		if idx := strings.Index(word, "("); idx != -1 {
			word = word[:idx]
		}

		phones := make([]string, len(fields)-1)
		copy(phones, fields[1:])

		entries = append(entries, Entry{Word: word, Phones: phones})
	}
	return entries, scanner.Err()
}

func getRhymingTail(phones []string) string {
	for i := len(phones) - 1; i >= 0; i-- {
		p := phones[i]
		if len(p) > 0 {
			lastChar := p[len(p)-1]
			if lastChar == '1' || lastChar == '2' {
				return strings.Join(phones[i:], " ")
			}
		}
	}
	return strings.Join(phones, " ")
}

func findRhymes(targetWord string, entries []Entry) []string {
	targetTails := make(map[string]struct{})
	for _, e := range entries {
		if e.Word == targetWord {
			tail := getRhymingTail(e.Phones)
			targetTails[tail] = struct{}{}
		}
	}

	if len(targetTails) == 0 {
		return nil
	}

	rhymeSet := make(map[string]struct{})
	for _, e := range entries {
		if e.Word == targetWord {
			continue
		}
		tail := getRhymingTail(e.Phones)
		if _, ok := targetTails[tail]; ok {
			rhymeSet[strings.ToLower(e.Word)] = struct{}{}
		}
	}

	rhymes := make([]string, 0, len(rhymeSet))
	for w := range rhymeSet {
		rhymes = append(rhymes, w)
	}
	return rhymes
}

// terminalWidth returns the current terminal width, falling back to 80
// columns if it can't be determined (e.g. output is piped/redirected).
func terminalWidth() int {
	if cols := os.Getenv("COLUMNS"); cols != "" {
		if n, err := strconv.Atoi(cols); err == nil && n > 0 {
			return n
		}
	}
	return 80
}

// printWrapped prints words space-separated, wrapping onto new lines so
// that no line exceeds width, similar to the fmt(1) command.
func printWrapped(words []string, width int) {
	lineLen := 0
	for i, w := range words {
		wLen := len(w)
		if lineLen == 0 {
			fmt.Print(w)
			lineLen = wLen
		} else if lineLen+1+wLen > width {
			fmt.Println()
			fmt.Print(w)
			lineLen = wLen
		} else {
			fmt.Print(" " + w)
			lineLen += 1 + wLen
		}
		if i == len(words)-1 {
			fmt.Println()
		}
	}
}