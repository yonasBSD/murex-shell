//go:build no_cachedb
// +build no_cachedb

package cache

import (
	"context"
	"sort"
	"time"
)

func SetPath(_ string) {
	// no nothing
}

func createDb(_ string) {
	// do nothing
}

func Read(namespace string, key string, ptr any) bool {
	return read(namespace, key, ptr)
}

func listDb(_ context.Context, _ string) (any, error) {
	return nil, nil
}

func Write(namespace string, key string, value any, ttl time.Time) {
	write(namespace, key, value, ttl)
}

func trimDb(_ context.Context, _ string) ([]string, error) {
	return nil, nil
}

func clearDb(_ context.Context, _ string) ([]string, error) {
	return nil, nil
}

func CloseDb() {
	// do nothing
}

func DbPath() string  { return "" }
func DbEnabled() bool { return false }

func ListNamespaces() []string {
	var (
		ret = make([]string, len(cache))
		i   int
	)

	for namespace := range cache {
		ret[i] = namespace
		i++
	}

	sort.Strings(ret)
	return ret
}
