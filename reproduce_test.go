package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
)

func TestReproduceHashPlaceholder(t *testing.T) {
	const sheetName = "Sheet"
	tests := []struct {
		fileName   string
		expectRows [][]string
	}{
		{
			fileName: "OK.xlsx",
			expectRows: [][]string{
				{"1", "2"},
				{"0", "12345"},
			},
		},
		{
			fileName: "HashPlaceholder.xlsx",
			expectRows: [][]string{
				{"1,234", "1"},
				{"3", "2"},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.fileName, func(t *testing.T) {
			f, err := excelize.OpenFile(tc.fileName)
			if err != nil {
				t.Fatal(err)
			}

			actualRows, err := f.GetRows(sheetName)
			if err != nil {
				t.Fatal(err)
			}

			if !assert.Equal(t, len(tc.expectRows), len(actualRows), "height mismatch") {
				return
			}
			for i := range tc.expectRows {
				if !assert.Equalf(t, len(tc.expectRows[i]), len(actualRows[i]), "row[%d] width mismatch", i) {
					continue
				}
				for j := range tc.expectRows[i] {
					assert.Equalf(t, tc.expectRows[i][j], actualRows[i][j], "cell[%d][%d] mismatch", i, j)
				}
			}
		})
	}
}
