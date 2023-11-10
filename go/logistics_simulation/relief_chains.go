package parasite_simulation

import (
	"github.com/umbralcalc/stochadex/pkg/phenomena"
	"github.com/umbralcalc/stochadex/pkg/simulator"
)

// ParasiteSpreadIteration defines an iteration for a model of parasitic infection spread
// through a spatially-distributed population which was defined in this chapter of the book
// Worlds Of Observation: https://umbralcalc.github.io/worlds-of-observation/controlling_parasitic_infections/chapter.pdf
type ParasiteSpreadIteration struct {
	coxProcess *phenomena.CoxProcessIteration
}

func (p *ParasiteSpreadIteration) Configure(
	partitionIndex int,
	settings *simulator.LoadSettingsConfig,
) {
	p.coxProcess.Configure(partitionIndex, settings)
}

func (p *ParasiteSpreadIteration) Iterate(
	otherParams *simulator.OtherParams,
	partitionIndex int,
	stateHistories []*simulator.StateHistory,
	timestepsHistory *simulator.CumulativeTimestepsHistory,
) []float64 {
	// do something
	return []float64{}
}
