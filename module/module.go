package module

import (
	wrap "example.com/go-wrap-test/module/wrap/types"
	types "example.com/go-wrap-test/module/types"
)

func CalcPriceImpactWithAmount(args *wrap.MethodArgsCalcPriceImpactWithAmount) string {
	spotPriceBefore, _ := types.NewDecFromStr(args.SpotPriceBefore)
	priceImpact, _ := types.NewDecFromStr(args.PriceImpact)
	effectivePrice := spotPriceBefore.Mul(priceImpact.Add(types.NewDecFromInt(types.NewInt(1))))
  tokenAmount, _ := types.NewDecFromStr(args.TokenAmount)
	return tokenAmount.Quo(effectivePrice).String()
}
