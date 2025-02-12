package common

import (
	"net/http"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/spf13/viper"
	rl "go.uber.org/ratelimit"
)

type ratelimit struct {
	Limiter           *limiter.Limiter
	RateLimit         bool
	RequestsPerSecond float64
}

var userRateLimit *ratelimit

func (rl *ratelimit) init() {
	if rl.RequestsPerSecond == 0 {
		rl.RateLimit = false
		return
	}
	rl.RateLimit = true
	rl.Limiter = tollbooth.NewLimiter(rl.RequestsPerSecond, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour}).
		SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"}).
		SetMethods([]string{"GET", "POST", "PUT", "DELETE"})
}

const DefaultRequestPerSecond = 100000

//ConfigRateLimits - configure the rate limits
func ConfigRateLimits() *GRPCRateLimiter {
	userRl := viper.GetFloat64("handlers.rate_limit")

	if userRl == 0 {
		userRl = DefaultRequestPerSecond
	}

	userRateLimit = &ratelimit{RequestsPerSecond: userRl}
	userRateLimit.init()

	return &GRPCRateLimiter{rl.New(int(userRl))}
}

type GRPCRateLimiter struct {
	rl.Limiter
}

func (r *GRPCRateLimiter) Limit() bool {
	r.Take()
	return false
}

//UserRateLimit - rate limiting for end user handlers
func UserRateLimit(handler ReqRespHandlerf) ReqRespHandlerf {
	if !userRateLimit.RateLimit {
		return handler
	}
	return func(writer http.ResponseWriter, request *http.Request) {
		tollbooth.LimitFuncHandler(userRateLimit.Limiter, handler).ServeHTTP(writer, request)
	}
}
