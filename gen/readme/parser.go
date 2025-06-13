package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type parseState int

const (
	stateIdle parseState = iota
	stateSecOp
	stateBSize
	stateBOp
	stateAllocsOp
)

// RawBenchmarkData holds all metrics for a single benchmark as strings.
// This is an intermediate representation used by the parser.
type RawBenchmarkData struct {
	FullName    string
	SecOpStr    string
	BSizeStr    string
	BOpStr      string
	AllocsOpStr string
}

// Parse reads the benchmark file content from an io.Reader and converts it into a
// structured map of raw string data. The map key is the full benchmark name
// (e.g., "Serializers/json-16").
func Parse(r io.Reader) (map[string]*RawBenchmarkData, error) {
	scanner := bufio.NewScanner(r)
	results := make(map[string]*RawBenchmarkData)
	var state parseState = stateIdle

	// Regex to capture "benchmarkName<space>value<space>±..."
	// It handles various units like µ, n, Ki.
	lineRegex := regexp.MustCompile(`^(\S+)\s+([0-9.µnKi]+).*`)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// State transitions based on headers
		if strings.Contains(line, "sec/op") {
			state = stateSecOp
			continue
		} else if strings.Contains(line, "B/size") {
			state = stateBSize
			continue
		} else if strings.Contains(line, "B/op") {
			state = stateBOp
			continue
		} else if strings.Contains(line, "allocs/op") {
			state = stateAllocsOp
			continue
		}

		// Ignore metadata, headers, and footers
		if strings.HasPrefix(line, "geomean") || strings.HasPrefix(line, "goos:") ||
			strings.HasPrefix(line, "goarch:") || strings.HasPrefix(line, "pkg:") ||
			strings.HasPrefix(line, "cpu:") || strings.HasPrefix(line, "│") ||
			strings.HasPrefix(line, "¹") {
			continue
		}

		matches := lineRegex.FindStringSubmatch(line)
		if len(matches) != 3 {
			continue // Not a data line we can parse
		}

		fullName := matches[1]
		value := matches[2]

		if _, ok := results[fullName]; !ok {
			results[fullName] = &RawBenchmarkData{FullName: fullName}
		}

		switch state {
		case stateSecOp:
			results[fullName].SecOpStr = value
		case stateBSize:
			results[fullName].BSizeStr = value
		case stateBOp:
			results[fullName].BOpStr = value
		case stateAllocsOp:
			results[fullName].AllocsOpStr = value
		default:
			// This can happen for lines before the first metric header, which is fine.
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return results, nil
}
