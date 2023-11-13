package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chenxuan520/seedcup-gosdk"
	"github.com/chenxuan520/seedcup-gosdk/elements"
	"github.com/nsf/termbox-go"
)

func init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

func Getch() (rune, error) {
	eventQueue := make(chan termbox.Event, 200)

	for {
		eventQueue <- termbox.PollEvent()
		ev := <-eventQueue
		switch ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				return 0, nil
			} else if ev.Key == termbox.KeySpace {
				return ' ', nil
			} else if ev.Key == termbox.KeyCtrlC {
				termbox.Close()
				os.Exit(0)
			} else {
				return ev.Ch, nil
			}
		case termbox.EventError:
			return 0, ev.Err
		}
	}
}

func main() {
	var seedgame seedcup.Game
	defer termbox.Close()
	err := seedgame.Init("../config.json", "gosdk")
	if err != nil {
		log.Println("err:", err)
		return
	}

	fmt.Println("init seedgame success")

	isGameOver := false
	go func() {
		for !isGameOver {
			key, _ := Getch()
			actions := []elements.ActionType{}
			switch key {
			case 'a':
				actions = append(actions, elements.MoveLeft)
			case 's':
				actions = append(actions, elements.MoveDown)
			case 'd':
				actions = append(actions, elements.MoveRight)
			case 'w':
				actions = append(actions, elements.MoveUp)
			case ' ':
				actions = append(actions, elements.Silent)
			}
			seedgame.TakeMutiAction(actions)
		}
	}()

	seedgame.RegisterCallBack(func(msg *elements.GameMsg, game *seedcup.Game) error {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		termbox.Flush()
		fmt.Println()
		fmt.Println()
		size := len(msg.GameMap)
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				area := msg.GameMap[i][j]
				if area.BlockID != -1 {
					if msg.Blocks[area.BlockID].Removable {
						fmt.Printf("🟩")
					} else {
						fmt.Printf("🧱")
					}
				} else if area.ItemType != elements.ItemNone {
					switch area.ItemType {
					case elements.ItemBombNum:
						fmt.Printf("💊")
					case elements.ItemBombRange:
						fmt.Printf("🧪")
					case elements.ItemInvincible:
						fmt.Printf("🗽")
					case elements.ItemHP:
						fmt.Printf("💖")
					case elements.ItemShield:
						fmt.Printf("🔰")
					}
				} else if area.BombID != -1 {
					fmt.Printf("💣")
				} else if len(area.Players) != 0 {
					player := msg.Players[area.Players[0]]
					if player.InvincibleTime > 0 {
						fmt.Printf("👼")
					} else if player.ShieldTime > 0 {
						fmt.Printf("👤")
					} else {
						fmt.Printf("👤")
					}
				} else {
					fmt.Printf("◻️ ")
				}
			}
			fmt.Println("")
		}
		return nil
	}, func(playerID int32, winners []int32, _ []elements.Scores) error {
		isGameOver = true
		for _, winner := range winners {
			if playerID == winner {
				fmt.Println("You Win")
				return nil
			}
		}
		fmt.Println("You Lose")
		return nil
	})
	log.Println(seedgame.Run())
	return
}
