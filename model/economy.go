package model

//SafeLeveler struct
type SafeLeveler struct {
	BusinessLevel int    `json:"businessLevel" bson:"businessLevel"`
	LevelRange    []int  `json:"levelRange" bson:"levelRange"`
	CapacityRange []int  `json:"capacityRange" bson:"capacityRange"`
	UpgradeTime   int    `json:"upgradeTime" bson:"upgradeTime"`
	UpgradePrice  *Price `json:"upgradePrice" bson:"upgradePrice"`

	/// only in pipe requests
	Last int `json:"-" bson:"last"`
}

// PRLeveler struct
type PRLeveler struct {
	BusinessLevel int          `json:"businessLevel" bson:"businessLevel"`
	LevelRange    []int        `json:"levelRange" bson:"levelRange"`
	PRMG          *Incrementer `json:"pr_mg" bson:"pr_mg"`
	PRTV          *Incrementer `json:"pr_tv" bson:"pr_tv"`
	PRSM          *Incrementer `json:"pr_sm" bson:"pr_sm"`
}

// Incrementer struct
type Incrementer struct {
	UpgradeTime    int    `json:"upgradeTime" bson:"upgradeTime"`
	LevelIncrement int    `json:"levelIncrement" bson:"levelIncrement"`
	SpeedIncrement int    `json:"speedIncrement" bson:"speedIncrement"`
	UpgradePrice   *Price `json:"upgradePrice" bson:"upgradePrice"`
}

// TechLeveler struct
type TechLeveler struct {
	BusinessLevel int          `json:"businessLevel" bson:"businessLevel"`
	LevelRange    []int        `json:"levelRange" bson:"levelRange"`
	TEPR          *Incrementer `json:"te_pr" bson:"te_pr"`
	TERE          *Incrementer `json:"te_re" bson:"te_re"`
	TERO          *Incrementer `json:"te_ro" bson:"te_ro"`
}
