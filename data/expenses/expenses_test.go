package expenses

//func TestExpenseInsertAndFetch(t *testing.T) {
//	entry := ExpenseEntry{
//		Name:      "Dinner A",
//		Amount:    25,
//		Category:  1,
//		UserId:    1,
//		Shared:    true,
//		Date:      time.Now().UTC(),
//		CreatedAt: time.Now().UTC(),
//		UpdatedAt: time.Now().UTC(),
//	}
//	e := entry.Insert()
//	assert.Nil(t, e)
//	entries := GetExpenseEntries()
//	assert.NotEqual(t, len(entries), 0)
//	e = entry.Delete()
//	assert.Nil(t, e)
//}
//
//func TestExpenseSave(t *testing.T) {
//	entry := ExpenseEntry{
//		Name:      "Dinner A",
//		Amount:    25,
//		Category:  1,
//		UserId:    1,
//		Shared:    true,
//		Date:      time.Now().UTC(),
//		CreatedAt: time.Now().UTC(),
//		UpdatedAt: time.Now().UTC(),
//	}
//	e := entry.Insert()
//	assert.Nil(t, e)
//	entry.Name = "Mercadona"
//	entry.Save()
//	assert.Equal(t, entry.Name, "Mercadona")
//	e = entry.Delete()
//	assert.Nil(t, e)
//}

//func TestExpenseSaveSeedData(t *testing.T) {
//	layout := "2006-01-02T15:04:05.000Z"
//	strStart := "2020-12-12T08:00:00.371Z"
//	t1, _ := time.Parse(layout, strStart)
//	table := []ExpenseEntry{
//		{Name: "Weekly big food", Amount: 22.8, Date: t1},
//		{Name: "Weekly big food", Amount: 11.65, Category: 1, Date: time.Now().UTC()},
//		{Name: "Sushi Wednesday", Amount: 17.87, Category: 2, Date: time.Now().UTC()},
//		{Name: "Lunch", Amount: 2.7, Category: 2, Date: time.Now().UTC()},
//		{Name: "B-day gift Roser", Amount: 35, Category: 3, Date: time.Now().UTC()},
//		{Name: "Weekly big food", Amount: 47.66, Category: 1, Date: time.Now().UTC()},
//		{Name: "Weekly big food", Amount: 1.1, Category: 1, Date: time.Now().UTC()},
//		{Name: "Burger Wednesday", Amount: 32.5, Category: 2, Date: time.Now().UTC()},
//		{Name: "Lunch", Amount: 8.9, Category: 2, Date: time.Now().UTC()},
//		{Name: "B-day gift Gabri", Amount: 102.77, Category: 3, Date: time.Now().UTC()}}
//
//	for _, row := range table {
//		row.Insert()
//	}
//	entries := GetExpenseEntries()
//	assert.Equal(t, len(entries), 10)
//}

//func TestGetExpenseEntries(t *testing.T) {
//	entries := GetExpenseEntries()
//	for _, entry := range entries {
//		fmt.Printf("%+v\n", entry)
//	}
//}
