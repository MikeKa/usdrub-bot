package ex

import (
	"fmt"
	"sync"
)

// Currency of exchange rate
type Currency struct {
	sync.RWMutex
	rate     float64
	err      error
	rateFunc func() (float64, error)
}

func New(rateFunc func() (float64, error)) *Currency {
	return &Currency{rateFunc: rateFunc}
}

// Update
func (c *Currency) Update(wg *sync.WaitGroup) {
	update := func() {
		c.Lock()
		defer c.Unlock()
		c.rate, c.err = c.rateFunc()
	}

	if wg == nil {
		update()
		return
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		update()
	}()
}

// Get exchange rate
func (c *Currency) Rate() (float64, error) {
	c.RLock()
	defer c.RUnlock()
	return c.rate, c.err
}

// Get formatted exchange rate
func (c *Currency) String() string {
	c.RLock()
	defer c.RUnlock()
	return fmt.Sprintf("%.2f RUB", c.rate)
}