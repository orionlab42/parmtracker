import http from "./httpService";

const  apiEndpointRegister = '/register';
const  apiEndpointLogin = '/login';
const  apiEndpointGetUser = '/user';
const  apiEndpointLogout = '/logout';
const  apiEndpointAllUsers = '/all-users';
const  apiEndpointUserUpdate = '/user/update-settings/';

export function getUsers() {
    return http.get(apiEndpointAllUsers);
}

export function register(user) {
    return http.post(apiEndpointRegister, {
            user_name: user.username,
            password: user.password,
            email: user.email
        });
}

export function login(user) {
    return http.post(apiEndpointLogin, {
        user_name: user.username,
        password: user.password
    });
}

export async function  getCurrentUser() {
    try {
        const { data: user } = await http.get(apiEndpointGetUser);
        if ((user === "") || (user=== null)) {
          return "";
        } else {
            return user;
        }
    } catch (ex) {
        return "";
    }
}

export function logout() {
    return http.post(apiEndpointLogout, {});
}

function userUrl(id) {
    return `${apiEndpointUserColor}/${id}`;
}

export function updateUserSettings(user) {
    return http.post(apiEndpointUserColor, user);
}

export function saveEntry(entry) {
    entry.category = parseInt(entry.category)
    entry.user_id = parseInt(entry.user_id)
    entry.amount = parseFloat(entry.amount)
    if (entry.id) {
        const body = { ...entry };
        delete body.id;
        body.category = parseInt(body.category)
        body.user_id = parseInt(body.user_id)
        body.amount = parseFloat(body.amount)
        return http.put(entryUrl(entry.id), body);
    }
    return http.post(apiEndpoint, entry);
}

