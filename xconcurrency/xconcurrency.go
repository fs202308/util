package xconcurrency

import (
	"context"
	"time"

	"github.com/fs202308/util/xparallelizer"
)

// Parallelize parallelizes the function calls
func Parallelize(functions ...func()) error {
	return ParallelizeContext(context.Background(), functions...)
}

// ParallelizeTimeout parallelizes the function calls with a timeout
func ParallelizeTimeout(timeout time.Duration, functions ...func()) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return ParallelizeContext(ctx, functions...)
}

// ParallelizeContext parallelizes the function calls with a context
func ParallelizeContext(ctx context.Context, functions ...func()) error {
	group := xparallelizer.NewGroup()
	for _, function := range functions {
		group.Add(function)
	}

	return group.Wait(xparallelizer.WithContext(ctx))
}
