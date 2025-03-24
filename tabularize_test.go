package tabularize

import (
	"testing"
)

func TestRender_BasicTable(t *testing.T) {
	table := NewTable()
	data := [][]string{
		{"Name", "Age", "City"},
		{"Alice", "30", "New York"},
		{"Bob", "25", "Los Angeles"},
	}
	expected := `+-------+-----+-------------+
| Name  | Age | City        |
+-------+-----+-------------+
| Alice | 30  | New York    |
| Bob   | 25  | Los Angeles |
+-------+-----+-------------+
`
	table.useTableBorders = true
	result := table.Render(data)
	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

func TestRender_NoBordersOrSeparators(t *testing.T) {
	table := NewTable()
	data := [][]string{
		{"Name", "Age", "City"},
		{"Alice", "30", "New York"},
		{"Bob", "25", "Los Angeles"},
	}
	expected := `Name  Age City       
Alice 30  New York   
Bob   25  Los Angeles
`
	table.useColumnSeparators = false
	table.useHeaderSeparator = false
	table.usePadding = false
	table.useTableBorders = false
	result := table.Render(data)
	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

func TestRender_NoPadding(t *testing.T) {
	table := NewTable()
	data := [][]string{
		{"Name", "Age", "City"},
		{"Alice", "30", "New York"},
		{"Bob", "25", "Los Angeles"},
	}
	expected := `+-----+---+-----------+
|Name |Age|City       |
+-----+---+-----------+
|Alice|30 |New York   |
|Bob  |25 |Los Angeles|
+-----+---+-----------+
`
	table.useTableBorders = true
	table.usePadding = false
	result := table.Render(data)
	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

func TestRender_EmptyData(t *testing.T) {
	table := NewTable()
	data := [][]string{}
	expected := ""
	result := table.Render(data)
	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

func TestRender_HeaderSeparator(t *testing.T) {
	table := NewTable()
	data := [][]string{
		{"Name", "Age", "City"},
		{"Alice", "30", "New York"},
		{"Bob", "25", "Los Angeles"},
	}
	expected := ` Name  | Age | City        
-------+-----+-------------
 Alice | 30  | New York    
 Bob   | 25  | Los Angeles 
`
	table.useTableBorders = false
	table.useHeaderSeparator = true
	result := table.Render(data)
	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}
