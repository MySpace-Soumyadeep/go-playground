package main

import (
	"fmt"
)

const SIZE = 5

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Hash map[string]*Node

type Cache struct {
	MyQueue Queue
	MyHash  Hash
}

func NewCache() Cache {
	return Cache{MyQueue: NewQueue(), MyHash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

func (c *Cache) Check(str string) {
	node := &Node{}

	if val, ok := c.MyHash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: str}
	}
	c.Add(node)
	c.MyHash[str] = node
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove: %s\n", n.Val)

	left := n.Left
	right := n.Right

	right.Left = left
	left.Right = right
	c.MyQueue.Length -= 1
	delete(c.MyHash, n.Val)
	return n
}

func (c *Cache) Add(n *Node) {
	fmt.Printf("add: %s\n", n.Val)

	temp := c.MyQueue.Head.Right

	c.MyQueue.Head.Right = n
	n.Left = c.MyQueue.Head
	n.Right = temp
	temp.Left = n
	c.MyQueue.Length++

	if c.MyQueue.Length > SIZE {
		c.Remove((c.MyQueue.Tail.Left))
	}

}

func (c *Cache) Display() {
	c.MyQueue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < q.Length-1 {
			fmt.Printf("<--->")
		}
		node = node.Right
	}
	fmt.Printf("]\n")
}

func main() {
	fmt.Println("START CACHE")
	cache := NewCache()
	for _, word := range []string{"chicken", "paneer", "mango", "salad", "pasta", "pizza", "paneer", "chicken"} {
		cache.Check(word)
		cache.Display()
	}
}
