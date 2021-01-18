package controller

import (
	"encoding/json"
	"io/ioutil"
	"time"
	"tycon_seka/model"
	"tycon_seka/utils"
)

// PRcreateTemp creates temp file for PR
func PRcreateTemp(BID int, TTime int) string {
	ntime := time.Now()
	fullDir := utils.ProjectPath() + "tmp/pr/dat" + ntime.String() + ".dat"
	info := model.RuntimeUpgrade{StartTime: ntime, EndTime: ntime.Add(time.Duration(TTime) * time.Second), BusinessID: BID, UpgradeObject: "up_pr"}
	formated, _ := json.Marshal(info)
	ioutil.WriteFile(fullDir, formated, 0644)
	return fullDir
}

// SafeCreateTemp creates temp file for CEO
func SafeCreateTemp(BID int, TTime int) string {
	ntime := time.Now()
	fullDir := utils.ProjectPath() + "tmp/ceo/dat" + ntime.String() + ".dat"
	info := model.RuntimeUpgrade{StartTime: ntime, EndTime: ntime.Add(time.Duration(TTime) * time.Second), BusinessID: BID, UpgradeObject: "up_ceo"}
	formated, _ := json.Marshal(info)
	ioutil.WriteFile(fullDir, formated, 0644)
	return fullDir
}

// CTOcreateTemp creates temp file for CTO
func CTOcreateTemp(BID int, TTime int) string {
	ntime := time.Now()
	fullDir := utils.ProjectPath() + "tmp/ceo/dat" + ntime.String() + ".dat"
	info := model.RuntimeUpgrade{StartTime: ntime, EndTime: ntime.Add(time.Duration(TTime) * time.Second), BusinessID: BID, UpgradeObject: "up_cto"}
	formated, _ := json.Marshal(info)
	ioutil.WriteFile(fullDir, formated, 0644)
	return fullDir
}
