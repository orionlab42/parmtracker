

export function searchKeyword(items, keyword) {
    return items.filter(entry => entry.entry_name.toUpperCase().includes(keyword.toUpperCase()));
};