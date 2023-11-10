package logistics_simulation

import (
	"github.com/umbralcalc/stochadex/pkg/phenomena"
	"github.com/umbralcalc/stochadex/pkg/simulator"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

// ReliefChainIteration defines an iteration for a model of disaster relief chain logistics
// which was defined in this chapter of the book Worlds Of Observation:
// https://umbralcalc.github.io/worlds-of-observation/optimising_relief_chain_logistics/chapter.pdf
type ReliefChainIteration struct {
	coxProcess      *phenomena.CoxProcessIteration
	categoricalDist *distuv.Categorical
}

func (r *ReliefChainIteration) Configure(
	partitionIndex int,
	settings *simulator.LoadSettingsConfig,
) {
	r.coxProcess.Configure(partitionIndex, settings)
	seed := settings.Seeds[partitionIndex]
	weights := make([]float64, 0)
	for i := 0; i < 15; i++ {
		weights = append(weights, 1.0)
	}
	catDist := distuv.NewCategorical(weights, rand.NewSource(seed))
	rand.Seed(seed)
	r.categoricalDist = &catDist
}

func (r *ReliefChainIteration) Iterate(
	otherParams *simulator.OtherParams,
	partitionIndex int,
	stateHistories []*simulator.StateHistory,
	timestepsHistory *simulator.CumulativeTimestepsHistory,
) []float64 {
	// do something
	return []float64{}
}
