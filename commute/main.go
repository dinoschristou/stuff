package main

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"

	nr "github.com/martinsirbe/go-national-rail-client/nationalrail"
)

func main() {
	client, err := nr.NewClient(
		nr.AccessTokenOpt(""),
	)
	if err != nil {
		panic(err)
	}

	board, err := client.GetDeparturesWithDetails(nr.StationCodeOrpington, nr.NumRowsOpt(5))
	if err != nil {
		panic(err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Time", "Platform", "Departure Status", "Destination"})

	for _, s := range board.TrainServices {

		std := "?"
		if s.STD != "" {
			std = s.STD
		}

		platform := s.Platform

		etd := "?"
		if s.ETD != "" {
			etd = s.ETD
		}

		t.AppendRow(table.Row{std, platform, etd, fmt.Sprintf("%s [%s]", s.Destination.Name, s.Destination.CRS)})
	}

	t.Render()
}
