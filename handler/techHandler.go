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
	localcontroller "tycon_seka/controller"
	"tycon_seka/model"
	localutils "tycon_seka/utils"
)

// GetPTechHandler gives business's tech
func GetPTechHandler(w http.ResponseWriter, r *http.Request) {
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

	business, err := localcontroller.GetBusinessByID(BID, database)

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

	result, err := localcontroller.GetPTech(BID, database)

	if err != nil {
		res.Message = "Error Get Tech page. Not Found"
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

// UpgradeTechHandler allows to upgrade Tech
func UpgradeTechHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var res gmodel.Response

	userToken, err := gutils.CheckTokens(r, utils.TOKEN)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	var upgradeRegist *model.UpgradeRegist
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &upgradeRegist)

	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	if upgradeRegist.PayMethod != localutils.PaymentChips && upgradeRegist.PayMethod != localutils.PaymentGolds {
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
	localutils.UpdateLeftTimes(business, user)

	err = controller.UpgradeTech(upgradeRegist, business, user, database)

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
