import http from "./httpService";

const  apiEndpoint = '/user';

function userUrl(id) {
    return `${apiEndpoint}/${id}`;
}

export function getUsers() {
    return http.get(apiEndpoint);
}

export function deleteUser(id) {
    return http.delete(userUrl(id));
}

export function getUser(id) {
    return http.get(userUrl(id));
}

export function saveUser(user) {
    if (user.id) {
        const body = { ...user };
        delete body.id;
        return http.put(userUrl(user.id), body);
    }
    return http.post(apiEndpoint, user);
}

