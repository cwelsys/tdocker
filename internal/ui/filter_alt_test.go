package ui

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/pivovarit/tdocker/internal/docker"
)

func TestFEntersFilterAndDoesNotPage(t *testing.T) {
	m := newWithClient(newStubClient(), "")
	m2, _ := m.Update(tea.WindowSizeMsg{Width: 120, Height: 20})
	m = m2.(App)
	m3, _ := m.Update(docker.ContainersMsg(mkMixed(40)))
	m = m3.(App)

	before := m.table.Cursor()
	m4, _ := m.Update(tea.KeyPressMsg{Code: 'f', Text: "f"})
	m = m4.(App)

	if !m.filtering {
		t.Errorf("f did not enter filter mode")
	}
	if got := m.table.Cursor(); got != before {
		t.Errorf("f still paged the table: cursor %d -> %d", before, got)
	}
	// typing after f should build the query, not trigger actions
	m5, _ := m.Update(tea.KeyPressMsg{Code: 'a', Text: "a"})
	m = m5.(App)
	if m.filterQuery != "a" {
		t.Errorf("filterQuery = %q, want %q", m.filterQuery, "a")
	}
}
