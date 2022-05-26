package expenses

import (
	"fmt"
	"time"
)

const (
	CurrentWeek  = "current week"
	CurrentMonth = "current month"
	CurrentYear  = "current year"
	LastWeek     = "last week"
	LastMonth    = "last month"
	LastYear     = "last year"
)

// GetExpenseEntriesMergedByWeek returns a slice of struct of ExpenseEntry (struct of Expenses) in which the entries of
// a week are merged into one entry, so for each week it returns one entry with the name being the dates of the week (ex. 23-29 May 22).
// The filter is to select a certain category, if the filter is 0 than all categories are added.
func GetExpenseEntriesMergedByWeek(filter int) Expenses {
	expenses := GetExpenseEntries()
	var expensesNew Expenses
	startDate, endDate := GetFilterDateLastTwoYears()
	for _, val := range expenses {
		isSaved := false
		isFilter := false
		if filter == 0 || val.Category == filter {
			isFilter = true
		}
		year1, week1 := val.Date.ISOWeek()
		for i, _ := range expensesNew {
			year2, week2 := expensesNew[i].Date.ISOWeek()
			if year1 == year2 && week1 == week2 && startDate.Before(val.Date) && endDate.After(val.Date) && isFilter {
				expensesNew[i].Amount = expensesNew[i].Amount + val.Amount
				isSaved = true
				break
			}
		}
		if isSaved == false && startDate.Before(val.Date) && endDate.After(val.Date) && isFilter {
			//val.Name = "Week nr." + fmt.Sprint(week1) + "/" + fmt.Sprint(year1)
			val.Name = fmt.Sprint(FirstDayOfISOWeek(year1, week1).Day()) + "-" + fmt.Sprint(FirstDayOfISOWeek(year1, week1).AddDate(0, 0, 6).Day()) + " " + fmt.Sprint(FirstDayOfISOWeek(year1, week1).Format("Jan 06"))
			expensesNew = append(expensesNew, val)
		}
	}
	return expensesNew
}

// GetExpenseEntriesMergedByMonth returns a slice of struct of ExpenseEntry (struct of Expenses) in which the entries of
// a month are merged into one entry, so for each month it returns one entry with the name being the month and year (ex. May 2022).
// The filter is to select a certain category, if the filter is 0 than all categories are added.
func GetExpenseEntriesMergedByMonth(filter int) Expenses {
	expenses := GetExpenseEntries()
	var expensesNew Expenses
	startDate, endDate := GetFilterDateLastTwoYears()
	for _, val := range expenses {
		isSaved := false
		isFilter := false
		if filter == 0 || val.Category == filter {
			isFilter = true
		}
		for i, _ := range expensesNew {
			if val.Date.Month() == expensesNew[i].Date.Month() && val.Date.Year() == expensesNew[i].Date.Year() && startDate.Before(val.Date) && endDate.After(val.Date) && isFilter {
				expensesNew[i].Amount = expensesNew[i].Amount + val.Amount
				//fmt.Printf("Added from date%v the value %v\n", val.Date, val.Amount)
				//fmt.Printf("With new expense %v\n", expensesNew)
				isSaved = true
				break
			}
		}
		if isSaved == false && startDate.Before(val.Date) && endDate.After(val.Date) && isFilter {
			val.Name = "Total expenses of " + fmt.Sprint(val.Date.Month()) + " " + fmt.Sprint(val.Date.Year())
			//fmt.Printf("Saved as %+v\n", val)
			expensesNew = append(expensesNew, val)
		}
	}
	return expensesNew
}

// GetExpenseEntriesMergedByDate returns a slice of struct of ExpenseEntry (struct of Expenses) in which the entries of
// a day are merged into one entry, so for each day it returns one entry.
// THIS FUNCTION IS NOT IN USE & IT IS FROM AN OLDER VERSION
func GetExpenseEntriesMergedByDate() Expenses {
	expenses := GetExpenseEntries()
	var expensesNew Expenses
	for _, val := range expenses {
		isSaved := false
		for i, _ := range expensesNew {
			if val.Date == expensesNew[i].Date {
				expensesNew[i].Name = "Total expenses of the day: " + fmt.Sprint(expensesNew[i].Date)
				expensesNew[i].Amount = expensesNew[i].Amount + val.Amount
				isSaved = true
				break
			}
		}
		if isSaved == false {
			expensesNew = append(expensesNew, val)
		}
	}
	return expensesNew
}

// GetExpenseEntriesMergedByCategory returns a slice of struct of ExpenseEntry (struct of Expenses) in which the entries of
// a category are merged into one entry, so for each category it returns one entry with the name of the category.
// The filter is to select a certain time frame, if the filter is "" than it gets the last 2 years.
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

// GetExpenseEntriesPieByCategory returns a slice of struct of ExpenseEntry (struct of Expenses) in which the results of
// GetExpenseEntriesMergedByCategory() are turned into a percentage.
// The filter is to select a certain time frame, if the filter is "" than it gets the last 2 years.
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

// GetFilterDate returns for a certain time filter(string) a start and an end date, in case there is no recognized filter
// it will return the oldest expense entry's date from the expenses table as starting date and now as end date.
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
	default:
		startDate, endDate = GetFilterDateAll()
	}
	return startDate, endDate
}

// GetFilterDateAll returns the oldest expense entry's date from the expenses table as starting date and now as end date.
func GetFilterDateAll() (time.Time, time.Time) {
	now := time.Now()
	endCurrentDay := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC)
	exp := GetExpenseEntries()
	startAll := exp[len(exp)-1].Date
	startAll = time.Date(startAll.Year(), startAll.Month(), startAll.Day(), 0, 0, 0, 0, time.UTC)
	return startAll, endCurrentDay
}

// GetFilterDateCurrentWeek returns this Monday's date as starting date and now as end date.
func GetFilterDateCurrentWeek() (time.Time, time.Time) {
	now := time.Now()
	endCurrentDay := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC)
	startCurrentWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	for startCurrentWeek.Weekday() != time.Monday {
		startCurrentWeek = startCurrentWeek.AddDate(0, 0, -1)
	}
	return startCurrentWeek, endCurrentDay
}

// GetFilterDateCurrentMonth returns this month first day's date as starting date and now as end date.
func GetFilterDateCurrentMonth() (time.Time, time.Time) {
	now := time.Now()
	startCurrentMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	endCurrentDay := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC)
	return startCurrentMonth, endCurrentDay
}

// GetFilterDateCurrentYear returns this year first day's date as starting date and now as end date.
func GetFilterDateCurrentYear() (time.Time, time.Time) {
	now := time.Now()
	startCurrentYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	endCurrentDay := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC)
	return startCurrentYear, endCurrentDay
}

// GetFilterDateLastWeek returns last Monday's date as starting date and this Monday's date as end date.
func GetFilterDateLastWeek() (time.Time, time.Time) {
	now := time.Now()
	endLastWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	for endLastWeek.Weekday() != time.Monday {
		endLastWeek = endLastWeek.AddDate(0, 0, -1)
	}
	startLastWeek := endLastWeek.AddDate(0, 0, -7)
	return startLastWeek, endLastWeek
}

// GetFilterDateLastMonth returns last month first day's date as starting date and this month first day's date as end date.
func GetFilterDateLastMonth() (time.Time, time.Time) {
	now := time.Now()
	startLastMonth := time.Date(now.Year(), now.Month()-1, 1, 0, 0, 0, 0, time.UTC)
	endLastMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	return startLastMonth, endLastMonth
}

// GetFilterDateLastYear returns last year's first day's date as starting date and this year first day's date as end date.
func GetFilterDateLastYear() (time.Time, time.Time) {
	now := time.Now()
	startLastYear := time.Date(now.Year()-1, 1, 1, 0, 0, 0, 0, time.UTC)
	endLastYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	return startLastYear, endLastYear
}

// GetFilterDateLastTwoYears returns two years ago this day's date as starting date and now as end date.
func GetFilterDateLastTwoYears() (time.Time, time.Time) {
	now := time.Now()
	startOfMaxTime := time.Date(now.Year()-2, now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	endCurrentDay := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC)
	return startOfMaxTime, endCurrentDay
}

// FirstDayOfISOWeek returns the date of the Monday of a certain iso week, based on the year and week number.
func FirstDayOfISOWeek(year int, week int) time.Time {
	date := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	isoYear, isoWeek := date.ISOWeek()
	// iterate back to Monday
	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}
	// iterate forward to the first day of the first week
	for isoYear < year {
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}
	// iterate forward to the first day of the given week
	for isoWeek < week {
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}
	return date
}
