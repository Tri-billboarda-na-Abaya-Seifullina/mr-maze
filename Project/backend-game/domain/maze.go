package domain

const (
	GENERATING = "GENERATING"
	READING    = "READING"
)

type Maze struct {
	Id   int      `json:"id"`
	Rows []string `json:"rows"`
}

type Map struct {
	Id    int      `json:"id"`
	Cells [][]Cell `json:"map"`
}

type Cell struct {
	Up    bool `json:"up"`
	Right bool `json:"right"`
	Down  bool `json:"down"`
	Left  bool `json:"left"`
}

type DSU struct {
	Size   int
	parent []int
	rank   []int
}

func NewDSU(size int) *DSU {
	ret := &DSU{
		Size:   size,
		parent: make([]int, size),
		rank:   make([]int, size),
	}

	for i := 0; i < size; i++ {
		ret.MakeSet(i)
	}

	return ret
}
func (d *DSU) MakeSet(v int) {
	d.parent[v] = v
	d.rank[v] = 0
}

func (d *DSU) FindSet(v int) int {
	if v == d.parent[v] {
		return v
	}

	d.parent[v] = d.FindSet(d.parent[v])

	return d.parent[v]
}

func (d *DSU) UnionSets(a, b int) {
	a = d.FindSet(a)
	b = d.FindSet(b)

	if a != b {
		if d.rank[a] < d.rank[b] {
			a, b = b, a
		}
		d.parent[b] = a
		if d.rank[a] == d.rank[b] {
			d.rank[a]++
		}
	}
}
