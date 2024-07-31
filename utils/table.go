package utils

import (
	"github.com/jedib0t/go-pretty/table"
	"os"
	"coffee-online-cli/entity"
)

func LoyalTable(loyals []entity.Loyal) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(
		table.Row{
			"No.",
			"Nama Customer",
			"Total Pesanan",
			"Total Pengeluaran",
		},
	)
	for i, loyal := range loyals {
		t.AppendRow(
			table.Row{
				i + 1,
				loyal.Name,
				loyal.TotalOrder,
				"Rp " + PriceFormat(loyal.TotalSpending),
			},
		)
	}
	t.Render()
}
