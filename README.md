> [\!WARNING]
> This project is still under construction. Don't fork! :-)

# pr-faq-validator

AI-powered tool to analyze and score [PR-FAQ documents](https://github.com/bordenet/Engineering_Culture/blob/main/SDLC/The_PR-FAQ.md) with comprehensive quality metrics and intelligent section detection.

## Overview

pr-faq-validator analyzes PR-FAQ (Press Release - Frequently Asked Questions) documents and provides comprehensive quality scoring with detailed feedback. The tool uses intelligent fuzzy logic to automatically detect press release and FAQ sections, even when not explicitly labeled, and evaluates content against journalistic best practices.

## Features

### Core Analysis
- **Intelligent Section Detection**: Uses fuzzy logic to identify press releases and FAQs without explicit headers
- **Comprehensive Quality Scoring**: 100-point scoring system across multiple dimensions
- **Journalistic Standards**: Evaluates against press release best practices for media pickup

### Press Release Quality Assessment (100 points)
- **Structure & Hook (25 pts)**: Headline quality, newsworthy hook with specificity
- **Content Quality (35 pts)**: 5 Ws coverage, credibility, inverted pyramid structure  
- **Professional Quality (25 pts)**: Tone, readability, marketing fluff detection
- **Customer Evidence (15 pts)**: Quote quality with quantitative metrics

### Advanced Features
- **Fuzzy Content Recognition**: Detects press releases by date patterns, announcement language, and business wire format
- **Quote Metric Analysis**: Identifies and scores quantitative metrics in customer testimonials
- **Marketing Fluff Detection**: Flags hyperbolic language and unsubstantiated claims
- **5 Ws Validation**: Ensures WHO, WHAT, WHEN, WHERE, WHY are clearly addressed
- **AI-Powered Feedback**: GPT-4 provides expert analysis and improvement suggestions
- **Flexible Section Support**: Handles various naming conventions and document structures

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

The tool supports both explicitly labeled and implicitly structured PR-FAQ documents:

### Explicit Headers (Traditional)
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

### Implicit Recognition (Fuzzy Logic)
The tool can automatically detect sections without explicit headers:

```markdown
# Product Launch Announcement

**Press Release Content** (detected by patterns):
- Date/location format: "SEATTLE, WA — December 1, 2016"
- Announcement language: "Today announces", "unveils", "launches"
- Business Wire format and structure
- Customer quotes with attribution

**FAQ Content** (detected by structure):
- Question/answer format
- Headers containing "FAQ", "Questions", "Q&A"
- Sequential numbered questions
```

### Supported Section Recognition

**Press Release Detection**:
- Explicit headers: `## Press Release`
- Content patterns: Date formats, business wire structure, announcement language
- Context clues: Company names, product launches, executive quotes

**FAQ Detection**:
- Explicit headers: `## FAQ`, `## FAQs`, `## Frequently Asked Questions`, `## Q&A`
- Content structure: Question/answer pairs, numbered sections
- Flexible naming: "Internal FAQ", "Common Questions", "Questions and Answers"

**Metrics Detection**:
- Headers: `## Success Metrics`, `## Key Metrics`
- Content analysis for KPI and measurement discussions

## Output

The tool provides comprehensive analysis with detailed scoring:

### Quality Score Breakdown (0-100 points)
```
== Press Release Quality Score: 51/100 ==

== Quality Breakdown ==
Structure & Hook:      6/25 points
  - Headline Quality:   0/10
  - Newsworthy Hook:    6/15
Content Quality:       28/35 points
  - 5 Ws Coverage:      15/15
  - Credibility:        8/10
  - Structure:          5/10
Professional Quality:  22/25 points
  - Tone & Readability: 8/10
  - Fluff Avoidance:    14/15
Customer Evidence:     3/15 points
  - Quote Quality:      3/15
```

### Actionable Feedback
- **Strengths**: What's working well (e.g., "Clear 5 Ws coverage", "Avoids marketing fluff")
- **Areas for Improvement**: Specific recommendations (e.g., "Add quantifiable metrics to quotes")
- **Quote Analysis**: Individual scoring of customer testimonials with metric detection

### Detailed Analysis
1. **Section Recognition**: Automatically identified content types
2. **Metric Detection**: Quantitative data found in quotes (percentages, improvements, KPIs)
3. **Professional Assessment**: Tone, readability, and journalistic standards evaluation
4. **AI Feedback**: GPT-4 powered insights for each section

## Project Structure

```
├── main.go                    # Main application entry point
├── internal/
│   ├── llm/
│   │   └── llm.go            # OpenAI API integration with retry logic
│   └── parser/
│       └── parser.go         # Fuzzy logic parsing, quality analysis, and scoring
├── testdata/
│   └── example_prfaq.md      # Sample PR-FAQ document with mixed explicit/implicit sections
└── go.mod                    # Go module dependencies
```

## Quality Scoring Methodology

The tool evaluates press releases using journalistic best practices:

### Structure & Hook (25 points)
- **Headline Quality**: Length, strong verbs, specificity, avoids generic language
- **Newsworthy Hook**: Timeliness, specific metrics, problem-solving focus

### Content Quality (35 points)  
- **5 Ws Coverage**: WHO, WHAT, WHEN, WHERE, WHY clearly addressed
- **Credibility**: Professional tone, supporting details, data backing claims
- **Structure**: Inverted pyramid, proper lead paragraph, logical flow

### Professional Quality (25 points)
- **Tone & Readability**: Sentence length, active voice, jargon management
- **Fluff Avoidance**: Penalizes hyperbolic language, vague claims, emotional filler

### Customer Evidence (15 points)
- **Quote Quality**: Presence of quantitative metrics in customer testimonials
  - 0 points: "This new capability is great!"
  - High points: "This reduces our processing time by 50% and improves NPS by 12 points"

## Dependencies

- [go-openai](https://github.com/sashabaranov/go-openai) - OpenAI API client for Go

## License

This project is licensed under the terms specified in the LICENSE file.
