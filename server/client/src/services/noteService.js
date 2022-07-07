import http from './httpService';

const  apiEndpointNotes = '/notes';
const  apiEndpointItem = '/notes/item';
const  apiEndpointItems = '/notes/items';

function noteUrl(id) {
    return `${apiEndpointNotes}/${id}`;
}

function itemUrl(id) {
    return `${apiEndpointItem}/${id}`;
}

function itemsUrl(id) {
    return `${apiEndpointItems}/${id}`;
}

export function getNotes() {
    return http.get(apiEndpointNotes);
}

export function deleteNote(id) {
    return http.delete(noteUrl(id));
}

// export function getNote(id) {
//     return http.get(noteUrl(id));
// }

export function saveNote(note) {
    if (typeof note.note_id !== "string") {
    // if (note.note_id) {
        const body = { ...note };
        delete body.id;
        // console.log("Service update", body);
        return http.put(noteUrl(note.note_id), body);
    }
    // console.log("Service save", note);
    return http.post(apiEndpointNotes, note);
}

export function saveItem(item) {
    if (typeof item.item_id !== "string") {
    // if (note.note_id) {
        const body = { ...item };
        delete body.id;
        // console.log("Service update title", body);
        return http.put(itemUrl(item.item_id), body);
    }
    // console.log("Service save", item);
    return http.post(apiEndpointItem, item);
}

export function deleteItem(id) {
    return http.delete(itemUrl(id));
}

export function getItems(noteId) {
    return http.get(itemsUrl(noteId));
}

export function deleteItems(noteId) {
    return http.delete(itemsUrl(noteId));
}
