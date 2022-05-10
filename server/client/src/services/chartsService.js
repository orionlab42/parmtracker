import http from './httpService';

const  apiEndpointExpByDate = '/charts-expenses-by-date';
const  apiEndpointExpByCat = '/charts-expenses-by-category';
const  apiEndpointExpByWeek = '/charts-expenses-by-week';
const  apiEndpointExpByMonth = '/charts-expenses-by-month';
const  apiEndpointExpPieByCat = '/charts-pie-expenses-by-category';


export function getEntriesByDate() {
    return http.get(apiEndpointExpByDate);
}

export function getEntriesByWeek(filter) {
    return http.get(`${apiEndpointExpByWeek}/${filter}`);
}

export function getEntriesByMonth(filter) {
    return http.get(`${apiEndpointExpByMonth}/${filter}`);
}

export function getEntriesByCategory(filter) {
    console.log("Chart filter:", filter);
    return http.get(`${apiEndpointExpByCat}/${filter}`);
}

export function getEntriesPieByCategory(filter) {
    return http.get(`${apiEndpointExpPieByCat}/${filter}`);
}
