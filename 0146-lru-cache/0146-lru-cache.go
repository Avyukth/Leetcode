type Node struct {
    key, value int
    prev, next *Node
}

type LRUCache struct {
    capacity int
    cache map[int]*Node
    head *Node
    tail *Node
}


func Constructor(capacity int) LRUCache {
    lru := LRUCache{
        capacity : capacity,
        cache :make(map[int]*Node),
        head: &Node{},
        tail : &Node{},
    }
    lru.head.next = lru.tail
    lru.tail.prev = lru.head

    return lru
}


func (this *LRUCache) Get(key int) int {
    if node , exists := this.cache[key]; exists{
        this.moveToFront(node)
        return node.value
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if node , exists := this.cache[key]; exists{
        node.value = value
        this.moveToFront(node)
    }else{
        if len(this.cache)>= this.capacity{
            this.removeLRU()
        }
        newNode := &Node{key: key, value: value}
        this.cache[key]= newNode
        this.addToFront(newNode)
    }
}

func (this *LRUCache)moveToFront(node *Node){
    this.removeNode(node)
    this.addToFront(node)
}

func (this *LRUCache)addToFront(node *Node){
    node.prev = this.head
    node.next = this.head.next
    this.head.next.prev =node
    this.head.next = node
}

func (this *LRUCache)removeNode(node *Node){
    node.prev.next = node.next
    node.next.prev = node.prev
}
func (this *LRUCache)removeLRU(){
    lru := this.tail.prev
    this.removeNode(lru)
    delete(this.cache, lru.key)
}
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */