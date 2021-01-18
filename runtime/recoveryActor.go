package runtime

import (
	aziutils "azi_api_v2/utils"
	"encoding/json"
	"fmt"
	"general_game/gcontroller"
	"io/ioutil"
	"log"
	"os"
	"time"
	"tycon_seka/controller"
	"tycon_seka/model"
	"tycon_seka/utils"

	"gopkg.in/mgo.v2"
)

// DoRecovery does general recovery
func DoRecovery() {
	database, session, err := gcontroller.GetDB(aziutils.DBNAME, aziutils.URI)

	if err != nil {
		log.Print("Error in establishing connectino with database")
		return
	}
	defer session.Close()
	DoPRRecovery(database)
	DoSafeRecovery(database)
	DoCTORecovery(database)
}

// DoPRRecovery does PR recovery
func DoPRRecovery(db *mgo.Database) {
	files, err := ioutil.ReadDir(utils.ProjectPath() + "tmp/pr/")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		// Open our jsonFile
		jsonFile, err := os.Open(utils.ProjectPath() + "tmp/pr/" + f.Name())
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var rec *model.RuntimeUpgrade
		json.Unmarshal(byteValue, &rec)
		jsonFile.Close()
		os.Remove(utils.ProjectPath() + "tmp/pr/" + f.Name())
		if time.Now().After(rec.EndTime) {
			// do manual update
			controller.UpgradePRSc(rec.BusinessID, 0, db)
		} else {
			// rerun goroutine
			TTime := int(rec.EndTime.Sub(rec.StartTime).Seconds())
			controller.UpgradePRSc(rec.BusinessID, TTime, db)
		}
		// defer the closing of our jsonFile so that we can parse it later on
	}
}

// DoSafeRecovery does Safe recovery
func DoSafeRecovery(db *mgo.Database) {
	files, err := ioutil.ReadDir(utils.ProjectPath() + "tmp/ceo/")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		// Open our jsonFile
		jsonFile, err := os.Open(utils.ProjectPath() + "tmp/ceo/" + f.Name())
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var rec *model.RuntimeUpgrade
		json.Unmarshal(byteValue, &rec)
		jsonFile.Close()
		os.Remove(utils.ProjectPath() + "tmp/ceo/" + f.Name())
		if time.Now().After(rec.EndTime) {
			// do manual update
			controller.UpgradeSafeSc(rec.BusinessID, 0, db)
		} else {
			// rerun goroutine
			TTime := int(rec.EndTime.Sub(rec.StartTime).Seconds())
			controller.UpgradeSafeSc(rec.BusinessID, TTime, db)
		}
	}
}

// DoCTORecovery does CTO recovery
func DoCTORecovery(db *mgo.Database) {
	files, err := ioutil.ReadDir(utils.ProjectPath() + "tmp/cto/")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		// Open our jsonFile
		jsonFile, err := os.Open(utils.ProjectPath() + "tmp/cto/" + f.Name())
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var rec *model.RuntimeUpgrade
		json.Unmarshal(byteValue, &rec)
		jsonFile.Close()
		os.Remove(utils.ProjectPath() + "tmp/cto/" + f.Name())
		if time.Now().After(rec.EndTime) {
			// do manual update
			controller.UpgradeTechSc(rec.BusinessID, 0, db)
		} else {
			// rerun goroutine
			TTime := int(rec.EndTime.Sub(rec.StartTime).Seconds())
			controller.UpgradeTechSc(rec.BusinessID, TTime, db)
		}
	}
}
