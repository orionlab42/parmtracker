import axios from 'axios';
import {toast} from 'react-toastify';

axios.defaults.baseURL = process.env.REACT_APP_API_URL;

// axios.interceptors.response.use(success, error)
axios.interceptors.response.use(null, error => {
    const expectedError = error.response &&
        error.response.status >= 404 &&
        error.response.status < 500

    if (!expectedError) {
        console.log('Logging the error', error);
        toast('An unexpected error occurred.');
    }

    return Promise.reject(error);
});

// eslint-disable-next-line import/no-anonymous-default-export
export  default {
    get: axios.get,
    post: axios.post,
    put: axios.put,
    delete: axios.delete,
};