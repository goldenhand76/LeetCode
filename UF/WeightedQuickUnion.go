// weighted QU + path compression

package main

type WeightedQuickUnion struct {
	id []int
	sz []int
}

func NewWeightedQuickUnionUF(n int) *WeightedQuickUnion {
	uf := &WeightedQuickUnion{}
	for i := range n {
		uf.id = append(uf.id, i)
		uf.sz = append(uf.sz, 0)
	}
	return uf
}

func (uf *WeightedQuickUnion) root(i int) int {
	for uf.id[i] != i {
		uf.id[i] = uf.id[uf.id[i]] // Path Compression
		i = uf.id[i]
	}
	return i
}

func (uf *WeightedQuickUnion) union(p int, q int) {
	i := uf.root(p)
	j := uf.root(q)
	if i == j {
		return
	}
	if uf.sz[i] < uf.sz[j] {
		uf.id[i] = j
		uf.sz[j] += uf.sz[i]
	} else {
		uf.id[j] = i
		uf.sz[i] += uf.sz[j]
	}

}

func (weighted *WeightedQuickUnion) connected(p int, q int) bool {
	return weighted.root(p) == weighted.root(q)
}
