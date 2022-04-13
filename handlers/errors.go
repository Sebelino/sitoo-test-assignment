package handlers

import "log"

func dbToCustomErrorCode(dbErrorCode uint16) int {
	if dbErrorCode > 10000 {
		log.Fatalf("Unexpected database error code: %d", dbErrorCode)
	}
	return 30000 + int(dbErrorCode)
}
