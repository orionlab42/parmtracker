import http from './httpService';

const  apiEndpointExpByDate = '/charts-expenses-by-date';
const  apiEndpointExpByCat = '/charts-expenses-by-category';
const  apiEndpointExpByWeek = '/charts-expenses-by-week';
const  apiEndpointExpByMonth = '/charts-expenses-by-month';
const  apiEndpointExpPieByCat = '/charts-pie-expenses-by-category';


export function getEntriesByDate() {
    return http.get(apiEndpointExpByDate);
}

export function getEntriesByCategory() {
    return http.get(apiEndpointExpByCat);
}

export function getEntriesByWeek() {
    return http.get(apiEndpointExpByWeek);
}

export function getEntriesByMonth() {
    return http.get(apiEndpointExpByMonth);
}

export function getEntriesPieByCategory() {
    return http.get(apiEndpointExpPieByCat);
}