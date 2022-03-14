import http from "./httpService";

const  apiEndpoint = '/auth';

export function login(user) {
    return http.post(apiEndpoint, {
        user_name: user.username,
        password: user.password
    });
}

