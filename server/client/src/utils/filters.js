
export function filterByTime(items, filter) {
    let today = new Date();
    switch(filter) {
        case 'Last week':
            let oneWeekAgo = new Date(today.getTime() - 7*24*60*60*1000);
            return  items.filter(entry => oneWeekAgo < new Date(entry.entry_date));
        case 'This month':
            return  items.filter(entry => today.getMonth() === new Date(entry.entry_date).getMonth());
        case 'Last month':
            return  items.filter(entry => today.getMonth() - 1 === new Date(entry.entry_date).getMonth());
        default:
            return items;
    }
}

export function filterByCategory(items, filter) {
    if (filter === 0) return items;
    return items.filter(entry => entry.category === filter);
}

export function filterByUser(items, filter) {
    if (filter === 0) return items;
    return items.filter(entry => entry.user_id === filter);
}

export function filterByShared(items, filter) {
    if (filter === false) return items;
    return items.filter(entry => entry.shared === true);
}