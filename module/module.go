package module

import (
	wrap "github.com/osmosis-labs/osmosis-math-wrap/module/wrap/types"
	types "github.com/osmosis-labs/osmosis-math-wrap/module/types"
)

func CalcPriceImpactWithAmount(args *wrap.ArgsCalcPriceImpactWithAmount) string {
	spotPriceBefore, _ := types.NewDecFromStr(args.SpotPriceBefore)
	priceImpact, _ := types.NewDecFromStr(args.PriceImpact)

	// Not working for some reason
	// one := types.NewDecFromInt(types.NewInt(1)) 
	oneDec, _ := types.NewDecFromStr("1")
	newImpact := priceImpact.Add(oneDec)
	effectivePrice := newImpact.Mul(spotPriceBefore)
  tokenAmount, _ := types.NewDecFromStr(args.TokenAmount)
	return tokenAmount.Quo(effectivePrice).String()
}
