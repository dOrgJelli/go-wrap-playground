package module_wrapped

import (
	methods "example.com/go-wrap-test/module"
)

func CalcPriceImpactWithAmountWrapped(argsBuf []byte, envSize uint32) []byte {

	args := DeserializeCalcPriceImpactWithAmountArgs(argsBuf)

	result := methods.CalcPriceImpactWithAmount(args)
	return SerializeCalcPriceImpactWithAmountResult(result)
}
