package controller

import (
	"errors"
	exmodel "general_game/gmodel"
	"general_game/gutils"
	"time"
	"tycon_seka/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GetCTOs is used to get ctos of certain level
func GetCTOs(Level int, db *mgo.Database) ([]*model.Worker, error) {
	ctosCollection := db.C(gutils.CTOS)
	workers := []*model.Worker{}
	pipea := []bson.M{bson.M{"$match": bson.M{"Level": Level}}}
	err := ctosCollection.Pipe(pipea).All(&workers)
	if err != nil {
		return nil, err
	}
	return workers, err
}

// GetCEOs is used to get ceos of certain level
func GetCEOs(Level int, db *mgo.Database) ([]*model.Worker, error) {
	ceosCollection := db.C(gutils.CEOS)
	workers := []*model.Worker{}
	pipea := []bson.M{bson.M{"$match": bson.M{"Level": Level}}}
	err := ceosCollection.Pipe(pipea).All(&workers)
	if err != nil {
		return nil, err
	}
	return workers, err
}

// GetPRs is used to get ctos of certain level
func GetPRs(Level int, db *mgo.Database) ([]*model.Worker, error) {
	prsCollection := db.C(gutils.PRS)
	workers := []*model.Worker{}
	pipea := []bson.M{bson.M{"$match": bson.M{"Level": Level}}}
	err := prsCollection.Pipe(pipea).All(&workers)
	if err != nil {
		return nil, err
	}
	return workers, err
}

// HireCEO is used to hire ceo
func HireCEO(WorkerID int, Business *model.Business, User *exmodel.User, db *mgo.Database) error {
	var ceo *model.Worker
	currentCeo := Business.CEO

	ceosCollection := db.C(gutils.CEOS)
	businessCollection := db.C(gutils.BUSINESS)
	usersCollection := db.C(gutils.USERS)

	err := ceosCollection.Find(bson.M{"ID": WorkerID}).One(&ceo)

	if err != nil {
		return err
	}

	// if ceo.WCID != 0 {
	// 	return errors.New("tceo")
	// }

	if User.Chips < ceo.Price {
		return errors.New("ner")
	}
	User.Chips = User.Chips - ceo.Price
	Business.Chips = User.Chips

	if len(currentCeo.Jobs) > 1 {
		firstJob := &model.JobHistory{StartYear: 2016, EndYear: 2018, Company: currentCeo.Jobs[1].Company, Location: currentCeo.Jobs[1].Location, Comment: currentCeo.Jobs[1].Comment, Title: currentCeo.Jobs[1].Title}
		secondJob := &model.JobHistory{StartYear: 2019, EndYear: time.Now().Year(), Company: Business.Name, Location: Business.Location.Name, Comment: "Random generate", Title: "Random Title"}
		currentCeo.Jobs = []*model.JobHistory{firstJob, secondJob}
	}

	businessCollection.Update(bson.M{"id": Business.ID}, bson.M{"$set": bson.M{"ceo": ceo}})
	usersCollection.Update(bson.M{"userId": User.UserID}, bson.M{"$set": bson.M{"chips": User.Chips}})
	//ceosCollection.Update(bson.M{"ID": WorkerID}, bson.M{"$set": bson.M{"WCID": Business.ID}})
	ceosCollection.Update(bson.M{"ID": currentCeo.ID}, bson.M{"$set": bson.M{"WCID": 0, "Jobs": currentCeo.Jobs}})

	Business.CEO = ceo

	return nil
}

// HireCTO is used to hire cto
func HireCTO(WorkerID int, Business *model.Business, User *exmodel.User, db *mgo.Database) error {
	var cto *model.Worker
	currentCto := Business.CTO

	ctosCollection := db.C(gutils.CTOS)
	businessCollection := db.C(gutils.BUSINESS)
	usersCollection := db.C(gutils.USERS)

	err := ctosCollection.Find(bson.M{"ID": WorkerID}).One(&cto)

	if err != nil {
		return err
	}

	// if cto.WCID != 0 {
	// 	return errors.New("tcto")
	// }

	if User.Chips < cto.Price {
		return errors.New("ner")
	}
	User.Chips = User.Chips - cto.Price
	Business.Chips = User.Chips

	if len(currentCto.Jobs) > 1 {
		firstJob := &model.JobHistory{StartYear: 2016, EndYear: 2018, Company: currentCto.Jobs[1].Company, Location: currentCto.Jobs[1].Location, Comment: currentCto.Jobs[1].Comment, Title: currentCto.Jobs[1].Title}
		secondJob := &model.JobHistory{StartYear: 2019, EndYear: time.Now().Year(), Company: Business.Name, Location: Business.Location.Name, Comment: "Random generate", Title: "Random Title"}
		currentCto.Jobs = []*model.JobHistory{firstJob, secondJob}
	}

	businessCollection.Update(bson.M{"id": Business.ID}, bson.M{"$set": bson.M{"cto": cto}})
	usersCollection.Update(bson.M{"userId": User.UserID}, bson.M{"$set": bson.M{"chips": User.Chips}})
	// ctosCollection.Update(bson.M{"ID": WorkerID}, bson.M{"$set": bson.M{"WCID": Business.ID}})
	ctosCollection.Update(bson.M{"ID": currentCto.ID}, bson.M{"$set": bson.M{"WCID": 0, "Jobs": currentCto.Jobs}})

	Business.CTO = cto

	return nil
}

// HirePR is used to hire pr
func HirePR(WorkerID int, Business *model.Business, User *exmodel.User, db *mgo.Database) error {
	var pr *model.Worker
	currentPr := Business.PR

	prsCollection := db.C(gutils.PRS)
	businessCollection := db.C(gutils.BUSINESS)
	usersCollection := db.C(gutils.USERS)

	err := prsCollection.Find(bson.M{"ID": WorkerID}).One(&pr)

	if err != nil {
		return err
	}

	// if pr.WCID != 0 {
	// 	return errors.New("tpr")
	// }

	if User.Chips < pr.Price {
		return errors.New("ner")
	}
	User.Chips = User.Chips - pr.Price
	Business.Chips = User.Chips

	if len(currentPr.Jobs) > 1 {
		firstJob := &model.JobHistory{StartYear: 2016, EndYear: 2018, Company: currentPr.Jobs[1].Company, Location: currentPr.Jobs[1].Location, Comment: currentPr.Jobs[1].Comment, Title: currentPr.Jobs[1].Title}
		secondJob := &model.JobHistory{StartYear: 2019, EndYear: time.Now().Year(), Company: Business.Name, Location: Business.Location.Name, Comment: "Random generate", Title: "Random Title"}
		currentPr.Jobs = []*model.JobHistory{firstJob, secondJob}
	}

	businessCollection.Update(bson.M{"id": Business.ID}, bson.M{"$set": bson.M{"pr": pr}})
	usersCollection.Update(bson.M{"userId": User.UserID}, bson.M{"$set": bson.M{"chips": User.Chips}})
	//prsCollection.Update(bson.M{"ID": WorkerID}, bson.M{"$set": bson.M{"WCID": Business.ID}})
	prsCollection.Update(bson.M{"ID": currentPr.ID}, bson.M{"$set": bson.M{"WCID": 0, "Jobs": currentPr.Jobs}})

	Business.PR = pr

	return nil
}
