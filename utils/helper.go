package utils

import (
	exmodel "general_game/gmodel"
	"log"
	"time"
	"tycon_seka/model"
)

// UpdateLeftTimes used to update left times in timer
func UpdateLeftTimes(Business *model.Business, User *exmodel.User) {
	if Business.PRLevel.LeftSeconds[0] != 0 {
		deltaTime := Business.PRLevel.EndTime.Sub(time.Now())
		t := int(deltaTime.Seconds())
		if t == 0 {
			t = 1
		}
		t = Business.PRLevel.LeftSeconds[1] - t
		Business.PRLevel.LeftSeconds[0] = t
	}
	if Business.TechLevel.LeftSeconds[0] != 0 {
		deltaTime := Business.TechLevel.EndTime.Sub(time.Now())
		t := int(deltaTime.Seconds())
		if t == 0 {
			t = 1
		}
		t = Business.TechLevel.LeftSeconds[1] - t
		Business.TechLevel.LeftSeconds[0] = t
	}
	if Business.SafeLevel.LeftSeconds[0] != 0 {
		deltaTime := Business.SafeLevel.EndTime.Sub(time.Now())
		t := int(deltaTime.Seconds())
		if t == 0 {
			t = 1
		}
		t = Business.SafeLevel.LeftSeconds[1] - t
		Business.SafeLevel.LeftSeconds[0] = t
		//Business.SafeLevel.LeftTime = []int{0, 0}
	} else {
		UpdateLeftTimeSafe(Business)
	}

	b := Business
	wl := b.Level + 1
	if b.PR.Level == wl && b.CTO.Level == wl && b.CEO.Level == wl && b.PRLevel.Level == b.PRLevel.MaxLevel && b.TechLevel.Level == b.TechLevel.MaxLevel && b.SafeLevel.Level == b.SafeLevel.MaxLevel && b.BattleLevel.Level == b.BattleLevel.MaxLevel {
		Business.Upgradable = true
	}

	b.Chips = User.Chips
	b.Golds = User.Golds
}

// UpdateLeftTimeSafe updates left time only for Business
func UpdateLeftTimeSafe(Business *model.Business) {
	sp := Business.SafeLevel.SafeParts
	sum := 0
	for i := 0; i < len(sp); i++ {
		if i == len(sp)-1 {
			if time.Now().After(sp[i].EndTime) {
				Business.SafeLevel.Sum = Business.SafeLevel.CurrentCapacity
				//Business.SafeLevel.LeftTime = []int{0, 0}
			} else {
				dt := time.Now().Sub(sp[i].StartTime)
				sf := dt.Seconds() * sp[i].Velocity //already collected

				sum = sum + int(sf)
				Business.SafeLevel.Sum = sum

				//td := sp[i].EndTime.Sub(sp[i].StartTime) //total needed time

				log.Print(Business.SafeLevel.CurrentCapacity)
				log.Print(sum)
				//Business.SafeLevel.LeftTime = []int{int(dt.Seconds()), int(td.Seconds())}
			}
		} else {
			sum = sum + sp[i].Collect
		}
	}
}

// UpdateCollectSumLast updates collect sum for last element in safe parts
func UpdateCollectSumLast(Business *model.Business) {
	sl := Business.SafeLevel
	sp := sl.SafeParts
	i := len(sp) - 1
	sp[i].Collect = int(time.Now().Sub(sp[i].StartTime).Seconds() * sl.CurrentVelocity)
}

// GetLeftCapacity gives left capacity
func GetLeftCapacity(Business *model.Business) int {
	sl := Business.SafeLevel
	sp := sl.SafeParts
	curC := sl.CurrentCapacity
	for i := 0; i < len(sp); i++ {
		curC = curC - sp[i].Collect
	}
	return curC
}

// GetLeftTimeSafe gives left time of safe
func GetLeftTimeSafe(Business *model.Business) int {
	sp := Business.SafeLevel.SafeParts
	lp := sp[len(sp)-1]
	dTime := lp.EndTime.Sub(time.Now())
	totalLeftTime := int(dTime.Seconds())
	return totalLeftTime
}

// Itob integer to bool
func Itob(Number int) bool {
	if Number == 1 {
		return true
	}
	return false
}

// PageItemActive verifies if page item active or not
func PageItemActive(LocalLevel, RaiserLevel, MaxLevel int) bool {
	if LocalLevel+RaiserLevel <= MaxLevel {
		return true
	}
	return false
}

// CheckBusinessExist checks if the user have a business in list
func CheckBusinessExist(UserToken, UTBusiness string) bool {
	if UTBusiness == "" {
		return true
	}
	if UserToken == UTBusiness {
		return true
	}
	return false
}
