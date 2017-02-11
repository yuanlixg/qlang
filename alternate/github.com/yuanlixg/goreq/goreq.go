package goreq

import (
	"github.com/yuanlixg/goreq"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/yuanlixg/goreq",

	"DELETE":  goreq.DELETE,
	"GET":     goreq.GET,
	"HEAD":    goreq.HEAD,
	"OPTIONS": goreq.OPTIONS,
	"PATCH":   goreq.PATCH,
	"POST":    goreq.POST,
	"PUT":     goreq.PUT,

	"ShortContentTypes": goreq.ShortContentTypes,

	"GoReq":       qlang.StructOf((*goreq.GoReq)(nil)),
	"new":         goreq.New,
	"RetryConfig": qlang.StructOf((*goreq.RetryConfig)(nil)),
}
