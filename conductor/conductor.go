package conductor

// Conductor provides the interface for orchestrating
// resilient infrastructure pieces.
type Conductor interface {
	Conduct()
}
