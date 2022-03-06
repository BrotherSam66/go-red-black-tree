package bstmodels

import "fmt"

// Hero 英雄的结构体
type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

// NewHero 构造函数
func NewHero(no int, name string, left *Hero, right *Hero) *Hero {
	return &Hero{
		No:    no,
		Name:  name,
		Left:  left,
		Right: right,
	}

}

// PreOrder 前序遍历：中左右 就是先访问根节点，再访问左节点，最后访问右节点，
func PreOrder(node *Hero) {
	if node != nil {
		//fmt.Printf("No:%d;Label:%s;Left:%v;Right:%v\n", node.No, node.Label, node.Left, node.Right)
		fmt.Println(node.No, node.Name, node.Left, node.Right)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
	return
}

// InfixOrder 中序遍历：左中右 所谓的中序遍历就是先访问左节点，再访问根节点，最后访问右节点，
func InfixOrder(node *Hero) {
	if node != nil {
		InfixOrder(node.Left)
		//fmt.Printf("No:%d;Label:%s;Left:%v;Right:%v\n", node.No, node.Label, node.Left, node.Right)
		fmt.Println(node.No, node.Name, node.Left, node.Right)
		InfixOrder(node.Right)
	}
	return
}

// PostOrder 后序遍历：左右中 所谓的后序遍历就是先访问左节点，再访问右节点，最后访问根节点。
func PostOrder(node *Hero) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		//fmt.Printf("No:%d;Label:%s;Left:%v;Right:%v\n", node.No, node.Label, node.Left, node.Right)
		fmt.Println(node.No, node.Name, node.Left, node.Right)
	}
	return
}

// LevelOrder 层序遍历：按层，左右
// 弄一个指针切片，仿队列，①显示left，②left进队列，③显示right，④right进队列；取队列下一个指针；
func LevelOrder(node *Hero) {
	if node == nil {
		fmt.Println("这是个空树！")
		return
	}
	// 定义一些准全局便利性+函数
	nodeQueue := make([]Hero, 0, 100) // 切片仿队列
	queueHead := 0                    // 队列的头
	queueTail := 0                    // 队列的尾巴
	travel := func() {
		curNode := nodeQueue[queueHead] // 当前、刚取出来的节点
		queueHead++                     // 队列头修正
		fmt.Println(curNode)
		if curNode.Left != nil {
			nodeQueue = append(nodeQueue, *curNode.Left) // 压入队列
			queueTail++                                  // 队列尾巴修正
		}
		if curNode.Right != nil {
			nodeQueue = append(nodeQueue, *curNode.Right) // 压入队列
			queueTail++                                   // 队列尾巴修正
		}
	}

	// 开始程序
	nodeQueue = append(nodeQueue, *node) // 压入队列
	queueTail++                          // 队列尾巴修正
	for queueTail-queueHead > 0 {
		travel()
	}
}
