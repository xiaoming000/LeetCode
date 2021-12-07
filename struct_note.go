//节点结构体
type Node struct {
	data interface{}
	nextNode *Node
}
//链表结构体
type List struct {
	root *Node
	count int
}

func main(){
	var list1 List
	list1.add("hello")
	list1.add(520)
	list1.add('a')
	list1.add(true)
	list1.add(nil)
	fmt.Println(list1.toArray())
	fmt.Println(list1.contains("hello"))
	fmt.Println(list1.get(1))
	list1.set(2,5201314)
	fmt.Println(list1.toArray())
	fmt.Println(list1.isEmpty())
	fmt.Println(list1.size())
	list1.add("baby")
	fmt.Println(list1.toArray())
	list1.removeByIndex(4)
	fmt.Println(list1.toArray())
	list1.removeByData(5201314)
	fmt.Println(list1.toArray())
}


//添加节点
func (node *Node)addNode(newNode *Node){
	if node.nextNode==nil{
		node.nextNode=newNode
	}else{
		node.nextNode.addNode(newNode)
	}
}

func (list *List)add(data interface{}){
	if data==nil{
		return
	}
	var newNode Node
	newNode.data=data
	if list.root==nil{
		list.root=&newNode
	}else{
		list.root.addNode(&newNode)
	}
	list.count++
}

//返回所有节点的内容
func (list *List)toArray() (datas []interface{}){
	if list.root==nil{
		return nil
	}
	node:=list.root
	for{
		datas=append(datas,node.data)
		if node.nextNode!=nil{
			node=node.nextNode
		}else{
			break
		}
	}
	return datas
}
//查询链表中是否包含某数据
func (list *List)contains(data interface{}) (isCon bool){
	if data==nil || list.root==nil{
		isCon=false
	}
	node:=list.root
	if node.data==data{
		isCon=true
	}else {
		for{
			if node.nextNode!=nil{
				if node.nextNode.data==data{
					isCon=true
					break
				} else{
					node=node.nextNode
				}
			}else{
				isCon=false
				break
			}
		}
	}
	return
}
//根据索引查询数据(从1开始算)
func (list *List)get(index int) (data interface{}){
	if index>list.count{
		data=nil
	}else{
		node:=list.root
		for i:=1;i<index;i++{
			node=node.nextNode
		}
		data=node.data
	}
	return
}
//根据索引更新数据(从1开始算)
func (list *List)set(index int,data interface{}){
	if index>list.count{
		return
	}else{
		node:=list.root
		for i:=1;i<index;i++{
			node=node.nextNode
		}
		node.data=data
	}
}
//删除指定位置的节点(从1开始算)
func (list *List)removeByIndex(index int){
	if index>list.count{
		return
	}else {
		preNode:=list.root
		if index==1{
			list.root=preNode.nextNode
			return
		}
		for i:=1;i<index-1;i++{
			preNode=preNode.nextNode
		}
		preNode.nextNode=preNode.nextNode.nextNode
	}
}
//删除第一次出现的某个值
func (List *List)removeByData(data interface{}){
	if !List.contains(data){
		return
	}
	foot:=1
	node:=List.root
	for{
		if node.data==data{
			List.removeByIndex(foot)
			return
		}else{
			foot++
			node=node.nextNode
		}
	}
}
//查询链表是否为空
func (list *List)isEmpty() bool{
	return list.count==0;
}
//查询链表的长度
func (list *List)size() int{
	return list.count
}
