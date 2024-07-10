package ddate

import (
	"fmt"
	"time"
)

var dayNames = [...]string{"Sweetmorn", "Boomtime", "Pungenday", "Prickle-Prickle", "Setting Orange"}
var seasonNames = [...]string{"Chaos", "Discord", "Confusion", "Bureaucracy", "The Aftermath"}
var holydaysFive = [...]string{"Mungday", "Mojoday", "Syaday", "Zaraday", "Maladay"}
var holydaysFifty = [...]string{"Chaoflux", "Discoflux", "Confuflux", "Bureflux", "Afflux"}

const (
	daysInYear         = 365
	initialYearPadding = 1166
	hailEris           = daysInYear / len(seasonNames)
)

func Format(now time.Time) string {
	return fmt.Sprintf("%s %s %d", dayNames[(now.YearDay()-1)%hailEris%5], seasonNames[(now.YearDay()-1)/hailEris], now.Year()+initialYearPadding)
}

func IsHolyday(now time.Time) (bool, string) {
	dayNumber := now.YearDay() % 73
	if dayNumber == 5 {
		return true, holydaysFive[now.YearDay()%5]
	}

	if dayNumber == 50 {
		return true, holydaysFifty[now.YearDay()%5]
	}

	if int(now.Month()) == 2 && now.Day() == 29 {
		return true, "St. Tib's Day"
	}

	return false, ""
}
