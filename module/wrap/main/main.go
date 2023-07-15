package main

import (
	"github.com/osmosis-labs/osmosis-math-wrap/module/wrap/module_wrapped"
	"github.com/polywrap/go-wrap/wrap"
)

//export _wrap_invoke
func _wrap_invoke(methodSize, argsSize, envSize uint32) bool {
	args := wrap.WrapInvokeArgs(methodSize, argsSize)
	switch args.Method {
	case "calcPriceImpactWithAmount":
		return wrap.WrapInvoke(args, envSize, module_wrapped.CalcPriceImpactWithAmountWrapped)
	default:
		return wrap.WrapInvoke(args, envSize, nil)
	}
}

func main() {
}
