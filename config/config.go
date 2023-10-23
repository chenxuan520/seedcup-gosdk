package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	GameMaxRound           int32  `json:"game_max_round,omitempty"`
	GamePrintMap           bool   `json:"game_print_map,omitempty"`
	GamePrintMapASCII      bool   `json:"game_print_map_ascii,omitempty"`
	MapSize                int32  `json:"map_size,omitempty"`
	PlayerNum              int32  `json:"player_num,omitempty"`
	PlayerHp               int32  `json:"player_hp,omitempty"`
	PlayerMaxHp            int32  `json:"player_max_hp,omitempty"`
	PlayerSpeed            int32  `json:"player_speed,omitempty"`
	BombTime               int32  `json:"bomb_time,omitempty"`
	BombNum                int32  `json:"bomb_num,omitempty"`
	BombRange              int32  `json:"bomb_range,omitempty"`
	BombRandom             int32  `json:"bomb_random,omitempty"`
	ShieldTime             int32  `json:"shield_time,omitempty"`
	InvincibleTime         int32  `json:"invincible_time,omitempty"`
	MarkKill               int32  `json:"mark_kill,omitempty"`
	MarkDead               int32  `json:"mark_dead,omitempty"`
	MarkPickPotion         int32  `json:"mark_pick_potion,omitempty"`
	MarkBombMud            int32  `json:"mark_bomb_mud,omitempty"`
	PotionProbability      int32  `json:"potion_probability,omitempty"`
	MudNum                 int32  `json:"mud_num,omitempty"`
	WallRandom             int32  `json:"wall_random,omitempty"`
	SeedRandom             int32  `json:"seed_random,omitempty"`
	ResultPath             string `json:"result_path,omitempty"`
	TimerInitialValue      int32  `json:"timer_initial_value,omitempty"`
	RoundIntervalValue     int32  `json:"round_interval_value,omitempty"`
	ServerMaxConnectionNum int32  `json:"server_max_connection_num,omitempty"`
	EpollMaxEventsNum      int32  `json:"epoll_max_events_num,omitempty"`
	EpollTimeout           int32  `json:"epoll_timeout,omitempty"`
	LogPrintStdout         bool   `json:"log_print_stdout,omitempty"`
	Host                   string `json:"host,omitempty"`
	Port                   uint32 `json:"port,omitempty"`
}

func InitConfig(configPath string) (*Config, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = json.Unmarshal(data, config)
	return config, nil
}
