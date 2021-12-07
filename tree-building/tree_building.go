package tree

import (
	"fmt"
	"sort"
)

// Define the Record type
type Record struct {
	ID     int
	Parent int
}

// Define the Node type
type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	sort.Slice(records, func(cur, prev int) bool {
		return records[cur].ID < records[prev].ID
	})

	nodes := map[int]*Node{}
	for i, r := range records {
		if i != r.ID || r.ID < r.Parent || (r.ID > 0 && r.ID == r.Parent) {
			return nil, fmt.Errorf("bad record %#v", r)
		}

		nodes[r.ID] = &Node{ID: r.ID}
		if r.ID != 0 {
			par := nodes[r.Parent]
			par.Children = append(par.Children, nodes[r.ID])
		}
	}
	return nodes[0], nil
}
