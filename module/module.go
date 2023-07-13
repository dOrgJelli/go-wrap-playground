package module

import (
	wrap "example.com/go-wrap-test/module/wrap/types"
	types "example.com/go-wrap-test/module/types"
)

func CalcPriceImpactWithAmount(args *wrap.ArgsCalcPriceImpactWithAmount) string {
	spotPriceBefore, errSPB := types.NewDecFromStr(args.SpotPriceBefore)
	if errSPB != nil {
		return errSPB.Error()
	}
	priceImpact, errPI := types.NewDecFromStr(args.PriceImpact)
	if errPI != nil {
		return errPI.Error()
	}

	// Not working for some reason
	// one := types.NewDecFromInt(types.NewInt(1)) 

	one, _ := types.NewDecFromStr("1")
	newImpact := priceImpact.Add(one)
	effectivePrice := newImpact.Mul(spotPriceBefore)

  tokenAmount, errTA := types.NewDecFromStr(args.TokenAmount)
	if errTA != nil {
		return errTA.Error()
	}
	return tokenAmount.Quo(effectivePrice).String()
}
