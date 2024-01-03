package repository

import (
	"curved-crater/utils"
	"fmt"
	"time"
)

func TableCreateQuery() string {
	types := ""
	for i, t := range utils.EventTypes {
		types = fmt.Sprintf("%s type = \"%s\"", types, t)

		if i < len(utils.EventTypes) - 1 {
			types = fmt.Sprintf("%s OR ", types)
		}
	}

	products := ""
	for i, p := range utils.Products {
		products = fmt.Sprintf("%s product = \"%s\"", products, p)

		if i < len(utils.Products) - 1 {
			products = fmt.Sprintf("%s OR ", products)
		}
	}

	return fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS Event (
			id INTEGER NOT NULL PRIMARY KEY,
			day DATETIME NOT NULL,
			timestamp DATETIME NOT NULL,
			type TEXT NOT NULL CHECK(%s) DEFAULT "%s",
			product TEXT NOT NULL CHECK(%s) DEFAULT "%s"
		);
	`, types, utils.EventTypes[0], products, utils.Products[0])
}

func InsertEventQuery(eventType string, product string) string {
	return fmt.Sprintf(`
		INSERT INTO Event (day, timestamp, type, product) VALUES(
			'%s', '%s', '%s', '%s'
		);
	`, utils.TimeToSQLDateTime(utils.GetTodaysDate()), utils.TimeToSQLDateTime(time.Now()), eventType, product)
}
