package handlers

func dbToCustomErrorCode(dbErrorCode uint16) int {
	return 30000 + int(dbErrorCode)
}
