package module_wrapped

import (
	. "github.com/osmosis-labs/osmosis-math-wrap/module/wrap/types"
	"github.com/polywrap/go-wrap/msgpack"
)

func DeserializeCalcPriceImpactWithAmountArgs(argsBuf []byte) *ArgsCalcPriceImpactWithAmount {
	ctx := msgpack.NewContext("Deserializing module-type: CalcPriceImpactWithAmount")
	reader := msgpack.NewReadDecoder(ctx, argsBuf)

	var (
		_spotPriceBefore    string
		_spotPriceBeforeSet bool
		_tokenAmount        string
		_tokenAmountSet     bool
		_priceImpact        string
		_priceImpactSet     bool
	)

	for i := int32(reader.ReadMapLength()); i > 0; i-- {
		field := reader.ReadString()
		reader.Context().Push(field, "unknown", "searching for property type")
		switch field {
		case "spotPriceBefore":
			reader.Context().Push(field, "string", "type found, reading property")
			_spotPriceBefore = reader.ReadString()
			_spotPriceBeforeSet = true
			reader.Context().Pop()
		case "tokenAmount":
			reader.Context().Push(field, "string", "type found, reading property")
			_tokenAmount = reader.ReadString()
			_tokenAmountSet = true
			reader.Context().Pop()
		case "priceImpact":
			reader.Context().Push(field, "string", "type found, reading property")
			_priceImpact = reader.ReadString()
			_priceImpactSet = true
			reader.Context().Pop()
		}
		reader.Context().Pop()
	}

	if !_spotPriceBeforeSet {
		panic(reader.Context().PrintWithContext("Missing required property: 'spotPriceBefore: String'"))
	}
	if !_tokenAmountSet {
		panic(reader.Context().PrintWithContext("Missing required property: 'tokenAmount: String'"))
	}
	if !_priceImpactSet {
		panic(reader.Context().PrintWithContext("Missing required property: 'priceImpact: String'"))
	}

	return &ArgsCalcPriceImpactWithAmount{
		SpotPriceBefore: _spotPriceBefore,
		TokenAmount:     _tokenAmount,
		PriceImpact:     _priceImpact,
	}
}

func SerializeCalcPriceImpactWithAmountResult(value string) []byte {
	ctx := msgpack.NewContext("Serializing module-type: CalcPriceImpactWithAmount")
	encoder := msgpack.NewWriteEncoder(ctx)
	WriteCalcPriceImpactWithAmountResult(encoder, value);
	return encoder.Buffer()
}

func WriteCalcPriceImpactWithAmountResult(writer msgpack.Write, value string) {
	writer.Context().Push("calcPriceImpactWithAmount", "string", "writing property")
	{
		v := value
		writer.WriteString(v)
	}
	writer.Context().Pop()
}
