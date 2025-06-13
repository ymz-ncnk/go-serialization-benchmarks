package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Processor encapsulates the logic to process raw benchmark data into final tables.
type Processor struct {
	nameRegex *regexp.Regexp
}

// NewProcessor creates a new processor with a compiled regex for performance.
func NewProcessor() *Processor {
	return &Processor{
		nameRegex: regexp.MustCompile(`^Serializers/([a-zA-Z0-9_]+)`),
	}
}

// Process converts raw parsed data into two tables: one for the fastest safe
// benchmarks and one for the fastest unsafe benchmarks per serializer family.
func (p *Processor) Process(rawData map[string]*RawBenchmarkData) (BenchmarkTable, BenchmarkTable, error) {
	fastestSafe := make(map[string]*BenchmarkResult)
	fastestUnsafe := make(map[string]*BenchmarkResult)

	for _, raw := range rawData {
		// 1. Parse all values from strings to numbers
		result, err := p.parseRawData(raw)
		if err != nil {
			// Silently skip lines that fail parsing (e.g., geomean remnants)
			continue
		}

		// 2. Determine if it's safe or unsafe
		isUnsafe := strings.Contains(raw.FullName, "+unsafe") ||
			strings.Contains(raw.FullName, "+unsafestr") ||
			strings.Contains(raw.FullName, "+unsafeunm")

		// 3. Update the corresponding map if this result is faster
		if isUnsafe {
			p.updateFastest(fastestUnsafe, result)
		} else {
			p.updateFastest(fastestSafe, result)
		}
	}

	return p.mapToSortedTable(fastestSafe), p.mapToSortedTable(fastestUnsafe), nil
}

// updateFastest checks if the new result is faster than the existing one for the
// same serializer name and updates the map if it is.
func (p *Processor) updateFastest(fastestMap map[string]*BenchmarkResult, result *BenchmarkResult) {
	existing, ok := fastestMap[result.Name]
	if !ok || result.NsOp < existing.NsOp {
		fastestMap[result.Name] = result
	}
}

// parseRawData converts a single RawBenchmarkData struct into a BenchmarkResult.
func (p *Processor) parseRawData(raw *RawBenchmarkData) (*BenchmarkResult, error) {
	matches := p.nameRegex.FindStringSubmatch(raw.FullName)
	if len(matches) < 2 {
		return nil, fmt.Errorf("could not extract short name from %s", raw.FullName)
	}
	name := matches[1]

	nsOp, err := p.parseTime(raw.SecOpStr)
	if err != nil {
		return nil, fmt.Errorf("could not parse NsOp '%s': %w", raw.SecOpStr, err)
	}

	bSize, err := p.parseValue(raw.BSizeStr)
	if err != nil {
		return nil, fmt.Errorf("could not parse BSize '%s': %w", raw.BSizeStr, err)
	}

	bOp, err := p.parseValue(raw.BOpStr)
	if err != nil {
		return nil, fmt.Errorf("could not parse BOp '%s': %w", raw.BOpStr, err)
	}

	allocsOp, err := p.parseValue(raw.AllocsOpStr)
	if err != nil {
		return nil, fmt.Errorf("could not parse AllocsOp '%s': %w", raw.AllocsOpStr, err)
	}

	return &BenchmarkResult{
		Name:     name,
		NsOp:     nsOp,
		BSize:    bSize,
		BOp:      bOp,
		AllocsOp: allocsOp,
	}, nil
}

// parseTime converts time strings like "2.847µ" or "154.3n" into nanoseconds.
func (p *Processor) parseTime(s string) (float64, error) {
	if s == "" {
		return 0, nil
	}
	if strings.HasSuffix(s, "µ") {
		val, err := strconv.ParseFloat(strings.TrimSuffix(s, "µ"), 64)
		return val * 1000, err
	}
	if strings.HasSuffix(s, "n") {
		return strconv.ParseFloat(strings.TrimSuffix(s, "n"), 64)
	}
	// Fallback for values without units (e.g., seconds)
	val, err := strconv.ParseFloat(s, 64)
	return val * 1e9, err // Assume seconds and convert to nanoseconds
}

// parseValue handles generic numeric values, including "Ki" suffix for kilobytes.
func (p *Processor) parseValue(s string) (float64, error) {
	if s == "" {
		return 0, nil
	}
	if strings.HasSuffix(s, "Ki") {
		val, err := strconv.ParseFloat(strings.TrimSuffix(s, "Ki"), 64)
		return val * 1024, err
	}
	return strconv.ParseFloat(s, 64)
}

// mapToSortedTable converts the result map to a slice and sorts it by NsOp
// (time per operation, ascending).
func (p *Processor) mapToSortedTable(m map[string]*BenchmarkResult) BenchmarkTable {
	table := make(BenchmarkTable, 0, len(m))
	for _, v := range m {
		table = append(table, *v)
	}
	sort.Slice(table, func(i, j int) bool {
		// Sort by NsOp (nanoseconds per operation) ascending.
		return table[i].NsOp < table[j].NsOp
	})
	return table
}
