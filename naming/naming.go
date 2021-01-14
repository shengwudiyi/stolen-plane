package naming

import (
	"context"
)

type metakey string

const (
	MetaWeight metakey = "weight"
	MetaColor  metakey = "color"
)

type Instance struct {
	AppID    string             `json:"appid"`
	LastTs   int64              `json:"latest_timestamp"`
	Metadata map[metakey]string `json:"metadata"`
}

type InstancesInfo struct {
}

type Resolver interface {
	Fetch(context.Context) (*InstancesInfo, bool)
	Watch() <-chan struct{}
	Close() error
}

type Builder interface {
	Build(appid string) Resolver
	Scheme() string
}

const (
	subsetSize = 20
)

func defulatSubset(backends []*Instance, size int) []*Instance {
	if len(backends) <= int(size) {
		return backends
	}

	return backends[:size]
}
