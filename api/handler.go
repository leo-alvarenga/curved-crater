package api

import (
	"curved-crater/repository"
	"curved-crater/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func visitHandler(analytics *repository.AnalyticsInstance) http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request) {
		uuid, err := uuid.NewRandom()
		id := ""

		if err != nil {
			id = utils.ApiUUIDFailed
		} else {
			id = uuid.String()
		}

		id = fmt.Sprintf(utils.ApiUUIDFmt, id)

		analytics.Event.Log(utils.Low, fmt.Sprintf(utils.ApiIntercepting, id))
		
		if req.Method != http.MethodPost {
			analytics.Event.Log(utils.Low, fmt.Sprintf(utils.ApiReqMethodNA, id))
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		
		
		var analytic Analytic
		err = json.NewDecoder(req.Body).Decode(&analytic)

		if err != nil {
			analytics.Event.Log(utils.Low, fmt.Sprintf(utils.ApiReqBodyProb, id))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		analytics.InsertEvent(analytic.EventType, analytic.Product)
		w.WriteHeader(http.StatusAccepted)
	}
}

func Api(address string, analytics *repository.AnalyticsInstance) {
	analytics.Event.Log(utils.Low, utils.ApiHandlerSetup)
	http.HandleFunc("/visit", visitHandler(analytics))

	connectionOpenMsg := fmt.Sprintf(utils.ApiConnOpened, address)

	println(connectionOpenMsg)
	analytics.Event.Log(utils.Low, connectionOpenMsg)
	
	http.ListenAndServe(address, nil)
	analytics.Event.Log(utils.Low, utils.ApiConnClosed)
}
