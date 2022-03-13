package rbtutils

import (
	"errors"
	"fmt"
	"go-red-black-tree/bstmodels"
	"go-red-black-tree/global"
	"go-red-black-tree/rbtmodels"
	"math"
	"math/rand"
)

// RBTDemo 红黑树演示
func RBTDemo() {
	RBTCreat()
	err := errors.New("出错，本节点是空！")
	//global.Root = nil
	//ShowTreeColor(global.Root)
	//ShowTree(global.Root)

	//err = LeftRotate(global.Root.Left) // 测通
	//err = LeftRotate(global.Root.Right) // 测通
	//err = LeftRotate(global.Root) // 测试根

	//err = RightRotate(global.Root.Left) // 测通
	//err = RightRotate(global.Root.Right) // 测通
	//err = RightRotate(global.Root) // 测试根

	//err = RightRotate(global.Root) // 测通
	//err = RightRotate(global.Root) // 测通
	//err = RightRotate(global.Root) // 测通
	//err = RightRotate(global.Root.Right) // 测通
	//err = RightRotate(global.Root.Right) // 测通

	if err != nil {
		fmt.Println(err)
	}

	ShowTreeColor(global.Root)
	//ShowTree(global.Root)
	//global.Root.Left.ShowTree()
}

// RBTCreat 红黑树创建
func RBTCreat() {
	// 定义
	global.Root = rbtmodels.NewRBTNode(false, 1, "1宋江", nil, nil, nil)
	DemoPush(global.Root, true, rbtmodels.NewRBTNode(false, 2, "2卢俊义", nil, nil, nil))
	DemoPush(global.Root, false, rbtmodels.NewRBTNode(true, 3, "3吴用", nil, nil, nil))
	DemoPush(global.Root.Left, true, rbtmodels.NewRBTNode(false, 4, "4公孙胜", nil, nil, nil))
	DemoPush(global.Root.Left, false, rbtmodels.NewRBTNode(true, 5, "5关胜", nil, nil, nil))
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

// RBTInputs 红黑树连续插入节点
func RBTInputs() {
	//RBTCreat()
	for {
		var key int
		fmt.Println("请输入KEY，按回车键(空按回车随机，-1退出)：")
		fmt.Scanln(&key)

		if key == -1 {
			return
		}
		if key == 0 {
			key = rand.Intn(global.MaxKey)
			fmt.Println(key)
		}
		if key > 99 || key < 1 {
			fmt.Println("必须是0~~99")
			continue
		}
		Put(key, "")
		ShowTreeColor(global.Root)
	}
}

// RBTDeletes 红黑树连续删除节点
func RBTDeletes() {

	for {
		var key int
		fmt.Println("请输入KEY，按回车键(空按回车随机，-1退出)：")
		fmt.Scanln(&key)

		if key == -1 {
			return
		}
		if key == 0 {
			key = rand.Intn(global.MaxKey)
			fmt.Println(key)
		}
		if key > 99 || key < 1 {
			fmt.Println("必须是0~~99")
			continue
		}
		node, err := Find(key)
		if err != nil {
			fmt.Println("查找错误，error == ", err)
			continue
		}
		Delete(node)
		ShowTreeColor(global.Root)
	}
}

// Put 加入节点
// @key 插入的键值
// @label 插入的标签值
func Put(key int, label string) {
	if global.Root == nil { // 原树为空树，新加入的转为根、黑色
		global.Root = rbtmodels.NewRBTNode(false, key, label, nil, nil, nil)
		return
	}

	// 从root开始查找附加的位置
	tempParent := global.Root // 临时的父亲，移动的指针
	var isToLeft bool         // 新加节点在tempParent的左儿子吗？
	for {
		if tempParent.Key > key { // 新来数值小，向左搜索
			if tempParent.Left == nil { // 左为空，左就是new位置，跳出循环
				isToLeft = true
				break
			}
			tempParent = tempParent.Left
		} else if tempParent.Key < key { // 新来数值大，向右搜索
			if tempParent.Right == nil { // 右为空，右就是new位置，跳出循环
				isToLeft = false
				break
			}
			tempParent = tempParent.Right
		} else { // 相等，就更新标签，完成任务退出
			tempParent.Label = label
			return
		}
	}

	// 找到位置了，开始拼装。global.NewUpNode是拟增加的节点（也可能是下级旋转上升上来的随机色节点）
	global.NewUpNode = rbtmodels.NewRBTNode(true, key, label, tempParent, nil, nil)
	if isToLeft { // 拼装在左儿子
		tempParent.Left = global.NewUpNode
	} else { // 拼装在右儿子
		tempParent.Right = global.NewUpNode
	}
	ShowTreeColor(global.Root)
	FixAfterPut() // 拼装后，要调整，包括旋转+变色，可能递归

	return
}

// FixAfterPut  拼装后，要调整，包括旋转+变色，可能递归
// global.NewUpNode是拟增加的节点（也可能是下级旋转上升上来的随机色节点）
func FixAfterPut() {
	err := errors.New("出错，本节点是空！")

	// [1]新加节点or上升上来的节点是root，改黑==》结束
	if global.NewUpNode == global.Root {
		global.Root.IsRed = false
		return
	}

	// [2]（二三四树原来有1个节点），新加一个红，上黑下红，不变
	if global.NewUpNode.Parent.IsRed == false { // 新加节点or上升上来的 的父亲黑色就不用旋转 ==》结束
		return
	}

	// [3.1] 父红，叔红(不能空，空算黑)， ==》爷红，父叔黑，爷爷变为当前节点 ==》递归
	/*    gB            gR
	 *   /  \          /  \
	 * flR  urR  ==> flB  urB
	 *   \             \
	 *   srR           srR
	 */
	if global.NewUpNode.Parent.Parent.Left != nil && global.NewUpNode.Parent.Parent.Right != nil { // 确保有叔叔
		if global.NewUpNode.Parent.Parent.Left.IsRed && global.NewUpNode.Parent.Parent.Right.IsRed {
			global.NewUpNode.Parent.Parent.Left.IsRed = false  // 父叔黑
			global.NewUpNode.Parent.Parent.Right.IsRed = false // 父叔黑
			global.NewUpNode.Parent.Parent.IsRed = true        // 爷红
			global.NewUpNode = global.NewUpNode.Parent.Parent  // 爷爷变为当前节点 ==》递归
			FixAfterPut()                                      // 递归
			return
		}
	}

	// 到这里，叔叔必然黑(或空)
	if global.NewUpNode.Parent.Parent.Left == global.NewUpNode.Parent { // [4.1] 父在爷左手
		// [4.1] 父在爷左手
		// [4.1.1] 父flR红，叔黑(空也算黑)，我在右， ==》以父flR为P左旋，原父flR做当前系欸但 ==》递归
		// [4.1.2] 父srR红，叔黑(空也算黑)，我在左(其实[4.1.1]递归过来就是这个)， ==》父黑爷红，以爷爷gB为P右旋，
		//       ==》新爷爷变红，父叔边喝，爷爷作为新节点递归
		/*   gB                   gB              srR                 srR
		 *   /  \     flR左旋     /  \   gB右旋    /   \  父黑爷孙红    /   \
		 *  flR urB    ==>    srR   urB  ==>  flR     gB  ==>      flB   gB
		 *   \                /                        \                  \
		 *   srR            flR                        urB                urB
		 */
		if global.NewUpNode.Parent.Right == global.NewUpNode { // [2.1.1]我在爸爸右手，flR左旋
			err = LeftRotate(global.NewUpNode.Parent)
			if err != nil {
				fmt.Println(err)
			}
			global.NewUpNode = global.NewUpNode.Left // 模拟新加的基准点，向左下移一下，递归
			FixAfterPut()                            // 递归
			return
		}
		// 到这里一定是[4.1.2] 父srR红，叔黑(空也算黑)，我在左(其实[4.1.1]递归过来就是这个)， ==》父黑爷红，以爷爷gB为P右旋，
		//		==》原爷爷的右手做当前节点，黑就结束，红就递归原
		err = RightRotate(global.NewUpNode.Parent.Parent) // 以爷爷gB为P右旋。
		if err != nil {
			fmt.Println(err)
		}
		global.NewUpNode.IsRed = false             // =》我flR变黑
		global.NewUpNode = global.NewUpNode.Parent // 我的父亲(新的爷爷是红)，作为新节点递归
		FixAfterPut()                              // 递归
		return
	} else { // [4.2] 父在爷右手
		// [4.2] 父在爷左手
		// [4.2.1] 父frR红，叔黑(空也算黑)，我在左， ==》以父frR为P右旋，原父frR做当前系欸但 ==》递归
		// [4.2.2] 父srR红，叔黑(空也算黑)，我在左(其实[4.1.1]递归过来就是这个)， ==》父黑爷红，以爷爷gB为P右旋，
		//       ==》新爷爷变红，父叔边喝，爷爷作为新节点递归

		/* [2.2.1]右三，爷右左，黑红红=》父亲支点右旋=》爷右右，黑红红
		 * [2.2.2]右三，爷右右，黑红红=》爷爷支点左旋=》上黑两下红。
		 *
		 *    gB             gB             slR                  slR
		 *   / \   frR右旋   /  \    gB左旋  /   \   父黑爷孙红    /    \
		 * ulB  frR  ==>   ulB slR   ==>  gB     frR   ==>     gB     frB
		 *     /                \        /
		 *    slR               frR    ulB
		 */
		if global.NewUpNode.Parent.Left == global.NewUpNode { // [4.2.1] 父frR红，叔黑(空也算黑)，我在左，
			err = RightRotate(global.NewUpNode.Parent)
			if err != nil {
				fmt.Println(err)
			}
			global.NewUpNode = global.NewUpNode.Right // 模拟新加的基准点，向右下移一下
			FixAfterPut()                             // 递归
			return
		}
		// 到这里一定是[4.2.2] 父slR红，叔黑(空也算黑)，我在右(其实[4.2.1]递归过来就是这个)， ==》父黑爷红，以爷爷gB为P左旋，
		//		==》原爷爷的右手做当前节点，黑就结束，红就递归原
		err = LeftRotate(global.NewUpNode.Parent.Parent) // 以爷爷gB为P左旋，
		if err != nil {
			fmt.Println(err)
		}

		global.NewUpNode.IsRed = false             // =》我frR变黑
		global.NewUpNode = global.NewUpNode.Parent // 我的父亲(新的爷爷是红)，作为新节点递归
		FixAfterPut()                              // 递归
		return
	}

	//// 然后，看爷爷，和太爷爷，双红就递归。要提前定义好global.NewUpNode
	//global.NewUpNode = global.NewUpNode.Parent // 重新定义global.NewUpNode。四种情况，都是一样的语句
	//if !global.NewUpNode.IsRed {               // 新up节点黑，返回
	//	return
	//}
	//if global.NewUpNode.Parent == nil || global.NewUpNode.Parent.IsRed { // 新up节点是根 or up的父亲红，递归
	//	FixAfterPut() // 递归
	//}
	fmt.Println("意外，执行了FixAfterPut()递归的最后一个return")
	return
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

// ShowTreeColor 彩色逐层显示这个树，回头废止。感觉需要递归、队列
func ShowTreeColor(r *rbtmodels.RBTNode) {
	var data [10][1000]*rbtmodels.RBTNode     // 数据。数据可能是nil。最多10层，每层最多1000数据
	var dataTemp [10][1000]*rbtmodels.RBTNode // 临时数据，回头某列全nil就不显示。数据可能是nil。最多10层，每层最多1000数据
	totalLevel := 1                           // 总层数
	//nowLevel := 0               // 当前层数
	//nnn := Name
	//nowColumn := 0 // 当前列
	if r == nil {
		fmt.Println("这个树/分支是空的")
	}
	data[0][0] = r // 来的最高位指针

	// 循环，把每个节点指针放入对应层的队列，每个上级节点占死2个下级，没有这个子节点就空着
	for i := 1; i < len(data); i++ { // 循环每一层
		countNotNil := 0                                      // 本层非nil个数，==0 表示上一层是最后一层
		for j := 0; j < int(math.Pow(2, float64(i-1))); j++ { // 上层应有的元素数量，遍历，本层翻倍
			if data[i-1][j] != nil {
				if data[i-1][j].Left != nil {
					countNotNil++
					data[i][j*2] = data[i-1][j].Left // 上层左儿子，放入
				}
				if data[i-1][j].Right != nil {
					countNotNil++
					data[i][j*2+1] = data[i-1][j].Right // 上层右儿子，放入
				}
			}
		}
		if countNotNil == 0 { // 本层无元素，中断，退出
			break
		}
		totalLevel++ // 总层数
	}

	// 二次循环，数据导入dataTemp，一组占3列，父亲骑在2个儿子中间
	for i := 0; i < totalLevel; i++ { // 循环每一层
		blankMiddleLen := int(math.Pow(2, float64(totalLevel-i))) - 1 // 中间空
		blankLeftLen := (blankMiddleLen - 1) / 2

		for j := 0; j < int(math.Pow(2, float64(i))*1.5); j++ { // 上层应有的元素数量，遍历，本层翻倍
			dataTemp[i][blankLeftLen+j*(blankMiddleLen+1)+1] = data[i][j]
			//fmt.Printf("i== %d ,j== %d ,totalLevel== %d ,blankMiddleLen== %d ,blankLeftLen== %d \n", i, j, totalLevel, blankMiddleLen, blankLeftLen)
		}
	}

	// 三次循环，把每层数据展示出来，
	for i := 0; i < totalLevel; i++ { // 循环每一层
		for j := 0; j < int(math.Pow(2, float64(totalLevel))); j++ { // 应有的元素数量

			// 查本列是否全nil
			isAllNil := true // 本列全nil
			for k := 0; k < totalLevel; k++ {
				if dataTemp[k][j] != nil {
					isAllNil = false
				}
			}
			if !isAllNil { // 只有本列不是全nil，才show一下
				ShowOneNodeColorNew(dataTemp[i][j], totalLevel, i, j)
			}
			//ShowOneNodeColorNew(dataTemp[i][j], totalLevel, i, j)
		}

		fmt.Println()
	}

	//// 三次循环，把每层数据展示出来，
	//for i := 1; i < totalLevel+1; i++ { // 循环每一层
	//	//nowColumn = 0                                         // 当前列
	//	for j := 0; j < int(math.Pow(2, float64(i-1))); j++ { // 上层应有的元素数量，遍历，本层翻倍
	//		ShowOneNodeColor(data[i-1][j], totalLevel, i, j)
	//	}
	//	fmt.Println()
	//}
}

// ShowOneNodeColorNew 彩色展示单个节点
func ShowOneNodeColorNew(n *rbtmodels.RBTNode, totalLevel, i, j int) {

	// blank=blankLeft+n*(位长global.KeyLen+1)
	blankNil := "" // 空节点，也占位置
	for k := 0; k < global.KeyLen; k++ {
		blankNil = blankNil + " "
	}

	if n == nil {
		fmt.Printf("%s", blankNil)
	} else {
		//其中0x1B是标记，[开始定义颜色，1代表高亮，40代表黑色背景，32代表绿色前景，0代表恢复默认颜色。
		red := 31
		black := 0
		if n.IsRed {
			fmt.Printf("%c[1;0;%dm%02d%c[0m", 0x1B, red, n.Key, 0x1B)
		} else {
			fmt.Printf("%c[1;0;%dm%02d%c[0m", 0x1B, black, n.Key, 0x1B)
		}
	}

}

// ShowOneNodeColor 彩色展示单个节点
func ShowOneNodeColor(n *rbtmodels.RBTNode, totalLevel, i, j int) {
	// blank=blankLeft+n*(位长global.KeyLen+1)
	blankNil := "" // 空节点，也占位置
	for k := 0; k < global.KeyLen+1; k++ {
		blankNil = blankNil + " "
	}
	blankLeftHead := ""                                                             //  最左边                                                  // 总体左边空
	blankMiddleLen := int(math.Pow(2, float64(totalLevel-i)))*(global.KeyLen+1) - 3 // 中间空
	blankLeftLen := blankMiddleLen / 2
	blankLeft := blankLeftHead // 最左空的
	for k := 0; k < blankLeftLen; k++ {
		blankLeft = blankLeft + " "
	}
	blankMiddle := "" // 中间空的
	for k := 0; k < blankMiddleLen; k++ {
		blankMiddle = blankMiddle + " "
	}

	if j == 0 { // 本列第一个
		fmt.Printf("%s", blankLeft)
	} else {
		fmt.Printf("%s", blankMiddle)
	}

	if n == nil {
		fmt.Printf("%s", blankNil)
	} else {
		//其中0x1B是标记，[开始定义颜色，1代表高亮，40代表黑色背景，32代表绿色前景，0代表恢复默认颜色。
		red := 31
		black := 0
		if n.IsRed {
			fmt.Printf("%c[1;0;%dm%02d%c[0m", 0x1B, red, n.Key, 0x1B)
		} else {
			fmt.Printf("%c[1;0;%dm%02d%c[0m", 0x1B, black, n.Key, 0x1B)
		}
	}

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

			ShowOneNode(data[i-1][j]) // 显示遍历到的上一行的这个节点。显示没换行

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
	if n == nil {
		fmt.Println("[]nil[]")
		return
	}
	fmt.Printf(" , ")  // 左右分割
	if n.Left == nil { // 左儿子KEY
		fmt.Printf("[ ]")
	} else {
		fmt.Printf("[%d]", n.Left.Key)
	}

	fmt.Printf("%d", n.Key) // 本节点KEY
	if n.IsRed == true {    // 本节点是红色
		fmt.Printf("R")
	} else { // 黑色
		fmt.Printf("B")
	}

	if n.Parent == nil { // 父节点KEY
		fmt.Printf("( )")
	} else {
		fmt.Printf("(%d)", n.Parent.Key)
	}

	if n.Right == nil { // 右节点KEY
		fmt.Printf("[ ]")
	} else {
		fmt.Printf("[%d]", n.Right.Key)
	}
}

// LeftRotate 左旋+变色
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

	//// 下来的P ==》红 ；上去的pr ==》黑
	//p.IsRed = true
	//pr.IsRed = false

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

// RightRotate 右旋+变色
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

	//// 下来的P ==》红 ；上去的pl ==》黑
	//p.IsRed = true
	//pl.IsRed = false

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

// Find 查找节点
// @key 键值
// @n 找到的节点指针
func Find(key int) (ret *rbtmodels.RBTNode, err error) {
	tempNode := global.Root
	for { // 递归循环
		if tempNode.Key == key { // 等于，找到了
			return tempNode, nil
		} else if tempNode.Key > key { // 大于，向左
			if tempNode.Left == nil { // 到nil了，返没找到
				return nil, errors.New("没找到！")
			}
			tempNode = tempNode.Left // 向左下递归
		} else { // 小于，向右
			if tempNode.Right == nil { // 到nil了，返没找到
				return nil, errors.New("没找到！")
			}
			tempNode = tempNode.Right // 向右下递归
		}
	}
	return
}

// Delete 删除节点
// @key 键值
func Delete(deleteNode *rbtmodels.RBTNode) {
	/*
	 * 普通二叉树，删除的时候，用前驱节点or后继节点替换(内容)，然后删掉替代节点(替代节点可能有单叶子，把这个单叶子提上来)
	 * 找前驱，如没有左子树(删除时候用不到这种情形)，就向父亲不断找，直到向左拐弯的地方。
	 * 找后继，如没有右子树(删除时候用不到这种情形)，就向父亲不断找，直到向右拐弯的地方。
	 * 找后继，如没有右子树(删除时候用不到这种情形)，就向父亲不断找，直到向右拐弯的地方。
	 * [1]自己能搞定，相当于234树的三or四节点，删掉低端红色；或删掉高位黑色，唯一儿子变黑上来
	 * [2]自己搞不定，相当于234树的二节点，删掉唯一黑色后，把父亲借下来，234树的兄弟(多节点)顶上去父亲位置
	 * [2.1]自己搞不定，相当于234树的二节点，删掉唯一黑色后，把父亲借下来，234树的兄弟(多节点)顶上去父亲位置
	 * [3]自己搞不定，相当于234树的二节点，删掉唯一黑色后，把父亲借下来，234树的兄弟(二节点)也没得借

	 */
	//avatarNode := Predecessor(deleteNode) // 用前驱节点做替身
	//avatarNode := Successor(deleteNode) // 用后继节点做替身
	ShowOneNode(Predecessor(deleteNode))
	ShowOneNode(deleteNode)
	ShowOneNode(Successor(deleteNode))
	ShowOneNode(deleteNode)
	fmt.Println()
	return
}

// Predecessor 找前驱节点。比我稍小的最大节点
func Predecessor(n *rbtmodels.RBTNode) (ret *rbtmodels.RBTNode) {
	if n == nil {
		fmt.Println("纳尼？让我给一个nil找前驱节点？")
		return nil
	}
	if n.Left != nil { // 有左儿子，找左儿子下面的最大
		ret = n.Left
		for {
			if ret.Right == nil { // 右下nil就算找到了
				return ret
			}
			ret = ret.Right // 沿着右手一直向下
		}
	} else { // 没有左儿子，向上找，每个和我比较，到root，返nil
		if n.Parent == nil { // 没左儿子，自己又是root
			fmt.Println("这个真没有，没左儿子，自己又是root")
			return nil
		}
		ret = n
		for {
			if ret.Parent == nil { // 游标移到本身是root，还没有找到，就nil了
				return nil
			} else {
				if ret.Parent.Key < n.Key { // 某个父辈小于我了，这个就是
					return ret.Parent
				}
			}
			ret = ret.Parent // 向上一直找
		}
	}
	return
}

// Successor 找后继节点。比我稍大的最小节点
func Successor(n *rbtmodels.RBTNode) (ret *rbtmodels.RBTNode) {
	if n == nil {
		fmt.Println("纳尼？让我给一个nil找后继节点？")
		return nil
	}
	if n.Right != nil { // 有右儿子，找右儿子下面的最大
		ret = n.Right
		for {
			if ret.Left == nil { // 左下nil就算找到了
				return ret
			}
			ret = ret.Left // 沿着左手一直向下
		}
	} else { // 没有右儿子，向上找，每个和我比较，到root，返nil
		if n.Parent == nil { // 没右儿子，自己又是root
			fmt.Println("这个真没有，没右儿子，自己又是root")
			return nil
		}
		ret = n
		for {
			if ret.Parent == nil { // 游标移到本身是root，还没有找到，就nil了
				return nil
			} else {
				if ret.Parent.Key > n.Key { // 某个父辈小于我了，这个就是
					return ret.Parent
				}
			}
			ret = ret.Parent // 向上一直找
		}
	}
	return
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
