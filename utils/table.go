package utils

import (
	"coffee-online-cli/entity"
	"os"

	"github.com/jedib0t/go-pretty/table"
)

func LoyalTable(loyals []entity.UserLoyal) {
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

func PopularProductTable(popular []entity.ProductPopular) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(
		table.Row{
			"No.",
			"Nama Produk",
			"Total Pesanan",
			"Total Pendapatan",
		},
	)
	for i, popular := range popular {
		t.AppendRow(
			table.Row{
				i + 1,
				popular.Name,
				popular.TotalOrder,
				"Rp " + PriceFormat(popular.TotalSpending),
			},
		)
	}
	t.Render()
}
