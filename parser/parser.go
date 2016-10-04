package parser

type Parser interface {
	Output() <-chan string
}

// Parse file contents located at `src`.
func Parse(src string) (Parser, error) {
	if startsWith(src, "http://") || startsWith(src, "https://") {
		return parseHTTP(src)
	}
	return parseFile(src)
}

func startsWith(str, prefix string) bool {
	return len(str) >= len(prefix) && str[:len(prefix)] == prefix
}
