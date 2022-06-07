package filters

import (
	"fmt"
	"github.com/orionlab42/parmtracker/data/categories"
	"github.com/orionlab42/parmtracker/data/expenses"
	"github.com/orionlab42/parmtracker/data/users"
	"sort"
	"strconv"
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
func GetExpenseEntriesMergedByWeek(filter int) expenses.Expenses {
	expensesAll := expenses.GetExpenseEntries()
	var expensesNew expenses.Expenses
	startDate, endDate := GetFilterDateLastTwoYears()
	for _, val := range expensesAll {
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
func GetExpenseEntriesMergedByMonth(filter int) expenses.Expenses {
	expensesAll := expenses.GetExpenseEntries()
	var expensesNew expenses.Expenses
	startDate, endDate := GetFilterDateLastTwoYears()
	for _, val := range expensesAll {
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
func GetExpenseEntriesMergedByDate() expenses.Expenses {
	expensesAll := expenses.GetExpenseEntries()
	var expensesNew expenses.Expenses
	for _, val := range expensesAll {
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
func GetExpenseEntriesMergedByCategory(filter string, exp expenses.Expenses) expenses.Expenses {
	if len(exp) == 0 {
		exp = expenses.GetExpenseEntries()
	}
	var expensesNew expenses.Expenses
	startDate, endDate := GetFilterDate(filter)
	for _, val := range exp {
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
func GetExpenseEntriesPieByCategory(filter string) expenses.Expenses {
	expensesByCategory := GetExpenseEntriesMergedByCategory(filter, []expenses.ExpenseEntry{})
	var totalExpenses float64
	for _, val := range expensesByCategory {
		totalExpenses += val.Amount
	}
	for i, _ := range expensesByCategory {
		expensesByCategory[i].Amount = expensesByCategory[i].Amount * 100 / totalExpenses
	}
	return expensesByCategory
}

// CompleteEntriesMergedByCategories orders the slices of entries of the argument by the category ids in ascending order
// and completed with an extra entry with value 0, for each category which id is in the catIds list but has no entry in entries,
func CompleteEntriesMergedByCategories(entries expenses.Expenses, catIds []int) expenses.Expenses {
	var orderedEntries expenses.Expenses
	for _, id := range catIds {
		isCategoryHere := false
		for _, val := range entries {
			if val.Category == id {
				orderedEntries = append(orderedEntries, val)
				isCategoryHere = true
				break
			}
		}
		if isCategoryHere == false {
			orderedEntries = append(orderedEntries, expenses.ExpenseEntry{Name: "Zero entry for category: " + strconv.Itoa(id), Amount: 0.0, Category: id})
		}
	}
	return orderedEntries
}

// GetAllCategoryIdsFromFilledUsers returns a slice of all the ids of the categories present in the map which has the
// entries by users. The ids are ordered in ascending order.
func GetAllCategoryIdsFromFilledUsers(entries map[int]expenses.Expenses) []int {
	catIds := make(map[int]int)
	for _, val := range entries {
		for _, entries := range val {
			catIds[entries.Category]++
		}
	}
	var categoriesIds []int
	for id, _ := range catIds {
		categoriesIds = append(categoriesIds, id)
	}
	sort.Slice(categoriesIds, func(i, j int) bool {
		return categoriesIds[i] < categoriesIds[j]
	})
	return categoriesIds
}

// GetExpenseEntriesForEachUser returns a map where the keys are the ids of the users and the values are slices of entries
// from that specific user. The map will contain only users that have an expense entry in the expenses table.
func GetExpenseEntriesForEachUser() map[int]expenses.Expenses {
	expensesAll := expenses.GetExpenseEntries()
	filledUsers := GetFilledUsers()
	entriesByUser := make(map[int]expenses.Expenses)
	for _, val := range expensesAll {
		for _, user := range filledUsers {
			if val.UserId == user.UserId {
				entriesByUser[user.UserId] = append(entriesByUser[user.UserId], val)
			}
		}
	}
	return entriesByUser
}

type SeriesByUserAll struct {
	UserId int          `json:"user_id"`
	Series SeriesByUser `json:"series"`
}

type SeriesByUser struct {
	Name       string    `json:"name"`
	Data       []float64 `json:"data"`
	Categories []string  `json:"categories"`
}

// GetExpenseEntriesToSeriesByUser returns for each user the series what is necessary for highcharts. Each user has
// a data field which is a slice of all the values of the merged categories and a categories field which is a slice of all
// the names of the merged categories. The order in the two slices of course is the same. If there is a category which has
// an entry only for one of the users, the category will still appear for all users with the value 0.
// Example of the result: [{UserId:2 Series:{Name:Orion Data:[16 13.95 0 12.95] Categories:[groceries gift services leisure]}}, {UserId:3 Series:{Name:Atik Data:[199 24.27 21.9 0] Categories:[groceries gift services leisure]}}
func GetExpenseEntriesToSeriesByUser(filter string) []SeriesByUserAll {
	entriesForEachUser := GetExpenseEntriesForEachUser()
	catIds := GetAllCategoryIdsFromFilledUsers(entriesForEachUser)
	var seriesByUserAll []SeriesByUserAll
	for id, val := range entriesForEachUser {
		mergedEntriesOneUser := GetExpenseEntriesMergedByCategory(filter, val)
		mergedEntriesOneUser = CompleteEntriesMergedByCategories(mergedEntriesOneUser, catIds)
		var data []float64
		var cat []string
		for _, entries := range mergedEntriesOneUser {
			data = append(data, entries.Amount)
			cat = append(cat, categories.GetCategoryName(entries.Category))
		}
		seriesByUser := SeriesByUser{Name: users.GetUserName(id), Data: data, Categories: cat}
		seriesByUserAll = append(seriesByUserAll, SeriesByUserAll{UserId: id, Series: seriesByUser})
	}
	return seriesByUserAll
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
	exp := expenses.GetExpenseEntries()
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

// GetFilledUserIds creates a map with only the users which already have entries saved. The key will be the id
// of the user and the value is the count of how many entries are of this user.
func GetFilledUserIds() map[int]int {
	expensesAll := expenses.GetExpenseEntries()
	ids := make(map[int]int)
	for _, val := range expensesAll {
		ids[val.UserId]++
	}
	return ids
}

// GetFilledUsers returns a slice of struct User with only the users which already have entries saved.
func GetFilledUsers() users.Users {
	allUsers := users.GetUsers()
	ids := GetFilledUserIds()
	u := users.Users{}
	for _, val := range allUsers {
		for id, _ := range ids {
			if val.UserId == id {
				u = append(u, val)
			}
		}
	}
	return u
}

// GetExpenseEntriesByDate returns a slice of struct ExpenseEntry (a struct of Expenses) within a time frame(filter) all expense entries from the expenses table.
func GetExpenseEntriesByDate(filter string) expenses.Expenses {
	expensesAll := expenses.GetExpenseEntries()
	var expensesNew expenses.Expenses
	startDate, endDate := GetFilterDate(filter)
	for _, val := range expensesAll {
		if startDate.Before(val.Date) && endDate.After(val.Date) {
			expensesNew = append(expensesNew, val)
		}
	}
	return expensesNew
}
