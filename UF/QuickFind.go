// Quick Find Implementation
// Eager Approach

package main

type QuickFindUF struct {
	id []int
}

func NewQuickFindUF(n int) *QuickFindUF {
	uf := &QuickFindUF{
		id: []int{},
	}
	for i := range n {
		uf.id = append(uf.id, i)
	}
	return uf
}

func (uf *QuickFindUF) Union(p int, q int) {
	pid := uf.id[p]
	qid := uf.id[q]

	for i := range len(uf.id) {
		if uf.id[i] == pid {
			uf.id[i] = qid
		}
	}
}

func (uf *QuickFindUF) connected(p int, q int) bool {
	return uf.id[p] == uf.id[q]
}
