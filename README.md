> [\!WARNING]
> This project is still under construction. Don't fork! :-)

# pr-faq-validator

LLM-based tool to provide feedback on [PR-FAQ documents](https://github.com/bordenet/Engineering_Culture/blob/main/SDLC/The_PR-FAQ.md) using OpenAI's GPT-4.

## Overview

pr-faq-validator analyzes PR-FAQ (Press Release - Frequently Asked Questions) documents and provides expert feedback on their clarity, completeness, and effectiveness. The tool parses markdown documents to extract key sections and uses AI to provide actionable improvement suggestions.

## Features

- **Markdown Parsing**: Automatically extracts sections including Press Release, FAQs, and Success Metrics
- **AI-Powered Analysis**: Uses GPT-4 to provide expert feedback on each section
- **Retry Logic**: Built-in exponential backoff for API resilience
- **Flexible Section Detection**: Supports various section naming conventions (FAQ/FAQs, Success Metrics/Key Metrics)

## Installation

### Prerequisites

- Go 1.24.5 or later
- OpenAI API key

### Build from Source

```bash
git clone https://github.com/bordenet/pr-faq-validator.git
cd pr-faq-validator
go mod tidy
go build
```

## Usage

### Setup

Set your OpenAI API key as an environment variable:

```bash
export OPENAI_API_KEY=your_openai_api_key_here
```

### Running the Tool

```bash
./pr-faq-validator -file path/to/your/prfaq.md
```

### Example

```bash
./pr-faq-validator -file testdata/example_prfaq.md
```

## Input Format

Your PR-FAQ document should be structured as a markdown file with the following sections:

```markdown
# Your PR-FAQ Title

## Press Release
Your press release content here...

## FAQ
Q: Your question here?
A: Your answer here...

## Success Metrics
Your success metrics here...
```

### Supported Section Names

- **Press Release**: `## Press Release`
- **FAQs**: `## FAQ` or `## FAQs`
- **Metrics**: `## Success Metrics` or `## Key Metrics`

## Output

The tool provides:

1. **Extracted Title**: The main document title
2. **Section Analysis**: For each detected section:
   - Expert feedback on clarity and effectiveness
   - Actionable improvement suggestions
   - Quality assessment

## Project Structure

```
├── main.go                    # Main application entry point
├── internal/
│   ├── llm/
│   │   └── llm.go            # OpenAI API integration with retry logic
│   └── parser/
│       └── parser.go         # Markdown parsing and section extraction
├── testdata/
│   └── example_prfaq.md      # Sample PR-FAQ document
└── go.mod                    # Go module dependencies
```

## Dependencies

- [go-openai](https://github.com/sashabaranov/go-openai) - OpenAI API client for Go

## License

This project is licensed under the terms specified in the LICENSE file.
