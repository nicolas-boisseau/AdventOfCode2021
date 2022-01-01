package day18

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Node_String(t *testing.T) {

	n := Node{
		isLeaf: false,
		left: &Node{
			isLeaf: true,
			value:  4,
		},
		right: &Node{
			isLeaf: true,
			value:  6,
		},
	}

	assert.Equal(t, "[4,6]", n.String())
}

func Test_Node_String_Complex(t *testing.T) {

	n := Node{
		isLeaf: false,
		left: &Node{
			isLeaf: true,
			value:  4,
		},
		right: &Node{
			isLeaf: false,
			left: &Node{
				isLeaf: false,
				left: &Node{
					isLeaf: true,
					value:  45,
				},
				right: &Node{
					isLeaf: true,
					value:  17,
				},
			},
			right: &Node{
				isLeaf: true,
				value:  22,
			},
		},
	}

	assert.Equal(t, "[4,[[45,17],22]]", n.String())
}

func Test_ReadNode(t *testing.T) {

	sample := "[4,22]"

	n := ReadNode(sample)

	assert.Equal(t, sample, n.String())
}

func Test_ReadNode2(t *testing.T) {

	sample := "[4,[12,42]]"

	n := ReadNode(sample)

	fmt.Println(n)

	assert.Equal(t, sample, n.String())
}

func Test_ReadNode3(t *testing.T) {

	sample := "[[34,52],[12,42]]"

	n := ReadNode(sample)

	fmt.Println(n)

	assert.Equal(t, sample, n.String())
}

func Test_ReadNode4(t *testing.T) {

	sample := "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]"

	n := ReadNode(sample)

	fmt.Println(n)

	assert.Equal(t, sample, n.String())
}

func Test_Add_Nodes(t *testing.T) {
	n1Str := "[1,2]"
	n2Str := "[[3,4],5]"
	n1 := ReadNode(n1Str)
	n2 := ReadNode(n2Str)

	n3 := Add(n1, n2)

	fmt.Println(n3)
	assert.Equal(t, "[[1,2],[[3,4],5]]", n3.String())
}

func Test_Explode_Node(t *testing.T) {
	n := ReadNode("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]")

	fmt.Println("Before reduce:", n)
	changed := n.Explode(1)
	fmt.Println("After reduce :", n)

	assert.True(t, changed)
	assert.Equal(t, "[[[[0,7],4],[15,[0,13]]],[1,1]]", n.String())

	// Try reduce again and verify that nothing happens
	changed = n.Explode(1)
	assert.False(t, changed)
	assert.Equal(t, "[[[[0,7],4],[15,[0,13]]],[1,1]]", n.String())
}

func Test_Explode_Node2(t *testing.T) {
	n := ReadNode("[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]")

	fmt.Println("Before reduce:", n)
	changed := n.Explode(1)
	fmt.Println("After reduce :", n)

	assert.True(t, changed)
	assert.Equal(t, "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", n.String())

	// Try reduce again and verify that nothing happens
	changed = n.Explode(1)
	assert.False(t, changed)
}

func Test_SplitNode(t *testing.T) {
	n := ReadNode("[[[[0,7],4],[15,[0,13]]],[1,1]]")

	fmt.Println("Before split:", n)
	changed := n.Split()
	fmt.Println("After split :", n)

	assert.True(t, changed)
	assert.Equal(t, "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]", n.String())

	// Try split again and verify that nothing happens
	changed = n.Split()
	assert.False(t, changed)
}

func Test_AddCmplexNodes(t *testing.T) {
	n1Str := "[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]"
	n2Str := "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]"
	n1 := ReadNode(n1Str)
	n2 := ReadNode(n2Str)

	n3 := Add(n1, n2)

	fmt.Println(n3)
	assert.Equal(t, "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]", n3.String())
}
