package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	md "github.com/nao1215/markdown"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

const ReadmeFileName = "README.md"
const (
	BenchmarksSectionName        = "Benchmarks"
	FastestsSafeSubsectionName   = "Fastest Safe"
	FastestsUnsafeSubsectionName = "Fastest Unsafe"
	AllSubsectionName            = "All"
	FeaturesTableName            = "Features"
)

const (
	NameColumn            = "name"
	IterationsCountColumn = "iterations count"
	NsOpColumn            = "ns/op"
	BSizeColumn           = "B/size"
	BOpColumn             = "B/op"
	AllocsOpColumn        = "allocs/op"
)

const ResultsExplanations = ", where `iterations count`, `ns/op`, `B/op`, `allocs/op` are standard \n" +
	"`go test -bench=.` results and `B/size` - determines how many bytes were used on \n" +
	"average by the serializer to encode `Data`."

//go:generate go run ./...

func main() {
	out, err := RunBenchmarks()
	if err != nil {
		panic(err)
	}
	table, err := ParseBenchmarksTable(out)
	if err != nil {
		panic(err)
	}
	SortBenchmarksTable(table)

	src, err := os.Open(ReadmeFileName)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create("../../" + ReadmeFileName)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		panic(err)
	}

	AddBenchmarksSectionToReadme(dst, table)
	err = AddFastestSafeSubsectionToReadme(dst, table)
	if err != nil {
		panic(err)
	}
	err = AddFastestUnsafeSubsectionToReadme(dst, table)
	if err != nil {
		panic(err)
	}
	AddAllSubsectionToReadme(dst, table)
	err = AddFeaturesSectionToReadme(dst)
	if err != nil {
		panic(err)
	}
}

func RunBenchmarks() (out []byte, err error) {
	cmd := exec.Command("go", "test", "-bench=BenchmarkSerializers")
	cmd.Dir = "../"
	return cmd.Output()
}

func AddBenchmarksSectionToReadme(readmeFile *os.File, table BenchmarksTable) {
	md.NewMarkdown(readmeFile).LF().LF().
		H1(BenchmarksSectionName).LF().
		Build()
}

func AddFastestSafeSubsectionToReadme(readmeFile *os.File, table BenchmarksTable) (
	err error) {
	filter := func(item SerializerItem) bool {
		if strings.Contains(item.Name, string(serializer.Unsafe)) {
			return !strings.Contains(item.Name, string(serializer.NotUnsafe))
		}
		return false
	}
	return addFastestSubsectionToReadme(readmeFile, FastestsSafeSubsectionName,
		filter, table)
}

func AddFastestUnsafeSubsectionToReadme(readmeFile *os.File,
	table BenchmarksTable) (err error) {
	filter := func(item SerializerItem) bool {
		if strings.Contains(item.Name, string(serializer.Unsafe)) {
			return strings.Contains(item.Name, string(serializer.NotUnsafe))
		}
		return true
	}
	return addFastestSubsectionToReadme(readmeFile, FastestsUnsafeSubsectionName,
		filter, table)
}

func AddAllSubsectionToReadme(readmeFile *os.File, table BenchmarksTable) {
	tableSet := md.TableSet{
		Header: []string{NameColumn, IterationsCountColumn, NsOpColumn,
			BSizeColumn, BOpColumn, AllocsOpColumn},
		Rows: [][]string{},
	}
	for i := 0; i < len(table); i++ {
		item := table[i]
		tableSet.Rows = append(tableSet.Rows, []string{
			item.Name,
			fmt.Sprintf("%v", item.IterationsCount),
			fmt.Sprintf("%v", item.NsOp),
			fmt.Sprintf("%v", item.BSize),
			fmt.Sprintf("%v", item.BOp),
			fmt.Sprintf("%v", item.AllocsOp),
		})
	}
	md.NewMarkdown(readmeFile).LF().
		H2(AllSubsectionName).
		Table(tableSet).
		PlainText(ResultsExplanations).
		Build()
}

func addFastestSubsectionToReadme(readmeFile *os.File, sectionName string,
	filter func(item SerializerItem) bool,
	table BenchmarksTable,
) (err error) {
	tableSet := md.TableSet{
		Header: []string{NameColumn, IterationsCountColumn, NsOpColumn,
			BSizeColumn, BOpColumn, AllocsOpColumn},
		Rows: [][]string{},
	}
	var (
		fastests       = map[string]struct{}{}
		fastestsTable  = []SerializerItem{}
		item           SerializerItem
		serializerName string
	)
	for i := 0; i < len(table); i++ {
		item = table[i]
		if filter(item) {
			continue
		}
		serializerName, err = serializer.ResultName(item.Name).SerializerName()
		if err != nil {
			return
		}
		if _, pst := fastests[serializerName]; pst {
			continue
		}
		fastestsTable = append(fastestsTable, item)
		fastests[serializerName] = struct{}{}
	}
	var name string
	for i := 0; i < len(fastestsTable); i++ {
		item = fastestsTable[i]
		name, _ = serializer.ResultName(item.Name).SerializerName()
		tableSet.Rows = append(tableSet.Rows, []string{
			name,
			fmt.Sprintf("%v", item.IterationsCount),
			fmt.Sprintf("%v", item.NsOp),
			fmt.Sprintf("%v", item.BSize),
			fmt.Sprintf("%v", item.BOp),
			fmt.Sprintf("%v", item.AllocsOp),
		})
	}
	md.NewMarkdown(readmeFile).LF().
		H2(sectionName).
		Table(tableSet).
		Build()
	return
}

func AddFeaturesSectionToReadme(readmeFile *os.File) (err error) {
	featuresList, err := MakeFeaturesList()
	if err != nil {
		return
	}
	md.NewMarkdown(readmeFile).LF().LF().
		H1(FeaturesTableName).
		BulletList(featuresList...).LF().
		Build()
	return
}
