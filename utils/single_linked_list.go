package utils

import (
	"github.com/Edilberto-Vazquez/weathercloud-etl-process/db/models"
)

type Node struct {
	Value models.WeatherCloud
	Next  *Node
}

type Singlylinkedlist struct {
	Head   *Node
	Tail   *Node
	length int
}

func (sll *Singlylinkedlist) Length() int {
	return sll.length
}

func (sll *Singlylinkedlist) Append(value models.WeatherCloud) {
	var newNode *Node = &Node{Value: value}
	if sll.Head == nil {
		sll.Head = newNode
		sll.Tail = sll.Head
	} else {
		sll.Tail.Next = newNode
		sll.Tail = newNode
	}
	sll.length++
}
