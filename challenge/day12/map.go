package day12

type Coords struct {
	X int
	Y int
}

type Node struct {
	Val       rune
	StepCount int
	Children  []Node
	Coords
}

func (n *Node) ValStr() string {
	return string(n.Val)
}

type NodeGrid [][]*Node
type NodeMap map[Coords]*Node

func (nm NodeMap) GetNode(c Coords) *Node {
	return nm[c]
}

type MapGrid struct {
	Grid       NodeGrid
	StartNode  Node
	TargetNode Node
	NodeMap
}

func NewMapGrid(lines []string) *MapGrid {
	mg := &MapGrid{}
	nodeGrid := make(NodeGrid, 0)
	nodeMap := make(NodeMap)
	for i, row := range lines {
		currRow := make([]*Node, 0)
		for j, colVal := range row {
			node := &Node{Val: colVal, Coords: Coords{X: j, Y: i}}
			if string(colVal) == "E" {
				mg.TargetNode = *node
			}
			if string(colVal) == "S" {
				mg.StartNode = *node
			}
			nodeMap[node.Coords] = node
			currRow = append(currRow, node)
		}
		nodeGrid = append(nodeGrid, currRow)
	}
	mg.Grid = nodeGrid
	mg.NodeMap = nodeMap
	return mg
}

type Simulation struct {
	MapGrid
	VistedNodes    []*Node
	UnvisitedNodes []*Node
}

func NewSimulation(mg MapGrid) *Simulation {
	sim := &Simulation{MapGrid: mg}
	for _, n := range sim.MapGrid.NodeMap {
		sim.UnvisitedNodes = append(sim.UnvisitedNodes, n)
	}
	return sim
}
