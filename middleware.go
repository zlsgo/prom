package prom

import (
	"github.com/sohaha/zlsgo/znet"
)

// Default use default configuration
func Default() znet.HandlerFunc {
	return New(&Config{})
}

// New new middleware
func New(o *Config) znet.HandlerFunc {
	return middleware(o)
}
