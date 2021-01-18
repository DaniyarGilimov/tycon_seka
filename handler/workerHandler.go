package handler

import (
	"encoding/json"
	"general_game/gcontroller"
	"general_game/gmodel"
	"general_game/gutils"
	"io/ioutil"
	"net/http"
	excontroller "points/controller"
	"points/utils"
	"tycon_seka/controller"
	"tycon_seka/model"
	localutils "tycon_seka/utils"
)

// GetCTOsHandler gives all ctos in certain level
func GetCTOsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var res gmodel.Response

	userToken, err := gutils.CheckTokens(r, utils.TOKEN)

	if err != nil {
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

	business, err := controller.GetBusinessForWorker(BID, database)

	if err != nil {
		res.Message = "Error Get Business for worker. Not Found"
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

	result, err := controller.GetCTOs(business.Level+1, database)

	if err != nil {
		res.Message = "Error Get CTO. Not Found"
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

// GetCEOsHandler gives all ceos in certain level
func GetCEOsHandler(w http.ResponseWriter, r *http.Request) {
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

	business, err := controller.GetBusinessForWorker(BID, database)

	if err != nil {
		res.Message = "Error Get Business for worker. Not Found"
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

	result, err := controller.GetCEOs(business.Level+1, database)

	if err != nil {
		res.Message = "Error Get CEOs. Not Found"
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

// GetPRsHandler gives all prs in certain level
func GetPRsHandler(w http.ResponseWriter, r *http.Request) {
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

	business, err := controller.GetBusinessForWorker(BID, database)

	if err != nil {
		res.Message = "Error Get Business for worker. Not Found"
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

	result, err := controller.GetPRs(business.Level+1, database)

	if err != nil {
		res.Message = "Error Get PRs. Not Found"
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

// HairCTOHandler allows to hire PR
func HairCTOHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var res gmodel.Response

	userToken, err := gutils.CheckTokens(r, utils.TOKEN)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	var hireRegist *model.HireRegist
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &hireRegist)

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
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(res)
		return
	}

	user, err := excontroller.GetUserByToken(userToken, database)

	if err != nil {
		res.Message = "Error Get user by token. Not Found"
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

	err = controller.HireCTO(hireRegist.ID, business, user, database)

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

	localutils.UpdateLeftTimes(business, user)
	res.Result = business

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}

// HairCEOHandler allows to hire CEO
func HairCEOHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var res gmodel.Response

	userToken, err := gutils.CheckTokens(r, utils.TOKEN)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	var hireRegist *model.HireRegist
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &hireRegist)

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
		res.Message = "Error Get Business by ID. Not Found"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	user, err := excontroller.GetUserByToken(userToken, database)

	if err != nil {
		res.Message = "Error Get User By Token. Not Found"
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

	err = controller.HireCEO(hireRegist.ID, business, user, database)

	if err != nil {
		if err.Error() == "ner" {
			w.WriteHeader(gutils.StatusNotEnoughResources)
		} else if err.Error() == "tceo" {
			w.WriteHeader(gutils.StatusTakenCEO)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		res.Message = err.Error()

		json.NewEncoder(w).Encode(res)
		return
	}

	res.Message = "SUCCESS"
	localutils.UpdateLeftTimes(business, user)

	res.Result = business

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}

// HairPRHandler allows to hire PR
func HairPRHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var res gmodel.Response

	userToken, err := gutils.CheckTokens(r, utils.TOKEN)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	var hireRegist *model.HireRegist
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &hireRegist)

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
		res.Message = "Error Get Business by ID. Not Found"
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

	err = controller.HirePR(hireRegist.ID, business, user, database)

	if err != nil {
		if err.Error() == "ner" {
			w.WriteHeader(gutils.StatusNotEnoughResources)
		} else if err.Error() == "tpr" {
			w.WriteHeader(gutils.StatusTakenPR)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		res.Message = err.Error()

		json.NewEncoder(w).Encode(res)
		return
	}

	res.Message = "SUCCESS"

	localutils.UpdateLeftTimes(business, user)
	res.Result = business

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}
