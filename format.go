package main

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
	Blue  = "\033[34m"
)

func format(content string, length int, indices []int) string {
	var newContent string
	if indices != nil {
		prev := 0
		for _, i := range indices {
			newContent += content[prev:i]
			newContent += Red + content[i:i+length] + Reset
			prev = i + length
		}
		return newContent
	} else {
		return ""
	}
}
