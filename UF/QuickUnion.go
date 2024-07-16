// QU without path compression

package main

type QuickUnionUF struct {
	id []int
}

func NewQuickUnionUF(n int) *QuickUnionUF {
	uf := &QuickUnionUF{id: []int{}}
	for i := range n {
		uf.id = append(uf.id, i)
	}
	return uf
}

func (uf QuickUnionUF) root(i int) int {
	for uf.id[i] != i {
		i = uf.id[i]
	}
	return i
}

func (uf QuickUnionUF) connected(p int, q int) bool {
	return uf.root(p) == uf.root(q)
}

func (uf QuickUnionUF) union(p int, q int) {
	uf.id[uf.root(p)] = uf.root(q)
}
