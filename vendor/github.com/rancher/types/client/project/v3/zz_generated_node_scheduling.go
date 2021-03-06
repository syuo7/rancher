package client

const (
	NodeSchedulingType            = "nodeScheduling"
	NodeSchedulingFieldNodeID     = "nodeId"
	NodeSchedulingFieldPreferred  = "preferred"
	NodeSchedulingFieldRequireAll = "requireAll"
	NodeSchedulingFieldRequireAny = "requireAny"
)

type NodeScheduling struct {
	NodeID     string   `json:"nodeId,omitempty" yaml:"nodeId,omitempty"`
	Preferred  []string `json:"preferred,omitempty" yaml:"preferred,omitempty"`
	RequireAll []string `json:"requireAll,omitempty" yaml:"requireAll,omitempty"`
	RequireAny []string `json:"requireAny,omitempty" yaml:"requireAny,omitempty"`
}
