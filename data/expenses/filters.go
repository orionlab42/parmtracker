package expenses

import (
	"fmt"
	"time"
)

func GetExpenseEntriesMergedByCategory(filter string) Expenses {
	expenses := GetExpenseEntries()
	var expensesNew Expenses
	if filter == "Current week" {
		for _, val := range expenses {
			isSaved := false
			for i, _ := range expensesNew {
				if val.Category == expensesNew[i].Category {
					expensesNew[i].Amount = expensesNew[i].Amount + val.Amount
					isSaved = true
					break
				}
			}
			if isSaved == false {
				val.Name = "Total expenses of " + fmt.Sprint(val.Category)
				expensesNew = append(expensesNew, val)
			}
		}
	}
	return expensesNew
}

func GetFilterDate(filter string) (time.Time, time.Time) {
	var startDate, endDate time.Time
	switch filter {
	case "Current week":
		startDate, endDate = GetFilterDateCurrentWeek()
	case "Current year":
		startDate, endDate = GetFilterDateCurrentYear()
	case "Last week":
		startDate, endDate = GetFilterDateLastWeek()
	case "Last month":
		startDate, endDate = GetFilterDateLastMonth()
	case "Last year":
		startDate, endDate = GetFilterDateLastYear()
	default:
		startDate, endDate = GetFilterDateCurrentMonth()
	}
	return startDate, endDate
}

func GetFilterDateCurrentWeek() (time.Time, time.Time) {
	now := time.Now()
	endCurrentDay := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC)
	startCurrentWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	for startCurrentWeek.Weekday() != time.Monday {
		startCurrentWeek = startCurrentWeek.AddDate(0, 0, -1)
	}
	return startCurrentWeek, endCurrentDay
}

func GetFilterDateCurrentMonth() (time.Time, time.Time) {
	now := time.Now()
	startCurrentMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	endCurrentDay := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC)
	return startCurrentMonth, endCurrentDay
}

func GetFilterDateCurrentYear() (time.Time, time.Time) {
	now := time.Now()
	startCurrentYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	endCurrentDay := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC)
	return startCurrentYear, endCurrentDay
}

func GetFilterDateLastWeek() (time.Time, time.Time) {
	now := time.Now()
	endLastWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	for endLastWeek.Weekday() != time.Monday {
		endLastWeek = endLastWeek.AddDate(0, 0, -1)
	}
	startLastWeek := endLastWeek.AddDate(0, 0, -7)
	return startLastWeek, endLastWeek
}

func GetFilterDateLastMonth() (time.Time, time.Time) {
	now := time.Now()
	startLastMonth := time.Date(now.Year(), now.Month()-1, 1, 0, 0, 0, 0, time.UTC)
	endLastMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	return startLastMonth, endLastMonth
}

func GetFilterDateLastYear() (time.Time, time.Time) {
	now := time.Now()
	startLastYear := time.Date(now.Year()-1, 1, 1, 0, 0, 0, 0, time.UTC)
	endLastYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	return startLastYear, endLastYear
}
