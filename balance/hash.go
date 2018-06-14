package balance

import (
	"fmt"
	"hash/crc32"
	"log"
	"math/rand"
)

type HashBalance struct {
	Name string
	Age  int
}

func init() {
	log.Println("HashBalance init ...")
	RegisterBalancer("hash", &HashBalance{})
}
func (p *HashBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	var defKey string = fmt.Sprintf("%d", rand.Int())
	if len(key) > 0 {
		defKey = key[0]
	}
	lens := len(insts)
	if lens == 0 {
		err = fmt.Errorf("No backend instance")
		return
	}
	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(defKey), crcTable)
	index := int(hashVal) % lens
	inst = insts[index]
	return
}
