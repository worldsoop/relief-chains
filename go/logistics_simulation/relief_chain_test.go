package logistics_simulation

import (
	"testing"

	"github.com/umbralcalc/stochadex/pkg/simulator"
)

func TestReliefChain(t *testing.T) {
	t.Run(
		"test that the relief chain logistics simulation runs",
		func(t *testing.T) {
			settings := simulator.NewLoadSettingsConfigFromYaml("relief_chain_config.yaml")
			iterations := make([]simulator.Iteration, 0)
			for partitionIndex := range settings.StateWidths {
				iteration := &ReliefChainIteration{}
				iteration.Configure(partitionIndex, settings)
				iterations = append(iterations, iteration)
			}
			store := make([][][]float64, 1)
			implementations := &simulator.LoadImplementationsConfig{
				Iterations:      iterations,
				OutputCondition: &simulator.EveryStepOutputCondition{},
				OutputFunction:  &simulator.VariableStoreOutputFunction{Store: store},
				TerminationCondition: &simulator.NumberOfStepsTerminationCondition{
					MaxNumberOfSteps: 100,
				},
				TimestepFunction: &simulator.ConstantTimestepFunction{Stepsize: 1.0},
			}
			config := simulator.NewStochadexConfig(
				settings,
				implementations,
			)
			coordinator := simulator.NewPartitionCoordinator(config)
			coordinator.Run()
		},
	)
}
