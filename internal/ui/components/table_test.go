package components_test

import (
	"testing"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/x/exp/golden"

	"rfz-cli/internal/ui/components"
)

func sampleColumns() []table.Column {
	return []table.Column{
		{Title: "Name", Width: 20},
		{Title: "Type", Width: 10},
		{Title: "Status", Width: 10},
	}
}

func sampleRows() []table.Row {
	return []table.Row{
		{"rfz-core", "Core", "Running"},
		{"rfz-api", "Plugin", "Stopped"},
		{"rfz-web", "Plugin", "Running"},
	}
}

func TestTable_Basic(t *testing.T) {
	tbl := components.NewTuiTable(components.TuiTableConfig{
		Columns: sampleColumns(),
		Rows:    sampleRows(),
		Height:  5,
		Focused: true,
	})
	output := tbl.View()
	golden.RequireEqual(t, []byte(output))
}

func TestTable_Unfocused(t *testing.T) {
	tbl := components.NewTuiTable(components.TuiTableConfig{
		Columns: sampleColumns(),
		Rows:    sampleRows(),
		Height:  5,
		Focused: false,
	})
	output := tbl.View()
	golden.RequireEqual(t, []byte(output))
}

func TestTable_Empty(t *testing.T) {
	output := components.TuiTableEmpty(sampleColumns(), 0)
	golden.RequireEqual(t, []byte(output))
}

func TestTable_SingleRow(t *testing.T) {
	tbl := components.NewTuiTable(components.TuiTableConfig{
		Columns: sampleColumns(),
		Rows: []table.Row{
			{"rfz-core", "Core", "Running"},
		},
		Height:  5,
		Focused: true,
	})
	output := tbl.View()
	golden.RequireEqual(t, []byte(output))
}

func TestTable_ManyRows(t *testing.T) {
	rows := []table.Row{
		{"component-01", "Core", "Running"},
		{"component-02", "Plugin", "Stopped"},
		{"component-03", "Core", "Running"},
		{"component-04", "Plugin", "Error"},
		{"component-05", "Core", "Running"},
		{"component-06", "Plugin", "Stopped"},
		{"component-07", "Core", "Running"},
		{"component-08", "Plugin", "Running"},
		{"component-09", "Core", "Stopped"},
		{"component-10", "Plugin", "Error"},
	}
	tbl := components.NewTuiTable(components.TuiTableConfig{
		Columns: sampleColumns(),
		Rows:    rows,
		Height:  5,
		Focused: true,
	})
	output := tbl.View()
	golden.RequireEqual(t, []byte(output))
}

func TestTable_LongValues(t *testing.T) {
	cols := []table.Column{
		{Title: "Name", Width: 15},
		{Title: "Description", Width: 20},
	}
	rows := []table.Row{
		{"rfz-core-module", "A very long description text that exceeds column width"},
		{"short", "OK"},
	}
	tbl := components.NewTuiTable(components.TuiTableConfig{
		Columns: cols,
		Rows:    rows,
		Height:  5,
		Focused: true,
	})
	output := tbl.View()
	golden.RequireEqual(t, []byte(output))
}

func TestTable_ZebraStripe(t *testing.T) {
	tbl := components.NewTuiTable(components.TuiTableConfig{
		Columns:     sampleColumns(),
		Rows:        sampleRows(),
		Height:      5,
		Focused:     true,
		ZebraStripe: true,
	})
	output := tbl.View()
	golden.RequireEqual(t, []byte(output))
}

func TestTable_Styles(t *testing.T) {
	styles := components.TuiTableStyles()
	// Verify styles are not zero values by rendering through them
	output := styles.Header.Render("Test Header") + "\n" +
		styles.Cell.Render("Test Cell") + "\n" +
		styles.Selected.Render("Test Selected")
	golden.RequireEqual(t, []byte(output))
}

func TestTable_DefaultHeight(t *testing.T) {
	tbl := components.NewTuiTable(components.TuiTableConfig{
		Columns: sampleColumns(),
		Rows:    sampleRows(),
		Focused: true,
	})
	// Default height should be 10
	if tbl.Height() != 10 {
		t.Errorf("expected default height 10, got %d", tbl.Height())
	}
}

func TestTable_CustomWidth(t *testing.T) {
	tbl := components.NewTuiTable(components.TuiTableConfig{
		Columns: sampleColumns(),
		Rows:    sampleRows(),
		Width:   80,
		Height:  5,
		Focused: true,
	})
	if tbl.Width() != 80 {
		t.Errorf("expected width 80, got %d", tbl.Width())
	}
}
