// MyHashMap project main.go
package main

import (
	"MyHashMap/LinkNodes"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	//创建头结点
	LinkNodes.CreateHeadNode("头结点")
	LinkNodes.AddNode("第二节点")
	LinkNodes.AddNode("第三节点")
	LinkNodes.AddNode("第四节点")
	//LinkNodes.InsertNodeByIndex(3, "新节点")
	//LinkNodes.DeleteNodeByIndex(3)
	LinkNodes.UpdateNodeByIndex(1, "abc")
	LinkNodes.ShowNodes()
	fmt.Println(LinkNodes.NodeCnt())
}
