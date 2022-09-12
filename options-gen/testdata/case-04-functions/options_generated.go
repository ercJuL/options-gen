// Code generated by options-gen. DO NOT EDIT.
package testcase

import (
	"net/http"
)

type optOptionsSetter func(o *Options)

func NewOptions(
	fnTypeParam FnType,
	fnParam func(server *http.Server) error,
	handlerFunc http.HandlerFunc,
	middleware func(next http.HandlerFunc) http.HandlerFunc,
	local localFnType,
	options ...optOptionsSetter,
) Options {
	o := Options{}
	o.fnTypeParam = fnTypeParam
	o.fnParam = fnParam
	o.handlerFunc = handlerFunc
	o.middleware = middleware
	o.local = local

	for _, opt := range options {
		opt(&o)
	}
	return o
}

func WithOptFnTypeParam(opt FnType) optOptionsSetter {
	return func(o *Options) {
		o.optFnTypeParam = opt
	}
}

func WithOptFnParam(opt func(server *http.Server) error) optOptionsSetter {
	return func(o *Options) {
		o.optFnParam = opt
	}
}

func WithOptHandlerFunc(opt http.HandlerFunc) optOptionsSetter {
	return func(o *Options) {
		o.optHandlerFunc = opt
	}
}

func WithOptMiddleware(opt func(next http.HandlerFunc) http.HandlerFunc) optOptionsSetter {
	return func(o *Options) {
		o.optMiddleware = opt
	}
}

func WithOptLocal(opt localFnType) optOptionsSetter {
	return func(o *Options) {
		o.optLocal = opt
	}
}

func (o *Options) Validate() error {
	return nil
}
