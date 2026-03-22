package ds

import "errors"

// DAG represents a Directed Acyclic Graph.
type DAG[K comparable, V any] struct {
	nodes map[K]V
	edges map[K][]K
}

// NewDAG creates a new DAG.
func NewDAG[K comparable, V any]() *DAG[K, V] {
	return &DAG[K, V]{
		nodes: make(map[K]V),
		edges: make(map[K][]K),
	}
}

// AddNode adds a node to the DAG.
func (d *DAG[K, V]) AddNode(id K, val V) {
	d.nodes[id] = val
}

// AddEdge adds a directed edge from one node to another.
func (d *DAG[K, V]) AddEdge(from, to K) error {
	if _, ok := d.nodes[from]; !ok {
		return errors.New("from node does not exist")
	}
	if _, ok := d.nodes[to]; !ok {
		return errors.New("to node does not exist")
	}
	d.edges[from] = append(d.edges[from], to)
	return nil
}

// TopologicalSort returns the nodes in topological order.
// Returns an error if a cycle is detected.
func (d *DAG[K, V]) TopologicalSort() ([]V, error) {
	inDegree := make(map[K]int)
	for id := range d.nodes {
		inDegree[id] = 0
	}
	for _, targets := range d.edges {
		for _, to := range targets {
			inDegree[to]++
		}
	}

	var queue []K
	for id, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, id)
		}
	}

	var result []V
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		result = append(result, d.nodes[u])

		for _, v := range d.edges[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if len(result) != len(d.nodes) {
		return nil, errors.New("cycle detected in graph")
	}

	return result, nil
}
