package internal

import (
	"fmt"
	"log"
	"net"
	"net/url"

	"github.com/gomodule/redigo/redis"
	"github.com/kelseyhightower/envconfig"
	"github.com/ory/dockertest"
)

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

// NewRedis returns a new redis pool.
func NewRedis(cfg RedisConfig) (*redis.Pool, func()) {
	pool := &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", cfg.String())
		},
	}
	return pool, func() {
		_ = pool.Close()
	}
}

// InitTestCache returns a new test redis running in docker.
func InitTestCache() (*redis.Pool, func()) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("start docker failed: %s", err)
	}

	resource, err := pool.Run("redis", "5.0.5-alpine", nil)
	if err != nil {
		log.Fatalf("start resource failed: %s", err)
	}

	u, err := url.Parse(pool.Client.Endpoint())
	if err != nil {
		log.Fatalf("parse client endpoint failed: %s", pool.Client.Endpoint())
	}

	_ = resource.Expire(60)

	var client *redis.Pool
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
		log.Fatalf("connect to docker failed: %s", err)
	}

	return client, func() {
		if err := client.Close(); err != nil {
			log.Printf("close client failed: %s", err)
		}

		if err := pool.Purge(resource); err != nil {
			log.Fatalf("purge resource failed: %s", err)
		}
	}
}
