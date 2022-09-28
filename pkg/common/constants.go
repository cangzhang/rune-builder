package common

import _ "embed"

var TrinketItems = []string{
	"3340",
	"3363",
	"3364",
}

var ConsumableItems = []string{
	"2138",
	"2139",
	"2140",
	//"2033",
	//"2031",
}

var WardItems = []string{"2055"}

var Fragments = [][]int{
	{5008, 5005, 5007},
	{5008, 5002, 5003},
	{5001, 5002, 5003},
}

//go:embed tpl/package.json
var PkgJsonTpl string

var Positions = []string{"top", "jungle", "middle", "adc", "support"}

var StatePerks = []RespRuneItem{
	{
		Id:        5001,
		Name:      "HealthScaling",
		Key:       "HealthScaling",
		ShortDesc: "+15-90 Health (based on level)",
		LongDesc:  "+15-90 Health (based on level)",
		Icon:      "/lol-game-data/assets/v1/perk-images/StatMods/StatModsHealthScalingIcon.png",
	},
	{
		Id:        5002,
		Name:      "Armor",
		Key:       "Armor",
		ShortDesc: "+6 Armor",
		LongDesc:  "+6 Armor",
		Icon:      "/lol-game-data/assets/v1/perk-images/StatMods/StatModsArmorIcon.png",
	},
	{
		Id:        5003,
		Name:      "MagicResist",
		Key:       "MagicRes",
		ShortDesc: "+8 Magic Resist",
		LongDesc:  "+8 Magic Resist",
		Icon:      "/lol-game-data/assets/v1/perk-images/StatMods/StatModsMagicResIcon.png",
	},
	{
		Id:        5005,
		Name:      "AttackSpeed",
		Key:       "AttackSpeed",
		ShortDesc: "+10% Attack Speed",
		LongDesc:  "+10% Attack Speed",
		Icon:      "/lol-game-data/assets/v1/perk-images/StatMods/StatModsAttackSpeedIcon.png",
	},
	{
		Id:        5007,
		Name:      "CDRScaling",
		Key:       "CDRScaling",
		ShortDesc: "+1-10% <lol-uikit-tooltipped-keyword key='LinkTooltip_Description_CDR'>CDR</lol-uikit-tooltipped-keyword> (based on level)",
		LongDesc:  "+1-10% <lol-uikit-tooltipped-keyword key='LinkTooltip_Description_CDR'>CDR</lol-uikit-tooltipped-keyword> (based on level)",
		Icon:      "/lol-game-data/assets/v1/perk-images/StatMods/StatModsCDRScalingIcon.png",
	},
	{
		Id:        5008,
		Name:      "Adaptive",
		Key:       "Adaptive",
		ShortDesc: "+9 <lol-uikit-tooltipped-keyword key='LinkTooltip_Description_Adaptive'><font color='#48C4B7'>Adaptive Force</font></lol-uikit-tooltipped-keyword>",
		LongDesc:  "+9 <lol-uikit-tooltipped-keyword key='LinkTooltip_Description_Adaptive'><font color='#48C4B7'>Adaptive Force</font></lol-uikit-tooltipped-keyword>",
		Icon:      "/lol-game-data/assets/v1/perk-images/StatMods/StatModsAdaptiveForceIcon.png",
	},
}

type GameMode int

const (
	GameModeSR   GameMode = 0
	GameModeARAM          = 1
	GameModeURF           = 2
)
