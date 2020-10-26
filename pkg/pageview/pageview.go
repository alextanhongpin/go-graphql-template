package pageview

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

type PageView struct {
	p *redis.Pool
}

func New(p *redis.Pool) *PageView {
	return &PageView{
		p: p,
	}
}

func (p *PageView) SetDailyPageView(key, value string) error {
	today := time.Now().Format("20060102")
	key = fmt.Sprintf("%s:%s", key, today)
	return p.SetPageView(key, value)
}

func (p *PageView) GetDailyPageView(key string) (int32, error) {
	today := time.Now().Format("20060102")
	key = fmt.Sprintf("%s:%s", key, today)
	return p.GetPageView(key)
}

func (p *PageView) SetPageView(key, value string) error {
	c := p.p.Get()
	defer func() {
		if err := c.Close(); err != nil {
			log.Println(err)
		}
	}()

	_, err := c.Do("PFADD", key, value)
	return err
}

func (p *PageView) GetPageView(key string) (int32, error) {
	c := p.p.Get()
	defer func() {
		if err := c.Close(); err != nil {
			log.Println(err)
		}
	}()

	n, err := redis.Int(c.Do("PFCOUNT", key))
	if err != nil {
		return 0, err
	}
	return int32(n), nil
}
