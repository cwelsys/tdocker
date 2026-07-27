package ui

import (
	"fmt"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/pivovarit/tdocker/internal/docker"
)

var dimGray = seqDim // stoppedRowStyle foreground

func mkMixed(n int) []docker.Container {
	cs := make([]docker.Container, n)
	for i := range cs {
		state, status := docker.StateRunning, "Up 2 hours"
		if i%3 == 0 { // interleave exited so drift is always detectable
			state, status = "exited", "Exited (0) 1 hour ago"
		}
		cs[i] = docker.Container{
			ID:     fmt.Sprintf("%012x", 0xaa0000+i),
			Names:  "svc-" + string(rune('a'+i/26)) + string(rune('a'+i%26)),
			Image:  "img:latest",
			State:  state,
			Status: status,
		}
	}
	return docker.Sort(cs)
}

// checkStyling verifies every visible row's dimness matches its real state.
func checkStyling(t *testing.T, m App, tag string) int {
	t.Helper()
	byName := map[string]docker.Container{}
	for _, c := range m.sorted {
		byName[c.Names] = c
	}
	cursorName := ""
	if c, ok := m.selectedContainer(); ok {
		cursorName = c.Names
	}
	bad := 0
	for _, line := range strings.Split(m.View().Content, "\n") {
		plain := ansiRe.ReplaceAllString(line, "")
		var name string
		for _, f := range strings.Fields(plain) {
			if strings.HasPrefix(f, "svc-") {
				name = f
			}
		}
		if name == "" {
			continue
		}
		c := byName[name]
		isDim := strings.Contains(line, dimGray)
		wantDim := c.State != docker.StateRunning && name != cursorName
		if isDim != wantDim {
			bad++
			t.Errorf("%s: %s state=%q cursor=%v -> dim=%v want=%v", tag, name, c.State, name == cursorName, isDim, wantDim)
		}
	}
	return bad
}

func TestRowStylingMatchesStateDuringTraversal(t *testing.T) {
	m := newWithClient(newStubClient(), "")
	m2, _ := m.Update(tea.WindowSizeMsg{Width: 120, Height: 20})
	m = m2.(App)
	m3, _ := m.Update(docker.ContainersMsg(mkMixed(40)))
	m = m3.(App)

	for i := 0; i < 39; i++ {
		mm, _ := m.Update(tea.KeyPressMsg{Code: 'j', Text: "j"})
		m = mm.(App)
		checkStyling(t, m, "down-"+string(rune('0'+i%10)))
	}
	for i := 0; i < 39; i++ {
		mm, _ := m.Update(tea.KeyPressMsg{Code: 'k', Text: "k"})
		m = mm.(App)
		checkStyling(t, m, "up-"+string(rune('0'+i%10)))
	}
}
