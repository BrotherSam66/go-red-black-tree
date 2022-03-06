package bstutils

import (
	"fmt"
	"go-red-black-tree/bstmodels"
)

// BinaryTreeDemo 二叉树演示
func BinaryTreeDemo() {
	// 定义
	heroRoot := bstmodels.NewHero(1, "宋江", nil, nil)
	heroRoot.Left = bstmodels.NewHero(2, "卢俊义", nil, nil)
	heroRoot.Right = bstmodels.NewHero(3, "吴用", nil, nil)
	heroRoot.Left.Left = bstmodels.NewHero(4, "公孙胜", nil, nil)
	heroRoot.Left.Right = bstmodels.NewHero(5, "关胜", nil, nil)
	heroRoot.Right.Left = bstmodels.NewHero(6, "林冲", nil, nil)
	heroRoot.Right.Right = bstmodels.NewHero(7, "秦明", nil, nil)
	heroRoot.Left.Left.Left = bstmodels.NewHero(8, "呼延灼", nil, nil)
	heroRoot.Left.Left.Right = bstmodels.NewHero(9, "华融", nil, nil)
	//heroRoot = nil
	fmt.Println("heroRoot == ", heroRoot)
	fmt.Println("前序遍历")
	bstmodels.PreOrder(heroRoot)
	fmt.Println("中序遍历")
	bstmodels.InfixOrder(heroRoot)
	fmt.Println("后序遍历")
	bstmodels.PostOrder(heroRoot)
	fmt.Println("层序遍历")
	bstmodels.LevelOrder(heroRoot)
	//// 选择菜单
	//for {
	//	var inputStr string
	//	var inputID int
	//	var inputName string
	//	fmt.Println("	i(insert)插入一条数据")
	//	fmt.Println("	f(find)寻找一条数据")
	//	fmt.Println("	d(delete)删除一条数据")
	//	fmt.Println("	l(list)显示全部数据")
	//	fmt.Println("	q(quit)退出")
	//	fmt.Println("请入字母，按回车键：")
	//	fmt.Scanln(&inputStr)
	//	if strings.ToUpper(inputStr) == "I" {
	//		fmt.Println("请输入雇员ID，按回车键：")
	//		fmt.Scanln(&inputID)
	//		fmt.Println("请输入雇员Name，按回车键：")
	//		fmt.Scanln(&inputName)
	//		emp := &models.Employ{ID: inputID, Label: inputName}
	//		err := hashTable.Insert(emp)
	//		if err != nil {
	//			fmt.Println("出错了：", err)
	//		}
	//		fmt.Println("插入结束")
	//	} else if strings.ToUpper(inputStr) == "F" {
	//		fmt.Println("请输入查询的雇员ID，按回车键：")
	//		fmt.Scanln(&inputID)
	//		emp, err := hashTable.DeleteOrFind(false, inputID)
	//		if err != nil {
	//			fmt.Println("出错了：", err)
	//		} else {
	//			fmt.Println("查询到的雇员资料：", emp)
	//		}
	//	} else if strings.ToUpper(inputStr) == "D" {
	//		fmt.Println("请输入删除的雇员ID，按回车键：")
	//		fmt.Scanln(&inputID)
	//		_, err := hashTable.DeleteOrFind(true, inputID)
	//		if err != nil {
	//			fmt.Println("出错了：", err)
	//		} else {
	//			fmt.Println("删除成功")
	//		}
	//	} else if strings.ToUpper(inputStr) == "L" {
	//		hashTable.List()
	//	} else if strings.ToUpper(inputStr) == "Q" {
	//		fmt.Println("您选择了q(quit)退出")
	//		os.Exit(0)
	//		return
	//	} else {
	//		fmt.Println("输入的什么乱七八糟的？")
	//	}
	//}

}
