# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based CLI tool that analyzes PR-FAQ (Press Release - Frequently Asked Questions) documents using a hybrid approach: deterministic rule-based scoring (100% of numerical score) plus optional OpenAI GPT-4 qualitative feedback. The tool provides comprehensive quality scoring, interactive terminal UI, and detailed improvement recommendations.

## Development Commands

### Build and Run
```bash
# Install dependencies
go mod tidy

# Build the binary
go build

# Run the tool
./pr-faq-validator -file path/to/your/prfaq.md

# Test with example files
./pr-faq-validator -file testdata/example_prfaq_1.md
./pr-faq-validator -file testdata/example_prfaq_4.md  # High-quality example (77/100)

# Run without TUI (legacy output)
./pr-faq-validator -file testdata/example_prfaq_1.md -no-tui

# Generate markdown report
./pr-faq-validator -file testdata/example_prfaq_1.md -report
```

### Environment Setup
```bash
# Optional: Set OpenAI API key for AI feedback
export OPENAI_API_KEY=your_openai_api_key_here

# Tool works without API key - scoring is 100% deterministic
# AI feedback provides supplementary qualitative insights only
```

## Architecture

The codebase follows a clean modular structure:

- **main.go**: CLI entry point with interactive TUI (Bubble Tea) and legacy stdout modes
- **internal/parser/**: Document parsing, quality analysis, and deterministic scoring algorithms
- **internal/llm/**: OpenAI API integration for supplementary qualitative feedback
- **internal/ui/**: Interactive terminal UI components using Lip Gloss and Bubble Tea

### Key Components

**Parser (internal/parser/parser.go)**:
- `PRQualityBreakdown` struct with comprehensive scoring across 4 dimensions
- `ParsePRFAQ()` extracts sections with flexible header detection (H1/H2, various naming)
- `comprehensivePRAnalysis()` performs 100-point deterministic scoring:
  - Structure & Hook (30 pts): Headlines, newsworthy hooks, release dates
  - Content Quality (35 pts): 5 Ws coverage, credibility, structure
  - Professional Quality (20 pts): Tone, readability, fluff detection
  - Customer Evidence (15 pts): Quote quality with quantitative metrics
- Advanced features: quote metric extraction, marketing fluff detection, 5Ws validation, quote count optimization

**LLM Integration (internal/llm/llm.go)**:
- `AnalyzeSection()` provides qualitative feedback only (no score influence)
- Exponential backoff with jitter for API resilience
- Gracefully handles missing API keys

**UI Components (internal/ui/)**:
- `components.go`: Renders score breakdowns, strengths, improvements, quote analysis
- `model.go`: Bubble Tea model with tabbed interface (Overview, Breakdown, Quotes, AI Feedback)
- `styles.go`: Lip Gloss styling with color-coded scoring and progress bars

### Data Flow
1. CLI parses input file path and flags (`-no-tui`, `-report`)
2. Parser extracts structured sections and performs deterministic scoring
3. Interactive TUI displays results with tabbed navigation
4. Optional: LLM provides qualitative feedback (requires API key)
5. Results available in interactive UI, legacy stdout, or markdown report formats

## Dependencies

- **github.com/charmbracelet/bubbletea**: Terminal UI framework
- **github.com/charmbracelet/lipgloss**: Terminal styling
- **github.com/sashabaranov/go-openai**: OpenAI API client (optional)
- Standard library packages for file I/O, regex, and text analysis

## Recent Updates

### Session Summary (August 2025)
- **Fixed TUI alignment issues**: Resolved emoji rendering problems in warning boxes
- **Enhanced README.md**: Simplified structure, updated sample document references
- **Created high-quality example**: `example_prfaq_1.md` scoring 77/100 with:
  - Perfect headline scoring (10/10)
  - Complete newsworthy hook (15/15)
  - Multiple high-scoring customer quotes with quantitative metrics
  - Comprehensive company information and structured sections
- **Added quote count feedback**: Documents with >4 quotes now get guidance to limit to 3-4 focused testimonials

### Key Architectural Decisions
- **Hybrid scoring approach**: 100% deterministic algorithms for consistency, AI for insights
- **Interactive-first design**: TUI as primary interface, legacy modes for compatibility
- **Flexible document parsing**: Supports various header styles and naming conventions
- **Comprehensive feedback**: Detailed breakdowns help teams improve document quality

## Testing

Sample documents demonstrate scoring range:
- `example_prfaq_4.md`: ~51/100 (typical first draft)
- `example_prfaq_3.md`: ~38/100 (needs significant improvement)
- `example_prfaq_1.md`: ~77/100 (high-quality example with metrics-rich quotes)

The validator is intentionally demanding - scores above 80 are rare and indicate publication-ready quality.
