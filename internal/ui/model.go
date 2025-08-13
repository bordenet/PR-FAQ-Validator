package ui

import (
	"fmt"
	"strings"

	"github.com/bordenet/pr-faq-validator/internal/parser"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Tab represents different views in the TUI
type Tab int

const (
	TabOverview Tab = iota
	TabBreakdown
	TabQuotes
	TabFeedback
)

// Model represents the TUI application state
type Model struct {
	// Core data
	sections parser.SpecSections
	prFeedback string
	faqFeedback string
	
	// UI state
	activeTab   Tab
	showHelp    bool
	windowWidth int
	windowHeight int
	
	// Navigation
	tabs        []string
	scrollPos   int
	maxScroll   int
	
	// Status
	status      string
	loading     bool
}

// NewModel creates a new TUI model
func NewModel(sections parser.SpecSections) Model {
	return Model{
		sections:    sections,
		activeTab:   TabOverview,
		showHelp:    false,
		tabs:        []string{"Overview", "Breakdown", "Quotes", "AI Feedback"},
		windowWidth: 80,
		windowHeight: 24,
		status:      "Ready",
	}
}

// Init initializes the TUI model
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles TUI events and state changes
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		return m, nil
		
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
			
		case "?":
			m.showHelp = !m.showHelp
			return m, nil
			
		case "left", "h":
			if m.activeTab > 0 {
				m.activeTab--
				m.scrollPos = 0
				m.status = fmt.Sprintf("Switched to %s", m.tabs[m.activeTab])
			}
			return m, nil
			
		case "right", "l":
			if int(m.activeTab) < len(m.tabs)-1 {
				m.activeTab++
				m.scrollPos = 0
				m.status = fmt.Sprintf("Switched to %s", m.tabs[m.activeTab])
			}
			return m, nil
			
		case "up", "k":
			if m.scrollPos > 0 {
				m.scrollPos--
			}
			return m, nil
			
		case "down", "j":
			if m.scrollPos < m.maxScroll {
				m.scrollPos++
			}
			return m, nil
		}
		
	case SetFeedbackMsg:
		if msg.Section == "Press Release" {
			m.prFeedback = msg.Feedback
		} else if msg.Section == "FAQs" {
			m.faqFeedback = msg.Feedback
		}
		m.loading = false
		m.status = "Analysis complete"
		return m, nil
		
	case SetStatusMsg:
		m.status = string(msg)
		return m, nil
		
	case SetLoadingMsg:
		m.loading = bool(msg)
		if m.loading {
			m.status = "Analyzing with AI..."
		}
		return m, nil
	}
	
	return m, nil
}

// View renders the TUI interface
func (m Model) View() string {
	var content []string
	
	// Header
	header := RenderHeader(m.sections.Title, m.sections.PRScore.OverallScore)
	content = append(content, header)
	content = append(content, "") // Add spacing
	
	// Tabs
	tabs := RenderTabs(m.tabs, int(m.activeTab))
	content = append(content, tabs)
	content = append(content, "") // Add spacing
	
	// Content based on active tab
	var tabContent string
	switch m.activeTab {
	case TabOverview:
		tabContent = m.renderOverview()
	case TabBreakdown:
		tabContent = m.renderBreakdown()
	case TabQuotes:
		tabContent = m.renderQuotes()
	case TabFeedback:
		tabContent = m.renderFeedback()
	}
	
	// Apply scrolling to content
	lines := strings.Split(tabContent, "\n")
	if len(lines) > m.windowHeight-10 { // Reserve space for header, tabs, status
		m.maxScroll = len(lines) - (m.windowHeight - 10)
		if m.scrollPos > m.maxScroll {
			m.scrollPos = m.maxScroll
		}
		
		endPos := m.scrollPos + (m.windowHeight - 10)
		if endPos > len(lines) {
			endPos = len(lines)
		}
		
		if m.scrollPos < len(lines) {
			lines = lines[m.scrollPos:endPos]
		}
	} else {
		m.maxScroll = 0
	}
	
	content = append(content, strings.Join(lines, "\n"))
	
	// Help section
	if m.showHelp {
		content = append(content, "")
		content = append(content, RenderHelp())
	}
	
	// Status line  
	content = append(content, "")
	statusLine := RenderStatus(m.status)
	if m.loading {
		statusLine = RenderStatus("🔄 " + m.status)
	}
	content = append(content, statusLine)
	
	return lipgloss.JoinVertical(lipgloss.Left, content...)
}

// renderOverview renders the overview tab
func (m Model) renderOverview() string {
	var sections []string
	
	// Quick summary with better formatting
	scoreText := GetScoreStyle(m.sections.PRScore.OverallScore).Render(fmt.Sprintf("%d/100", m.sections.PRScore.OverallScore))
	
	summaryContent := lipgloss.JoinVertical(lipgloss.Left,
		SubtitleStyle.Render("📊 Summary"),
		ListItemStyle.Render("Overall Score: "+scoreText),
		ListItemStyle.Render(fmt.Sprintf("Press Release: %s", m.getStatusText(len(m.sections.PressRelease) > 0))),
		ListItemStyle.Render(fmt.Sprintf("FAQ Section: %s", m.getStatusText(len(m.sections.FAQs) > 0))),
		ListItemStyle.Render(fmt.Sprintf("Quotes Found: %d", m.sections.PRScore.TotalQuotes)),
	)
	
	summary := CardStyle.Render(summaryContent)
	sections = append(sections, summary)
	
	// Top strengths
	if len(m.sections.PRScore.QualityBreakdown.Strengths) > 0 {
		topStrengths := m.sections.PRScore.QualityBreakdown.Strengths
		if len(topStrengths) > 3 {
			topStrengths = topStrengths[:3]
		}
		sections = append(sections, RenderStrengths(topStrengths))
	}
	
	// Top improvements
	if len(m.sections.PRScore.QualityBreakdown.Issues) > 0 {
		topImprovements := m.sections.PRScore.QualityBreakdown.Issues
		if len(topImprovements) > 3 {
			topImprovements = topImprovements[:3]
		}
		sections = append(sections, RenderImprovements(topImprovements))
	}
	
	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

// renderBreakdown renders the detailed score breakdown tab
func (m Model) renderBreakdown() string {
	return RenderScoreBreakdown(m.sections.PRScore.QualityBreakdown)
}

// renderQuotes renders the quotes analysis tab
func (m Model) renderQuotes() string {
	if len(m.sections.PRScore.MetricDetails) == 0 {
		return CardStyle.Render(
			SubtitleStyle.Render("💬 Quote Analysis") + "\n\n" +
			WarningListItemStyle.Render("No quotes found in the press release section."))
	}
	
	return RenderQuoteAnalysis(*m.sections.PRScore)
}

// renderFeedback renders the AI feedback tab
func (m Model) renderFeedback() string {
	var sections []string
	
	if m.prFeedback != "" {
		sections = append(sections, RenderLLMFeedback("Press Release", m.prFeedback))
	}
	
	if m.faqFeedback != "" {
		sections = append(sections, RenderLLMFeedback("FAQ", m.faqFeedback))
	}
	
	if len(sections) == 0 {
		return CardStyle.Render(
			SubtitleStyle.Render("🤖 AI Feedback") + "\n\n" +
			StatusStyle.Render("AI analysis will appear here once processing is complete."))
	}
	
	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

// getStatusText returns a colored status indicator
func (m Model) getStatusText(present bool) string {
	if present {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#10B981")).Render("✓ Found")
	}
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#F59E0B")).Render("✗ Not Found")
}

// Commands and Messages for handling async operations

type SetFeedbackMsg struct {
	Section  string
	Feedback string
}

type SetStatusMsg string

type SetLoadingMsg bool

// SetFeedback creates a command to set feedback
func SetFeedback(section, feedback string) tea.Cmd {
	return func() tea.Msg {
		return SetFeedbackMsg{Section: section, Feedback: feedback}
	}
}

// SetStatus creates a command to set status
func SetStatus(status string) tea.Cmd {
	return func() tea.Msg {
		return SetStatusMsg(status)
	}
}

// SetLoading creates a command to set loading state
func SetLoading(loading bool) tea.Cmd {
	return func() tea.Msg {
		return SetLoadingMsg(loading)
	}
}