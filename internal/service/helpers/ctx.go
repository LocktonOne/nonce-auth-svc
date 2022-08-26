package helpers

import (
	"context"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokene/nonce-auth-svc/internal/config"
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	dbCtxKey
	serviceConfigCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}
func CtxDB(entry data.MasterQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, dbCtxKey, entry)
	}
}
func DB(r *http.Request) data.MasterQ {
	return r.Context().Value(dbCtxKey).(data.MasterQ).New()
}
func CtxServiceConfig(entry *config.ServiceConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, serviceConfigCtxKey, entry)
	}
}

func ServiceConfig(r *http.Request) *config.ServiceConfig {
	return r.Context().Value(serviceConfigCtxKey).(*config.ServiceConfig)
}