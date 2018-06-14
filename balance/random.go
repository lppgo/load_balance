package balance

import (
	"errors"
	"log"
	"math/rand"
)

type RandomBalance struct {
}

func init() {
	log.Println("RandomBalance init ...")
	RegisterBalancer("random", &RandomBalance{})
}

// 随机算法
func (c *RandomBalance) DoBalance(insts []*Instance, key ...string) (in *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}
	lens := len(insts)
	index := rand.Intn(lens)
	in = insts[index] //随机取一个实例
	// log.Println(index)
	return
}
