package game

import (
	"fmt"
	"utils"
)

func LogList(fileName string, list []Node) {
	var listStr string
	listStr += fmt.Sprintf("List of 1000 first elements")

	for i, node := range list {
		if i > 1000 {
			break;
		}
		listStr += fmt.Sprintf("\n\n- Node %d: %d\n", i, node.Heuristic)
		for _, move := range node.Moves {
			listStr += fmt.Sprintf("(%d, %d) ", move.Y, move.X)
		}
	}

	utils.WriteStringToFile(fileName, listStr)
}

func LogNode(fileName string, node Node) {
	var solutionStr string
	solutionStr += fmt.Sprintf("Node\n")
	solutionStr += fmt.Sprintf("\n- Heuristic: %d ", node.Heuristic)
	solutionStr += fmt.Sprintf("\n- Final grid: %v ", node.Grid)
	solutionStr += fmt.Sprintf("\n- Moves %d: %v ", len(node.Moves), node.Moves)

	utils.WriteStringToFile(fileName, solutionStr)
}
