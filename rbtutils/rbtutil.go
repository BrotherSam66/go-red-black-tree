package rbtutils

import (
	"go-red-black-tree/rbtmodels"
)

var Root = new(rbtmodels.RBTNode)

// RBTDemo 红黑树演示
func RBTDemo() {
	RBTCreat()
	Root.ShowTree()
	//Root.LeftRotate(Root.Left) // 测通
	//Root.LeftRotate(Root.Right) // 测通
	Root.LeftRotate(Root)
	Root.ShowTree()
	//Root.Left.ShowTree()
}

// RBTCreat 红黑树创建
func RBTCreat() {
	// 定义
	Root = rbtmodels.NewRBTNode(false, 1, "1宋江", nil, nil, nil)
	Root.DemoPush(true, rbtmodels.NewRBTNode(false, 2, "2卢俊义", nil, nil, nil))
	Root.DemoPush(false, rbtmodels.NewRBTNode(false, 3, "3吴用", nil, nil, nil))
	Root.Left.DemoPush(true, rbtmodels.NewRBTNode(false, 4, "4公孙胜", nil, nil, nil))
	Root.Left.DemoPush(false, rbtmodels.NewRBTNode(false, 5, "5关胜", nil, nil, nil))
	Root.Right.DemoPush(true, rbtmodels.NewRBTNode(false, 6, "6林冲", nil, nil, nil))
	Root.Right.DemoPush(false, rbtmodels.NewRBTNode(false, 7, "7秦明", nil, nil, nil))
	Root.Left.Left.DemoPush(true, rbtmodels.NewRBTNode(false, 8, "8呼延灼", nil, nil, nil))
	Root.Left.Left.DemoPush(false, rbtmodels.NewRBTNode(false, 9, "9华融", nil, nil, nil))
	Root.Left.Right.DemoPush(true, rbtmodels.NewRBTNode(false, 10, "10柴进", nil, nil, nil))
	Root.Left.Right.DemoPush(false, rbtmodels.NewRBTNode(false, 11, "11李应", nil, nil, nil))

}
