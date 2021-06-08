package binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	examples = []struct {
		name     string
		root     *TreeNode
		target   int
		expected *TreeNode
	}{
		{
			name: "example 1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
			},
			target: 2,
			expected: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:  3,
					Left: nil,
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
		{
			name: "example 2",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			target: 3,
			expected: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  3,
					Left: nil,
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
				Right: nil,
			},
		},
		{
			name: "example 3",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   2,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: nil,
				},
				Right: nil,
			},
			target: 2,
			expected: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
	}
)

func Test_removeLeafNodes(t *testing.T) {
	for _, example := range examples {
		example := example
		t.Run(example.name, func(t *testing.T) {
			res := removeLeafNodes(example.root, example.target)
			assert.Equal(t, example.expected, res, "unexpected result")
		})
	}
}
