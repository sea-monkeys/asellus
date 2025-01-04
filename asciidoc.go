package asellus

// This is a work in progress. It is not yet complete. 🚧 

func SplitAsciiDocBySections(content string) []string {
	return SplitContentBySectionWithRegex(content, `^(=+|\#+)\s+(.*)`)
}