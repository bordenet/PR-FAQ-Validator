
# pr-faq-validator

AI-assisted tool to analyze and score [PR-FAQ documents](https://github.com/bordenet/Engineering_Culture/blob/main/SDLC/The_PR-FAQ.md) with comprehensive quality metrics.

## Overview

pr-faq-validator analyzes PR-FAQ (Press Release - Frequently Asked Questions) documents and provides comprehensive quality scoring with detailed feedback. The tool automatically detects press release and FAQ sections and evaluates content against journalistic best practices.

## Features

- **100-point scoring system** across 4 key dimensions
- **Automatic section detection** for any document structure
- **Press release standards** - evaluates whether documents read like authentic announcements
- **Quote metric analysis** - identifies quantitative data in customer testimonials
- **Marketing fluff detection** - flags hyperbolic language and vague claims
- **5 Ws validation** - ensures WHO, WHAT, WHEN, WHERE, WHY coverage
- **Interactive terminal UI** with detailed breakdowns and suggestions
- **AI-powered feedback** (optional) for expert analysis

## Installation

```bash
git clone https://github.com/bordenet/pr-faq-validator.git
cd pr-faq-validator
go mod tidy
go build
```

**Requirements:** Go 1.24.5+, OpenAI API key (optional, for AI feedback)

## Usage

```bash
# Set API key for AI feedback (optional)
export OPENAI_API_KEY=your_openai_api_key_here

# Analyze a document
./pr-faq-validator -file path/to/your/prfaq.md
```

### Examples

Analyze any of the included sample documents:

```bash
./pr-faq-validator -file testdata/example_prfaq_1.md
./pr-faq-validator -file testdata/example_prfaq_2.txt  
./pr-faq-validator -file testdata/example_prfaq_3.md
./pr-faq-validator -file testdata/example_prfaq_4.md
```

## Input Format

Works with any document structure. Recommended format:

```markdown
# Your PR-FAQ Title

## Press Release
Your press release content...

## FAQ
Q: Question?
A: Answer...
```

Automatically detects sections regardless of headers ("Press Release", "Announcement", "Q&A", etc.).

## Output

Provides interactive terminal UI with:
- **Score breakdown** across 4 categories (Structure, Content, Professional, Evidence)
- **Strengths & improvements** with specific recommendations
- **Quote analysis** with individual scoring and metric detection
- **AI feedback** for detailed insights (with OpenAI API key)


## Quality Scoring Methodology

### Scoring Methodology

**Deterministic Scoring (100% of numerical score):** Rule-based algorithms analyze text patterns for consistent, reproducible results. No AI influence on scores.

**AI Feedback (qualitative only):** Optional GPT-4 insights for improvement suggestions.

**Scoring Breakdown:**
- **Structure & Hook (30 pts):** Headline quality, newsworthy hook, release date
- **Content Quality (35 pts):** 5 Ws coverage, credibility, structure
- **Professional Quality (20 pts):** Tone, readability, fluff detection
- **Customer Evidence (15 pts):** Quote quality with quantitative metrics

**Note:** Intentionally demanding grader - perfect scores are rare. Focus on actionable feedback to improve document quality.

## License

MIT License - see LICENSE file for details.
