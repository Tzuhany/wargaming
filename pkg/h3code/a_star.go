package h3code

import (
	"container/heap"

	"github.com/uber/h3-go/v4"
)

type Tile struct {
	G, H       int     // 实际成本和预估成本
	Position   h3.Cell // 当前位置
	Adjacent   []*Tile // 相邻网格列表
	IsObstacle bool    // 是否是障碍物
	Index      int     // 当前网格索引
	Prev       *Tile   // 上一个节点
}

// F 总成本
func (t *Tile) F() int {
	return t.G + t.H
}

type AStarPathfinding struct{}

// TileHeap 构造 Tile 堆，按照 F 排序
type TileHeap []*Tile

func (h TileHeap) Len() int            { return len(h) }
func (h TileHeap) Less(i, j int) bool  { return h[i].F() < h[j].F() }
func (h TileHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *TileHeap) Push(x interface{}) { *h = append(*h, x.(*Tile)) }
func (h *TileHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func (a *AStarPathfinding) FindPath(startTile, endTile *Tile) []*Tile {
	openSet := make(TileHeap, 0)
	closedSet := make(map[*Tile]bool)

	// 起始网格
	startTile.G = 0
	startTile.H = a.getEstimatedPathCost(&startTile.Position, &endTile.Position)

	openSet = append(openSet, startTile)

	// 使用堆 按照 F 值最小排序
	heap.Init(&openSet)

	for len(openSet) != 0 {
		// 取出 F 值最小的进行处理
		currentTile := heap.Pop(&openSet).(*Tile)
		closedSet[currentTile] = true

		// 如果当前节点为最终节点，则返回最终结果
		if currentTile.Position == endTile.Position {
			return a.reconstructPath(startTile, currentTile)
		}

		// 遍历所有邻居节点
		for _, adjacentTile := range currentTile.Adjacent {
			// 如果当前节点在关闭列表中，或者当前节点为障碍物，则跳过
			if adjacentTile.IsObstacle || closedSet[adjacentTile] {
				continue
			}

			// 计算实际代价
			tentativeG := currentTile.G + 1

			if tentativeG < adjacentTile.G || !a.inHeap(&openSet, adjacentTile) {
				adjacentTile.G = tentativeG
				adjacentTile.H = a.getEstimatedPathCost(&adjacentTile.Position, &endTile.Position)
				adjacentTile.Prev = currentTile

				if !a.inHeap(&openSet, adjacentTile) {
					heap.Push(&openSet, adjacentTile)
				} else {
					heap.Fix(&openSet, adjacentTile.Index)
				}
			}
		}
	}
	return nil
}

// 获取两个单元格之间的大圆距离
func (a *AStarPathfinding) getEstimatedPathCost(startPos, endPos *h3.Cell) int {
	return len(startPos.GridPath(*endPos))
}

// 得到结果，回溯路径
func (a *AStarPathfinding) reconstructPath(startPoint, endPoint *Tile) []*Tile {
	var path []*Tile
	currentTile := endPoint

	for currentTile != startPoint {
		path = append([]*Tile{currentTile}, path...)
		currentTile = currentTile.Prev
	}

	return append([]*Tile{startPoint}, path...)
}

func (a *AStarPathfinding) inHeap(h *TileHeap, tile *Tile) bool {
	for _, t := range *h {
		if t == tile {
			return true
		}
	}
	return false
}

func InitTiles(cells, obstacles []h3.Cell) []*Tile {

	// 创建障碍物 map
	obstacleMap := make(map[h3.Cell]struct{})
	for _, obstacle := range obstacles {
		obstacleMap[obstacle] = struct{}{}
	}

	// 创建 cell map，用于过滤边界 cell 的相邻网格
	cellMap := make(map[h3.Cell]struct{})
	for _, cell := range cells {
		cellMap[cell] = struct{}{}
	}

	tileMap := make(map[h3.Cell]*Tile)

	var tiles []*Tile
	for _, cell := range cells {
		if _, exits := tileMap[cell]; !exits {
			tile := &Tile{
				Position: cell,
				Adjacent: make([]*Tile, 0),
			}

			tileMap[cell] = tile
		}

		currentTile := tileMap[cell]

		directedEdges := cell.DirectedEdges()

		for _, edge := range directedEdges {

			if edge.IsValid() {
				neighborCell := edge.Destination()

				if _, exit := cellMap[neighborCell]; !exit {
					continue
				}

				// 检查相邻单元格是否为障碍物
				isObstacle := false
				if _, exists := obstacleMap[neighborCell]; exists {
					isObstacle = true
				}

				if _, exists := tileMap[neighborCell]; !exists {
					neighborTile := &Tile{
						Position:   neighborCell,
						Adjacent:   make([]*Tile, 0),
						IsObstacle: isObstacle,
					}

					tileMap[neighborCell] = neighborTile
				}

				currentTile.Adjacent = append(currentTile.Adjacent, tileMap[neighborCell])
			}

		}
	}

	for _, tile := range tileMap {
		tiles = append(tiles, tile)
	}

	return tiles
}
