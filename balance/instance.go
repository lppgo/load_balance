package balance

import (
	"strconv"
)

type Instance struct {
	host string
	port int
}

func NewInstance(host string, port int) *Instance {
	return &Instance{
		host: host,
		port: port,
	}
}

func (c *Instance) GetHost() string {
	return c.host
}

func (c *Instance) String() string {
	return c.host + ":" + strconv.Itoa(c.port)
}
