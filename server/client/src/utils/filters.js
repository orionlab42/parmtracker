
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


