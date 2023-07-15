package module_wrapped

import (
	methods "github.com/osmosis-labs/osmosis-math-wrap/module"
)

func CalcPriceImpactWithAmountWrapped(argsBuf []byte, envSize uint32) []byte {

	args := DeserializeCalcPriceImpactWithAmountArgs(argsBuf)

	result := methods.CalcPriceImpactWithAmount(args)
	return SerializeCalcPriceImpactWithAmountResult(result)
}
