package out

import (
	"encoding/json"
	"fmt"
	"gitmap/sum"
	"log"

	"github.com/alexeyco/simpletable"
)

func Table(changes []sum.Change) string {
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
	return table.String()
}

func Json(changes []sum.Change) string {
	data, err := json.MarshalIndent(changes, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	content := string(data)
	return content
}
