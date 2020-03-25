package tmpcache

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrKeyEmpty = errors.New("key is empty")
)

type TmpCache struct {
	store   sync.Map
	closeCh chan struct{}
	isClose bool
}

type TmpValue struct {
	expire time.Time
	data   interface{}
}

func (v *TmpValue) isExpired() bool {
	return v.expire.Before(time.Now()) || v.expire.Equal(time.Now())
}

func NewCache() *TmpCache {
	cache := &TmpCache{}
	go cache.run()
	return cache
}

func (t *TmpCache) Close() {
	if !t.isClose {
		t.isClose = true
		close(t.closeCh)
	}
}

func (t *TmpCache) reStart() {
	t = &TmpCache{}
	go t.run()
}

func (t *TmpCache) run() {
	defer func() {
		if e := recover(); e != nil {
			// TODO log
			t.reStart()
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			t.store.Range(func(key, value interface{}) bool {
				data, ok := value.(*TmpValue)
				if !ok {
					return false
				}
				if data.isExpired() {
					t.store.Delete(key)
				}
				return true
			})
		case <-t.closeCh:
			return
		}
	}
}

func (t *TmpCache) Save(key string, value interface{}, expire uint) error {
	if key == "" {
		return ErrKeyEmpty
	}

	v := &TmpValue{
		data:   value,
		expire: time.Now().Add(time.Duration(expire) * time.Second),
	}

	t.store.Store(key, v)
	return nil
}

func (t *TmpCache) Load(key string) (interface{}, bool) {
	value, ok := t.store.Load(key)
	if !ok {
		return nil, false
	}

	data, ok := value.(*TmpValue)
	if !ok {
		return nil, false
	}

	return data.data, true
}

func (t *TmpCache) Delete(key string) {
	t.store.Delete(key)
}
