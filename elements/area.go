package elements

import "encoding/json"

type ObjType uint8

const (
	ObjPlayer ObjType = 1
	ObjBomb   ObjType = 2
	ObjBlock  ObjType = 3
	ObjItem   ObjType = 4
)

type Obj struct {
	Type     ObjType         `json:"type,omitempty"`
	Property json.RawMessage `json:"property,omitempty"`
}

type Area struct {
	X            int32  `json:"x,omitempty"`
	Y            int32  `json:"y,omitempty"`
	LastBombTime int32  `json:"last_bomb_time,omitempty"`
	Objs         []*Obj `json:"objs,omitempty"`
}
