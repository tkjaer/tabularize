package tabularize

import (
	"strings"
)

type Table struct {
	useColumnSeparators bool
	useHeaderSeparator  bool
	usePadding          bool
	useRowSeparators    bool
	useTableBorders     bool
}

func NewTable() *Table {
	return &Table{
		useColumnSeparators: true,
		useHeaderSeparator:  true,
		usePadding:          true,
		useTableBorders:     false,
		useRowSeparators:    false,
	}
}

// find the max length of each column
func maxLengths(data [][]string) []int {
	maxLengths := make([]int, len(data[0]))
	for _, row := range data {
		for i, col := range row {
			if len(col) > maxLengths[i] {
				maxLengths[i] = len(col)
			}
		}
	}
	return maxLengths
}

func (t *Table) horizontalSeparator(sb *strings.Builder, maxLengths []int) {
	if t.useTableBorders {
		sb.WriteString("+")
	}
	for col, length := range maxLengths {
		if t.usePadding {
			length += 2
		}
		sb.WriteString(strings.Repeat("-", length))
		if !(col == len(maxLengths)-1) || t.useTableBorders {
			sb.WriteString("+")
		}
	}
	sb.WriteString("\n")
}

func (t *Table) padCell(cell string, maxLength int) string {
	pad := ""
	if t.usePadding {
		pad = " "
	}
	return pad + cell + strings.Repeat(" ", maxLength-len(cell)) + pad
}

// Tabluarize takes a 2D slice of strings and returns a string
// representation of the data in a table format.
func (t *Table) Render(data [][]string) string {
	if len(data) == 0 {
		return ""
	}
	var sb strings.Builder
	maxLengths := maxLengths(data)
	if t.useTableBorders {
		t.horizontalSeparator(&sb, maxLengths)
	}
	for row, rowdata := range data {
		if t.useTableBorders {
			sb.WriteString("|")
		}
		for col, cell := range rowdata {
			sb.WriteString(t.padCell(cell, maxLengths[col]))
			if t.useColumnSeparators && (col < len(maxLengths)-1 || t.useTableBorders) {
				sb.WriteString("|")
			} else if col < len(maxLengths)-1 {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
		if row == 0 && t.useHeaderSeparator {
			t.horizontalSeparator(&sb, maxLengths)
		}
	}
	if t.useTableBorders {
		t.horizontalSeparator(&sb, maxLengths)
	}
	return sb.String()
}
