package filters_test

import (
	"fmt"
	"github.com/orionlab42/parmtracker/data/expenses"
	"github.com/orionlab42/parmtracker/data/filters"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetFilterDateAll(t *testing.T) {
	_, endCurrentDay := filters.GetFilterDateAll()
	now := time.Now()
	assert.Equal(t, now.Day()+1, endCurrentDay.Day())
}

func TestGetFilterDateCurrentWeek(t *testing.T) {
	startCurrentWeek, endCurrentDay := filters.GetFilterDateCurrentWeek()
	now := time.Now()
	assert.Equal(t, time.Monday, startCurrentWeek.Weekday())
	assert.Equal(t, now.Day()+1, endCurrentDay.Day())
}

func TestGetFilterDateCurrentMonth(t *testing.T) {
	startCurrentMonth, endCurrentDay := filters.GetFilterDateCurrentMonth()
	now := time.Now()
	assert.Equal(t, 1, startCurrentMonth.Day())
	assert.Equal(t, now.Day()+1, endCurrentDay.Day())
}

func TestGetFilterDateCurrentYear(t *testing.T) {
	startCurrentYear, endCurrentDay := filters.GetFilterDateCurrentYear()
	now := time.Now()
	assert.Equal(t, time.Month(1), startCurrentYear.Month())
	assert.Equal(t, 1, startCurrentYear.Day())
	assert.Equal(t, now.Day()+1, endCurrentDay.Day())
}

func TestGetFilterDateLastWeek(t *testing.T) {
	startLastWeek, endLastWeek := filters.GetFilterDateLastWeek()
	assert.Equal(t, time.Monday, startLastWeek.Weekday())
	assert.LessOrEqual(t, 7.0, time.Now().Sub(startLastWeek).Hours()/24)
	assert.GreaterOrEqual(t, 7.0, time.Now().Sub(endLastWeek).Hours()/24)
}

func TestGetFilterDateLastMonth(t *testing.T) {
	startLastMonth, endLastMonth := filters.GetFilterDateLastMonth()
	now := time.Now()
	assert.Equal(t, 1, startLastMonth.Day())
	assert.Equal(t, now.Month()-1, startLastMonth.Month())
	assert.Equal(t, now.Month(), endLastMonth.Month())
}

func TestGetFilterDateLastYear(t *testing.T) {
	startLastYear, endLastYear := filters.GetFilterDateLastYear()
	now := time.Now()
	assert.Equal(t, 1, startLastYear.Day())
	assert.Equal(t, time.Month(1), startLastYear.Month())
	assert.Equal(t, now.Year()-1, startLastYear.Year())
	assert.Equal(t, 1, endLastYear.Day())
	assert.Equal(t, time.Month(1), endLastYear.Month())
	assert.Equal(t, now.Year(), endLastYear.Year())
}

func TestGetFilterDateLastTwoYears(t *testing.T) {
	startLastYear, endNow := filters.GetFilterDateLastTwoYears()
	now := time.Now()
	fmt.Println(startLastYear)
	fmt.Println(endNow)
	assert.Equal(t, now.Day(), startLastYear.Day())
	assert.Equal(t, now.Month(), startLastYear.Month())
	assert.Equal(t, now.Year()-2, startLastYear.Year())
	assert.Equal(t, now.Day()+1, endNow.Day())
	assert.Equal(t, now.Month(), endNow.Month())
	assert.Equal(t, now.Year(), endNow.Year())
}

//func TestGetExpenseEntriesMergedByWeek(t *testing.T) {
//	entries := GetExpenseEntriesMergedByWeek(1)
//	for _, entry := range entries {
//		fmt.Printf("%+v\n", entry)
//	}
//	fmt.Println(entries)
//}

//func TestGetExpenseEntriesMergedByMonth(t *testing.T) {
//	entries := GetExpenseEntriesMergedByMonth(1)
//	//for _, entry := range entries {
//	//	fmt.Printf("%+v\n", entry)
//	//}
//	fmt.Println(entries)
//}

//func TestGetExpenseEntriesMergedByDate(t *testing.T) {
//	entries := GetExpenseEntriesMergedByDate()
//	for _, entry := range entries {
//		fmt.Printf("%+v\n", entry)
//	}
//}

func TestGetExpenseEntriesMergedByCategory(t *testing.T) {
	entries := filters.GetExpenseEntriesMergedByCategory(filters.LastMonth, expenses.Expenses{})
	now := time.Now()
	assert.Equal(t, now.Month()-1, entries[0].Date.Month())
	for _, entry := range entries {
		fmt.Printf("%+v\n", entry)
	}
}

func TestGetExpenseEntriesPieByCategory(t *testing.T) {
	entries := filters.GetExpenseEntriesPieByCategory(filters.LastMonth)
	var total float64
	for _, entry := range entries {
		total += entry.Amount
	}
	//assert.Equal(t, 100.00, total)
	fmt.Println(total)
	for _, entry := range entries {
		fmt.Printf("%+v\n", entry)
	}
}

//func TestFirstDayOfISOWeek(t *testing.T) {
//	week := FirstDayOfISOWeek(2021, 52)
//	fmt.Println(week)
//	fmt.Println(week.Format("2006-01-02"))
//	fmt.Println(week.Format("02/01/06"))
//}

func TestGetFilledUserIds(t *testing.T) {
	userIds := filters.GetFilledUserIds()
	fmt.Println(userIds)
	for i, u := range userIds {
		fmt.Printf("%+v\n", i)
		fmt.Printf("%+v\n", u)
	}
}

func TestGetFilledUsers(t *testing.T) {
	users := filters.GetFilledUsers()
	fmt.Println(users)
	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}
}

func TestGetAllCategoryIdsFromFilledUsers(t *testing.T) {
	expensesAll := filters.GetExpenseEntriesForEachUser()
	catIds := filters.GetAllCategoryIdsFromFilledUsers(expensesAll)
	for _, val := range catIds {
		fmt.Printf("%+v\n", val)
	}
}

func TestOrderAndCompleteEntriesByCategories(t *testing.T) {
	catIds := []int{1, 2, 3, 9, 12, 17}
	entries := filters.GetExpenseEntriesMergedByCategory(filters.LastMonth, expenses.Expenses{})
	orderedExpenses := filters.CompleteEntriesMergedByCategories(entries, catIds)
	for _, val := range orderedExpenses {
		fmt.Printf("%+v\n", val)
	}
}

func TestGetExpenseEntriesForEachUser(t *testing.T) {
	expensesAll := filters.GetExpenseEntriesForEachUser()
	for i, val := range expensesAll {
		fmt.Println("Id ", i)
		for _, val2 := range val {
			fmt.Printf("%+v\n", val2)
		}
	}
}

func TestConvertExpenseEntriesToSeriesByUser(t *testing.T) {
	expensesAll := filters.GetExpenseEntriesToSeriesByUser("")
	for _, val := range expensesAll {
		fmt.Printf("%+v\n", val)
	}
}

//func TestGetExpenseEntriesByDate(t *testing.T) {
//	entries := filters.GetExpenseEntriesByDate("")
//	fmt.Println("Length: ", len(entries))
//	for _, entry := range entries {
//		fmt.Printf("%+v\n", entry)
//	}
//}
