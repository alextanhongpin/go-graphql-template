package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/url"

	"github.com/alextanhongpin/pkg/grace"
	"github.com/gomodule/redigo/redis"
	"github.com/google/wire"
	"github.com/kelseyhightower/envconfig"
	"github.com/ory/dockertest"
)

var RedisSet = wire.NewSet(NewRedisConfig, NewCache)

// RedisConfig represents the config for Redis that is parsed from the environment variables.
type RedisConfig struct {
	Host string `envconfig:"REDIS_HOST" default:"127.0.0.1"`
	Port int    `envconfig:"REDIS_PORT" default:"6379"`
}

func (r *RedisConfig) String() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}

func NewRedisConfig() RedisConfig {
	var cfg RedisConfig
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}

// NewCache returns a new redis pool.
func NewCache(cfg RedisConfig, sg *grace.ShutdownGroup) *redis.Pool {
	pool := &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", cfg.String())
		},
	}
	sg.Add(func(context.Context) {
		if err := pool.Close(); err != nil {
			log.Fatal(err)
		}
	})
	return pool
}

// SetupTestCache returns a new test redis running in docker.
func SetupTestCache() (*redis.Pool, func()) {
	var client *redis.Pool
	var err error
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	resource, err := pool.Run("redis", "5.0.5-alpine", nil)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	u, err := url.Parse(pool.Client.Endpoint())
	if err != nil {
		log.Fatalf("Could not parse endpoint: %s", pool.Client.Endpoint())
	}
	_ = resource.Expire(60)
	if err := pool.Retry(func() error {
		client = &redis.Pool{
			MaxActive: 5,
			MaxIdle:   5,
			Wait:      true,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", net.JoinHostPort(u.Hostname(), resource.GetPort("6379/tcp")))
			},
		}
		conn := client.Get()
		defer func() {
			if err := conn.Close(); err != nil {
				log.Println(err)
			}
		}()
		_, err := conn.Do("PING")
		return err

	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	return client, func() {
		if err := client.Close(); err != nil {
			log.Printf("Close not close client: %s", err)
		}
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}
}
