package expenses

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetFilterDateAll(t *testing.T) {
	_, endCurrentDay := GetFilterDateAll()
	now := time.Now()
	assert.Equal(t, now.Day()+1, endCurrentDay.Day())
}

func TestGetFilterDateCurrentWeek(t *testing.T) {
	startCurrentWeek, endCurrentDay := GetFilterDateCurrentWeek()
	now := time.Now()
	assert.Equal(t, time.Monday, startCurrentWeek.Weekday())
	assert.Equal(t, now.Day()+1, endCurrentDay.Day())
}

func TestGetFilterDateCurrentMonth(t *testing.T) {
	startCurrentMonth, endCurrentDay := GetFilterDateCurrentMonth()
	now := time.Now()
	assert.Equal(t, 1, startCurrentMonth.Day())
	assert.Equal(t, now.Day()+1, endCurrentDay.Day())
}

func TestGetFilterDateCurrentYear(t *testing.T) {
	startCurrentYear, endCurrentDay := GetFilterDateCurrentYear()
	now := time.Now()
	assert.Equal(t, time.Month(1), startCurrentYear.Month())
	assert.Equal(t, 1, startCurrentYear.Day())
	assert.Equal(t, now.Day()+1, endCurrentDay.Day())
}

func TestGetFilterDateLastWeek(t *testing.T) {
	startLastWeek, endLastWeek := GetFilterDateLastWeek()
	assert.Equal(t, time.Monday, startLastWeek.Weekday())
	assert.LessOrEqual(t, 7.0, time.Now().Sub(startLastWeek).Hours()/24)
	assert.GreaterOrEqual(t, 7.0, time.Now().Sub(endLastWeek).Hours()/24)
}

func TestGetFilterDateLastMonth(t *testing.T) {
	startLastMonth, endLastMonth := GetFilterDateLastMonth()
	now := time.Now()
	assert.Equal(t, 1, startLastMonth.Day())
	assert.Equal(t, now.Month()-1, startLastMonth.Month())
	assert.Equal(t, now.Month(), endLastMonth.Month())
}

func TestGetFilterDateLastYear(t *testing.T) {
	startLastYear, endLastYear := GetFilterDateLastYear()
	now := time.Now()
	assert.Equal(t, 1, startLastYear.Day())
	assert.Equal(t, time.Month(1), startLastYear.Month())
	assert.Equal(t, now.Year()-1, startLastYear.Year())
	assert.Equal(t, 1, endLastYear.Day())
	assert.Equal(t, time.Month(1), endLastYear.Month())
	assert.Equal(t, now.Year(), endLastYear.Year())
}

func TestGetFilterDateLastTwoYears(t *testing.T) {
	startLastYear, endNow := GetFilterDateLastTwoYears()
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

//func TestGetExpenseEntriesMergedByCategory(t *testing.T) {
//	entries := GetExpenseEntriesMergedByCategory(CurrentMonth)
//	now := time.Now()
//	assert.Equal(t, now.Month(), entries[0].Date.Month())
//	//for _, entry := range entries {
//	//	fmt.Printf("%+v\n", entry)
//	//}
//}

//func TestGetExpenseEntriesPieByCategory(t *testing.T) {
//	entries := GetExpenseEntriesPieByCategory(LastMonth)
//	var total float64
//	for _, entry := range entries {
//		total += entry.Amount
//	}
//	assert.Equal(t, 100.00, total)
//	//fmt.Println(total)
//}

//func TestFirstDayOfISOWeek(t *testing.T) {
//	week := FirstDayOfISOWeek(2021, 52)
//	fmt.Println(week)
//	fmt.Println(week.Format("2006-01-02"))
//	fmt.Println(week.Format("02/01/06"))
//}
