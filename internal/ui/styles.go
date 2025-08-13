package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Color palette
	primaryColor    = lipgloss.Color("#7C3AED") // Purple
	successColor    = lipgloss.Color("#10B981") // Green
	warningColor    = lipgloss.Color("#F59E0B") // Orange
	errorColor      = lipgloss.Color("#EF4444") // Red
	mutedColor      = lipgloss.Color("#6B7280") // Gray
	backgroundColor = lipgloss.Color("#1F2937") // Dark gray
	textColor       = lipgloss.Color("#F9FAFB") // Light gray

	// Header styles
	TitleStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			Padding(0, 1).
			Align(lipgloss.Center).
			Width(25)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(textColor).
			Bold(true).
			MarginBottom(1)

	// Score styles
	ScoreStyle = lipgloss.NewStyle().
			Foreground(successColor).
			Bold(true).
			Padding(0, 1)

	ScoreLowStyle = lipgloss.NewStyle().
			Foreground(errorColor).
			Bold(true).
			Padding(0, 1)

	ScoreMediumStyle = lipgloss.NewStyle().
				Foreground(warningColor).
				Bold(true).
				Padding(0, 1)

	// Table styles
	TableHeaderStyle = lipgloss.NewStyle().
				Foreground(primaryColor).
				Bold(true).
				Border(lipgloss.NormalBorder(), false, false, true, false).
				BorderForeground(mutedColor).
				Padding(0, 1)

	TableRowStyle = lipgloss.NewStyle().
			Foreground(textColor).
			Padding(0, 1)

	TableRowAltStyle = lipgloss.NewStyle().
				Foreground(textColor).
				Background(lipgloss.Color("#374151")).
				Padding(0, 1)

	// Progress bar styles
	ProgressBarStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(mutedColor).
				Padding(0, 1)

	ProgressFillStyle = lipgloss.NewStyle().
				Background(successColor).
				Foreground(lipgloss.Color("#000000"))

	ProgressEmptyStyle = lipgloss.NewStyle().
				Background(mutedColor)

	// Card styles
	CardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(mutedColor).
			Padding(1, 2).
			MarginBottom(1)

	SuccessCardStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(successColor).
				Padding(1, 2).
				MarginBottom(1)

	WarningCardStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(warningColor).
				Padding(1, 2).
				MarginBottom(1)

	// List styles
	ListItemStyle = lipgloss.NewStyle().
			Foreground(textColor).
			PaddingLeft(2)

	SuccessListItemStyle = lipgloss.NewStyle().
				Foreground(successColor).
				PaddingLeft(2)

	WarningListItemStyle = lipgloss.NewStyle().
				Foreground(warningColor).
				PaddingLeft(2)

	// Status styles
	StatusStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Italic(true)

	// Help styles
	HelpStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(mutedColor).
			MarginTop(1)

	// Tab styles
	ActiveTabStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			Padding(0, 2).
			Border(lipgloss.Border{
				Top:         "─",
				Bottom:      "",
				Left:        "│",
				Right:       "│",
				TopLeft:     "╭",
				TopRight:    "╮",
				BottomLeft:  "│",
				BottomRight: "│",
			}).
			BorderForeground(primaryColor)

	InactiveTabStyle = lipgloss.NewStyle().
				Foreground(mutedColor).
				Padding(0, 2).
				Border(lipgloss.Border{
					Top:         "─",
					Bottom:      "",
					Left:        "│",
					Right:       "│",
					TopLeft:     "╭",
					TopRight:    "╮",
					BottomLeft:  "│",
					BottomRight: "│",
				}).
				BorderForeground(mutedColor)
)

// GetScoreStyle returns the appropriate style based on score
func GetScoreStyle(score int) lipgloss.Style {
	if score >= 70 {
		return ScoreStyle
	} else if score >= 40 {
		return ScoreMediumStyle
	}
	return ScoreLowStyle
}

// CreateProgressBar creates a styled progress bar
func CreateProgressBar(current, max int, width int) string {
	if max == 0 {
		return ""
	}
	
	percentage := float64(current) / float64(max)
	fillWidth := int(float64(width) * percentage)
	emptyWidth := width - fillWidth
	
	fill := ProgressFillStyle.Width(fillWidth).Render("")
	empty := ProgressEmptyStyle.Width(emptyWidth).Render("")
	
	return ProgressBarStyle.Render(fill + empty)
}

// FormatScore formats a score with appropriate styling
func FormatScore(score, max int) string {
	style := GetScoreStyle(score)
	return style.Render(lipgloss.JoinHorizontal(lipgloss.Center, 
		lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("%d", score)),
		"/",
		fmt.Sprintf("%d", max),
	))
}