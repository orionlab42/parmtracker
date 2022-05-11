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
    let entries = http.get(`${apiEndpointExpByWeek}/${filter}`);
    if (entries == null) {
        return [];
    }
    return entries;
}

export function getEntriesByMonth(filter) {
    let entries = http.get(`${apiEndpointExpByMonth}/${filter}`);
    if (entries == null) {
        return [];
    }
    return entries;
}

export function getEntriesByCategory(filter) {
    return http.get(`${apiEndpointExpByCat}/${filter}`);
}

export function getEntriesPieByCategory(filter) {
    return http.get(`${apiEndpointExpPieByCat}/${filter}`);
}
