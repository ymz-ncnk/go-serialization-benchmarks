package main

import (
	"fmt"
	"os"

	"github.com/nao1215/markdown"
)

const (
	IntroductionFile = "gen/readme/templates/introduction.md"
	AllResultsFile   = "gen/readme/templates/all_results.md"
	WhereFile        = "gen/readme/templates/where.md"
	TaileFile        = "gen/readme/templates/tail.md"
	ReadmeFile       = "README.md"
)

// ReadmeGenerator encapsulates the logic for creating the README.md file.
type ReadmeGenerator struct {
}

// NewReadmeGenerator creates a new generator instance.
func NewReadmeGenerator() *ReadmeGenerator {
	return &ReadmeGenerator{}
}

// Generate creates the README.md file from the provided benchmark tables and template files.
func (g *ReadmeGenerator) Generate(safeTable, unsafeTable BenchmarkTable) error {
	// Create a new markdown builder
	md := markdown.NewMarkdown(os.Stdout) // We will later write to a file, not os.Stdout

	// 1. Append ./templates/introduction.md
	if err := g.appendTemplate(md, IntroductionFile); err != nil {
		return fmt.Errorf("could not append introduction: %w", err)
	}

	// 2. Add "## Benchmarks" header
	md.H2("Benchmarks")

	// 3. Append ./templates/all_results.md
	if err := g.appendTemplate(md, AllResultsFile); err != nil {
		return fmt.Errorf("could not append 'all_results' content: %w", err)
	}

	// 4. Add "## Fastest Safe" header
	md.H2("Fastest Safe")

	// 5. Add the safe benchmark table
	g.appendBenchmarkTable(md, safeTable)

	// 6. Add "## Fastest Unsafe" header
	md.H2("Fastest Unsafe")

	// 7. Add the unsafe benchmark table
	g.appendBenchmarkTable(md, unsafeTable)

	// 8. Append ./templates/where.md
	if err := g.appendTemplate(md, WhereFile); err != nil {
		return fmt.Errorf("could not append 'where' content: %w", err)
	}

	// 8. Append ./templates/tail.md
	if err := g.appendTemplate(md, TaileFile); err != nil {
		return fmt.Errorf("could not append 'tail' section: %w", err)
	}

	// Create and write to the README.md file
	readmeFile, err := os.Create(ReadmeFile)
	if err != nil {
		return fmt.Errorf("failed to create README.md: %w", err)
	}
	defer readmeFile.Close()

	if _, err := readmeFile.WriteString(md.String()); err != nil {
		return fmt.Errorf("failed to write to README.md: %w", err)
	}
	return nil
}

// appendTemplate reads a template file and appends its content as raw text to the markdown builder.
func (g *ReadmeGenerator) appendTemplate(md *markdown.Markdown, path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	// Add newlines to ensure separation between sections
	md.PlainText(string(content)).LF()
	return nil
}

// appendBenchmarkTable creates a markdown table from a BenchmarkTable.
func (g *ReadmeGenerator) appendBenchmarkTable(md *markdown.Markdown, tableData BenchmarkTable) {
	// Define table headers
	headers := []string{"Name", "sec/Op", "B/Size", "B/Op", "allocs/Op"}
	table := markdown.TableSet{
		Header: headers,
		Rows:   [][]string{},
	}
	// table := markdown.Table(len(headers), len(tableData))
	// table.SetHeaders(headers)

	// Populate table rows
	for _, r := range tableData {
		row := []string{
			r.Name,
			fmt.Sprintf("%.2f", r.NsOp),
			fmt.Sprintf("%.2f", r.BSize),
			fmt.Sprintf("%.2f", r.BOp),
			fmt.Sprintf("%.0f", r.AllocsOp),
		}
		table.Rows = append(table.Rows, row)
		// table.SetRow(i, row)
	}

	// Add the completed table to the markdown builder
	md.Table(table)
}
