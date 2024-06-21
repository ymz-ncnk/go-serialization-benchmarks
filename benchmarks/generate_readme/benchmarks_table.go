package main

import (
	"bufio"
	"bytes"
	"errors"
	"regexp"
	"sort"
	"strconv"
)

var (
	NameRegexp            = regexp.MustCompile(`BenchmarkSerializers/(.+)-\d+`)
	IterationsCountRegexp = regexp.MustCompile(`\s(\d+)\s`)
	NsOpRegexp            = regexp.MustCompile(`\s(\d+.\d+) ns/op`)
	BSizeRegexp           = regexp.MustCompile(`\s(\d+.\d+) B/size`)
	BOpRegexp             = regexp.MustCompile(`\s(\d+) B/op`)
	AllocsOpRegexp        = regexp.MustCompile(`\s(\d+) allocs/op`)
)

type BenchmarksTable []SerializerItem

type SerializerItem struct {
	Name            string
	IterationsCount int
	NsOp            float64
	BSize           float64
	BOp             int
	AllocsOp        int
}

func SortBenchmarksTable(table BenchmarksTable) {
	sort.Slice(table, func(i, j int) bool {
		return table[i].NsOp < table[j].NsOp
	})
}

func ParseBenchmarksTable(bs []byte) (table BenchmarksTable, err error) {
	table = []SerializerItem{}
	var (
		scanner = bufio.NewScanner(bytes.NewReader(bs))
		text    string
		name    string
	)
	for scanner.Scan() {
		text = scanner.Text()
		if name, err = ParseName(text); err != nil {
			err = nil
			continue
		}
		item := SerializerItem{Name: name}
		item.IterationsCount, err = ParseIterationsCount(text)
		if err != nil {
			return
		}
		item.NsOp, err = ParseNsOp(text)
		if err != nil {
			return
		}
		item.BSize, err = ParseBSize(text)
		if err != nil {
			return
		}
		item.BOp, err = ParseBOp(text)
		if err != nil {
			return
		}
		item.AllocsOp, err = ParseAllocsOp(text)
		if err != nil {
			return
		}
		table = append(table, item)
	}
	return
}

func ParseName(str string) (name string, err error) {
	strs := NameRegexp.FindStringSubmatch(str)
	if len(strs) != 2 {
		err = errors.New("NameRegexp failed to find value")
		return
	}
	name = strs[1]
	return
}

func ParseIterationsCount(str string) (iterationsCount int, err error) {
	strs := IterationsCountRegexp.FindStringSubmatch(str)
	if len(strs) != 2 {
		err = errors.New("IterationsCountRegexp failed to find value")
		return
	}
	return strconv.Atoi(strs[1])
}

func ParseNsOp(str string) (nsOp float64, err error) {
	strs := NsOpRegexp.FindStringSubmatch(str)
	if len(strs) != 2 {
		err = errors.New("NsOpRegexp failed to find value")
		return
	}
	return strconv.ParseFloat(strs[1], 64)
}

func ParseBSize(str string) (bSize float64, err error) {
	strs := BSizeRegexp.FindStringSubmatch(str)
	if len(strs) != 2 {
		err = errors.New("BSizeRegexp failed to find value")
		return
	}
	return strconv.ParseFloat(strs[1], 64)
}

func ParseBOp(str string) (bOp int, err error) {
	strs := BOpRegexp.FindStringSubmatch(str)
	if len(strs) != 2 {
		err = errors.New("BOpRegexp failed to find value")
		return
	}
	return strconv.Atoi(strs[1])
}

func ParseAllocsOp(str string) (allocsOp int, err error) {
	strs := AllocsOpRegexp.FindStringSubmatch(str)
	if len(strs) != 2 {
		err = errors.New("AllocsOpRegexp failed to find value")
		return
	}
	return strconv.Atoi(strs[1])
}
