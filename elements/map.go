package elements

var globalMap [][]Map

type Map struct {
	BlockID  int32
	BombID   int32
	ItemType ItemType
	Players  []int32
}

func (this *Map) Clear() {
	this.BombID = -1
	this.BlockID = -1
	this.ItemType = ItemNone
	this.Players = []int32{}
}

func InitMap(mapSize int32) {
	for i := 0; i < int(mapSize); i++ {
		globalMap = append(globalMap, make([]Map, mapSize))
	}
}

func parseMap(areaMap []*Area, msg *GameMsg) error {
	for _, area := range areaMap {
		x, y := area.X, area.Y
		globalMap[x][y].Clear()
		for _, obj := range area.Objs {
			switch obj.Type {
			case ObjPlayer:
				player, err := parsePlayer(obj)
				if err != nil {
					return err
				}
				player.PosX, player.PoxY = x, y
				msg.Players[player.PlayerID] = player
				msg.GameMap[x][y].Players = append(msg.GameMap[x][y].Players, player.PlayerID)
			case ObjBlock:
				block, err := parseBlock(obj)
				if err != nil {
					return err
				}
				block.PosX, block.PoxY = x, y
				msg.Blocks[block.BlockID] = block
				msg.GameMap[x][y].BlockID = block.BlockID
			case ObjBomb:
				bomb, err := parseBomb(obj)
				if err != nil {
					return err
				}
				bomb.PosX, bomb.PoxY = x, y
				msg.Bombs[bomb.BombID] = bomb
				msg.GameMap[x][y].BombID = bomb.BombID
			case ObjItem:
				itemType, err := parseItem(obj)
				msg.GameMap[x][y].ItemType = itemType
				if err != nil {
					return nil
				}
			}
		}
	}
	return nil
}
