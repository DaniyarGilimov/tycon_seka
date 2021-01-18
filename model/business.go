package model

import (
	"time"
)

// Business struct
type Business struct {
	ID               int       `json:"id"`
	UserToken        string    `json:"-" bson:"userToken"`
	Name             string    `json:"name"`
	Sector           *Sector   `json:"sector"`
	Location         *Location `json:"location"`
	RegistrationTime time.Time `json:"registrationTime" bson:"registrationTime"`

	// Business affect
	Level       int              `json:"level"`
	PRLevel     *LevelInfoPR     `json:"prLevel" bson:"prLevel"`
	TechLevel   *LevelInfoTech   `json:"techLevel" bson:"techLevel"`
	SafeLevel   *LevelInfoSafe   `json:"safeLevel" bson:"safeLevel"`
	TestInfo    *TestInfo        `json:"testInfo" bson:"testInfo"`
	BattleLevel *LevelInfoBattle `json:"battleLevel" bson:"battleLevel"`

	// Workers
	CEO *Worker `json:"ceo"`
	CTO *Worker `json:"cto"`
	PR  *Worker `json:"pr"`

	// Runtime generated
	Chips      int  `json:"chips" bson:"-"`
	Golds      int  `json:"golds" bson:"-"`
	Upgradable bool `json:"upgradable" bson:"-"`
}

// Price is chips and golds
type Price struct {
	Chips int `json:"chips" bson:"chips"`
	Golds int `json:"golds" bson:"golds"`
}

// TestInfo struct
type TestInfo struct {
	TestAvailable bool      `json:"testAvailable" bson:"-"`
	TestLeftTime  []int     `json:"testLeftTime" bson:"-"`
	LastTestPass  time.Time `json:"-" bson:"lastTestPass"`
	Decrimented   bool      `json:"-" bson:"decrimented"`
}

// LevelInfoBattle struct
type LevelInfoBattle struct {
	Level       int `json:"level" bson:"level"`
	MaxLevel    int `json:"maxLevel" bson:"maxLevel"`
	BattlePrice int `json:"battlePrice" bson:"battlePrice"`
}

//LevelInfoPR struct
type LevelInfoPR struct {
	Level          int          `json:"level" bson:"level"`
	MaxLevel       int          `json:"maxLevel" bson:"maxLevel"`
	UpgardeLevel   int          `json:"-" bson:"upgradeLevel"`
	LeftSeconds    []int        `json:"leftSeconds" bson:"leftSeconds"`
	SpeedIncrement int          `json:"-" bson:"speedIncrement"`
	EndTime        time.Time    `json:"-" bson:"endTime"`
	PRMG           *Incrementer `json:"-" bson:"pr_mg"`
	PRTV           *Incrementer `json:"-" bson:"pr_tv"`
	PRSM           *Incrementer `json:"-" bson:"pr_sm"`
}

//LevelInfoTech struct
type LevelInfoTech struct {
	Level          int          `json:"level" bson:"level"`
	MaxLevel       int          `json:"maxLevel" bson:"maxLevel"`
	UpgardeLevel   int          `json:"-" bson:"upgradeLevel"`
	LeftSeconds    []int        `json:"leftSeconds" bson:"leftSeconds"`
	SpeedIncrement int          `json:"-" bson:"speedIncrement"`
	EndTime        time.Time    `json:"-" bson:"endTime"`
	TEPR           *Incrementer `json:"-" bson:"te_pr"`
	TERE           *Incrementer `json:"-" bson:"te_re"`
	TERO           *Incrementer `json:"-" bson:"te_ro"`
}

// LevelInfoSafe struct
type LevelInfoSafe struct {
	Level           int       `json:"level" bson:"level"`
	MaxLevel        int       `json:"maxLevel" bson:"maxLevel"`
	UpgradePrice    *Price    `json:"upgradePrice" bson:"upgradePrice"`
	LeftSeconds     []int     `json:"leftSeconds" bson:"leftSeconds"`
	UpgradeTakeTime int       `json:"upgradeTakeTime" bson:"upgradeTakeTime"`
	EndTime         time.Time `json:"-" bson:"endTime"`
	UpgradeCapacity int       `json:"upgradeCapacity" bson:"upgradeCapacity"`

	// Current Data
	CurrentCapacity int     `json:"currentCapacity" bson:"currentCapacity"`
	CurrentVelocity float64 `json:"currentVelocity" bson:"currentVelocity"`
	Sum             int     `json:"sum" bson:"-"`
	//LeftTime        []int   `json:"-" bson:"-"`

	// New version
	SafeParts []*SafePart `json:"-" bson:"safeParts"`
}

// SafePart struct
type SafePart struct {
	Velocity  float64   `json:"velocity" bson:"velocity"`
	StartTime time.Time `json:"startTime" bson:"startTime"`
	Collect   int       `json:"collect" bson:"collect"`
	EndTime   time.Time `json:"endTime" bson:"endTime"`
}

// BusinessRegist struct
type BusinessRegist struct {
	Name     string `json:"name"`
	Sector   int    `json:"sector"`
	Location int    `json:"location"`
	CEO      int    `json:"ceo"`
	PR       int    `json:"pr"`
	CTO      int    `json:"cto"`
}

// Sector struct
type Sector struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Population int    `json:"population"`
}

// Location struct
type Location struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Population int    `json:"population"`
}

// SecLoc struct
type SecLoc struct {
	Sectors   []*Sector   `json:"sectors"`
	Locations []*Location `json:"locations"`
	Price     int         `json:"price"`
}

// PaTPage struct
type PaTPage struct {
	Level int          `json:"level"`
	Items []*LevelItem `json:"items"`
}

// UpgradeRegist is request data while upgrading
type UpgradeRegist struct {
	ID        string `json:"id"`
	PayMethod string `json:"payMethod"`
}

// LevelItem struct
type LevelItem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Level       int    `json:"level"`
	Price       *Price `json:"price"`
	Time        int    `json:"time"` //time to do refactoring in seconds
	IsActive    bool   `json:"isActive"`
}

// TyconStatistics struct
type TyconStatistics struct {
	ID     int           `json:"id" bson:"id"`
	Graph  *TyconGraph   `json:"graph" bson:"-"`
	Events []*TyconEvent `json:"events" bson:"events"`
}

// TyconGraph struct
type TyconGraph struct {
	XAxis []int     `json:"xAxis" bson:"xAxis"` //time axis
	YAxis []float64 `json:"yAxis" bson:"yAxis"` //earn velocity axis
}

// TyconEvent struct
type TyconEvent struct {
	EventTitle    string    `json:"eventTitle" bson:"eventTitle"`
	EventSubTitle string    `json:"eventSubTitle" bson:"eventSubTitle"`
	Type          string    `json:"type" bson:"type"`
	EarnVelocity  float64   `json:"earnVelocity" bson:"earnVelocity"`
	Delta         float64   `json:"delta" bson:"delta"`
	Date          time.Time `json:"date" bson:"date"`
}

// Question struct
type Question struct {
	ID      int       `json:"id" bson:"id"`
	Title   string    `json:"title" bson:"title"`
	Answers []*Answer `json:"answers" bson:"answers"`
}

// Answer struct
type Answer struct {
	Text    string `json:"text" bson:"text"`
	Correct bool   `json:"correct" bson:"correct"`
}

// PostResult struct
type PostResult struct {
	CorrectCount int `json:"correctCount" bson:"-"`
}

// TestResponse struct
type TestResponse struct {
	CorrectCount        int     `json:"correctCount"`
	CurrentIncome       float64 `json:"currentIncome"`
	IncomeChangePercent int     `json:"incomeChangePercent"`
}
