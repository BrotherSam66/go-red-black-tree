package rbtmodels

import (
	"errors"
	"fmt"
	"go-red-black-tree/bstmodels"
)

/*
《红黑树》

前序遍历：左子树=》根节点=》右子树=》（逐个向下级递归）
中序遍历：根节点=》左子树=》右子树=》（逐个向下级递归）
后序遍历：左子树=》右子树=》根节点=》（逐个向下级递归）

前驱节点：小于当前节点的最大值。
后继节点：大于当前节点额最小值。
删除当前节点，可用前驱/后继节点替换上来。

BST 二叉树：可能不平很
AVL 高度平衡树：左右子树高度差不大于1

Tree234树：每个节点最多3个元素，每个元素也分左右儿子。234树每个叶子到根的路径长度相等。
234树映射红黑树：①一元素黑色。②二元素一上一下，上黑下红，可能左倾右倾。③三元素中间升起，上黑下红。
④三元素加人，相当于中间升起的变红，左右变黑，加入的新元素红色。

红黑树：①节点分红黑色。②根是黑色。③所有叶子都是黑色(叶子是nil，这类节点不可忽视，否则代码看不懂)。
④每个红色必须下挂2个黑色(必须2个，也可以说红色不可上下相连)。⑤任何节点到属下所有叶子路径上黑色节点数量相同(黑色平衡)。
操作①变色：节点颜色红《==》黑变色。
操作②左旋：以某节点A左旋，A右儿子成为父亲，右儿子的左孙子成为A的右儿子，A的左儿子不变。
操作③右旋：以某节点B右旋，B左儿子成为父亲，左儿子的右孙子成为B的左儿子，B的右儿子不变。

*/

// RBTNode 英雄的结构体
type RBTNode struct {
	IsRed  bool     // red=true;black=false
	Key    int      // 排序序号
	Label  string   // 标签，本节点说明
	Parent *RBTNode // 父节点
	Left   *RBTNode // 左儿子节点
	Right  *RBTNode // 右儿子节点
}

// NewRBTNode 构造函数
func NewRBTNode(isRed bool, key int, label string, parent *RBTNode, left *RBTNode, right *RBTNode) *RBTNode {
	return &RBTNode{
		IsRed:  isRed,
		Key:    key,
		Label:  label,
		Parent: parent,
		Left:   left,
		Right:  right,
	}
}

// DemoPush 简易附加在尾部，回头废止
func (r *RBTNode) DemoPush(isLeft bool, san *RBTNode) {
	if isLeft { // 附加在左边
		r.Left = san
	} else { // 附加在右边
		r.Right = san
	}
	san.Parent = r
}

// ShowTree 逐层显示这个树，回头废止。感觉需要递归、队列
func (r *RBTNode) ShowTree() {
	var data [10][1000]*RBTNode // 数据。数据可能是nil。最多10层，每层最多1000数据
	//totalLevel := 0             // 总层数
	//nowLevel := 0               // 当前层数
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
func ShowOneNode(n *RBTNode) {
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

/*
 * 以p为支点左旋
 *   parent                 parent
 *     |                      |
 *     p                      pr
 *    / \        ===>     	 /  \
 *  pl   pr                 p    rr
 *      /  \               / \
 *     rl   rr            pl  rl
 *
 */

// LeftRotate 左旋转
// @param root 树的根
// @param p point 旋转的支点
func (r *RBTNode) LeftRotate(p *RBTNode) (err error) {
	if p == nil {
		return errors.New("出错，本节点是空！")
	}
	if p.Right == nil {
		return errors.New("出错，本节点右儿子是空！")
	}

	parent := p.Parent // 父亲
	pSelf := p         // 本身
	pr := p.Right      // 右儿子
	rl := pr.Left      // 右儿子的左孙子

	// 下方p是否root都要执行
	// 2.3 (升级的)pr右指向：不动
	pSelf.Parent = pr // 3.1 pSelf上指向：
	// 3.2 pSelf左指向：不动
	pSelf.Right = rl // 3.3 pSelf右指向：rl
	if rl != nil {
		rl.Parent = pSelf // 4.1 pr.Left上指向：
	}
	// 4.2 pr.Left左指向：不动
	// 4.3 pr.Left右指向：不动

	// 下方需要判断 p 是否root
	if parent != nil { // p不是root，，还需要分析，p在父亲的左还是右
		if parent.Left == pSelf { // p 在父亲左手
			parent.Left = pr // 1.1 父亲左指向：
		} else { // p 在父亲右手
			parent.Right = pr // 1.2 父亲右指向：
		}
		pr.Parent = parent // 2.1 (升级的)pr上指向：
		pr.Left = pSelf    // 2.2 (升级的)pr左指向：
	} else { // p是root
		pr.Left = pSelf // 2.2 (升级的)pr左指向：
		*r = *pr        // 1.1 + 1.2 父亲的指向。其实是root指向pr
		r.Parent = nil  // 2.1 (升级的)pr上指向：
	}
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
