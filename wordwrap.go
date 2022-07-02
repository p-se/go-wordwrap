package wordwrap

import (
	"strings"
	"unicode/utf8"
)

// Dedent removes all spaces and tabs from the start and end of each line.
// Everything that is separated by newline characters counts as a line.
func Dedent(s string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = strings.Trim(line, " \t")
	}
	return strings.Join(lines, "\n")
}

// Join joins all sentences separated by a single newline into a paragraph
// without newlines. If the string to be joined is indented, `Dedent` should be
// called first. Otherwise the joined text will include the indentation.
func Join(s string) string {
	paragraphs := strings.Split(s, "\n\n")
	for i := range paragraphs {
		paragraphs[i] = strings.ReplaceAll(paragraphs[i], "\n", " ")
		paragraphs[i] = strings.ReplaceAll(paragraphs[i], "  ", " ")
		paragraphs[i] = strings.Trim(paragraphs[i], " \t\n")
	}
	return strings.Join(paragraphs, "\n\n")
}

// Prepare is the preparation step necessary before being able to word-wrap a
// text.
func Prepare(s string) string {
	return Join(Dedent(s))
}

// WordWrapParagraph wraps a given paragraph by the amount of characters. A
// character is considered a rune, not byte. It assumes a paragraph of text
// (single line) without any newlines.
func WordWrapParagraph(p string, width int) string {
	var (
		lines        [][]string
		current_line []string
	)
	words := strings.Split(p, " ")
	for i, word := range words {
		if utf8.RuneCountInString(strings.Join(append(current_line, word), " ")) < width {
			current_line = append(current_line, word)
		} else {
			lines = append(lines, current_line)
			current_line = []string{word}
		}
		if len(words)-1 == i {
			lines = append(lines, current_line)
		}
	}
	var strLines []string
	for _, line := range lines {
		strLines = append(strLines, strings.Join(line, " "))
	}
	return strings.Join(strLines, "\n")
}

// WrapText wraps a complete text which may consist of several paragraphs. A
// paragraph is considered a text separated by two consecutive newlines. The
// text is wrapped at a given amount of characters. A character is considered a
// rune, not a byte.
func WrapText(s string, width int) string {
	paragraphs := strings.Split(s, "\n\n")
	for i, paragraph := range paragraphs {
		paragraphs[i] = WordWrapParagraph(Prepare(paragraph), width)
	}
	return strings.Join(paragraphs, "\n\n")
}

// Indent indents each line of a string by the given indent string. Empty lines
// are not indented.
func Indent(s string, indent string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if line != "" {
			lines[i] = indent + line
		}
	}
	return strings.Join(lines, "\n")
}

// Wrap is the combination of all functions in this module and enables to wrap
// and indent a text properly by just using this function.
func Wrap(text string, width int, indent string) string {
	return Indent(WrapText(text, width-utf8.RuneCountInString(indent)), indent)
}
