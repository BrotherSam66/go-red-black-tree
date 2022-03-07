package rbtutils

import (
	"errors"
	"fmt"
	"go-red-black-tree/bstmodels"
	"go-red-black-tree/global"
	"go-red-black-tree/rbtmodels"
)

// RBTDemo 红黑树演示
func RBTDemo() {
	RBTCreat()
	err := errors.New("出错，本节点是空！")

	ShowTree(global.Root)

	//err = LeftRotate(global.Root.Left) // 测通
	//err = LeftRotate(global.Root.Right) // 测通
	//err = LeftRotate(global.Root) // 测试根

	//err = RightRotate(global.Root.Left) // 测通
	//err = RightRotate(global.Root.Right) // 测通
	//err = RightRotate(global.Root) // 测试根

	err = RightRotate(global.Root.Right.Right.Right) // 测通

	if err != nil {
		fmt.Println(err)
	}

	ShowTree(global.Root)
	//global.Root.Left.ShowTree()
}

// RBTCreat 红黑树创建
func RBTCreat() {
	// 定义
	global.Root = rbtmodels.NewRBTNode(false, 1, "1宋江", nil, nil, nil)
	DemoPush(global.Root, true, rbtmodels.NewRBTNode(false, 2, "2卢俊义", nil, nil, nil))
	DemoPush(global.Root, false, rbtmodels.NewRBTNode(false, 3, "3吴用", nil, nil, nil))
	DemoPush(global.Root.Left, true, rbtmodels.NewRBTNode(false, 4, "4公孙胜", nil, nil, nil))
	DemoPush(global.Root.Left, false, rbtmodels.NewRBTNode(false, 5, "5关胜", nil, nil, nil))
	DemoPush(global.Root.Right, true, rbtmodels.NewRBTNode(false, 6, "6林冲", nil, nil, nil))
	DemoPush(global.Root.Right, false, rbtmodels.NewRBTNode(false, 7, "7秦明", nil, nil, nil))
	DemoPush(global.Root.Left.Left, true, rbtmodels.NewRBTNode(false, 8, "8呼延灼", nil, nil, nil))
	DemoPush(global.Root.Left.Left, false, rbtmodels.NewRBTNode(false, 9, "9华融", nil, nil, nil))
	DemoPush(global.Root.Left.Right, true, rbtmodels.NewRBTNode(false, 10, "10柴进", nil, nil, nil))
	DemoPush(global.Root.Left.Right, false, rbtmodels.NewRBTNode(false, 11, "11李应", nil, nil, nil))
	DemoPush(global.Root.Right.Left, true, rbtmodels.NewRBTNode(false, 12, "12朱仝", nil, nil, nil))
	DemoPush(global.Root.Right.Left, false, rbtmodels.NewRBTNode(false, 13, "13鲁智深", nil, nil, nil))
	DemoPush(global.Root.Right.Right, true, rbtmodels.NewRBTNode(false, 14, "14武松", nil, nil, nil))
	DemoPush(global.Root.Right.Right, false, rbtmodels.NewRBTNode(false, 15, "15董平", nil, nil, nil))

}

// DemoPush 简易附加在尾部，回头废止
func DemoPush(r *rbtmodels.RBTNode, isLeft bool, san *rbtmodels.RBTNode) {
	if isLeft { // 附加在左边
		r.Left = san
	} else { // 附加在右边
		r.Right = san
	}
	san.Parent = r
}

// ShowTree 逐层显示这个树，回头废止。感觉需要递归、队列
func ShowTree(r *rbtmodels.RBTNode) {
	var data [10][1000]*rbtmodels.RBTNode // 数据。数据可能是nil。最多10层，每层最多1000数据
	//totalLevel := 0             // 总层数
	//nowLevel := 0               // 当前层数
	//nnn := Name
	nowColumn := 0 // 当前列
	fmt.Printf("\n展示树：[左子]本(父)[右子]")
	if r == nil {
		fmt.Println("这个树/分支是空的")
	}
	data[0][0] = r // 来的最高位指针

	for i := 1; i < len(data); i++ { // 循环每一层
		fmt.Println("") // 先来一个换行
		nowColumn = 0   // 当前列
		for j := 0; j < len(data[0]); j++ {
			if data[i-1][j] == nil { // 本行遍历结束
				break
			}

			ShowOneNode(data[i-1][j])     // 显示遍历到的上一行的这个节点。显示没换行
			if data[i-1][j].Left != nil { // 如果有，在下一行填入左节点
				data[i][nowColumn] = data[i-1][j].Left
				nowColumn++
			}
			if data[i-1][j].Right != nil { // 如果有，在下一行填入右节点
				data[i][nowColumn] = data[i-1][j].Right
				nowColumn++
			}
		}
	}

}

// ShowOneNode 展示单个节点
func ShowOneNode(n *rbtmodels.RBTNode) {
	fmt.Printf(" , ") // 左右分割
	if n.Left == nil {
		fmt.Printf("[ ]")
	} else {
		fmt.Printf("[%d]", n.Left.Key)
	}
	fmt.Printf("%d", n.Key)
	if n.Parent == nil {
		fmt.Printf("( )")
	} else {
		fmt.Printf("(%d)", n.Parent.Key)
	}
	if n.Right == nil {
		fmt.Printf("[ ]")
	} else {
		fmt.Printf("[%d]", n.Right.Key)
	}
}

// LeftRotate 左旋
// @param p point 旋转的支点
/*
 *   parent                 parent
 *     |                      |
 *     p                      pr
 *    / \        ===>     	 /  \
 *  pl   pr                 p    rr
 *      /  \               / \
 *     rl   rr            pl  rl
 */
func LeftRotate(p *rbtmodels.RBTNode) (err error) {
	if p == nil {
		return errors.New("出错，本节点是空！")
	}
	if p.Right == nil {
		return errors.New("出错，本节点右儿子是空！")
	}

	parent := p.Parent // 父亲
	pSelf := p         // 本身
	pr := p.Right      // 右儿子（升级）
	rl := pr.Left      // 右儿子的左孙子（断枝重连）

	// 下方需要判断 p 是否root
	if parent != nil { // p不是root，，还需要分析，p在父亲的左还是右
		if parent.Left == pSelf { // p 在父亲左手
			parent.Left = pr // 1.1 父亲左指向：
		} else { // p 在父亲右手
			parent.Right = pr // 1.2 父亲右指向：
		}
		pr.Parent = parent // 2.1 (升级的)pr上指向：
	} else { // p是root
		global.Root = pr         // 1.1 + 1.2 父亲的指向。其实是root指向pr
		global.Root.Parent = nil // 2.1 (升级的)pr上指向：
	}

	// 下方p是否root都要执行
	pr.Left = pSelf // 2.2 (升级的)pr左指向：
	// 2.3 (升级的)pr右指向：不动
	pSelf.Parent = pr // 3.1 pSelf上指向：
	// 3.2 pSelf左指向：不动
	pSelf.Right = rl // 3.3 pSelf右指向：rl
	if rl != nil {
		rl.Parent = pSelf // 4.1 rl上指向：
	}
	// 4.2 rl左指向：不动
	// 4.3 rl右指向：不动

	return err
}

// RightRotate 右旋
// @param p point 旋转的支点
/*
 *     parent                 parent
 *       |                      |
 *       p                     pl
 *      / \        ===>    	  /  \
 *    pl   pr                LL   P
 *   /  \                        / \
 *  ll   lr                     lr  pr
 */
func RightRotate(p *rbtmodels.RBTNode) (err error) {
	if p == nil {
		return errors.New("出错，本节点是空！")
	}
	if p.Left == nil {
		return errors.New("出错，本节点左儿子是空！")
	}

	parent := p.Parent // 父亲
	pSelf := p         // 本身
	pl := p.Left       // 左儿子（升级）
	lr := pl.Right     // 左儿子的右孙子（断枝重连）

	// 下方需要判断 p 是否root
	if parent != nil { // p不是root，，还需要分析，p在父亲的左还是右
		if parent.Left == pSelf { // p 在父亲左手
			parent.Left = pl // 1.1 父亲左指向：
		} else { // p 在父亲右手
			parent.Right = pl // 1.2 父亲右指向：
		}
		pl.Parent = parent // 2.1 (升级的)pl上指向：
	} else { // p是root
		global.Root = pl         // 1.1 + 1.2 父亲的指向。其实是root指向pr
		global.Root.Parent = nil // 2.1 (升级的)pr上指向：
	}

	// 下方p是否root都要执行
	pl.Right = pSelf // 2.2 (升级的)pl右指向：
	// 2.3 (升级的)pl左指向：不动
	pSelf.Parent = pl // 3.1 pSelf上指向：
	// 3.2 pSelf右指向：不动
	pSelf.Left = lr // 3.3 pSelf左指向：lr
	if lr != nil {
		lr.Parent = pSelf // 4.1 lr上指向：
	}
	// 4.2 lr左指向：不动
	// 4.3 lr右指向：不动

	return err
}

// PreOrder 前序遍历：中左右 就是先访问根节点，再访问左节点，最后访问右节点，
func PreOrder(node *bstmodels.Hero) {
	if node != nil {
		//fmt.Printf("No:%d;Label:%s;Left:%v;Right:%v\n", node.No, node.Label, node.Left, node.Right)
		fmt.Println(node.No, node.Name, node.Left, node.Right)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
	return
}

// InfixOrder 中序遍历：左中右 所谓的中序遍历就是先访问左节点，再访问根节点，最后访问右节点，
func InfixOrder(node *bstmodels.Hero) {
	if node != nil {
		InfixOrder(node.Left)
		//fmt.Printf("No:%d;Label:%s;Left:%v;Right:%v\n", node.No, node.Label, node.Left, node.Right)
		fmt.Println(node.No, node.Name, node.Left, node.Right)
		InfixOrder(node.Right)
	}
	return
}

// PostOrder 后序遍历：左右中 所谓的后序遍历就是先访问左节点，再访问右节点，最后访问根节点。
func PostOrder(node *bstmodels.Hero) {
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
func LevelOrder(node *bstmodels.Hero) {
	if node == nil {
		fmt.Println("这是个空树！")
		return
	}
	// 定义一些准全局便利性+函数
	nodeQueue := make([]bstmodels.Hero, 0, 100) // 切片仿队列
	queueHead := 0                              // 队列的头
	queueTail := 0                              // 队列的尾巴
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
