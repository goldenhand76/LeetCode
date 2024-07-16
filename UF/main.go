package main

func main() {
	uf := NewQuickFindUF(10)
	uf.Union(0, 1)
	uf.Union(2, 4)
	uf.connected(0, 4)
}
