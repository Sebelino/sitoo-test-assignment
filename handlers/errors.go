package handlers

import "log"

const mysqlDuplicateEntry = 1062

func dbToCustomErrorCode(dbErrorCode uint16) int {
	if dbErrorCode > 10000 {
		log.Fatalf("Unexpected database error code: %d", dbErrorCode)
	}
	return 30000 + int(dbErrorCode)
}
