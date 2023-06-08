package pterm

import (
	"atomicgo.dev/keyboard/keys"
	"github.com/pterm/pterm"
)

func Confirm(desc string) bool {
	result, _ := pterm.DefaultInteractiveConfirm.WithDefaultText(desc).Show()
	return result
}

func Multiselect(desc string, options []string) (selectedOptions []string) {
	printer := pterm.DefaultInteractiveMultiselect.WithOptions(options).WithDefaultText(desc)
	printer.Filter = false
	printer.KeyConfirm = keys.Enter
	printer.KeySelect = keys.Space
	selectedOptions, _ = printer.Show()
	return
}

func Info(a ...interface{}) {
	pterm.Info.Println(a...)
}

func Error(a ...interface{}) {
	pterm.Error.Println(a...)
}

func Success(a ...interface{}) {
	pterm.Success.Println(a...)
}

func Warning(a ...interface{}) {
	pterm.Warning.Println(a...)
}

func DefaultLogger() pterm.Logger {
	return pterm.DefaultLogger
}
