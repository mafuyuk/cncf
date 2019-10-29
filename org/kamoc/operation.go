package kamoc

type containerID string

// @see https://github.com/opencontainers/runtime-spec/blob/master/runtime.md
type Operator interface {
	QueryState(containerID) (State, error)
	Create(containerID, pathToBundle string)
	Start(containerID) error
	Kill(containerID, signal string) error
	Delete(containerID) error
}
