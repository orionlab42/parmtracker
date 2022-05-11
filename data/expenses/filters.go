package expenses

import (
	"fmt"
	"time"
)

const (
	CurrentWeek  = "Current week"
	CurrentMonth = "Current month"
	CurrentYear  = "Current year"
	LastWeek     = "Last week"
	LastMonth    = "Last month"
	LastYear     = "Last year"
)

func GetExpenseEntriesMergedByCategory(filter string) Expenses {
	expenses := GetExpenseEntries()
	var expensesNew Expenses
	startDate, endDate := GetFilterDate(filter)
	for _, val := range expenses {
		isSaved := false
		for i, _ := range expensesNew {
			if val.Category == expensesNew[i].Category && startDate.Before(val.Date) && endDate.After(val.Date) {
				expensesNew[i].Amount = expensesNew[i].Amount + val.Amount
				isSaved = true
				break
			}
		}
		if isSaved == false && startDate.Before(val.Date) && endDate.After(val.Date) {
			val.Name = "Total expenses of " + fmt.Sprint(val.Category)
			expensesNew = append(expensesNew, val)
		}
	}
	return expensesNew
}

func GetExpenseEntriesPieByCategory(filter string) Expenses {
	expensesByCategory := GetExpenseEntriesMergedByCategory(filter)
	var totalExpenses float64
	for _, val := range expensesByCategory {
		totalExpenses += val.Amount
	}
	for i, _ := range expensesByCategory {
		expensesByCategory[i].Amount = expensesByCategory[i].Amount * 100 / totalExpenses
	}
	return expensesByCategory
}

func GetFilterDate(filter string) (time.Time, time.Time) {
	var startDate, endDate time.Time
	switch filter {
	case CurrentWeek:
		startDate, endDate = GetFilterDateCurrentWeek()
	case CurrentMonth:
		startDate, endDate = GetFilterDateCurrentMonth()
	case CurrentYear:
		startDate, endDate = GetFilterDateCurrentYear()
	case LastWeek:
		startDate, endDate = GetFilterDateLastWeek()
	case LastMonth:
		startDate, endDate = GetFilterDateLastMonth()
	case LastYear:
		startDate, endDate = GetFilterDateLastYear()
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
