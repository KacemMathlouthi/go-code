package utils

import (
	"fmt"
	"strings"
)

// ANSI color codes for terminal formatting
const (
	ColorReset   = "\033[0m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorWhite   = "\033[37m"
	ColorBold    = "\033[1m"
	ColorDim     = "\033[2m"
)

// Box drawing characters
const (
	TopLeft     = "╭"
	TopRight    = "╮"
	BottomLeft  = "╰"
	BottomRight = "╯"
	Horizontal  = "─"
	Vertical    = "│"
)

// FormatUserInput formats user input in a styled box
func FormatUserInput(input string) string {
	lines := wrapLines(strings.Split(input, "\n"), 80) // wrap at 80 chars
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}
	if maxWidth < 20 {
		maxWidth = 20
	}
	var result strings.Builder
	result.WriteString(ColorCyan + ColorBold + "👤 User" + ColorReset + "\n")
	result.WriteString(ColorCyan + TopLeft + strings.Repeat(Horizontal, maxWidth+2) + TopRight + ColorReset + "\n")
	for _, line := range lines {
		padding := maxWidth - len(line)
		result.WriteString(ColorCyan + Vertical + ColorReset + " " + line + strings.Repeat(" ", padding) + " " + ColorCyan + Vertical + ColorReset + "\n")
	}
	result.WriteString(ColorCyan + BottomLeft + strings.Repeat(Horizontal, maxWidth+2) + BottomRight + ColorReset + "\n")
	return result.String()
}

// FormatAIResponse prints only the bot tag and the raw output, no box or formatting
func FormatAIResponse(response string) string {
	return "🤖 " + ColorMagenta + ColorBold + "AI Assistant" + ColorReset + "\n" + response
}

// wrapLines wraps each line in the input slice to the given width
func wrapLines(lines []string, width int) []string {
	var wrapped []string
	for _, line := range lines {
		for len(line) > width {
			wrapped = append(wrapped, line[:width])
			line = line[width:]
		}
		wrapped = append(wrapped, line)
	}
	return wrapped
}

// FormatError formats error messages in a styled box
func FormatError(err string) string {
	lines := strings.Split(err, "\n")
	maxWidth := 0

	// Find the maximum line width
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	// Ensure minimum width
	if maxWidth < 20 {
		maxWidth = 20
	}

	// Build the formatted output
	var result strings.Builder

	// Top border with error icon
	result.WriteString(ColorRed + ColorBold + "❌ Error" + ColorReset + "\n")
	result.WriteString(ColorRed + TopLeft + strings.Repeat(Horizontal, maxWidth+2) + TopRight + ColorReset + "\n")

	// Content lines
	for _, line := range lines {
		padding := maxWidth - len(line)
		result.WriteString(ColorRed + Vertical + ColorReset + " " + line + strings.Repeat(" ", padding) + " " + ColorRed + Vertical + ColorReset + "\n")
	}

	// Bottom border
	result.WriteString(ColorRed + BottomLeft + strings.Repeat(Horizontal, maxWidth+2) + BottomRight + ColorReset + "\n")

	return result.String()
}

// FormatPrompt formats the input prompt with styling
func FormatPrompt() string {
	return ColorCyan + ColorBold + "> " + ColorReset
}

// ClearScreen clears the terminal screen with a nice message
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
	fmt.Println(ColorGreen + ColorBold + "✨ Terminal cleared! Ready for new conversation." + ColorReset)
	fmt.Println()
}
