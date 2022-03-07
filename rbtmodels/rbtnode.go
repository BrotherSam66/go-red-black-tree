package rbtmodels

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
