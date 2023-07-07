package main

import (
	"example.com/go-wrap-test/module/wrap/module_wrapped"
	"github.com/polywrap/go-wrap/polywrap"
)

//export _wrap_invoke
func _wrap_invoke(methodSize, argsSize, envSize uint32) bool {
	args := polywrap.WrapInvokeArgs(methodSize, argsSize)
	switch args.Method {
	case "calcPriceImpactWithAmount":
		return polywrap.WrapInvoke(args, envSize, module_wrapped.CalcPriceImpactWithAmountWrapped)
	default:
		return polywrap.WrapInvoke(args, envSize, nil)
	}
}

func main() {
}
