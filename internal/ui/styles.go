package ui

import "charm.land/lipgloss/v2"

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colAccent)

	titleHintStyle = lipgloss.NewStyle().
			Foreground(colDim)

	tableStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(colBorder)

	helpStyle = lipgloss.NewStyle().
			Foreground(colDim).
			MarginTop(1)

	keyStyle = lipgloss.NewStyle().
			Foreground(colKey).
			Bold(true)

	errorStyle = lipgloss.NewStyle().
			Foreground(colDanger).
			Bold(true)

	stoppedRowStyle = lipgloss.NewStyle().
			Foreground(colDim)

	pausedRowStyle = lipgloss.NewStyle().
			Foreground(colPaused)

	emptyStyle = lipgloss.NewStyle().
			Foreground(colDim).
			MarginLeft(2).
			MarginTop(1)

	confirmStyle = lipgloss.NewStyle().
			Foreground(colWarn).
			Bold(true)

	confirmNameStyle = lipgloss.NewStyle().
				Foreground(colBright).
				Bold(true)

	logsDividerStyle = lipgloss.NewStyle().
				Foreground(colBorder)

	logsTitleStyle = lipgloss.NewStyle().
			Foreground(colAccent).
			Bold(true)

	logsLineStyle = lipgloss.NewStyle().
			Foreground(colText)

	logsTimestampStyle = lipgloss.NewStyle().
				Foreground(colDim)

	logsHighlightStyle = lipgloss.NewStyle().
				Foreground(colBlack).
				Background(colWarn).
				Bold(true)

	inspectSectionStyle = lipgloss.NewStyle().
				Foreground(colSection).
				Bold(true)

	inspectValueStyle = lipgloss.NewStyle().
				Foreground(colText)

	contextActiveStyle = lipgloss.NewStyle().
				Foreground(colOK).
				Bold(true)

	contextCursorStyle = lipgloss.NewStyle().
				Foreground(colBright).
				Background(colBorder)

	trendUpStyle = lipgloss.NewStyle().
			Foreground(colDanger).
			Bold(true)

	trendDownStyle = lipgloss.NewStyle().
			Foreground(colOK).
			Bold(true)

	trendSteadyStyle = lipgloss.NewStyle().
				Foreground(colDim)

	sparklineStyle = lipgloss.NewStyle().
			Foreground(colDim)

	diagnosisWarnStyle = lipgloss.NewStyle().
				Foreground(colWarn).Bold(true)

	diagnosisErrorStyle = lipgloss.NewStyle().
				Foreground(colDanger).Bold(true)

	eventStartStyle = lipgloss.NewStyle().
			Foreground(colOK).Bold(true)

	eventStopStyle = lipgloss.NewStyle().
			Foreground(colDanger).Bold(true)

	eventWarnStyle = lipgloss.NewStyle().
			Foreground(colWarn).Bold(true)

	eventDimStyle = lipgloss.NewStyle().
			Foreground(colDim)

	eventTimeStyle = lipgloss.NewStyle().
			Foreground(colDim)

	eventTypeStyle = lipgloss.NewStyle().
			Foreground(colKey)

	eventNameStyle = lipgloss.NewStyle().
			Foreground(colText)

	collapsedRowStyle = lipgloss.NewStyle().
				Foreground(colDim)

	detailRowStyle = lipgloss.NewStyle().
			Foreground(colDim)

	detailRowSelectedStyle = lipgloss.NewStyle().
				Foreground(colBright).
				Background(colBorder)
)
