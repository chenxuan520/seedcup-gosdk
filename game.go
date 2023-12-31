package seedcup

import (
	"github.com/chenxuan520/seedcup-gosdk/config"
	"github.com/chenxuan520/seedcup-gosdk/elements"
	"github.com/chenxuan520/seedcup-gosdk/server"
)

type Game struct {
	playerID         int32
	playerName       string
	conn             *server.Conn
	config           *config.Config
	mapCallBack      func(gameMsg *elements.GameMsg, game *Game) error
	gameOverCallBack func(playerID int32, winners []int32, scores []elements.Scores) error
}

func (game *Game) Init(configPath, playerName string) error {
	var err error
	game.config, err = config.InitConfig(configPath)
	if err != nil {
		return err
	}
	game.conn, err = server.CreateConn(game.config)
	if err != nil {
		return err
	}
	game.playerName = playerName
	elements.InitMap(game.config.MapSize)
	return nil
}

func (game *Game) RegisterCallBack(mapCallBack func(*elements.GameMsg, *Game) error, gameOverCallBack func(int32, []int32, []elements.Scores) error) {
	game.mapCallBack = mapCallBack
	game.gameOverCallBack = gameOverCallBack
}

func (game *Game) Run() error {
	defer game.conn.Close()
	err := game.conn.UpstreamInit(game.playerName)
	if err != nil {
		return err
	}
	for {
		packet, err := game.conn.RecvPacket()
		if err != nil {
			return err
		}
		switch packet.Type {
		case elements.ActionResp:
			result, err := elements.ParseRespData(&packet.Data)
			if err != nil {
				return err
			}
			game.playerID = result.PlayerID
			if game.mapCallBack != nil {
				err = game.mapCallBack(result, game)
			}
			if err != nil {
				return err
			}
		case elements.GameOverResp:
			if err != nil {
				return err
			}
			if game.gameOverCallBack != nil {
				err = game.gameOverCallBack(game.playerID, packet.Data.WinnerIds, packet.Data.Scores)
			}
			return err
		}
	}
}

func (game *Game) TakeAction(action elements.ActionType) error {
	return game.conn.UpstreamAction(game.playerID, action)
}

func (game *Game) TakeMutiAction(actions []elements.ActionType) error {
	return game.conn.UpstreamMutiAction(game.playerID, actions)
}
