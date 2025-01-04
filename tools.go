package asellus

import (
	"bufio"
	"regexp"
	"strings"
)

func SplitContentBySectionWithRegex(content string, regexDelimiter string) []string {
	var sections []string
	var currentSection []string

	// Regex to detect Markdown/AsciiDoc titles
	re := regexp.MustCompile(regexDelimiter)

	// use a scanner to read the content line by line
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line is a title
		if matches := re.FindStringSubmatch(line); matches != nil {
			// If a new title is found, ends the current section and starts a new section
			if len(currentSection) > 0 {
				sections = append(sections, strings.Join(currentSection, "\n"))
			}
			currentSection = []string{line}
		} else {
			// Adds the content lines to the current section
			currentSection = append(currentSection, line)
		}
	}

	// Add the last section if it exists
	if len(currentSection) > 0 {
		sections = append(sections, strings.Join(currentSection, "\n"))
	}

	return sections
}


func CreateChunksFrom(fragmentsOfDoc []string) []Chunk {
	var chunks []Chunk
	for _, fragment := range fragmentsOfDoc {
		chunks = append(chunks, Chunk{
			Content: fragment,
		})
	}
	return chunks
}