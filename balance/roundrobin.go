package balance

import (
	"errors"
	"log"
)

type RoundRobinBalance struct {
	curIndex int
}

func init() {
	log.Println("RoundRobinBalance init ...")
	RegisterBalancer("roundRobin", &RoundRobinBalance{})
}

// 轮询算法
func (c *RoundRobinBalance) DoBalance(insts []*Instance, key ...string) (in *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}
	lens := len(insts)
	if c.curIndex >= len(insts) {
		c.curIndex = 0
	}
	in = insts[c.curIndex]
	c.curIndex = (c.curIndex + 1) % lens
	return
}
