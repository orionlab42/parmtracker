import http from './httpService';

const  apiEndpoint = '/notes';

function noteUrl(id) {
    return `${apiEndpoint}/${id}`;
}

export function getNotes() {
    return http.get(apiEndpoint);
}

export function deleteNote(id) {
    return http.delete(noteUrl(id));
}

// export function getNote(id) {
//     return http.get(noteUrl(id));
// }

export function saveNote(note) {
    // if (typeof note.note_id !== "string") {
    if (note.note_id) {
        const body = { ...note };
        delete body.id;
        console.log("Service update", body);
        return http.put(noteUrl(note.note_id), body);
    }
    console.log("Service save", note);
    return http.post(apiEndpoint, note);
}