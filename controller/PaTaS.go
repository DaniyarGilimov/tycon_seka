package controller

import (
	"errors"
	exmodel "general_game/gmodel"
	"general_game/gutils"
	"log"
	"os"
	"time"
	"tycon_seka/model"
	"tycon_seka/utils"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GetPPR used to get pr page data
func GetPPR(BID int, db *mgo.Database) (*model.PaTPage, error) {
	businessCollection := db.C(gutils.BUSINESS)
	var business *model.Business
	err := businessCollection.Find(bson.M{"id": BID}).One(&business)

	if err != nil {
		return nil, err
	}

	page := &model.PaTPage{Level: business.PRLevel.Level}

	prTVItem := &model.LevelItem{ID: "pr_tv", Title: "PR TV", Description: "TV is like a dinasour", Level: business.PRLevel.PRTV.LevelIncrement, Price: &model.Price{Chips: business.PRLevel.PRTV.UpgradePrice.Chips, Golds: business.PRLevel.PRTV.UpgradePrice.Golds}, Time: business.PRLevel.PRTV.UpgradeTime, IsActive: utils.PageItemActive(business.PRLevel.Level, business.PRLevel.PRTV.LevelIncrement, business.PRLevel.MaxLevel)}
	prSMItem := &model.LevelItem{ID: "pr_sm", Title: "PR SM", Description: "SM is like a dinasour", Level: business.PRLevel.PRSM.LevelIncrement, Price: &model.Price{Chips: business.PRLevel.PRSM.UpgradePrice.Chips, Golds: business.PRLevel.PRSM.UpgradePrice.Golds}, Time: business.PRLevel.PRSM.UpgradeTime, IsActive: utils.PageItemActive(business.PRLevel.Level, business.PRLevel.PRSM.LevelIncrement, business.PRLevel.MaxLevel)}
	prMGItem := &model.LevelItem{ID: "pr_mg", Title: "PR MG", Description: "MG is like a dinasour", Level: business.PRLevel.PRMG.LevelIncrement, Price: &model.Price{Chips: business.PRLevel.PRMG.UpgradePrice.Chips, Golds: business.PRLevel.PRMG.UpgradePrice.Golds}, Time: business.PRLevel.PRMG.UpgradeTime, IsActive: utils.PageItemActive(business.PRLevel.Level, business.PRLevel.PRMG.LevelIncrement, business.PRLevel.MaxLevel)}

	page.Items = append(page.Items, prTVItem)
	page.Items = append(page.Items, prSMItem)
	page.Items = append(page.Items, prMGItem)

	return page, nil
}

// GetPTech used to get tech page data
func GetPTech(BID int, db *mgo.Database) (*model.PaTPage, error) {
	businessCollection := db.C(gutils.BUSINESS)
	var business *model.Business
	err := businessCollection.Find(bson.M{"id": BID}).One(&business)

	if err != nil {
		return nil, err
	}

	page := &model.PaTPage{Level: business.TechLevel.Level}

	teREItem := &model.LevelItem{ID: "te_re", Title: "TE RE", Description: "RE is like a dinasour", Level: business.TechLevel.TERE.LevelIncrement, Price: &model.Price{Chips: business.TechLevel.TERE.UpgradePrice.Chips, Golds: business.TechLevel.TERE.UpgradePrice.Golds}, Time: business.TechLevel.TERE.UpgradeTime, IsActive: utils.PageItemActive(business.TechLevel.Level, business.TechLevel.TERE.LevelIncrement, business.TechLevel.MaxLevel)}
	teROItem := &model.LevelItem{ID: "te_ro", Title: "TE RO", Description: "RO is like a dinasour", Level: business.TechLevel.TERO.LevelIncrement, Price: &model.Price{Chips: business.TechLevel.TERO.UpgradePrice.Chips, Golds: business.TechLevel.TERO.UpgradePrice.Golds}, Time: business.TechLevel.TERO.UpgradeTime, IsActive: utils.PageItemActive(business.TechLevel.Level, business.TechLevel.TERO.LevelIncrement, business.TechLevel.MaxLevel)}
	tePRItem := &model.LevelItem{ID: "te_pr", Title: "TE PR", Description: "PR is like a dinasour", Level: business.TechLevel.TEPR.LevelIncrement, Price: &model.Price{Chips: business.TechLevel.TEPR.UpgradePrice.Chips, Golds: business.TechLevel.TEPR.UpgradePrice.Golds}, Time: business.TechLevel.TEPR.UpgradeTime, IsActive: utils.PageItemActive(business.TechLevel.Level, business.TechLevel.TEPR.LevelIncrement, business.TechLevel.MaxLevel)}

	page.Items = append(page.Items, teREItem)
	page.Items = append(page.Items, teROItem)
	page.Items = append(page.Items, tePRItem)

	return page, nil
}

// UpgradePR is used to upgrade pr
func UpgradePR(UReg *model.UpgradeRegist, Business *model.Business, User *exmodel.User, db *mgo.Database) error {
	businessCollection := db.C(gutils.BUSINESS)
	usersCollection := db.C(gutils.USERS)

	neededPriceChips := 0
	neededPriceGolds := 1
	levelUp := 0
	takeTime := 0
	speedIncrement := 0

	switch UReg.ID {
	case "pr_tv":
		neededPriceChips = Business.PRLevel.PRTV.UpgradePrice.Chips
		neededPriceGolds = Business.PRLevel.PRTV.UpgradePrice.Golds
		levelUp = Business.PRLevel.PRTV.LevelIncrement
		takeTime = Business.PRLevel.PRTV.UpgradeTime
		speedIncrement = Business.PRLevel.PRTV.SpeedIncrement
		break
	case "pr_sm":
		neededPriceChips = Business.PRLevel.PRSM.UpgradePrice.Chips
		neededPriceGolds = Business.PRLevel.PRSM.UpgradePrice.Golds
		levelUp = Business.PRLevel.PRSM.LevelIncrement
		takeTime = Business.PRLevel.PRSM.UpgradeTime
		speedIncrement = Business.PRLevel.PRSM.SpeedIncrement
		break
	case "pr_mg":
		neededPriceChips = Business.PRLevel.PRMG.UpgradePrice.Chips
		neededPriceGolds = Business.PRLevel.PRMG.UpgradePrice.Golds
		levelUp = Business.PRLevel.PRMG.LevelIncrement
		takeTime = Business.PRLevel.PRMG.UpgradeTime
		speedIncrement = Business.PRLevel.PRMG.SpeedIncrement
		break
	default:
		return errors.New("id of PR is not correct")
	}

	if levelUp+Business.PRLevel.Level > Business.PRLevel.MaxLevel {
		return errors.New("level is exceeding! its impossible")
	}

	if UReg.PayMethod == utils.PaymentChips {
		if User.Chips < neededPriceChips {
			return errors.New("ner")
		}
		User.Chips = User.Chips - neededPriceChips
	} else {
		if User.Golds < neededPriceGolds {
			return errors.New("ner")
		}
		User.Golds = User.Golds - neededPriceGolds
	}

	Business.Chips = User.Chips
	Business.Golds = User.Golds
	Business.PRLevel.SpeedIncrement = speedIncrement
	Business.PRLevel.UpgardeLevel = Business.PRLevel.Level + levelUp

	if UReg.PayMethod == utils.PaymentGolds {
		// immidiate upgrade
		UpgradePRIm(Business, db)
	} else {
		StartTime := time.Now()
		EndTime := StartTime.Add(time.Second * time.Duration(takeTime))

		Business.PRLevel.EndTime = EndTime
		Business.PRLevel.LeftSeconds = []int{1, takeTime}

		businessCollection.Update(bson.M{"id": Business.ID}, bson.M{"$set": bson.M{"prLevel": Business.PRLevel}})
		// Goroutine to schedule the update
		go UpgradePRSc(Business.ID, takeTime, db)
	}

	// leveling up in collection
	usersCollection.Update(bson.M{"userId": User.UserID}, bson.M{"$set": bson.M{"chips": User.Chips, "golds": User.Golds}})
	return nil
}

// UpgradePRSc scheduled upgrade of pr
func UpgradePRSc(BID int, TTime int, db *mgo.Database) {
	if TTime != 0 {
		dirr := PRcreateTemp(BID, TTime)
		// TODO: create temp file to recover from fault
		<-time.After(time.Duration(TTime) * time.Second)
		os.Remove(dirr)
	}

	var business *model.Business
	businessCollection := db.C(gutils.BUSINESS)

	err := businessCollection.Find(bson.M{"id": BID}).One(&business)
	if err != nil {
		return
	}

	UpgradePRIm(business, db)
}

// UpgradePRIm immidiate upgrade
func UpgradePRIm(business *model.Business, db *mgo.Database) {
	businessCollection := db.C(gutils.BUSINESS)
	tyconstatisticCollection := db.C(gutils.TYCONSTATISTICS)
	sp := business.SafeLevel.SafeParts

	// Example last part
	upVelocity := business.SafeLevel.CurrentVelocity + float64(business.PRLevel.SpeedIncrement)/3600.0

	if sp[len(sp)-1].EndTime.After(time.Now()) {
		// case when safe is collected
		utils.UpdateCollectSumLast(business)
		leftC := utils.GetLeftCapacity(business)
		dt := float64(leftC) / upVelocity

		sp = append(sp, &model.SafePart{Velocity: upVelocity, StartTime: time.Now(), Collect: 0, EndTime: time.Now().Add(time.Duration(dt) * time.Second)})
	}

	deltaVelocity := upVelocity - business.SafeLevel.CurrentVelocity
	business.SafeLevel.CurrentVelocity = upVelocity
	business.SafeLevel.SafeParts = sp
	business.PRLevel.Level = business.PRLevel.UpgardeLevel

	event := &model.TyconEvent{EventTitle: utils.EventTitlePR, EventSubTitle: utils.EventSubTitlePR, Type: utils.EventTypePR, EarnVelocity: upVelocity, Delta: deltaVelocity, Date: time.Now()}

	businessCollection.Update(bson.M{"id": business.ID}, bson.M{"$set": bson.M{"prLevel.level": business.PRLevel.UpgardeLevel, "prLevel.leftSeconds": []int{0, 0}, "safeLevel.currentVelocity": upVelocity, "safeLevel.safeParts": sp}})
	tyconstatisticCollection.Update(bson.M{"id": business.ID}, bson.M{"$push": bson.M{"events": bson.M{"$each": []*model.TyconEvent{event}, "$slice": -5}}})

}

// UpgradeTech is used to upgrade tech
func UpgradeTech(UReg *model.UpgradeRegist, Business *model.Business, User *exmodel.User, db *mgo.Database) error {
	businessCollection := db.C(gutils.BUSINESS)
	usersCollection := db.C(gutils.USERS)

	neededPriceChips := 0
	neededPriceGolds := 1
	levelUp := 0
	takeTime := 0
	speedIncrement := 0

	switch UReg.ID {
	case "te_pr":
		neededPriceChips = Business.TechLevel.TEPR.UpgradePrice.Chips
		neededPriceGolds = Business.TechLevel.TEPR.UpgradePrice.Golds
		levelUp = Business.TechLevel.TEPR.LevelIncrement
		takeTime = Business.TechLevel.TEPR.UpgradeTime
		speedIncrement = Business.TechLevel.TEPR.SpeedIncrement
		break
	case "te_ro":
		neededPriceChips = Business.TechLevel.TERO.UpgradePrice.Chips
		neededPriceGolds = Business.TechLevel.TERO.UpgradePrice.Golds
		levelUp = Business.TechLevel.TERO.LevelIncrement
		takeTime = Business.TechLevel.TERO.UpgradeTime
		speedIncrement = Business.TechLevel.TERO.SpeedIncrement
		break
	case "te_re":
		neededPriceChips = Business.TechLevel.TERE.UpgradePrice.Chips
		neededPriceGolds = Business.TechLevel.TERE.UpgradePrice.Golds
		levelUp = Business.TechLevel.TERE.LevelIncrement
		takeTime = Business.TechLevel.TERE.UpgradeTime
		speedIncrement = Business.TechLevel.TERE.SpeedIncrement
		break
	default:
		return errors.New("id of PR is not correct")
	}

	if levelUp+Business.TechLevel.Level > Business.TechLevel.MaxLevel {
		return errors.New("level is exceeding! its impossible")
	}

	if UReg.PayMethod == utils.PaymentChips {
		if User.Chips < neededPriceChips {
			return errors.New("ner")
		}
		User.Chips = User.Chips - neededPriceChips
	} else {
		if User.Golds < neededPriceGolds {
			return errors.New("ner")
		}
		User.Golds = User.Golds - neededPriceGolds
	}

	Business.Chips = User.Chips
	Business.Golds = User.Golds
	Business.TechLevel.UpgardeLevel = Business.TechLevel.Level + levelUp
	Business.TechLevel.SpeedIncrement = speedIncrement

	if UReg.PayMethod == utils.PaymentGolds {
		UpgradeTechIm(Business, db)
	} else {
		EndTime := time.Now().Add(time.Second * time.Duration(takeTime))
		Business.TechLevel.EndTime = EndTime
		Business.TechLevel.LeftSeconds = []int{1, takeTime}

		// leveling up in collection
		businessCollection.Update(bson.M{"id": Business.ID}, bson.M{"$set": bson.M{"techLevel": Business.TechLevel}})
		go UpgradeTechSc(Business.ID, takeTime, db)
	}

	usersCollection.Update(bson.M{"userId": User.UserID}, bson.M{"$set": bson.M{"chips": User.Chips, "golds": User.Golds}})
	return nil
}

// UpgradeTechSc scheduled upgrade of tech
func UpgradeTechSc(BID int, TTime int, db *mgo.Database) {
	if TTime != 0 {
		dirr := CTOcreateTemp(BID, TTime)
		<-time.After(time.Duration(TTime) * time.Second)
		os.Remove(dirr)
	}
	var business *model.Business
	businessCollection := db.C(gutils.BUSINESS)

	err := businessCollection.Find(bson.M{"id": BID}).One(&business)
	if err != nil {
		return
	}

	UpgradeTechIm(business, db)
}

// UpgradeTechIm used for immidiate update of tech
func UpgradeTechIm(business *model.Business, db *mgo.Database) {
	businessCollection := db.C(gutils.BUSINESS)
	tyconstatisticCollection := db.C(gutils.TYCONSTATISTICS)

	sp := business.SafeLevel.SafeParts

	// Example last part
	upVelocity := business.SafeLevel.CurrentVelocity + float64(business.TechLevel.SpeedIncrement)/3600.0

	if sp[len(sp)-1].EndTime.After(time.Now()) {
		// case when safe is collected
		utils.UpdateCollectSumLast(business)
		leftC := utils.GetLeftCapacity(business)
		dt := float64(leftC) / upVelocity

		sp = append(sp, &model.SafePart{Velocity: upVelocity, StartTime: time.Now(), Collect: 0, EndTime: time.Now().Add(time.Duration(dt) * time.Second)})
	}

	deltaVelocity := upVelocity - business.SafeLevel.CurrentVelocity
	business.SafeLevel.CurrentVelocity = upVelocity
	business.SafeLevel.SafeParts = sp
	business.TechLevel.Level = business.TechLevel.UpgardeLevel

	event := &model.TyconEvent{EventTitle: utils.EventTitleTech, EventSubTitle: utils.EventSubTitleTech, Type: utils.EventTypeTech, EarnVelocity: upVelocity, Delta: deltaVelocity, Date: time.Now()}

	businessCollection.Update(bson.M{"id": business.ID}, bson.M{"$set": bson.M{"techLevel.level": business.TechLevel.UpgardeLevel, "techLevel.leftSeconds": []int{0, 0}, "safeLevel.currentVelocity": upVelocity, "safeLevel.safeParts": sp}})
	tyconstatisticCollection.Update(bson.M{"id": business.ID}, bson.M{"$push": bson.M{"events": bson.M{"$each": []*model.TyconEvent{event}, "$slice": -5}}})

}

// UpgradeSafe is used to upgrade safe
func UpgradeSafe(UReg *model.UpgradeRegist, Business *model.Business, User *exmodel.User, db *mgo.Database) error {
	businessCollection := db.C(gutils.BUSINESS)
	usersCollection := db.C(gutils.USERS)

	neededPrice := Business.SafeLevel.UpgradePrice
	takeTime := Business.SafeLevel.UpgradeTakeTime

	if Business.SafeLevel.Level+1 > Business.SafeLevel.MaxLevel {
		return errors.New("level is exceeding! its impossible")
	}

	if UReg.PayMethod == utils.PaymentChips {
		if User.Chips < neededPrice.Chips {
			return errors.New("ner")
		}
		User.Chips = User.Chips - neededPrice.Chips
	} else {
		if User.Golds < neededPrice.Golds {
			return errors.New("ner")
		}
		User.Golds = User.Golds - neededPrice.Golds
	}

	Business.Chips = User.Chips
	Business.Golds = User.Golds

	if UReg.PayMethod == utils.PaymentGolds {
		// Example last part
		nsl, err := GetNextSafe(Business, db)
		if err != nil {
			//TODO: this is fatal error
			log.Print(err.Error())
		}

		Business.SafeLevel = nsl

		businessCollection.Update(bson.M{"id": Business.ID}, bson.M{"$set": bson.M{"safeLevel": Business.SafeLevel}})
	} else {
		EndTime := time.Now().Add(time.Second * time.Duration(takeTime))

		Business.SafeLevel.EndTime = EndTime
		Business.SafeLevel.LeftSeconds = []int{1, takeTime}
		//Business.SafeLevel.LeftTime = []int{0, 0}

		// leveling up in collection
		businessCollection.Update(bson.M{"id": Business.ID}, bson.M{"$set": bson.M{"safeLevel": Business.SafeLevel}})
		go UpgradeSafeSc(Business.ID, takeTime, db)
	}

	usersCollection.Update(bson.M{"userId": User.UserID}, bson.M{"$set": bson.M{"chips": User.Chips, "golds": User.Golds}})
	return nil
}

// UpgradeSafeSc scheduled upgrade of safe
func UpgradeSafeSc(BID int, TTime int, db *mgo.Database) {
	if TTime != 0 {
		dirr := SafeCreateTemp(BID, TTime)
		<-time.After(time.Duration(TTime) * time.Second)
		os.Remove(dirr)
	}

	var business *model.Business
	businessCollection := db.C(gutils.BUSINESS)

	err := businessCollection.Find(bson.M{"id": BID}).One(&business)
	if err != nil {
		return
	}

	// Example last part
	business.SafeLevel, err = GetNextSafe(business, db)
	if err != nil {
		//TODO: this is fatal error
		log.Print(err.Error())
	}

	businessCollection.Update(bson.M{"id": BID}, bson.M{"$set": bson.M{"safeLevel": business.SafeLevel}})
}

// CollectSafe collects safe chips
func CollectSafe(Business *model.Business, User *exmodel.User, db *mgo.Database) error {
	businessCollection := db.C(gutils.BUSINESS)
	usersCollection := db.C(gutils.USERS)

	utils.UpdateLeftTimeSafe(Business)

	User.Chips = User.Chips + Business.SafeLevel.Sum
	Business.Chips = User.Chips

	dTime := int(float64(Business.SafeLevel.CurrentCapacity) / Business.SafeLevel.CurrentVelocity)
	endTime := time.Now().Add(time.Duration(dTime) * time.Second)

	Business.SafeLevel.SafeParts = []*model.SafePart{{Velocity: Business.SafeLevel.CurrentVelocity, StartTime: time.Now(), EndTime: endTime}}
	Business.SafeLevel.Sum = 0
	//Business.SafeLevel.LeftTime = []int{0, int(float64(Business.SafeLevel.CurrentCapacity) / Business.SafeLevel.CurrentVelocity)}

	// leveling up in collection
	businessCollection.Update(bson.M{"id": Business.ID}, bson.M{"$set": bson.M{"safeLevel.safeParts": []bson.M{{"velocity": Business.SafeLevel.CurrentVelocity, "endTime": endTime, "startTime": time.Now()}}}})
	usersCollection.Update(bson.M{"userId": User.UserID}, bson.M{"$set": bson.M{"chips": User.Chips}})

	return nil
}

// RefreshSafeParts used to add a new safe part into parts
func RefreshSafeParts(Business *model.Business, upVelocity float64) {
	sp := Business.SafeLevel.SafeParts
	if sp[len(sp)-1].EndTime.After(time.Now()) {
		// case when safe is collected
		utils.UpdateCollectSumLast(Business)
		leftC := utils.GetLeftCapacity(Business)
		dt := float64(leftC) / upVelocity

		sp = append(sp, &model.SafePart{Velocity: upVelocity, StartTime: time.Now(), Collect: 0, EndTime: time.Now().Add(time.Duration(dt) * time.Second)})
	}
	Business.SafeLevel.SafeParts = sp
}
