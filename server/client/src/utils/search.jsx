

export function searchKeyword(items, keyword) {
    return items.filter(entry => entry.entry_name.toUpperCase().includes(keyword.toUpperCase()));
};

// export function searchKeywordNotes(items, keyword) {
//     return items.filter(note => {
//         if (note.type === "simple-note") {
//             console.log("Note", note.text.toUpperCase().includes(keyword.toUpperCase()));
//             return note.text.toUpperCase().includes(keyword.toUpperCase())
//         }
//         if (note.type === "checklist") {
//             const itemsSearched = note.list.map(item => item.text.toUpperCase().includes(keyword.toUpperCase()));
//             if (itemsSearched.length > 0 ) {
//                 return true;
//             }
//         }
//     });
// };

