package asellus

import (
	"regexp"
	"strings"
)

func ParseMarkdownWithLineage(content string) []Chunk {
	lines := strings.Split(content, "\n")
	var chunks []Chunk
	var stack []Chunk

	headerRegex := regexp.MustCompile(`^(#+)\s+(.*)$`)

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if matches := headerRegex.FindStringSubmatch(line); matches != nil {
			level := len(matches[1])
			header := matches[2]
			prefix := matches[1]

			// Find content for this header
			contentLines := []string{}
			for j := i + 1; j < len(lines); j++ {
				if headerRegex.MatchString(lines[j]) {
					break
				}
				contentLines = append(contentLines, lines[j])
			}
			content := strings.Join(contentLines, "\n")

			// Determine parent header
			var parent Chunk
			for len(stack) > 0 && stack[len(stack)-1].Level >= level {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				parent = stack[len(stack)-1]
			}

			// Build lineage
			lineage := buildLineage(stack, header)

			chunk := Chunk{
				Level:        level,
				Prefix:       prefix,
				Header:       header,
				Content:      strings.TrimSpace(content),
				ParentPrefix: parent.Prefix,
				ParentLevel:  parent.Level,
				ParentHeader: parent.Header,
				Lineage:      lineage,
			}
			/*
				if chunk.Content == "" { // empty content because only title

				}
				if chunk.Content != "" {

				}
			*/
			// use even the chunks with empty content to keep the lineage
			// but do not create embeddings for the chnuks with empty content
			chunks = append(chunks, chunk)
			stack = append(stack, chunk)
		}
	}

	return chunks
}

func buildLineage(stack []Chunk, currentHeader string) string {
	var lineage []string
	for _, chunk := range stack {
		lineage = append(lineage, chunk.Header)
	}
	lineage = append(lineage, currentHeader)
	return strings.Join(lineage, " > ")
}

func SplitMarkdownBySections(content string) []string {
	return SplitContentBySectionWithRegex(content, `^(#+)\s+(.*)`)
}

// SplitMarkdownWithDelimiter
// SplitMarkdownWithRegex
