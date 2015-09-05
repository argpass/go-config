package config_test

import (
	"log"
	"os"
	"time"

	"github.com/tj/go-config"
)

type ResolverOptions struct {
	Timeout     time.Duration `desc:"message timeout"`
	Concurrency uint          `desc:"max in-flight messages"`
	CacheSize   config.Bytes  `desc:"cache size in bytes"`
	BatchSize   uint          `desc:"batch size" validate:"min=1,max=1000"`
	LogLevel    string        `desc:"set the log severity" from:"env,flag"`
}

// ExampleResolvers illustrates how you may initialize a Config
// struct in order to provide custom resolvers for more flexibility.
func ExampleResolvers() {
	options := &ResolverOptions{
		Timeout:     5 * time.Second,
		Concurrency: 5,
		CacheSize:   config.ParseBytes("150mb"),
		BatchSize:   1000,
		LogLevel:    "info",
	}

	c := config.Config{
		Options: options,
		Resolvers: []config.Resolver{
			&config.FlagResolver{Args: os.Args},
			&config.EnvResolver{},
		},
	}

	err := c.Resolve()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}
