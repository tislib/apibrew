// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis

package module

import (
	"github.com/apibrew/apibrew/pkg/service/impl"
	module1 "github.com/apibrew/nano/pkg"
)

var Modules = map[string]string{
	"github.com/apibrew/nano/pkg": "v1.0.5",
}

func RegisterModules(app *impl.App) {
	app.RegisterModule(module1.NewModule)
}