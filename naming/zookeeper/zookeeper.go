package zookeeper

import (
	"github.com/shengwudiyi/stolen-plane/naming"

	"github.com/go-zookeeper/zk"
)

type Config struct {
	Prefix    string
	Endpoints []string
}

type Zookeeper struct {
	c   *Config
	cli *zk.Conn
}

func New() (builder naming.Builder, err error) {
	return
}

func (z *Zookeeper) Schema() string {
	return "zookeeper"
}
