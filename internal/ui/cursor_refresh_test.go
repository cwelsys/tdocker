package ui

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/pivovarit/tdocker/internal/docker"
)

func cursorVisible(m App) (bool, string) {
	c, ok := m.selectedContainer()
	if !ok {
		return false, "<no selection>"
	}
	for _, line := range strings.Split(m.View().Content, "\n") {
		if strings.Contains(ansiRe.ReplaceAllString(line, ""), c.ID) {
			return true, c.Names
		}
	}
	return false, c.Names
}

func visibleNames(m App) []string {
	var out []string
	for _, line := range strings.Split(m.View().Content, "\n") {
		plain := ansiRe.ReplaceAllString(line, "")
		for _, f := range strings.Fields(plain) {
			if strings.HasPrefix(f, "svc-") {
				out = append(out, f)
			}
		}
	}
	return out
}

func TestCursorStaysVisibleAcrossRefresh(t *testing.T) {
	m := newWithClient(newStubClient(), "")
	m2, _ := m.Update(tea.WindowSizeMsg{Width: 120, Height: 20})
	m = m2.(App)
	cs := mkMixed(40)
	m3, _ := m.Update(docker.ContainersMsg(cs))
	m = m3.(App)

	for i := 0; i < 25; i++ {
		mm, _ := m.Update(tea.KeyPressMsg{Code: 'j', Text: "j"})
		m = mm.(App)
	}
	t.Logf("window before: %v", visibleNames(m))
	vis, name := cursorVisible(m)
	t.Logf("before refresh: cursor=%d on %s visible=%v", m.table.Cursor(), name, vis)

	// identical data, exactly what the 10s periodic refresh delivers
	m4, _ := m.Update(docker.ContainersMsg(cs))
	m = m4.(App)
	t.Logf("window after:  %v", visibleNames(m))
	vis2, name2 := cursorVisible(m)
	t.Logf("after refresh:  cursor=%d on %s visible=%v", m.table.Cursor(), name2, vis2)
	if !vis2 {
		t.Errorf("cursor went offscreen after refresh (container %s)", name2)
	}
}

func TestCursorStableAfterScrollUpRefresh(t *testing.T) {
	m := newWithClient(newStubClient(), "")
	m2, _ := m.Update(tea.WindowSizeMsg{Width: 120, Height: 20})
	m = m2.(App)
	cs := mkMixed(40)
	m3, _ := m.Update(docker.ContainersMsg(cs))
	m = m3.(App)

	for i := 0; i < 38; i++ {
		mm, _ := m.Update(tea.KeyPressMsg{Code: 'j', Text: "j"})
		m = mm.(App)
	}
	for i := 0; i < 18; i++ {
		mm, _ := m.Update(tea.KeyPressMsg{Code: 'k', Text: "k"})
		m = mm.(App)
	}
	before := visibleNames(m)
	_, name := cursorVisible(m)
	t.Logf("before: cursor=%d on %s window=%v", m.table.Cursor(), name, before)

	m4, _ := m.Update(docker.ContainersMsg(cs))
	m = m4.(App)
	after := visibleNames(m)
	vis, name2 := cursorVisible(m)
	t.Logf("after:  cursor=%d on %s window=%v", m.table.Cursor(), name2, after)
	if !vis {
		t.Errorf("cursor offscreen after refresh (%s)", name2)
	}
}
