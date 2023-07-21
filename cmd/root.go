package cmd

import (
	JSON "encoding/json"
	"fmt"
	"gitmap/sum"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

var json bool
var limit int
var threshold int
var ignoreList []string = []string{
	".gitignore",
	".md",
	"go.mod",
	"go.sum",
	"package.json",
	"pnpm-lock.yaml",
	"poetry.lock",
}

var rootCmd = &cobra.Command{
	Use:   "gitmap",
	Short: "gitmap",
	Long:  "git heat map",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := gitLog(limit)

		if err != nil {
			log.Fatal(err)
		}

		content := string(data)

		lines := strings.Split(content, "\n")

		changes := sum.HeatMap(lines, threshold, ignoreList)
		sort.Slice(changes, func(i, j int) bool {
			return changes[i].Count < changes[j].Count
		})

		if json {
			printJson(changes)
		} else {
			printTable(changes)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&json, "json", false, "json")
	rootCmd.PersistentFlags().IntVar(&limit, "limit", 0, "limit")
	rootCmd.PersistentFlags().IntVar(&threshold, "threshold", 1, "threshold")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printTable(changes []sum.Change) {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Count"},
			{Align: simpletable.AlignCenter, Text: "Path"},
		},
	}
	for _, change := range changes {
		row := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%d", change.Count)},
			{Align: simpletable.AlignLeft, Text: change.Path},
		}

		table.Body.Cells = append(table.Body.Cells, row)
	}
	table.SetStyle(simpletable.StyleDefault)
	fmt.Println(table.String())
}

func printJson(changes []sum.Change) {
	data, err := JSON.Marshal(changes)

	if err != nil {
		log.Fatal(err)
	}

	content := string(data)
	fmt.Println(content)
}

func gitLog(limit int) ([]byte, error) {
	cwd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	args := []string{"--no-pager", "log", "--name-only", "--format="}

	if limit > 0 {
		args = append(args, fmt.Sprintf("-%d", limit))
	}

	cmd := exec.Command("git", args...)
	cmd.Dir = cwd

	return cmd.Output()
}
