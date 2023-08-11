package syncPoolUtil

import (
	"sync"
	"sync/atomic"
)

// StringSlicePool represents a pool of string slices.
type StringSlicePool struct {
	Pool   sync.Pool
	Inited atomic.Bool
}

// NewStringSlicePool creates a new StringSlicePool with the specified capacity.
func NewStringSlicePool(capacity int) (ssp *StringSlicePool) {
	ssp = &StringSlicePool{
		Pool: sync.Pool{
			// Create a new sync.Pool to hold string slices.
			// The New function initializes each new string slice with a capacity.
			New: func() interface{} {
				return make([]string, 0, capacity)
			},
		},
		// Initialize the inited flag as false.
		Inited: atomic.Bool{},
	}

	// Mark the pool as initialized.
	ssp.Inited.Store(true)

	return
}

// Get retrieves a string slice from the pool.
//
//go:inline
func (ssp *StringSlicePool) Get() (slice []string) {
	if ssp.Inited.Load() {
		slice = ssp.Pool.Get().([]string)
		return
	}

	return
}

// Put returns a string slice to the pool.
//
//go:inline
func (ssp *StringSlicePool) Put(strSlice *[]string) {
	// It's necessary to lock this section again.
	// (这里要不要上锁，有必要时再说)
	*strSlice = (*strSlice)[:0] // reset and clean slice (清空)
	/*
		Here is the issue:
		it's crucial to pass a pointer as the parameter here, as otherwise the data won't be cleared promptly.
		If another goroutine reads it while it's being modified, errors may occur.
		(这里一定要传指标进入，不然资料不会被及时清空，如果被其他协程读取时会发生错误)
	*/
	if ssp.Inited.Load() {
		ssp.Pool.Put((*strSlice)[:0])
	}
}
