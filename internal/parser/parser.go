package parser

import (
	"bufio"
	"os"
	"strings"
)

type SpecSections struct {
	Title         string
	PressRelease  string
	FAQs          string
	Metrics       string
	OtherSections map[string]string
}

// ParsePRFAQ reads a markdown file and extracts key sections
func ParsePRFAQ(path string) (*SpecSections, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sections := &SpecSections{
		OtherSections: make(map[string]string),
	}

	var currentSection string
	var sectionBuffer strings.Builder
	var titleSet bool

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Extract the title (first H1)
		if !titleSet && strings.HasPrefix(line, "# ") {
			sections.Title = strings.TrimPrefix(line, "# ")
			titleSet = true
			continue
		}

		// Detect section heading
		if strings.HasPrefix(line, "## ") {
			// Save the previous section's content
			if currentSection != "" {
				content := strings.TrimSpace(sectionBuffer.String())
				switch currentSection {
				case "Press Release":
					sections.PressRelease = content
				case "FAQ", "FAQs":
					sections.FAQs = content
				case "Success Metrics", "Key Metrics":
					sections.Metrics = content
				default:
					sections.OtherSections[currentSection] = content
				}
				sectionBuffer.Reset()
			}

			currentSection = strings.TrimSpace(strings.TrimPrefix(line, "## "))
			continue
		}

		// Accumulate section content
		if currentSection != "" {
			sectionBuffer.WriteString(line + "\n")
		}
	}

	// Capture last section
	if currentSection != "" {
		content := strings.TrimSpace(sectionBuffer.String())
		switch currentSection {
		case "Press Release":
			sections.PressRelease = content
		case "FAQ", "FAQs":
			sections.FAQs = content
		case "Success Metrics", "Key Metrics":
			sections.Metrics = content
		default:
			sections.OtherSections[currentSection] = content
		}
	}

	return sections, nil
}