package cache

import "time"

type Storager interface {
	Get(key string) ([]byte, error)
	Set(key string, content []byte, duration time.Duration)
}
