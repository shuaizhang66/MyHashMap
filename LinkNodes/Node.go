package LinkNodes

import (
	"fmt"
)
//用结构体作为数据域的类型
type DM struct{
	k string
	v string
}

//声明全局变量，保存头结点
var head *Node
var curr *Node

//声明节点类型
type Node struct {
	//数据域
	Data string
	//地址域
	NextNode *Node
}

//创建头结点
func CreateHeadNode(data string) *Node {
	var node *Node = new(Node)
	node.Data = data
	node.NextNode = nil
	//保存头结点
	head = node
	curr = node

	return node

}

//添加新节点
func AddNode(data string) *Node {
	var newNode *Node = new(Node)
	newNode.Data = data
	newNode.NextNode = nil
	//挂接节点
	curr.NextNode = newNode
	curr = newNode
	//返回值
	return newNode
}

//遍历链表
func ShowNodes() {
	var node = head
	for {
		if node.NextNode == nil {
			fmt.Print(node.Data)
			break
		} else {
			fmt.Println(node.Data)
			node = node.NextNode
		}
	}

}

//计算节点的个数
func NodeCnt() int {
	var cnt int = 1
	var node = head
	for {
		if node.NextNode == nil {
			break
		} else {
			node = node.NextNode
			cnt = cnt + 1
		}

	}
	return cnt
}

//插入节点
func InsertNodeByIndex(index int, data string) *Node {
	if index == 0 {
		//添加的为新的头节点
		var node *Node = new(Node)
		node.Data = data
		node.NextNode = head
		head = node

	} else if index > NodeCnt()-1 {
		//添加节点
		AddNode(data)

	} else {
		//中间插入节点
		var n = head
		for i := 0; i < index-1; i++ {
			n = n.NextNode
		}

		var newNode *Node = new(Node)
		newNode.Data = data
		newNode.NextNode = n.NextNode
		n.NextNode = newNode

	}

	return nil

}

//删除节点
func DeleteNodeByIndex(index int) {

	var node = head
	if index == 0 {
		//删除头节点，就是第二个节点为头结点
		head = node.NextNode

	} else {
		for i := 0; i < index-1; i++ {
			node = node.NextNode
		}
		node.NextNode = node.NextNode.NextNode
	}

}

//修改指定下标的节点内容
func UpdateNodeByIndex(index int, data string) {
	var node = head
	if index == 0 {
		head.Data = data
	} else {

		for i := 0; i < index; i++ {
			node = node.NextNode
		}
		node.Data = data
	}

}
