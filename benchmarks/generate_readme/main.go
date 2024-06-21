package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	md "github.com/nao1215/markdown"
)

const ReadmeFileName = "README.md"
const (
	BenchmarksTableName = "Benchmarks"
	FeaturesTableName   = "Features"
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
		H1(BenchmarksTableName).
		Table(tableSet).
		PlainText(ResultsExplanations).LF().
		Build()
}

func AddFeaturesSectionToReadme(readmeFile *os.File) (err error) {
	md.NewMarkdown(readmeFile).LF().
		H1(FeaturesTableName).
		BulletList(MakeFeaturesList()...).LF().
		Build()
	return
}
