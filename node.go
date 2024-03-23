package art

type NodeType int

const (
	typeNode4   NodeType = 1
	typeNode16  NodeType = 2
	typeNode48  NodeType = 3
	typeNode256 NodeType = 4
)

const (
	keyLengthOfNode4   = 4
	keyLengthOfNode16  = 16
	keyLengthOfNode48  = 256
	keyLengthOfNode256 = keyLengthOfNode48
)

type node struct {
}
