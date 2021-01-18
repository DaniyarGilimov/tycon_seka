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

// CollectSafeHandler allows to collect safe
func CollectSafeHandler(w http.ResponseWriter, r *http.Request) {
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

	err = controller.CollectSafe(business, user, database)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		res.Message = err.Error()

		json.NewEncoder(w).Encode(res)
		return
	}
	localutils.UpdateLeftTimes(business, user)

	res.Message = "SUCCESS"

	res.Result = business

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}

// UpgradeSafeHandler allows to upgrade Safe
func UpgradeSafeHandler(w http.ResponseWriter, r *http.Request) {

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

	err = controller.UpgradeSafe(upgradeRegist, business, user, database)

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
