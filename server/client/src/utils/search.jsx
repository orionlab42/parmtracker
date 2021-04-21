

export function searchKeyword(items, keyword) {
    return items.filter(entry => entry.title.toUpperCase().includes(keyword.toUpperCase()));
};