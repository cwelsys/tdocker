package ui

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/pivovarit/tdocker/internal/docker"
)

func screenRowOf(m App, name string) int {
	row := -1
	i := 0
	for _, line := range strings.Split(m.View().Content, "\n") {
		plain := ansiRe.ReplaceAllString(line, "")
		if !strings.Contains(plain, "svc-") {
			continue
		}
		for _, f := range strings.Fields(plain) {
			if f == name {
				row = i
			}
		}
		i++
	}
	return row
}

func TestCursorScreenRowDriftAcrossRefresh(t *testing.T) {
	cs := mkMixed(40)
	worst := 0
	for _, up := range []int{0, 1, 2, 3, 5, 8, 12, 18, 25, 30} {
		m := newWithClient(newStubClient(), "")
		m2, _ := m.Update(tea.WindowSizeMsg{Width: 120, Height: 20})
		m = m2.(App)
		m3, _ := m.Update(docker.ContainersMsg(cs))
		m = m3.(App)
		for i := 0; i < 38; i++ {
			mm, _ := m.Update(tea.KeyPressMsg{Code: 'j', Text: "j"})
			m = mm.(App)
		}
		for i := 0; i < up; i++ {
			mm, _ := m.Update(tea.KeyPressMsg{Code: 'k', Text: "k"})
			m = mm.(App)
		}
		c, _ := m.selectedContainer()
		before := screenRowOf(m, c.Names)
		m4, _ := m.Update(docker.ContainersMsg(cs))
		m = m4.(App)
		after := screenRowOf(m, c.Names)
		d := after - before
		if d < 0 {
			d = -d
		}
		if after < 0 {
			t.Errorf("up=%2d: cursor went OFFSCREEN", up)
		}
		if d > worst {
			worst = d
		}
		t.Logf("up=%2d  %-8s screenRow %2d -> %2d  drift=%d", up, c.Names, before, after, after-before)
	}
	t.Logf("worst drift: %d row(s)", worst)
	if worst > 1 {
		t.Errorf("drift %d exceeds 1 row", worst)
	}
}

func TestChangedDataStillRebuildsAndKeepsCursorVisible(t *testing.T) {
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
	sel, _ := m.selectedContainer()

	// flip a running container to exited, as a real state change would
	changed := append([]docker.Container(nil), cs...)
	var target string
	for i := range changed {
		if changed[i].State == docker.StateRunning && changed[i].ID != sel.ID {
			changed[i].State = "exited"
			changed[i].Status = "Exited (0) just now"
			target = changed[i].Names
			break
		}
	}
	m4, _ := m.Update(docker.ContainersMsg(changed))
	m = m4.(App)

	got := ""
	for _, c := range m.sorted {
		if c.Names == target {
			got = c.State
		}
	}
	if got != "exited" {
		t.Errorf("table did not pick up state change for %s: got %q", target, got)
	}
	if row := screenRowOf(m, sel.Names); row < 0 {
		t.Errorf("cursor %s went offscreen after a real change", sel.Names)
	}
}
