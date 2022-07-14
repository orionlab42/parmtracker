import http from './httpService';

const  apiEndpointNotes = '/notes';
const  apiEndpointItem = '/note_item';
const  apiEndpointItems = '/note_items';
const  apiEndpointNoteUser = '/notes_user';

function noteUrl(id) {
    return `${apiEndpointNotes}/${id}`;
}

function itemUrl(id) {
    return `${apiEndpointItem}/${id}`;
}

function itemsUrl(id) {
    return `${apiEndpointItems}/${id}`;
}

function noteUserUrl(noteId, userId) {
    return `${apiEndpointNoteUser}/${noteId}/${userId}`;
}

export function getNotes(user_id) {
    return http.get(noteUrl(user_id));
}

export function deleteNote(noteId, userId) {
    let url = `${apiEndpointNotes}/${noteId}/${userId}`;
    return http.delete(url);
}

// export function getNote(id) {
//     return http.get(noteUrl(id));
// }

export function saveNote(note, userId) {
    if (typeof note.note_id !== "string") {
        const body = { ...note };
        delete body.id;
        return http.put(noteUrl(note.note_id), body);
    }
    return http.post(noteUrl(userId), note);
}

export function saveItem(item) {
    if (typeof item.item_id !== "string") {
        const body = { ...item };
        delete body.id;
        return http.put(itemUrl(item.item_id), body);
    }
    return http.post(apiEndpointItem, item);
}

export function deleteItem(id) {
    return http.delete(itemUrl(id));
}

export function getItems(noteId) {
    return http.get(itemsUrl(noteId));
}

export function saveItems(noteId, startDate, endDate) {
    if (startDate !== null && endDate !== null) {
        const params = new URLSearchParams({
            note_id: noteId,
            start_date: startDate.toISOString(),
            end_date: endDate.toISOString(),
        });
        let url = `${apiEndpointItems}/?${params}`;
        return http.get(url);
    } else {
        const params = new URLSearchParams({
            note_id: noteId,
            start_date: startDate,
            end_date: endDate,
        });
        let url = `${apiEndpointItems}/?${params}`;
        return http.get(url);
    }
}

export function deleteItems(noteId) {
    return http.delete(itemsUrl(noteId));
}

export function saveNoteUser(noteId, userId) {
    return http.post(noteUserUrl(noteId, userId));
}
