# go-red-black-tree

# A Red Black Tree coding by golang
```
I Insert插入数据
S Show完整的树
F Find查找数据
D Delete删除数据
Q PreOrder 前序遍历
Z InfixOrder 中序遍历
H PostOrder 后序遍历
E Exit退出
```
###Can Show The As Color<br/>
　　　　　22                  
　　04　　　　　　　　66      
　03　05　　　　<font color=red>45</font>　　　　　<font color=red>69</font>  
     <font color=red>02</font>　　　<font color=red>21</font>　　37　48　　　68　78<br/>
   　　　　　  <font color=red>36</font>　　　　<font color=red>54</font>  


```
// [3.1]avatar黑兄红侄
/* [3.1.1]avatar在右，黑兄红侄.侄子左红右随意。==》原兄升父位继承父的颜色/原父下来变黑(填补双黑值)/原红侄黑==》父轴，右旋，==》结束
*  [3.1.1.J]镜像，avatar在左
*          (50Y)       |        (30Y)           ||        (50Y)           |          (70Y)       |
*        /       \     |      /      \          ||      /      \          |        /       \     |
*     (30B)     A(60B) |   (20B)     (50B)      ||  A(40B)     (70B)      |     (50B)      (80B) |
*     /    \     / \   |   / \      /   \       ||   / \      /   \       |     /    \     / \   |
*  (20R)  (40X)(?5)(?6)|(?1)(?2) (40X)  (60B)   ||(?1)(?2) (60X)  (80R)   |  (40R)  (60X)(?5)(?6)|
*   / \    / \         |         / \      / \   ||         / \      / \   |   / \    / \         |
*(?1)(?2)(?3)(?4)      |       (?3)(?4) (?5)(?6)||       (?3)(?4) (?5)(?6)|(?1)(?2)(?3)(?4)      |
 */
 
 
/* [3.1.2]avatar在右，黑兄红侄.侄子左黑右红。==>原父下来变黑(填补双黑值)/原红侄升父位继承父的颜色==》兄轴左旋==》父轴，右旋，==》结束
*  [3.1.2.J]镜像avatar在左，黑兄红侄.侄子右黑左红。==>原父下来变黑(填补双黑值)/原红侄升父位继承父的颜色==》兄轴右旋==》父轴，左旋，==》结束
*          (50Y)       |          (50Y)      |         (40Y)        ||       (50Y)           |       (50Y)         |         (60Y)        |
*        /       \     |        /      \     |       /      \       ||     /      \          |     /     \         |       /      \       |
*     (30B)     A(60B) |     (40R)    A(60B) |    (30B)     (50B)   || A(40B)     (70B)      | A(40B)   (60R)      |    (5B)     (70B) c  |
*     /    \     / \   |     /    \    / \   |    / \      /   \    ||  / \       /    \     |  /  \    / \        |    / \      /   \    |
*  (20B)  (40R)(?5)(?6)|  (30B)  (?4)(?5)(?6)| (20B)(?3) (?4) (60B) ||(?1)(?2) (60R)  (80B)  |(?1)(?2)(?3)(70B)    | (40B)(?3) (?4) (80B) |
*   / \    / \         |   / \               |  / \           /  \  ||          / \    / \   |             / \     |  / \           /  \  |
*(?1)(?2)(?3)(?4)      |(20B)(?3)            |(?1)(?2)      (?5)(?6)||       (?3)(?4)(?5)(?6)|          (?4)(80B)  |(?1)(?2)      (?5)(?6)|
*                      |  / \                |                      ||                       |               / \   |                      |
*                      |(?1)(?2)             |                      ||                       |             (?5)(?6)|                      |
 */
 
 /* [3.2]avatar在左右都行，黑兄黑侄红父。==》父黑兄红==》结束
* 对应2-3-4树删除操作中兄弟节点为2节点，父节点至少是个3节点，父节点key下移与兄弟节点合并。
* 黑兄红侄.侄子左红右随意。==》父轴，右旋，父下来变黑(填补双黑值)/升父位的兄继承父的颜色/左侄黑==》结束
*          (50R)       |          (50B)       |
*        /       \     |        /       \     |
*     (30B)     A(60B) |     (30R)     A(60B) |
*     /    \     / \   |     /    \     / \   |
*  (20B)  (40B)(?5)(?6)|  (20B)  (40B)(?5)(?6)|
*   / \    / \         |   / \    / \         |
*(?1)(?2)(?3)(?4)      |(?1)(?2)(?3)(?4)      |
 */
 
 /* [3.3]avatar在左右都行，黑兄黑侄黑父。==》兄红==》父亲做为avatar==》递归
*          (50B)       |         A(50B)       ||        (30B)           |       A(30B)           |
*        /       \     |        /       \     ||      /      \          |      /      \          |
*     (30B)     A(60B) |     (30R)      (60B) ||  A(20B)     (50B)      |   (20B)     (50R)      |
*     /    \     / \   |     /    \     / \   ||   / \      /   \       |   / \      /   \       |
*  (20B)  (40B)(?5)(?6)|  (20B)  (40B)(?5)(?6)||(?1)(?2) (40B)  (60B)   |(?1)(?2) (40B)  (60B)   |
*   / \    / \         |   / \    / \         ||         / \      / \   |         / \      / \   |
*(?1)(?2)(?3)(?4)      |(?1)(?2)(?3)(?4)      ||       (?3)(?4) (?5)(?6)|       (?3)(?4) (?5)(?6)|
 */
 
/* [3.4]avatar在右，红兄黑侄黑父。==>原兄黑，原父红==》父轴右旋==》avatar的新兄弟就是黑的==》用avatar递归（红父黑兄[3.2]）
*  [3.4J]镜像avatar在左，红兄黑侄黑父。==>原兄黑，原父红==》父轴左旋==》avatar的新兄弟就是黑的==》用avatar递归（红父黑兄[3.2]）
*  [注意]就地讨论avatar侄子是不是存在，否则递归下去默认侄子是全的
*          (50B)       |        (30B)           ||        (50B)           |          (70B)       |
*        /       \     |      /      \          ||      /      \          |        /       \     |
*     (30R)     A(60B) |   (20B)     (50R)      ||  A(40B)     (70R)      |     (50R)      (80B) |
*     /    \     / \   |   / \      /   \       ||   / \      /   \       |     /    \     / \   |
*  (20B)  (40B)(?5)(?6)|(?1)(?2) (40B) A(60B)   ||(?1)(?2) (60B)  (80B)   | A(40B)  (60B)(?5)(?6)|
*   / \    / \         |         / \      / \   ||         / \      / \   |   / \    / \         |
*(?1)(?2)(?3)(?4)      |       (?3)(?4) (?5)(?6)||       (?3)(?4) (?5)(?6)|(?1)(?2)(?3)(?4)      |
 */
```
### Can Show a RBT color
### Include  a binary-tree

