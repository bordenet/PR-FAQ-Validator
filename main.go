package main

import (
	"github.com/bordenet/pr-faq-validator/internal/llm"
	"github.com/bordenet/pr-faq-validator/internal/parser"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	inputFile := flag.String("file", "", "Path to the PR-FAQ markdown file")
	reportFile := flag.String("report", "", "Optional: Output markdown report file (default: stdout)")
	flag.Parse()

	if *inputFile == "" {
		log.Fatal("Please provide a markdown file with -file")
	}

	sections, err := parser.ParsePRFAQ(*inputFile)
	if err != nil {
		log.Fatalf("Failed to parse PR-FAQ: %v", err)
	}

	// Generate comprehensive markdown report
	report := parser.GenerateMarkdownReport(sections, sections.PRScore)
	
	// Output report
	if *reportFile != "" {
		err := writeReportToFile(*reportFile, report)
		if err != nil {
			log.Fatalf("Failed to write report: %v", err)
		}
		fmt.Printf("Report generated: %s\n", *reportFile)
		fmt.Printf("Overall Score: %d/100\n", sections.PRScore.OverallScore)
	} else {
		fmt.Print(report)
	}

	// Original detailed analysis follows for reference
	fmt.Printf("\n---\n\n== Detailed Analysis ==\n\n")
	fmt.Printf("== PR-FAQ Title ==\n%s\n\n", sections.Title)

	// Display comprehensive PR scoring results
	if sections.PressRelease != "" {
		fmt.Printf("== Press Release Quality Score: %d/100 ==\n\n", sections.PRScore.OverallScore)
		
		// Quality breakdown
		breakdown := sections.PRScore.QualityBreakdown
		fmt.Println("== Quality Breakdown ==")
		fmt.Printf("Structure & Hook:      %d/25 points\n", breakdown.HeadlineScore + breakdown.HookScore)
		fmt.Printf("  - Headline Quality:   %d/10\n", breakdown.HeadlineScore)
		fmt.Printf("  - Newsworthy Hook:    %d/15\n", breakdown.HookScore)
		fmt.Printf("Content Quality:       %d/35 points\n", breakdown.FiveWsScore + breakdown.CredibilityScore + breakdown.StructureScore)
		fmt.Printf("  - 5 Ws Coverage:      %d/15\n", breakdown.FiveWsScore)
		fmt.Printf("  - Credibility:        %d/10\n", breakdown.CredibilityScore)
		fmt.Printf("  - Structure:          %d/10\n", breakdown.StructureScore)
		fmt.Printf("Professional Quality:  %d/25 points\n", breakdown.ToneScore + breakdown.FluffScore)
		fmt.Printf("  - Tone & Readability: %d/10\n", breakdown.ToneScore)
		fmt.Printf("  - Fluff Avoidance:    %d/15\n", breakdown.FluffScore)
		fmt.Printf("Customer Evidence:     %d/15 points\n", breakdown.QuoteScore)
		fmt.Printf("  - Quote Quality:      %d/15\n\n", breakdown.QuoteScore)
		
		// Strengths
		if len(breakdown.Strengths) > 0 {
			fmt.Println("== Strengths ==")
			for _, strength := range breakdown.Strengths {
				fmt.Printf("✓ %s\n", strength)
			}
			fmt.Println()
		}
		
		// Issues to address
		if len(breakdown.Issues) > 0 {
			fmt.Println("== Areas for Improvement ==")
			for _, issue := range breakdown.Issues {
				fmt.Printf("⚠ %s\n", issue)
			}
			fmt.Println()
		}
		
		// Detailed quote analysis if present
		if len(sections.PRScore.MetricDetails) > 0 {
			fmt.Printf("== Quote Analysis (%d quotes found) ==\n", sections.PRScore.TotalQuotes)
			for i, detail := range sections.PRScore.MetricDetails {
				fmt.Printf("\nQuote %d (Score: %d/10):\n", i+1, detail.Score)
				fmt.Printf("\"%s\"\n", detail.Quote)
				if len(detail.Metrics) > 0 {
					fmt.Printf("Metrics detected: %v\n", detail.Metrics)
					fmt.Printf("Metric types: %v\n", detail.MetricTypes)
				} else {
					fmt.Println("No quantitative metrics detected")
				}
			}
			fmt.Println()
		}

		fmt.Println("Analyzing Press Release...")
		feedback, err := llm.AnalyzeSection("Press Release", sections.PressRelease)
		if err != nil {
			log.Fatalf("LLM error: %v", err)
		}
		fmt.Printf("== Feedback for Press Release ==\n%s\n\n", feedback.Comments)
	}

	if sections.FAQs != "" {
		fmt.Println("Analyzing FAQs...")
		feedback, err := llm.AnalyzeSection("FAQs", sections.FAQs)
		if err != nil {
			log.Fatalf("LLM error: %v", err)
		}
		fmt.Printf("== Feedback for FAQs ==\n%s\n\n", feedback.Comments)
	}
}

func writeReportToFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	_, err = file.WriteString(content)
	return err
}

func countNumberedQuestions(faqContent string) int {
	count := 0
	lines := strings.Split(faqContent, "\n")
	for _, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "## ") {
			header := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), "## "))
			if matched, _ := regexp.MatchString(`^\d+\.`, header); matched {
				count++
			}
		}
	}
	return count
}
