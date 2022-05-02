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
	Id   int   `json:"id"`
	Rows []Row `json:"row"`
}

type Row struct {
	Cells []Cell `json:"cells"`
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
		ret.MakeSet(i)
	}

	return ret
}
func (d *DSU) MakeSet(v int) {
	d.parent[v] = v
	d.rank[v] = 0
	d.used[v] = true
}

func (d *DSU) FindSet(v int) int {
	if v == d.parent[v] {
		return v
	}

	d.used[d.parent[v]] = false
	d.parent[v] = d.FindSet(d.parent[v])

	return d.parent[v]
}

func (d *DSU) UnionSets(a, b int) {
	a = d.FindSet(a)
	b = d.FindSet(b)

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
