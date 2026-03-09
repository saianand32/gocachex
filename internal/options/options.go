package options

import "time"

type Options struct {
	TTL time.Duration
}

type Option func(*Options)

func WithTTL(d time.Duration) Option {
	return func(o *Options) { o.TTL = d }
}
