package controller

import (
	"errors"
	exmodel "general_game/gmodel"
	"general_game_azi/gutils"
	"time"
	"tycon_seka/model"
	"tycon_seka/utils"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// IsBusinessUnique shows if business is unique or not
func IsBusinessUnique(Name string, db *mgo.Database) error {
	businessCollection := db.C(gutils.BUSINESS)
	var res *model.Business
	err := businessCollection.Find(
		bson.M{"name": bson.M{"$regex": bson.RegEx{Pattern: Name, Options: "i"}}},
	).One(&res)

	if err == nil {
		return errors.New("Business name should be unique")
	}
	return nil
}

// InsertBusiness is used to insert business
func InsertBusiness(BusinessRegist *model.BusinessRegist, Business *model.Business, db *mgo.Database) error {
	businessCollection := db.C(gutils.BUSINESS)
	ctosCollection := db.C(gutils.CTOS)
	ceosCollection := db.C(gutils.CEOS)
	prsCollection := db.C(gutils.PRS)
	sectorCollection := db.C(gutils.SECTOR)
	locationCollection := db.C(gutils.LOCATION)
	tyconstatsCollection := db.C(gutils.TYCONSTATISTICS)

	var cto *model.Worker
	var ceo *model.Worker
	var pr *model.Worker
	var sector *model.Sector
	var location *model.Location

	err := ctosCollection.Find(bson.M{"ID": Business.CTO.ID}).One(&cto)
	if err != nil {
		return err
	}

	err = ceosCollection.Find(bson.M{"ID": Business.CEO.ID}).One(&ceo)
	if err != nil {
		return err
	}

	err = prsCollection.Find(bson.M{"ID": Business.PR.ID}).One(&pr)
	if err != nil {
		return err
	}

	err = sectorCollection.Find(bson.M{"id": BusinessRegist.Sector}).One(&sector)
	if err != nil {
		return err
	}

	err = locationCollection.Find(bson.M{"id": BusinessRegist.Location}).One(&location)
	if err != nil {
		return err
	}

	// if ceo.WCID != 0 {
	// 	return errors.New("tceo")
	// }
	// if cto.WCID != 0 {
	// 	return errors.New("tcto")
	// }
	// if pr.WCID != 0 {
	// 	return errors.New("tpr")
	// }

	// ctosCollection.Update(bson.M{"ID": Business.CTO.ID}, bson.M{"$set": bson.M{"WCID": Business.ID}})
	// ceosCollection.Update(bson.M{"ID": Business.CEO.ID}, bson.M{"$set": bson.M{"WCID": Business.ID}})
	// prsCollection.Update(bson.M{"ID": Business.PR.ID}, bson.M{"$set": bson.M{"WCID": Business.ID}})

	Business.CEO = ceo
	Business.CTO = cto
	Business.PR = pr
	Business.Sector = sector
	Business.Location = location

	event := &model.TyconEvent{EventTitle: "Поддержка от государства", EventSubTitle: "Малый бизнес кредит", Type: utils.EventTypeNews, EarnVelocity: Business.SafeLevel.CurrentVelocity, Delta: Business.SafeLevel.CurrentVelocity, Date: time.Now()}
	tyconStatistics := &model.TyconStatistics{ID: Business.ID, Events: []*model.TyconEvent{event}}

	err = businessCollection.Insert(Business)
	if err != nil {
		return err
	}

	err = tyconstatsCollection.Insert(tyconStatistics)
	if err != nil {
		return err
	}

	err = sectorCollection.Update(bson.M{"id": sector.ID}, bson.M{"$set": bson.M{"population": sector.Population + 1}})
	if err != nil {
		return err
	}

	err = locationCollection.Update(bson.M{"id": location.ID}, bson.M{"$set": bson.M{"population": location.Population + 1}})

	return err
}

// GetMyBusinesses is used to get businesses
func GetMyBusinesses(UserToken string, db *mgo.Database) ([]*model.Business, error) {
	businessCollection := db.C(gutils.BUSINESS)
	business := []*model.Business{}
	err := businessCollection.Find(bson.M{"userToken": UserToken}).All(&business)
	if err != nil {
		return nil, err
	}
	return business, err
}

// GetMyStatistics is used to get statistics
func GetMyStatistics(BID int, db *mgo.Database) (*model.TyconStatistics, error) {
	tyconstatsCollection := db.C(gutils.TYCONSTATISTICS)
	businessCollection := db.C(gutils.BUSINESS)
	var stat *model.TyconStatistics
	var business *model.Business

	err := tyconstatsCollection.Find(bson.M{"id": BID}).One(&stat)
	if err != nil {
		return nil, err
	}

	err = businessCollection.Find(bson.M{"id": BID}).One(&business)
	if err != nil {
		return nil, err
	}

	stat.Graph = &model.TyconGraph{XAxis: []int{}, YAxis: []float64{}}
	for i := 0; i < len(stat.Events); i++ {
		stat.Graph.XAxis = append(stat.Graph.XAxis, i)
		stat.Graph.YAxis = append(stat.Graph.YAxis, stat.Events[i].EarnVelocity)
	}
	return stat, err
}

// GetQuestions is used to get questions
func GetQuestions(db *mgo.Database) ([]*model.Question, error) {
	testCollection := db.C(gutils.TESTQUESTIONS)
	var questions []*model.Question

	pipea := []bson.M{bson.M{"$sample": bson.M{"size": 5}}}

	err := testCollection.Pipe(pipea).All(&questions)
	if err != nil {
		return nil, err
	}
	return questions, err
}

// GetBusinessByID is used to get business by its token
func GetBusinessByID(BID int, db *mgo.Database) (*model.Business, error) {
	businessCollection := db.C(gutils.BUSINESS)
	tyconstatsCollection := db.C(gutils.TYCONSTATISTICS)

	var business *model.Business
	err := businessCollection.Find(bson.M{"id": BID}).One(&business)
	if err != nil {
		return nil, err
	}

	//Defining if test is available
	if !business.TestInfo.LastTestPass.Add(1 * time.Hour).After(time.Now()) {
		business.TestInfo.TestAvailable = true
		// Calculating left time for next test
		business.TestInfo.TestLeftTime = []int{0, 3600}
		//check for decriment
		if !business.TestInfo.LastTestPass.Add(2 * time.Hour).After(time.Now()) {
			//check if decriment had set, otherwise set it
			if !business.TestInfo.Decrimented {
				// do decriment, and update the table and set event
				cVelocity := business.SafeLevel.CurrentVelocity
				upVelocity := cVelocity - cVelocity*0.2
				delta := upVelocity - cVelocity
				if upVelocity < 0 {
					upVelocity = 0
				}

				// Setting safe part
				RefreshSafeParts(business, upVelocity)

				event := &model.TyconEvent{EventTitle: "Missing test", EventSubTitle: "Dont miss test anymore", Type: utils.EventTypeNews, EarnVelocity: upVelocity, Delta: delta, Date: business.TestInfo.LastTestPass.Add(2 * time.Hour)}

				business.TestInfo.Decrimented = true

				businessCollection.Update(bson.M{"id": BID}, bson.M{"$set": bson.M{"safeLevel.currentVelocity": upVelocity, "safeLevel.safeParts": business.SafeLevel.SafeParts, "testInfo.decrimented": true}})
				tyconstatsCollection.Update(bson.M{"id": BID}, bson.M{"$push": bson.M{"events": bson.M{"$each": []*model.TyconEvent{event}, "$slice": -5}}})
			}
		}
	} else {
		// Calculating left time for next test
		d := time.Now().Sub(business.TestInfo.LastTestPass)
		l := int(d.Seconds())

		if l > 3600 {
			l = 3600
		}

		business.TestInfo.TestLeftTime = []int{l, 3600}
	}

	return business, err
}

// GetBusinessForWorker is used to get business by its token for worker case
func GetBusinessForWorker(BID int, db *mgo.Database) (*model.Business, error) {
	businessCollection := db.C(gutils.BUSINESS)
	var business *model.Business
	if BID != -1 {
		err := businessCollection.Find(bson.M{"id": BID}).One(&business)
		if err != nil {
			return nil, err
		}
		return business, err
	}

	business = &model.Business{Level: 0}
	return business, nil
}

//GetNewIDBusiness gives new id from business collection
func GetNewIDBusiness(database *mgo.Database) (int, error) {

	businessCollection := database.C(gutils.BUSINESS)

	type MaxID struct {
		ID  string
		Max int
	}

	var result MaxID
	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id": "_id",
				"max": bson.M{"$max": "$id"},
			},
		},
	}

	err := businessCollection.Pipe(pipeline).One(&result)
	result.Max++
	return result.Max, err
}

// GetSectors is used to get sectors
func GetSectors(db *mgo.Database) ([]*model.Sector, error) {
	sectorsCollection := db.C(gutils.SECTOR)
	sectors := []*model.Sector{}
	err := sectorsCollection.Find(bson.M{}).All(&sectors)
	if err != nil {
		return nil, err
	}
	return sectors, err
}

// GetLocations is used to get locations
func GetLocations(db *mgo.Database) ([]*model.Location, error) {
	locationsCollection := db.C(gutils.LOCATION)
	locations := []*model.Location{}
	err := locationsCollection.Find(bson.M{}).All(&locations)
	if err != nil {
		return nil, err
	}
	return locations, err
}

// GetSecLocs is used to get sectors and locations
func GetSecLocs(db *mgo.Database) (*model.SecLoc, error) {
	locationsCollection := db.C(gutils.LOCATION)
	sectorsCollection := db.C(gutils.SECTOR)
	locations := []*model.Location{}
	sectors := []*model.Sector{}

	err := locationsCollection.Find(bson.M{}).All(&locations)
	if err != nil {
		return nil, err
	}
	err = sectorsCollection.Find(bson.M{}).All(&sectors)
	if err != nil {
		return nil, err
	}
	secLoc := &model.SecLoc{Sectors: sectors, Locations: locations, Price: utils.InitialCreatePrice}

	return secLoc, err
}

// UpgradeBusiness is used to upgrade business
func UpgradeBusiness(Business *model.Business, User *exmodel.User, db *mgo.Database) error {
	businessCollection := db.C(gutils.BUSINESS)

	if Business.Level+1 > 15 {
		return errors.New("level is exceeding! its impossible")
	}

	nL := Business.Level + 1

	if Business.CEO.Level != nL || Business.CTO.Level != nL || Business.PR.Level != nL {
		return errors.New("workers are not upgraded")
	}

	if Business.PRLevel.Level != Business.PRLevel.MaxLevel {
		return errors.New("PR need to be upgraded")
	}

	if Business.TechLevel.Level != Business.TechLevel.MaxLevel {
		return errors.New("Tech need to be upgraded")
	}

	if Business.SafeLevel.Level != Business.SafeLevel.MaxLevel {
		return errors.New("Safe need to be upgraded")
	}

	Business.Level = Business.Level + 1
	Business.Chips = User.Chips
	Business.Golds = User.Golds

	// leveling up in collection
	businessCollection.Update(bson.M{"id": Business.ID}, bson.M{"$set": bson.M{"level": nL}})
	return nil
}

// SubmitQuesitons is used to submit question result
func SubmitQuesitons(TestResult *model.PostResult, Business *model.Business, db *mgo.Database) (*model.TestResponse, error) {
	tyconstatsCollection := db.C(gutils.TYCONSTATISTICS)
	businessCollection := db.C(gutils.BUSINESS)

	cVelocit := Business.SafeLevel.CurrentVelocity
	upVelocity := Business.SafeLevel.CurrentVelocity
	var delta float64
	percentage := 0
	// Defining up velocity for each correct answer
	switch TestResult.CorrectCount {
	case 0:
		upVelocity = upVelocity - Business.SafeLevel.CurrentVelocity*0.15
		percentage = -15
	case 1:
		upVelocity = upVelocity - Business.SafeLevel.CurrentVelocity*0.1
		percentage = -10
	case 2:
		upVelocity = upVelocity - Business.SafeLevel.CurrentVelocity*0.05
		percentage = -5
	case 3:
		upVelocity = upVelocity + Business.SafeLevel.CurrentVelocity*0.2
		percentage = 20
	case 4:
		upVelocity = upVelocity + Business.SafeLevel.CurrentVelocity*0.3
		percentage = 30
	case 5:
		upVelocity = upVelocity + Business.SafeLevel.CurrentVelocity*0.4
		percentage = 40
	}
	delta = upVelocity - cVelocit

	sp := Business.SafeLevel.SafeParts

	if sp[len(sp)-1].EndTime.After(time.Now()) {
		// case when safe is collected
		utils.UpdateCollectSumLast(Business)
		leftC := utils.GetLeftCapacity(Business)
		dt := float64(leftC) / upVelocity

		sp = append(sp, &model.SafePart{Velocity: upVelocity, StartTime: time.Now(), Collect: 0, EndTime: time.Now().Add(time.Duration(dt) * time.Second)})
	}

	event := &model.TyconEvent{EventTitle: "Test results", EventSubTitle: "test affect on graph", Type: utils.EventTypeNews, EarnVelocity: upVelocity, Delta: delta, Date: time.Now()}

	tyconstatsCollection.Update(bson.M{"id": Business.ID}, bson.M{"$push": bson.M{"events": bson.M{"$each": []*model.TyconEvent{event}, "$slice": -5}}})
	businessCollection.Update(bson.M{"id": Business.ID}, bson.M{"$set": bson.M{"testInfo.lastTestPass": time.Now(), "safeLevel.currentVelocity": upVelocity, "safeLevel.safeParts": sp}})

	resp := &model.TestResponse{CorrectCount: TestResult.CorrectCount, CurrentIncome: upVelocity, IncomeChangePercent: percentage}
	return resp, nil
}

// SubmitTimer is used to set timer for passing the test
func SubmitTimer(BID int, db *mgo.Database) {
	<-time.After(time.Duration(utils.TimerTest) * time.Second)
	tyconstatsCollection := db.C(gutils.TYCONSTATISTICS)
	businessCollection := db.C(gutils.BUSINESS)

	var Business *model.Business
	businessCollection.Find(bson.M{"id": BID}).One(&Business)

	if time.Now().Sub(Business.TestInfo.LastTestPass) > time.Duration(utils.TimerTest*time.Second) {
		cVelocit := Business.SafeLevel.CurrentVelocity
		upVelocity := cVelocit - Business.SafeLevel.CurrentVelocity*0.15
		var delta float64

		delta = upVelocity - cVelocit

		sp := Business.SafeLevel.SafeParts

		if sp[len(sp)-1].EndTime.After(time.Now()) {
			// case when safe is collected
			utils.UpdateCollectSumLast(Business)
			leftC := utils.GetLeftCapacity(Business)
			dt := float64(leftC) / upVelocity

			sp = append(sp, &model.SafePart{Velocity: upVelocity, StartTime: time.Now(), Collect: 0, EndTime: time.Now().Add(time.Duration(dt) * time.Second)})
		}

		event := &model.TyconEvent{EventTitle: "Test results", EventSubTitle: "test affect on graph", Type: utils.EventTypeNews, EarnVelocity: upVelocity, Delta: delta, Date: time.Now()}

		tyconstatsCollection.Update(bson.M{"id": Business.ID}, bson.M{"$push": bson.M{"events": bson.M{"$each": []*model.TyconEvent{event}, "$slice": -5}}})
		businessCollection.Update(bson.M{"id": Business.ID}, bson.M{"$set": bson.M{"testInfo.lastTestPass": time.Now(), "safeLevel.currentVelocity": upVelocity, "safeLevel.safeParts": sp}})
	}
	return
}
