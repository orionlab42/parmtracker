import http from "./httpService";

const  apiEndpointRegister = '/register';
const  apiEndpointLogin = '/login';
const  apiEndpointGetUser = '/user';
const  apiEndpointLogout = '/logout';
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

