

export function searchKeyword(items, keyword) {
    return items.filter(entry => entry.entry_name.toUpperCase().includes(keyword.toUpperCase()));
};

export function searchKeywordNotes(items, keyword) {
    return items.filter(note => note.text.toUpperCase().includes(keyword.toUpperCase()));
};