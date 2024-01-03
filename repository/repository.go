package repository

import (
	"curved-crater/utils"
	"database/sql"
	"errors"
	"fmt"
)

type AnalyticsInstance struct {
	db *sql.DB
	Event *utils.Logger
	Analytic *utils.Logger
}

func (a *AnalyticsInstance) OpenDb(databaseFile string, openDb func (name string) (*sql.DB, error)) (*sql.DB, error) {
	a.Event.Log(utils.Low, fmt.Sprintf(utils.DBOpenMsg, databaseFile))
	db, err := openDb(databaseFile)

	if err != nil {
		a.Event.Log(utils.High, utils.DBOpenedFailure, err.Error())
		return nil, err
	} else {
		a.Event.Log(utils.Success, utils.DBOpenedSuccess)
	}

	a.Event.Log(utils.Low, utils.DbTableMsg)
	_, err = db.Exec(TableCreateQuery())

	if err != nil {
		a.Event.Log(utils.High, utils.DBTableFailure, err.Error())
	} else {
		a.Event.Log(utils.Low, utils.DBTableSuccess)
	}
	
	a.db = db
	a.Event.Log(utils.Low, utils.DBReadyMsg)
	return db, nil
}

func (a *AnalyticsInstance) CloseDb() {
	a.Event.Log(utils.Low, utils.DBCloseMsg)
	
	if a != nil && a.db != nil {
		err := a.db.Close()

		if err == nil {
			a.Event.Log(utils.Success, utils.DBCloseSuccess)
		} else {
			a.Event.Log(utils.Medium, utils.DBCloseFailure)
		}
	} else {
		a.Event.Log(utils.Medium, utils.DBNotOpen)
	}
}

func (a *AnalyticsInstance) InsertEvent(eventType string, product string) (bool, error) {
	if !utils.IsThisAnEventType(eventType) {
		a.Event.Log(utils.Medium, utils.NotAValidEventTypeError, eventType, product)
		return false, errors.New(utils.NotAValidEventTypeError)
	}

	if !utils.IsThisAProduct(product) {
		a.Event.Log(utils.Medium, utils.NotAValidProductError, eventType, product)
		return false, errors.New(utils.NotAValidProductError)
	}

	_, err := a.db.Exec(InsertEventQuery(eventType, product))

	if err != nil {
		a.Event.Log(utils.Medium, err.Error(), eventType, product)
		return false, err
	}

	a.Analytic.LogEntry(product, eventType)
	return true, nil
}

func NewAnalyticsInstance(EventLogger *utils.Logger, AnalyticsLogger *utils.Logger) *AnalyticsInstance {
	EventLogger.Log(utils.Low, "Initializing AnalyticsInstance...")
	
	a := new(AnalyticsInstance)

	a.Event = EventLogger
	a.Analytic = AnalyticsLogger

	EventLogger.Log(utils.Success, "AnalyticsInstance successfully initialized")

	return a
}
