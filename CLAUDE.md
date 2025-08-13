# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based CLI tool that analyzes PR-FAQ (Press Release - Frequently Asked Questions) documents using OpenAI's GPT-4. The tool parses markdown files to extract key sections and provides AI-powered feedback on their quality and effectiveness.

## Development Commands

### Build and Run
```bash
# Install dependencies
go mod tidy

# Build the binary
go build

# Run the tool
./pr-faq-validator -file path/to/your/prfaq.md

# Test with example file
./pr-faq-validator -file testdata/example_prfaq.md
```

### Environment Setup
```bash
# Required: Set OpenAI API key
export OPENAI_API_KEY=your_openai_api_key_here
```

## Architecture

The codebase follows a clean modular structure:

- **main.go**: CLI entry point that handles argument parsing and orchestrates the analysis flow
- **internal/parser/**: Markdown parsing logic that extracts structured sections from PR-FAQ documents
- **internal/llm/**: OpenAI API integration with retry logic and error handling

### Key Components

**Parser (internal/parser/parser.go)**:
- `SpecSections` struct represents the parsed document structure
- `ParsePRFAQ()` function extracts title, Press Release, FAQs, Success Metrics, and other sections
- Supports flexible section naming (FAQ/FAQs, Success Metrics/Key Metrics)

**LLM Integration (internal/llm/llm.go)**:
- `AnalyzeSection()` function sends content to GPT-4 for analysis
- Implements exponential backoff with jitter for API resilience
- Handles retryable HTTP errors (429, 500, 502, 503, 504)
- Uses GPT-4o model for analysis

### Data Flow
1. CLI parses input file path
2. Parser extracts structured sections from markdown
3. Each section is analyzed by LLM with specific prompts
4. Feedback is formatted and displayed to user

## Dependencies

- **go-openai**: OpenAI API client library for Go
- Standard library packages for file I/O, CLI flags, and HTTP handling