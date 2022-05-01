package domain

type Maze struct {
	Rows []string `json:"rows"`
}

type DSU struct {
	Size   int
	parent []int
	rank   []int
	used   []bool
}

func NewDSU(size int) *DSU {
	ret := &DSU{
		Size:   size,
		parent: make([]int, size),
		rank:   make([]int, size),
		used:   make([]bool, size),
	}

	for i := 0; i < size; i++ {
		ret.makeSet(i)
	}

	return ret
}
func (d *DSU) makeSet(v int) {
	d.parent[v] = v
	d.rank[v] = v
	d.used[v] = true
}

func (d *DSU) findSet(v int) int {
	if v == d.parent[v] {
		return v
	}

	d.used[d.parent[v]] = false
	d.parent[v] = d.findSet(d.parent[v])

	return d.parent[v]
}

func (d *DSU) unionSets(a, b int) {
	a = d.findSet(a)
	b = d.findSet(b)

	if a != b {
		if d.rank[a] > d.rank[b] {
			a, b = b, a
		}
		d.used[d.parent[b]] = false
		d.parent[b] = a
		if d.rank[a] == d.rank[b] {
			d.rank[a]++
		}
	}
}

func (d *DSU) GetFree() []int {
	ret := []int{}
	for i := 0; i < d.Size; i++ {
		if !d.used[i] {
			ret = append(ret, i)
		}
	}

	return ret
}
