package utils

import (
	"errors"
	"os"
	"strings"
	"time"
)

const (
	NotAValidEventTypeError string = "The event type received does not exist"
	NotAValidProductError   string = "The product received does not exist"
)

const (
	DBOpenMsg       string = "Trying to open the DB '%s'..."
	DBOpenedSuccess string = "DB successfully opened"
	DBOpenedFailure string = "DB does not exist and/or could not be created"
	DbTableMsg      string = "Checking 'Event' table"
	DBTableSuccess  string = "'Event' table OK"
	DBTableFailure  string = "Could not create 'Event' table"
	DBReadyMsg		string = "DB ready for use"
	DBCloseMsg		string = "Trying to close the DB..."
	DBCloseSuccess  string = "DB closed successfully"
	DBCloseFailure  string = "Could not close the DB; Proceed carefuly"
	DBNotOpen		string = "DB was not correctly opened and could not be closed"
	ApiHandlerSetup string = "Setting up the API endpoints..."
	ApiConnOpened   string = "Connection opened. Listening to %s"
	ApiConnClosed   string = "Connection closed. Exited"
	ApiUUIDFailed   string = "NO ID - Failed to generate one"
	ApiUUIDFmt		string = "(UUID: %s)"
	ApiIntercepting string = "Intercepting request %s"
	ApiReqMethodNA  string = "Request method not allowed %s"
	ApiReqBodyProb  string = "Request body could not be interpreted and/or was empty %s"
)

var Products = []string{
	"leoalvarenga", "jenifferlaila", "darkness-within",
}

var EventTypes = []string{
	"visit",
}

func IsThisAProduct(product string) bool {
	for i := 0; i < len(Products); i++ {
		if Products[i] == product {
			return true
		}
	}

	return false
}

func IsThisAnEventType(event string) bool {
	for i := 0; i < len(EventTypes); i++ {
		if EventTypes[i] == event {
			return true
		}
	}

	return false
}

func DoesThisFileExist(filename string) bool {
	_, err := os.Stat(filename)

	return !errors.Is(err, os.ErrNotExist)
}

func NormalizeDatetime(datetime time.Time) time.Time {
	return time.Date(datetime.Year(), datetime.Month(), datetime.Day(), 0, 0, 0, 0, datetime.Location())
}

func GetTodaysDate() time.Time {
	return NormalizeDatetime(time.Now())
}

func TimeToSQLDateTime(t time.Time) string {
	res := t.Format(time.RFC3339)
	res = strings.Replace(res, "T", "", -1)

	res = strings.Split(res, "Z")[0]

	return res
}
