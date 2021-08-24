package main

import "fmt"

type AstNode struct {
	left  *AstNode
	mid   *AstNode
	right *AstNode
	op    int
	v     AstNodeValue
}

type AstNodeValue struct {
	intval int
	id     int
}

func makeAstNode(op int, left, mid, right *AstNode, intval, id int) *AstNode {
	n := AstNode{
		v: AstNodeValue{
			intval: intval,
			id:     id,
		},
		left:  left,
		mid:   mid,
		right: right,
		op:    op,
	}
	return &n
}

func makeLeaf(op int, intval, id int) *AstNode {
	return makeAstNode(op, nil, nil, nil, intval, id)
}

func makeAstUnary(op int, left *AstNode, intval, id int) *AstNode {
	return makeAstNode(op, left, nil, nil, intval, id)
}

func printAstDepth(n *AstNode) {
	fmt.Printf(" Height %d \n", getDepthStupid(n))

}
func getDepthStupid(n *AstNode) int {
	if n == nil {
		return 0
	}
	fmt.Printf("%v ", n)

	depth := 1
	depthLeft := getDepthStupid(n.left)
	depthRight := getDepthStupid(n.right)

	if depthLeft > depthRight {
		depth += depthLeft
	} else {
		depth += depthRight
	}

	return depth
}
