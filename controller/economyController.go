package controller

import (
	"general_game_azi/gutils"
	"log"
	"time"
	"tycon_seka/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//GetDefaultSafe method
func GetDefaultSafe(db *mgo.Database) (*model.LevelInfoSafe, error) {
	safeLevelerCollection := db.C(gutils.SAFELEVELER)

	var fsl *model.SafeLeveler
	err := safeLevelerCollection.Find(bson.M{"businessLevel": 1}).One(&fsl)
	if err != nil {
		return nil, err
	}

	log.Print(fsl.CapacityRange[0])
	fVel := float64(fsl.CapacityRange[0]) / 3600.0
	log.Print(fVel)
	maxTime := float64(fsl.CapacityRange[0]) / fVel
	log.Print(maxTime)

	// below LeftTime: []int{0, int(maxTime)}
	defSafe := &model.LevelInfoSafe{Level: 1, MaxLevel: fsl.LevelRange[1], UpgradePrice: &model.Price{Chips: fsl.UpgradePrice.Chips, Golds: fsl.UpgradePrice.Golds}, UpgradeTakeTime: fsl.UpgradeTime, UpgradeCapacity: fsl.CapacityRange[1], CurrentVelocity: fVel, CurrentCapacity: fsl.CapacityRange[0], Sum: fsl.CapacityRange[0], LeftSeconds: []int{0, 0}}
	defSafe.SafeParts = []*model.SafePart{&model.SafePart{Velocity: defSafe.CurrentVelocity, StartTime: time.Now(), EndTime: time.Now().Add(time.Duration(float64(defSafe.CurrentCapacity)/defSafe.CurrentVelocity) * time.Second)}}

	return defSafe, nil
}

// GetNextSafe method
func GetNextSafe(Business *model.Business, db *mgo.Database) (*model.LevelInfoSafe, error) {
	safeLevelerCollection := db.C(gutils.SAFELEVELER)

	sl := Business.SafeLevel
	nowLevel := sl.Level + 1

	var fsl *model.SafeLeveler
	pipe := []bson.M{
		bson.M{"$project": bson.M{"businessLevel": 1, "levelRange": 1, "capacityRange": 1, "upgradeTime": 1, "upgradePrice": 1, "last": bson.M{"$arrayElemAt": []interface{}{"$levelRange", -1}}}},
		bson.M{
			"$match": bson.M{
				"$or": []bson.M{
					bson.M{"levelRange": bson.M{"$gt": nowLevel, "$lt": nowLevel + 2}},
					bson.M{"last": nowLevel + 1},
				},
			},
		},
	}

	err := safeLevelerCollection.Pipe(pipe).One(&fsl)
	if err != nil {
		return nil, err
	}

	currentCapacity := sl.UpgradeCapacity
	maxLevel := fsl.LevelRange[1]
	upgradeCapacit := fsl.CapacityRange[1]

	if fsl.Last != nowLevel+1 { //Takes next, maximum in our set
		maxLevel = nowLevel
		upgradeCapacit = (fsl.CapacityRange[0] + fsl.CapacityRange[1]) / 2
		//maxLevel
	}

	//maxTime := float64(currentCapacity) / sl.CurrentVelocity
	//below LeftTime: []int{0, int(maxTime)}
	defSafe := &model.LevelInfoSafe{Level: nowLevel, MaxLevel: maxLevel, UpgradePrice: fsl.UpgradePrice, UpgradeTakeTime: fsl.UpgradeTime, UpgradeCapacity: upgradeCapacit, CurrentVelocity: sl.CurrentVelocity, CurrentCapacity: currentCapacity, Sum: 0, LeftSeconds: []int{0, 0}}
	defSafe.SafeParts = []*model.SafePart{&model.SafePart{Velocity: defSafe.CurrentVelocity, StartTime: time.Now(), EndTime: time.Now().Add(time.Duration(float64(defSafe.CurrentCapacity)/defSafe.CurrentVelocity) * time.Second)}}

	return defSafe, nil
}

// GetDefaultPR method
func GetDefaultPR(db *mgo.Database) (*model.LevelInfoPR, error) {
	prLevelerCollection := db.C(gutils.PRLEVELER)

	var fpl *model.PRLeveler
	err := prLevelerCollection.Find(bson.M{"businessLevel": 1}).One(&fpl)
	if err != nil {
		return nil, err
	}

	defaultLevel := &model.LevelInfoPR{Level: 1, MaxLevel: fpl.LevelRange[1], LeftSeconds: []int{0, 0}, PRMG: fpl.PRMG, PRTV: fpl.PRTV, PRSM: fpl.PRSM}

	return defaultLevel, nil
}

// GetDefaultTech method
func GetDefaultTech(db *mgo.Database) (*model.LevelInfoTech, error) {
	techLevelerCollection := db.C(gutils.TECHLEVELER)

	var ftl *model.TechLeveler
	err := techLevelerCollection.Find(bson.M{"businessLevel": 1}).One(&ftl)
	if err != nil {
		return nil, err
	}

	defaultLevel := &model.LevelInfoTech{Level: 1, MaxLevel: ftl.LevelRange[1], LeftSeconds: []int{0, 0}, TEPR: ftl.TEPR, TERE: ftl.TERE, TERO: ftl.TERO}

	return defaultLevel, nil
}
