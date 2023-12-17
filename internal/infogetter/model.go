package infogetter

type Response struct {
	Result     string     `json:"result"`
	Code       int        `json:"code"`
	PlayerInfo PlayerInfo `json:"data"`
}

type PVE struct {
	GamePlayed     int            `json:"gamePlayed"`
	AttackLevel    int            `json:"unlimPve_playerAttackLevel"`
	DefenceLevel   int            `json:"unlimPve_playerDefenceLevel"`
	MissionLevels  map[string]int `json:"unlimPve_missionLevels"`
	WavePveMaxWave int            `json:"wavePve_maxWave"`
}

type PVP struct {
	GamePlayed       int     `json:"gamePlayed"`
	GameWin          int     `json:"gameWin"`
	TotalAssists     int     `json:"totalAssists"`
	TotalBattleTime  int64   `json:"totalBattleTime"`
	TotalDeath       int     `json:"totalDeath"`
	TotalDmgDone     float64 `json:"totalDmgDone"`
	TotalHealingDone float64 `json:"totalHealingDone"`
	TotalKill        int     `json:"totalKill"`
	TotalVpDmgDone   float64 `json:"totalVpDmgDone"`
}

type COOP struct {
	GameWin         int `json:"gameWin"`
	GamePlayed      int `json:"gamePlayed"`
	TotalBattleTime int `json:"totalBattleTime"`
}

type OpenWorld struct {
	Karma int `json:"karma"`
}

type Clan struct {
	Name      string `json:"name"`
	Tag       string `json:"tag"`
	Cid       int    `json:"cid"`
	PvpRating int    `json:"pvpRating"`
	PveRating int    `json:"pveRating"`
}

type PlayerInfo struct {
	EffRating     float64   `json:"effRating"`
	NickName      string    `json:"nickName"`
	PrestigeBonus float64   `json:"prestigeBonus"`
	UID           int       `json:"uid"`
	AccountRank   int       `json:"accountRank"`
	PVE           PVE       `json:"pve"`
	PVP           PVP       `json:"pvp"`
	COOP          COOP      `json:"coop"`
	OpenWorld     OpenWorld `json:"openWorld"`
	Clan          Clan      `json:"clan"`
}
