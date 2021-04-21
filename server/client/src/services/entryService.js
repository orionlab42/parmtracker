import http from './httpService';

const  apiEndpoint = '/expenses';


function entryUrl(id) {
    return `${apiEndpoint}/${id}`;
}

export function getEntries() {
    return http.get(apiEndpoint);
}

export function deleteEntry(id) {
    return http.delete(entryUrl(id));
}

export function getEntry(id) {
    return http.get(entryUrl(id));
}

export function saveEntry(entry) {
    entry.category = parseInt(entry.category)
    entry.amount = parseInt(entry.amount)
    if (entry.id) {
        const body = { ...entry };
        delete body.id;
        body.category = parseInt(body.category)
        body.amount = parseInt(body.amount)
        return http.put(entryUrl(entry.id), body);
    }
    return http.post(apiEndpoint, entry);
}

