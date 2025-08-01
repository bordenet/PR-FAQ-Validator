package main

import (
	"github.com/bordenet/pr-faq-validator/internal/llm"
	"github.com/bordenet/pr-faq-validator/internal/parser"
	"flag"
	"fmt"
	"log"
)

func main() {
	inputFile := flag.String("file", "", "Path to the PR-FAQ markdown file")
	flag.Parse()

	if *inputFile == "" {
		log.Fatal("Please provide a markdown file with -file")
	}

	sections, err := parser.ParsePRFAQ(*inputFile)
	if err != nil {
		log.Fatalf("Failed to parse PR-FAQ: %v", err)
	}

	fmt.Printf("== PR-FAQ Title ==\n%s\n\n", sections.Title)

	if sections.PressRelease != "" {
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
