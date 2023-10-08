package elements

type PacketType uint8

const (
	InitReq      PacketType = 1
	ActionReq    PacketType = 2
	ActionResp   PacketType = 3
	GameOverResp PacketType = 4
)

type RespPacket struct {
	Type PacketType `json:"type,omitempty"`
	Data RespData   `json:"data,omitempty"`
}

type RespData struct {
	PlayerID  int32    `json:"player_id,omitempty"`
	Round     int32    `json:"round,omitempty"`
	Map       []*Area  `json:"map,omitempty"`
	WinnerIds []int32  `json:"winner_ids,omitempty"`
	Scores    []Scores `json:"scores,omitempty"`
}

type Scores struct {
	PlayerID int32 `json:"player_id,omitempty"`
	Score    int32 `json:"score,omitempty"`
}

type GameMsg struct {
	PlayerID int32
	Round    int32
	GameMap  [][]Map
	Players  map[int32]*Player
	Blocks   map[int32]*Block
	Bombs    map[int32]*Bomb
}

type ActionType uint8

const (
	Silent    ActionType = 0
	MoveLeft  ActionType = 1
	MoveRight ActionType = 2
	MoveUp    ActionType = 3
	MoveDown  ActionType = 4
	Pleaced   ActionType = 5
)

type ReqAction struct {
	PlayerID   int32      `json:"playerID"`
	ActionType ActionType `json:"actionType"`
}

type ReqPacket struct {
	Type PacketType  `json:"type,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type ReqInit struct{}

func ParseRespData(packet *RespData) (*GameMsg, error) {
	msg := GameMsg{
		PlayerID: packet.PlayerID,
		Round:    0,
		GameMap:  globalMap,
		Players:  make(map[int32]*Player),
		Blocks:   make(map[int32]*Block),
		Bombs:    make(map[int32]*Bomb),
	}
	err := parseMap(packet.Map, &msg)
	return &msg, err
}
