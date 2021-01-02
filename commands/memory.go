package commands

// Memory usage representation.
type Memory struct {
	Used  uint64 `json:"used,omitempty"`
	Total uint64 `json:"total,omitempty"`
}
