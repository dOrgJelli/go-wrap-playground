package types

type ArgsCalcPriceImpactWithAmount struct {
	SpotPriceBefore string `json:"spotPriceBefore"`
	TokenAmount     string `json:"tokenAmount"`
	PriceImpact     string `json:"priceImpact"`
}
