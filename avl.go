package avl

import (
	"fmt"
)

const avl_DEBUG bool = false

type avlNode[K Key] struct {
	pLeft, pRight, pParent *avlNode[K]
	key                    K
	payload                any
}

// Returns found node
func find_node[K Key](key K, pNode *avlNode[K]) *avlNode[K] {
	var node *avlNode[K]

	if pNode == nil {
		return nil
	} else if pNode.key == key {
		return pNode
	} else if key < pNode.key {
		if pNode.pLeft != nil {
			node = find_node(key, pNode.pLeft)
			return node
		} else {
			return nil
		}
	} else {
		if pNode.pRight != nil {
			node = find_node(key, pNode.pRight)
			return node
		} else {
			return nil
		}
	}
}

func find_min[K Key](pNode *avlNode[K]) *avlNode[K] {
	if pNode == nil {
		return nil
	} else if pNode.pLeft == nil {
		return pNode
	} else {
		return find_min(pNode.pLeft)
	}
}

func find_max[K Key](pNode *avlNode[K]) *avlNode[K] {
	if pNode == nil {
		return nil
	} else if pNode.pRight == nil {
		return pNode
	} else {
		return find_max(pNode.pRight)
	}
}

// Used for previous node determination
func first_left_ancestor[K Key](pNode *avlNode[K]) *avlNode[K] {
	if pNode == nil || pNode.pParent == nil {
		return nil
	} else if pNode.pParent.pRight == pNode {
		return pNode.pParent
	} else {
		return first_left_ancestor(pNode.pParent)
	}
}

// Used for next node determination
func first_right_ancestor[K Key](pNode *avlNode[K]) *avlNode[K] {
	if pNode == nil || pNode.pParent == nil {
		return nil
	} else if pNode.pParent.pLeft == pNode {
		return pNode.pParent
	} else {
		return first_right_ancestor(pNode.pParent)
	}
}

// Calculate Balance Factor
func calc_bf[K Key](pNode *avlNode[K]) int {
	var leftDepth int = 0
	var rightDepth int = 0

	if pNode.pLeft != nil {
		leftDepth = node_depth(pNode.pLeft, 0)
	}
	if pNode.pRight != nil {
		rightDepth = node_depth(pNode.pRight, 0)
	}
	return rightDepth - leftDepth
}

func node_depth[K Key](pNode *avlNode[K], startCount int) int {
	var leftDepth int = 0
	var rightDepth int = 0

	if pNode.pLeft != nil {
		leftDepth = node_depth(pNode.pLeft, 0)
	}
	if pNode.pRight != nil {
		rightDepth = node_depth(pNode.pRight, 0)
	}

	if leftDepth > rightDepth {
		leftDepth++
		return leftDepth - startCount
	} else {
		rightDepth++
		return rightDepth - startCount
	}
}

func rotate_right[K Key](pHead *avlNode[K], pNode *avlNode[K]) *avlNode[K] {
	var parent_parent, pLocLeft, floater *avlNode[K]
	var was_i_left_or_right byte

	if avl_DEBUG {
		fmt.Printf("Rotating right on node: %vi\n", pNode.key)
	}

	// Do I have a parent, if so am I the left or right child?
	parent_parent = pNode.pParent
	if parent_parent != nil {
		if parent_parent.pLeft == pNode {
			was_i_left_or_right = 'L'
		} else {
			was_i_left_or_right = 'R'
		}
	}

	// floater is the right child of the left child
	pLocLeft = pNode.pLeft
	floater = pLocLeft.pRight

	if parent_parent != nil {
		if was_i_left_or_right == 'L' {
			parent_parent.pLeft = pLocLeft
		} else {
			parent_parent.pRight = pLocLeft
		}
	} else {
		pHead = pLocLeft
	}
	pLocLeft.pParent = parent_parent

	pLocLeft.pRight = pNode
	pNode.pParent = pLocLeft

	pNode.pLeft = floater
	if floater != nil {
		floater.pParent = pNode
	}

	return pHead
}

func rotate_left[K Key](pHead *avlNode[K], pNode *avlNode[K]) *avlNode[K] {
	var parent_parent, pLocRight, floater *avlNode[K]
	var was_i_left_or_right byte

	if avl_DEBUG {
		fmt.Printf("Rotating left on node: %vi\n", pNode.key)
	}

	// Do I have a parent, if so am I the left or right child?
	parent_parent = pNode.pParent
	if parent_parent != nil {
		if parent_parent.pLeft == pNode {
			was_i_left_or_right = 'L'
		} else {
			was_i_left_or_right = 'R'
		}
	}

	// floater is the left child of the right child
	pLocRight = pNode.pRight
	floater = pLocRight.pLeft

	if parent_parent != nil {
		if was_i_left_or_right == 'L' {
			parent_parent.pLeft = pLocRight
		} else {
			parent_parent.pRight = pLocRight
		}
	} else {
		pHead = pLocRight
	}
	pLocRight.pParent = parent_parent

	pLocRight.pLeft = pNode
	pNode.pParent = pLocRight

	pNode.pRight = floater
	if floater != nil {
		floater.pParent = pNode
	}

	return pHead
}

func abs_int(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// starting with this node, continue up the tree finding an abs(BF) > 1
func balance_node[K Key](pHead *avlNode[K], pNode *avlNode[K]) *avlNode[K] {
	var myBF int = calc_bf(pNode)
	if avl_DEBUG {
		fmt.Printf("BF for node %vi = %d\n", pNode.key, myBF)
	}

	if abs_int(myBF) > 1 {
		if myBF < -1 {
			if calc_bf(pNode.pLeft) <= 0 {
				pHead = rotate_right(pHead, pNode)
			} else {
				pHead = rotate_left(pHead, pNode.pLeft)
				pHead = rotate_right(pHead, pNode)
			}
		} else {
			if calc_bf(pNode.pRight) >= 0 {
				pHead = rotate_left(pHead, pNode)
			} else {
				pHead = rotate_right(pHead, pNode.pRight)
				pHead = rotate_left(pHead, pNode)
			}
		}
	} else if pNode.pParent != nil {
		pHead = balance_node(pHead, pNode.pParent)
	}
	return pHead
}

// returns the number of added nodes
func add_node_rec[K Key](ppHead **avlNode[K], pNode *avlNode[K], pKey K, pPayload any) int {
	var pHead *avlNode[K] = *ppHead
	var myCount int

	if pNode == nil {
		pNode = pHead
	}

	if pHead == nil {
		// If no head then create it
		pHead = new(avlNode[K])
		pHead.key = pKey
		pHead.payload = pPayload
		pHead.pLeft = nil
		pHead.pRight = nil
		pHead.pParent = nil
		*ppHead = pHead
		return 1
	} else if pKey == pNode.key {
		// If equal than no add
		return 0
	} else if pKey < pNode.key {
		if pNode.pLeft != nil {
			myCount = add_node_rec(&pHead, pNode.pLeft, pKey, pPayload)
			*ppHead = pHead
			return myCount
		} else {
			var node *avlNode[K] = new(avlNode[K])
			node.key = pKey
			node.payload = pPayload
			node.pLeft = nil
			node.pRight = nil
			node.pParent = pNode
			pNode.pLeft = node
			pHead = balance_node(pHead, node)
			*ppHead = pHead
			return 1
		}
	} else {
		// pKey > pNode->key
		if pNode.pRight != nil {
			myCount = add_node_rec(&pHead, pNode.pRight, pKey, pPayload)
			*ppHead = pHead
			return myCount
		} else {
			var node *avlNode[K] = new(avlNode[K])
			node.key = pKey
			node.payload = pPayload
			node.pLeft = nil
			node.pRight = nil
			node.pParent = pNode
			pNode.pRight = node
			pHead = balance_node(pHead, node)
			*ppHead = pHead
			return 1
		}
	}
}

// returns pointer to the last node being affected
func delete_node_rec[K Key](ppHead **avlNode[K], key K, pNode *avlNode[K]) *avlNode[K] {
	var pHead *avlNode[K] = *ppHead
	var pTmpCell *avlNode[K]

	// Dead end and key not found
	if pNode == nil {
		return nil
	}

	if key < pNode.key {
		// If key is less then this node, traverse left
		pNode.pLeft = delete_node_rec(&pHead, key, pNode.pLeft)
	} else if key > pNode.key {
		// If key is greater then this node, traverse right
		pNode.pRight = delete_node_rec(&pHead, key, pNode.pRight)
	} else if pNode.pLeft != nil && pNode.pRight != nil {
		// Key found and both children exist, go find minimal node on right
		// and replace this node (the one being deleted) with that one.
		// This can be recursive.
		pTmpCell = find_min(pNode.pRight)
		pNode.key = pTmpCell.key
		pNode.payload = pTmpCell.payload
		pNode.pRight = delete_node_rec(&pHead, pNode.key, pNode.pRight)
	} else {
		// Key found but at least one of the children is missing.
		pTmpCell = pNode
		if pNode.pLeft == nil && pNode.pRight == nil && pNode == pHead {
			// This was the last node in the tree
			pHead = nil
			*ppHead = pHead
		}
		if pNode.pLeft == nil {
			//
			if pNode.pRight != nil { // reset parent
				pNode.pRight.pParent = pNode.pParent
			}
			if pNode == pHead {
				pHead = pNode.pRight
				*ppHead = pHead
			}
			pNode = pNode.pRight
		} else if pNode.pRight == nil {
			if pNode.pLeft != nil { // reset parent
				pNode.pLeft.pParent = pNode.pParent
			}
			if pNode == pHead {
				pHead = pNode.pLeft
				*ppHead = pHead
			}
			pNode = pNode.pLeft
		}
		pTmpCell = nil
	}
	return pNode
}

func count_nodes_rec[K Key](pNode *avlNode[K], pCount *int) {
	if pNode != nil {
		*pCount = *pCount + 1
	}
	if pNode.pLeft != nil {
		count_nodes_rec(pNode.pLeft, pCount)
	}
	if pNode.pRight != nil {
		count_nodes_rec(pNode.pRight, pCount)
	}
}
