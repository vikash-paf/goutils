package ds

import "testing"

func TestDAG(t *testing.T) {
	dag := NewDAG[string, string]()
	dag.AddNode("A", "Task A")
	dag.AddNode("B", "Task B")
	dag.AddNode("C", "Task C")

	dag.AddEdge("A", "B")
	dag.AddEdge("B", "C")
	dag.AddEdge("A", "C")

	order, err := dag.TopologicalSort()
	if err != nil {
		t.Fatalf("TopologicalSort failed: %v", err)
	}

	// Order should be Task A, Task B, Task C
	if order[0] != "Task A" || order[1] != "Task B" || order[2] != "Task C" {
		t.Errorf("Unexpected order: %v", order)
	}

	// Test cycle
	dag.AddEdge("C", "A")
	_, err = dag.TopologicalSort()
	if err == nil {
		t.Error("Expected error due to cycle, got nil")
	}
}
