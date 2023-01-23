package datastructures

import "testing"

func TestCompareBinaryTrees(t *testing.T) {
	type args struct {
		a *NodeTree
		b *NodeTree
	}
	trueTree := basicTree().root
	falseTree := basicTree().root
	falseTree.left.left.right.value = 33
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true",
			args: args{
				a: trueTree,
				b: trueTree,
			},
			want: true,
		},
		{
			name: "should return false",
			args: args{
				a: trueTree,
				b: falseTree,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareBinaryTrees(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("CompareBinaryTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func basicTree() *Tree {
	return &Tree{
		root: &NodeTree{
			value: 10,
			right: &NodeTree{
				value: 8,
				left: &NodeTree{
					value: 1,
					left: &NodeTree{
						value: 20,
					},
				},
			},
			left: &NodeTree{
				value: 15,
				left: &NodeTree{
					value: 2,
					right: &NodeTree{
						value: 7,
					},
				},
			},
		},
	}
}
