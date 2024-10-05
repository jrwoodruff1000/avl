/*
# avl

Package creates and maintains an AVL (self-balancing) Binary Search Tree data structure

	    struct		Avl(Name string)
		interface		Key [int|string]
		func		NewAvl [int|string] (pAvlName string) *Avl
		func 		Add_node(pKey Key, pPayload any) error
		func 		Delete_node(key Key) error
		func		Update_payload(key Key, pPayload any) error
		func		Count_nodes() int
		func		Get_payload(key Key) (any, error)
		func		Get_metadata() (string)
		func		Max_node() (K, error)
		func		Min_node() (K, error)
		func 		Previous_node(inpKey K) (K, error)
		func		Next_node(inpKey K) (K, error) {
*/
package avl

import (
	"errors"
	"log"
	"os"
)

type Key interface {
	int | string
}

type Avl[K Key] struct {
	avlName string
	pTop    *avlNode[K]
}

// Constructor (key can only be of type int or string)
func NewAvl[K Key](inpAvlName string) *Avl[K] {

	// Example error handler
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error while creating new node: ", err)
			os.Exit(1)
		}
	}()

	newAvl := new(Avl[K])
	newAvl.avlName = inpAvlName
	newAvl.pTop = nil
	return newAvl
}

// Returns error
func (a *Avl[K]) Add_node(inpKey K, inpPayload any) error {
	var ppHead **avlNode[K] = &a.pTop
	if add_node_rec(ppHead, nil, inpKey, inpPayload) == 0 {
		return errors.New("Node already exists")
	} else {
		a.pTop = *ppHead
		return nil
	}
}

// Returns error
func (a *Avl[K]) Delete_node(inpKey K) error {
	var nodeToDelete *avlNode[K] = nil
	var tempParent *avlNode[K] = nil
	var ppHead **avlNode[K] = &a.pTop

	nodeToDelete = find_node(inpKey, a.pTop)

	if nodeToDelete == nil {
		return errors.New("Node not found")
	}

	tempParent = nodeToDelete.pParent
	delete_node_rec(ppHead, inpKey, a.pTop)
	a.pTop = *ppHead

	if tempParent != nil {
		a.pTop = balance_node(a.pTop, tempParent)
	} else if a.pTop != nil {
		a.pTop = balance_node(a.pTop, a.pTop)
	}

	return nil
}

// Returns error
func (a *Avl[K]) Update_payload(inpKey K, inpPayload any) error {
	var foundNode *avlNode[K] = find_node(inpKey, a.pTop)

	if foundNode == nil {
		return errors.New("Node not found")
	} else {
		foundNode.payload = inpPayload
		return nil
	}
}

// Returns count
func (a *Avl[K]) Count_nodes() int {
	var myCount int = 0

	if a.pTop == nil {
		myCount = 0
	} else {
		count_nodes_rec(a.pTop, &myCount)
	}
	return myCount
}

// Returns payload with error
func (a *Avl[K]) Get_payload(inpKey K) (any, error) {
	foundNode := find_node(inpKey, a.pTop)
	if foundNode == nil {
		return nil, errors.New("Node not found")
	} else {
		return foundNode.payload, nil
	}
}

// Returns metadata
func (a *Avl[K]) Get_metadata() string {
	return a.avlName
}

// Returns key of maximum node or zero value plus error
func (a *Avl[K]) Max_node() (K, error) {
	var zeroValueKey K

	maxNode := find_max(a.pTop)
	if maxNode == nil {
		return zeroValueKey, errors.New("Tree is empty")
	} else {
		return maxNode.key, nil
	}
}

// Returns key of minimum node or zero value plus error
func (a *Avl[K]) Min_node() (K, error) {
	var zeroValueKey K

	minNode := find_min(a.pTop)
	if minNode == nil {
		return zeroValueKey, errors.New("Tree is empty")
	} else {
		return minNode.key, nil
	}
}

// Returns key of previous node or zero value plus error
func (a *Avl[K]) Previous_node(inpKey K) (K, error) {
	var zeroValueKey K

	foundNode := find_node(inpKey, a.pTop)
	if foundNode == nil {
		return zeroValueKey, errors.New("Node not found")
	} else {
		if foundNode.pLeft != nil {
			previousNode := find_max(foundNode.pLeft)
			return previousNode.key, nil
		} else {
			// First up traversal when moving left
			previousNode := first_left_ancestor(foundNode)
			if previousNode == nil {
				return zeroValueKey, errors.New("Current node is minimum node")
			}
			return previousNode.key, nil
		}
	}
}

// Returns key of next node or zero value plus error
func (a *Avl[K]) Next_node(inpKey K) (K, error) {
	var zeroValueKey K

	foundNode := find_node(inpKey, a.pTop)
	if foundNode == nil {
		return zeroValueKey, errors.New("Node not found")
	} else {
		if foundNode.pRight != nil {
			nextNode := find_min(foundNode.pRight)
			return nextNode.key, nil
		} else {
			// First up traversal when moving right
			nextNode := first_right_ancestor(foundNode)
			if nextNode == nil {
				return zeroValueKey, errors.New("Current node is maximum node")
			}
			return nextNode.key, nil
		}
	}
}
