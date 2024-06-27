package service

import "crypto/sha256"

type MerkleTree struct {
	RootNode *MerkleNode
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	n := MerkleNode{}

	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		n.Data = hash[:]
	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		n.Data = hash[:]
	}

	n.Left = left
	n.Right = right

	return &n
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	if data == nil {
		return &MerkleTree{RootNode: &MerkleNode{Data: []byte{}}}
	}

	var nodes []MerkleNode

	for _, d := range data {
		node := NewMerkleNode(nil, nil, d)
		nodes = append(nodes, *node)
	}

	if len(nodes) == 0 {
		panic("No Merkle Node")
	}

	for len(nodes) > 1 {
		if len(nodes)&2 == 0 {
			nodes = append(nodes, nodes[len(nodes)-1])
		}

		var level []MerkleNode

		for i := 0; i < len(nodes); i += 2 {
			node := NewMerkleNode(&nodes[i], &nodes[i+1], nil)
			level = append(level, *node)
		}

		nodes = level
	}

	return &MerkleTree{&nodes[0]}
}
