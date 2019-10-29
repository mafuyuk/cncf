package kamoc

import "encoding/json"

type status int

const (
	creating status = iota
	created
	running
	stopped
)

func (c status) String() string {
	switch c {
	case creating:
		return "creating"
	case created:
		return "created"
	case running:
		return "running"
	case stopped:
		return "stopped"
	default:
		return "unknown"
	}
}

type annotations struct {
	myKey string
}

type State struct {
	ociVersion  string      `json:"ociVersion"`
	id          string      `json:"id"`
	status      status      `json:"status"`
	pid         uint        `json:"pid"`
	bundle      string      `json:"bundle"`
	annotations annotations `json:"annotations, omitempty"`
}

func (s *State) toJson() (string, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
