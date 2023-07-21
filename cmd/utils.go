package cmd

import (
	JSON "encoding/json"
	"fmt"
	"gitmap/sum"
	"log"
	"os"
	"os/exec"

	"github.com/alexeyco/simpletable"
)

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
