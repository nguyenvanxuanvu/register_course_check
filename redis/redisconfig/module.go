package redisconfig

import "go.uber.org/fx"

var Module = fx.Provide(NewCacheSingle)
