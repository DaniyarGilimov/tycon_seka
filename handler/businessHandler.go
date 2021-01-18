package handler

import (
	"encoding/json"
	"general_game_azi/gcontroller"
	"general_game_azi/gmodel"
	"general_game_azi/gutils"
	"io/ioutil"
	"log"
	"net/http"
	excontroller "points/controller"
	"points/utils"
	"sync"
	"time"
	"tycon_seka/controller"
	localcontroller "tycon_seka/controller"
	"tycon_seka/model"
	localutils "tycon_seka/utils"
)

var registrationLock = sync.RWMutex{}

// GetBusinessesHandler gives all businesses
func GetBusinessesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var res gmodel.Response

	userToken, err := gutils.CheckTokens(r, utils.TOKEN)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	database, session, err := gcontroller.GetDB(utils.DBNAME, utils.URI)
	defer session.Close()

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	result, err := localcontroller.GetMyBusinesses(userToken, database)

	if err != nil {
		res.Message = "Error Get Business. Not Found"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Message = "SUCCESS"

	res.Result = result

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}

// GetStatisticsHandler gives all businesses
func GetStatisticsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var res gmodel.Response

	userToken, err := gutils.CheckTokens(r, utils.TOKEN)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	database, session, err := gcontroller.GetDB(utils.DBNAME, utils.URI)
	defer session.Close()

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	BID, err := GetBID(r)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	business, err := controller.GetBusinessByID(BID, database)

	if err != nil {
		res.Message = "Error Get Business. Not Found"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	if !localutils.CheckBusinessExist(userToken, business.UserToken) {
		res.Message = "Error such user hasn't such business"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	result, err := localcontroller.GetMyStatistics(BID, database)

	if err != nil {
		res.Message = "Error Get Business. Not Found"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Message = "SUCCESS"

	res.Result = result

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}

// GetTestHandler gives 5 test questions
func GetTestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var res gmodel.Response

	userToken, err := gutils.CheckTokens(r, utils.TOKEN)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	database, session, err := gcontroller.GetDB(utils.DBNAME, utils.URI)
	defer session.Close()

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	BID, err := GetBID(r)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	business, err := controller.GetBusinessByID(BID, database)

	if err != nil {
		res.Message = "Error Get Business. Not Found"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	if !localutils.CheckBusinessExist(userToken, business.UserToken) {
		res.Message = "Error such user hasn't such business"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	result, err := localcontroller.GetQuestions(database)

	if err != nil {
		res.Message = "Error Get Questions. Not Found"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	// Timer for passing test
	go localcontroller.SubmitTimer(BID, database)

	res.Message = "SUCCESS"

	res.Result = result

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}

//RegisterHandler is used to register user in db
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res gmodel.Response

	userToken, err := gutils.CheckTokens(r, utils.TOKEN)

	if err != nil {
		if err.Error() == "Access denied" {
			log.Print("Access denied")
			res.Message = "Access denied"
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(res)
			return
		}
	}

	var businessRegist *model.BusinessRegist
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &businessRegist)

	if err != nil {
		log.Print("read error")
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	BID, err := GetBID(r)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	database, session, err := gcontroller.GetDB(utils.DBNAME, utils.URI)
	defer session.Close()

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	user, err := excontroller.GetUserByToken(userToken, database)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	var business *model.Business

	//version, _ := controller.GetVersion(database)

	if BID == -1 {
		err := controller.IsBusinessUnique(businessRegist.Name, database)

		if err != nil {
			res.Message = err.Error()
			w.WriteHeader(gutils.StatusUniqueNameRequired)
			json.NewEncoder(w).Encode(res)
			return
		}

		business := &model.Business{UserToken: userToken, CEO: &model.Worker{}, CTO: &model.Worker{}, PR: &model.Worker{}}

		registrationLock.Lock()
		id, _ := localcontroller.GetNewIDBusiness(database)
		registrationLock.Unlock()

		business.ID = id
		business.Name = businessRegist.Name
		business.Level = 1

		business.RegistrationTime = time.Now()

		business.CEO.ID = businessRegist.CEO
		business.CTO.ID = businessRegist.CTO
		business.PR.ID = businessRegist.PR
		business.Chips = user.Chips
		business.Golds = user.Golds

		business.Level = 1

		defaultPRLevel, err := controller.GetDefaultPR(database)

		if err != nil {
			res.Message = err.Error() + "GetDefaultPR"
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(res)
			return
		}

		defaultTechLevel, err := controller.GetDefaultTech(database)

		if err != nil {
			res.Message = err.Error() + "GetDefaultTech"
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(res)
			return
		}

		defaultSafeLevel, err := controller.GetDefaultSafe(database)

		if err != nil {
			res.Message = err.Error() + "GetDefaultSafe"
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(res)
			return
		}

		business.PRLevel = defaultPRLevel
		business.TechLevel = defaultTechLevel
		business.SafeLevel = defaultSafeLevel
		business.TestInfo = &model.TestInfo{LastTestPass: time.Now().Add(time.Duration(-1) * time.Hour), TestLeftTime: []int{3600, 3600}, Decrimented: false}

		err = localcontroller.InsertBusiness(businessRegist, business, database)

		if err != nil {
			res.Message = err.Error()
			if err.Error() == "tceo" {
				w.WriteHeader(gutils.StatusTakenCEO)
			} else if err.Error() == "tcto" {
				w.WriteHeader(gutils.StatusTakenCTO)
			} else if err.Error() == "tpr" {
				w.WriteHeader(gutils.StatusTakenPR)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}

			json.NewEncoder(w).Encode(res)
			return
		}

		res.Message = "SUCCESS"
		res.Result = business
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
		return
	}

	business, err = localcontroller.GetBusinessByID(BID, database)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	if !localutils.CheckBusinessExist(userToken, business.UserToken) {
		res.Message = "Error such user hasn't such business"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	localutils.UpdateLeftTimes(business, user)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Message = "SUCCESS"
	res.Result = business
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}

// GetSecotrsHandler gives all sectors
func GetSecotrsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var res gmodel.Response

	_, err := gutils.CheckTokens(r, utils.TOKEN)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	database, session, err := gcontroller.GetDB(utils.DBNAME, utils.URI)
	defer session.Close()

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	result, err := localcontroller.GetSectors(database)

	if err != nil {
		res.Message = "Error Get sectors. Not Found"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Message = "SUCCESS"

	res.Result = result

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}

// GetLocationsHandler gives all locations
func GetLocationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var res gmodel.Response

	_, err := gutils.CheckTokens(r, utils.TOKEN)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	database, session, err := gcontroller.GetDB(utils.DBNAME, utils.URI)
	defer session.Close()

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	result, err := localcontroller.GetLocations(database)

	if err != nil {
		res.Message = "Error Get locations. Not Found"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Message = "SUCCESS"

	res.Result = result

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}

// GetSeclocsHandler gives merged sectors and locations
func GetSeclocsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var res gmodel.Response

	_, err := gutils.CheckTokens(r, utils.TOKEN)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	database, session, err := gcontroller.GetDB(utils.DBNAME, utils.URI)
	defer session.Close()

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	result, err := localcontroller.GetSecLocs(database)

	if err != nil {
		res.Message = "Error Get Sectors and Locations. Not Found"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Message = "SUCCESS"

	res.Result = result

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}

// UpgradeBusinessHandler allows to upgrade Business
func UpgradeBusinessHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var res gmodel.Response

	userToken, err := gutils.CheckTokens(r, utils.TOKEN)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	database, session, err := gcontroller.GetDB(utils.DBNAME, utils.URI)
	defer session.Close()

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	BID, err := GetBID(r)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	business, err := controller.GetBusinessByID(BID, database)

	if err != nil {
		res.Message = "Error Get Business. Not Found"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	user, err := excontroller.GetUserByToken(userToken, database)

	if err != nil {
		res.Message = "Error Get User by token. Not Found"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	if !localutils.CheckBusinessExist(userToken, business.UserToken) {
		res.Message = "Error such user hasn't such business"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	err = controller.UpgradeBusiness(business, user, database)

	if err != nil {
		if err.Error() == "ner" {
			w.WriteHeader(gutils.StatusNotEnoughResources)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		res.Message = err.Error()

		json.NewEncoder(w).Encode(res)
		return
	}

	res.Message = "SUCCESS"

	res.Result = business

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}

// PostQuestionsHandler after test completion user posts the results
func PostQuestionsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var res gmodel.Response

	userToken, err := gutils.CheckTokens(r, utils.TOKEN)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	var post *model.PostResult
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &post)

	if err != nil {
		res.Message = err.Error() + " error in read, json format is incorrect"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	database, session, err := gcontroller.GetDB(utils.DBNAME, utils.URI)
	defer session.Close()

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	BID, err := GetBID(r)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	business, err := controller.GetBusinessByID(BID, database)

	if err != nil {
		res.Message = "Error Get Business. Not Found"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	// user, err := excontroller.GetUserByToken(userToken, database)

	// if err != nil {
	// 	res.Message = "Error Get User by token. Not Found"
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(res)
	// 	return
	// }

	if !localutils.CheckBusinessExist(userToken, business.UserToken) {
		res.Message = "Error such user hasn't such business"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	result, err := controller.SubmitQuesitons(post, business, database)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		res.Message = err.Error()

		json.NewEncoder(w).Encode(res)
		return
	}

	res.Message = "SUCCESS"

	res.Result = result

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}
