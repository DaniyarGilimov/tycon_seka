package main

import (
	"log"
	"net/http"
	"points/utils"
	"tycon_seka/handler"
	"tycon_seka/runtime"
	localutils "tycon_seka/utils"

	"github.com/gorilla/mux"
)

func main() {

	go runtime.DoRecovery()

	r := mux.NewRouter()

	// Worker Handler
	r.HandleFunc("/api/v1/get/ceo/{bid}", handler.GetCEOsHandler).
		Methods("GET")
	r.HandleFunc("/api/v1/get/cto/{bid}", handler.GetCTOsHandler).
		Methods("GET")
	r.HandleFunc("/api/v1/get/pr/{bid}", handler.GetPRsHandler).
		Methods("GET")

	r.HandleFunc("/api/v1/hire/ceo/{bid}", handler.HairCEOHandler).
		Methods("POST")
	r.HandleFunc("/api/v1/hire/cto/{bid}", handler.HairCTOHandler).
		Methods("POST")
	r.HandleFunc("/api/v1/hire/pr/{bid}", handler.HairPRHandler).
		Methods("POST")

	//Business handler
	r.HandleFunc("/api/v1/get/businesses", handler.GetBusinessesHandler).
		Methods("GET")
	r.HandleFunc("/api/v1/auth/{bid}", handler.RegisterHandler).
		Methods("POST")
	r.HandleFunc("/api/v1/get/sectors", handler.GetSecotrsHandler).
		Methods("GET")
	r.HandleFunc("/api/v1/get/locations", handler.GetLocationsHandler).
		Methods("GET")
	r.HandleFunc("/api/v1/get/secloc", handler.GetSeclocsHandler).
		Methods("GET")

	r.HandleFunc("/api/v1/get/business/pr/{bid}", handler.GetPPRHandler).
		Methods("GET")
	r.HandleFunc("/api/v1/upgrade/pr/{bid}", handler.UpgradePRHandler).
		Methods("POST")
	r.HandleFunc("/api/v1/get/business/tech/{bid}", handler.GetPTechHandler).
		Methods("GET")
	r.HandleFunc("/api/v1/upgrade/tech/{bid}", handler.UpgradeTechHandler).
		Methods("POST")
	r.HandleFunc("/api/v1/upgrade/safe/{bid}", handler.UpgradeSafeHandler).
		Methods("POST")

	// Collect safe
	r.HandleFunc("/api/v1/collect/safe/{bid}", handler.CollectSafeHandler).
		Methods("GET")

	// Upgrade business
	r.HandleFunc("/api/v1/upgrade/business/{bid}", handler.UpgradeBusinessHandler).
		Methods("POST")

	// Get Statistics API
	r.HandleFunc("/api/v1/get/statistics/{bid}", handler.GetStatisticsHandler).
		Methods("GET")

	// Get and Post Test API
	r.HandleFunc("/api/v1/get/test/{bid}", handler.GetTestHandler).
		Methods("GET")
	r.HandleFunc("/api/v1/post/test/{bid}", handler.PostQuestionsHandler).
		Methods("POST")

	// Image handlers
	r.HandleFunc("/workers/ceo/men", handler.HandleCeoMenImage)
	r.HandleFunc("/workers/ceo/women", handler.HandleCeoWomenImage)
	r.HandleFunc("/workers/pr/men", handler.HandlePrMenImage)
	r.HandleFunc("/workers/pr/women", handler.HandlePrWomenImage)
	r.HandleFunc("/workers/cto/men", handler.HandleCtoMenImage)
	r.HandleFunc("/workers/cto/women", handler.HandleCtoMenImage)

	// new branch

	if err := http.ListenAndServeTLS(localutils.PortAPITLS, utils.SertificateName, utils.SertificateKey, r); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
