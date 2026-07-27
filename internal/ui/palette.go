package ui

import (
	"strconv"

	"charm.land/lipgloss/v2"
)

// Colors resolve against the terminal's own 16-color palette rather than fixed
// hex values, so tdocker follows whatever theme the terminal is already using.
const (
	idxBlack   = 0
	idxPaused  = 3
	idxBorder  = 4
	idxText    = 7
	idxDim     = 8
	idxDanger  = 9
	idxOK      = 10
	idxWarn    = 11
	idxAccent  = 12
	idxSection = 13
	idxKey     = 14
	idxBright  = 15
)

var (
	colBlack   = lipgloss.ANSIColor(idxBlack)
	colPaused  = lipgloss.ANSIColor(idxPaused)
	colBorder  = lipgloss.ANSIColor(idxBorder)
	colText    = lipgloss.ANSIColor(idxText)
	colDim     = lipgloss.ANSIColor(idxDim)
	colDanger  = lipgloss.ANSIColor(idxDanger)
	colOK      = lipgloss.ANSIColor(idxOK)
	colWarn    = lipgloss.ANSIColor(idxWarn)
	colAccent  = lipgloss.ANSIColor(idxAccent)
	colSection = lipgloss.ANSIColor(idxSection)
	colKey     = lipgloss.ANSIColor(idxKey)
	colBright  = lipgloss.ANSIColor(idxBright)
)

// fgSeq is the raw SGR sequence for an indexed foreground, matching how lipgloss
// encodes ANSIColor, for the places color must be embedded in rendered output.
func fgSeq(idx int) string {
	return "\x1b[38;5;" + strconv.Itoa(idx) + "m"
}

var (
	seqSelectedFg = fgSeq(idxBlack)
	seqDim        = fgSeq(idxDim)
	seqPaused     = fgSeq(idxPaused)
	seqReset      = "\x1b[39m"
)
