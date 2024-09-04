package slogcontext

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	ctx := context.Background()

	var values []string
	for i := 0; i < 100; i++ {
		values = append(values, fmt.Sprintf("value-%d", i))
	}

	for _, val := range values {
		newCtx := WithValue(ctx, "key", val)

		go func(ctx context.Context, val string) {
			m, ok := newCtx.Value(fields).(*sync.Map)
			if !ok {
				return
			}

			valOfKey, ok := m.Load("key")
			if !ok {
				return
			}
			t.Logf("key: %v - should be: %s", valOfKey, val)
		}(ctx, val)

	}
	time.Sleep(1 * time.Second)
}
