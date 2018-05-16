package HashMp

import (
	"MyHashMap/LinkNodes"
)

func CreateBulet() [16]*LinkNodes.Node {

	var arr = [16]*LinkNodes.Node{}
	for i := 0; i < 16; i++ {
		arr[i] = LinkNodes.CreateHeadNode("头节点")

	}
	return arr
}
