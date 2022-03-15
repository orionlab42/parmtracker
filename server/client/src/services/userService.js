import http from "./httpService";

const  apiEndpointRegister = '/register';
const  apiEndpointLogin = '/login"';
const  apiEndpointAllUsers = '/all-users';

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

