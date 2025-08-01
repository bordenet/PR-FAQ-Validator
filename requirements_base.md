# Purpose of file
Yes-- this file has been included on purpose to illustrate vibe coding from a guy who doesn't have a lot of free time to code from scratch. Here, I'm using ChatGPT to bootstrap project creation.
# Body
🧾 Requirements Doc: PR-FAQ / Spec Reviewer AI

⸻

📌 One-Liner

Build a developer-first tool that uses an LLM to analyze product documents (PR-FAQs, specs, RFCs) for clarity, completeness, and rigor—suggesting improvements and flagging ambiguous or vague content.

⸻

🎯 Goals
	•	Help PMs and engineers improve written specs via LLM-generated feedback
	•	Encourage a culture of clarity, ownership, and measurable outcomes
	•	Provide a structured critique rather than a fuzzy rewrite
	•	Fit into a CLI-based SDLC toolchain; possible GitHub Action or VS Code extension later

⸻

💼 Users
	•	Primary: Staff+ PMs, Staff+ Engineers writing PR-FAQs/specs
	•	Secondary: VPs/Directors reviewing early planning docs

⸻

🧠 Core Use Case

User runs:

prfaq-reviewer input.md

And gets back:

== Summary ==

This PR-FAQ presents a compelling customer problem, but lacks measurable success criteria and assumes internal technical knowledge.

== Flags ==

[Line 12] "Significant performance improvement" → vague term  
[Line 47] "Some teams may..." → who exactly?  

== Suggestions ==

- Clarify the FAQ response to "What does success look like?"
- Add real metrics: latency? conversion? revenue?

== Overall Readiness Score: 7.5 / 10 ==


⸻

🧱 Key Components

⸻

✅ Input
	•	One .md file (PR-FAQ or spec)
	•	Optionally: section heuristics to help LangGraph route analysis (see below)

⸻

✅ Output
	•	Structured report, including:
	•	Summary
	•	Flags (vague, inconsistent, missing ownership, passive voice, weasel words, unclear success metrics)
	•	Suggestions (what to add/clarify)
	•	Optional score (for CI)

⸻

✅ Functional Requirements

Feature	Description
Markdown parsing	Parse sections: Press Release, FAQs, Metrics
Content analysis	LLM checks for vague terms, missing owners, unclear outcomes
Heuristic tagging	Use regex/NLP to pre-label weasel words or passive constructions
LLM-based critique	Summarize issues with clarity, conciseness, impact
CLI interface	Run from terminal and output to stdout or JSON
Output mode	Color-coded terminal output + --json export option


⸻

✅ Non-Functional Requirements
	•	Low latency (under 10s per spec)
	•	Works offline with cache or mock LLM
	•	Composability — extensible with new checks, possibly as plug-ins

⸻

🧠 LangGraph & MCP: Does it Fit?

✅ When LangGraph makes sense:
	•	You want multi-step analysis, with routing logic (e.g. split PR and FAQ and Metrics into separate paths)
	•	You may want modularity: swap in different agents per section
	•	You want state tracking or conversational memory across chunks

Verdict: LangGraph is a good fit if you want composable, section-specific analysis and reusability.

⸻

🪜 Proposed MVP Plan

⸻

🗂️ Phase 1: Skeleton MVP (Tonight)
	•	cli.py script: input .md, output text
	•	Use OpenAI or Groq LLM for a single-shot prompt:
“Act as a Staff PM reviewing this PR-FAQ…”
	•	Parse sections with simple regex
	•	Return structured summary + flags + suggestions

⸻

🏗️ Phase 2: LangGraph Modular Pass
	•	Break into nodes:
	•	parse_sections
	•	analyze_press_release
	•	analyze_faq
	•	analyze_success_metrics
	•	merge_results
	•	Add scoring logic or plug-in slots for extra checks (e.g., detect_passive_voice)
	•	Export to JSON for GitHub bot compatibility

⸻

🧰 Phase 3: CI + Collaboration-Ready
	•	Wrap as GitHub Action
	•	Add inline comment suggestions (via PR)
	•	Slack/Teams integration optional

⸻

🧠 Prompt Design (First Draft)

System prompt:

You are a Staff Product Manager reviewing a written product document. Your job is to flag areas that are vague, unclear, or lack measurable success criteria. Provide actionable suggestions and identify language that introduces uncertainty or lack of ownership.

User prompt (template):

Here's a PR-FAQ document. Review each section (Press Release, FAQs, Success Metrics) and return:

1. Summary
2. Flags (with line # if possible)
3. Suggestions
4. Overall readiness score (1–10)

```markdown
<contents of input.md>
