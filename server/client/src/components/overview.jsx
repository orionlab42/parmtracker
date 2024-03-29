import React, { useState, useEffect } from "react";
import Highcharts from 'highcharts';
import HighchartsReact from 'highcharts-react-official';
import {getFilledCategories} from "../services/categoryService";
import {
    getEntriesByWeek,
    getEntriesByMonth,
    getEntriesByCategory,
    getEntriesPieByCategory,
    getEntriesByCategoryAndUser
} from "../services/chartsService";
import FilterTime from "./common/filterTime";
import FilterCategory from "./common/filterCategories";


const Overview = (props) => {
    // const [entriesByDate, setEntriesByDate] = useState([]);
    const [entriesByWeek, setEntriesByWeek] = useState([]);
    const [entriesByMonth, setEntriesByMonth] = useState([]);
    const [categories, setCategories] = useState([]);
    const [entriesByCat, setEntriesByCat] = useState([]);
    const [entriesByCatAndUser, setEntriesByCatAndUser] = useState([]);
    const [entriesPieByCat, setEntriesPieByCat] = useState([]);
    const [filterTime, setFilterTime] = useState("get all");
    const [filterTimeForUsers, setFilterTimeForUsers] = useState("get all");
    const [filterCategory, setFilterCategory] = useState(0);

    useEffect( () => {
        async function getEntriesByTime() {
            // const { data: entriesDate } = await getEntriesByDate();
            // setEntriesByDate(entriesDate);
            const { data: entriesWeek } = await getEntriesByWeek(filterCategory);
            if (entriesWeek != null) {
                setEntriesByWeek(entriesWeek);
            }
            const { data: entriesMonth } = await getEntriesByMonth(filterCategory);
            if (entriesMonth != null) {
                setEntriesByMonth(entriesMonth);
            }
        }
        getEntriesByTime();
    }, [filterCategory]);

    useEffect( () => {
        async function getEntriesByCat() {
            const { data } = await getFilledCategories();
            const categories = [{id:0 , category_name: "all categories"}, ...data];
            setCategories(categories);
            const { data: entriesCat } = await getEntriesByCategory(filterTime);
            setEntriesByCat(entriesCat);
            const { data: entriesCatAndUser } = await getEntriesByCategoryAndUser(filterTimeForUsers);
            setEntriesByCatAndUser(entriesCatAndUser);
            const { data: entriesPieByCat } = await getEntriesPieByCategory(filterTime);
            setEntriesPieByCat(entriesPieByCat);
        }
        getEntriesByCat();
    }, [filterTime]);

    useEffect( () => {
        async function getEntriesByCatAndUser() {
            const { data: entriesCatAndUser } = await getEntriesByCategoryAndUser(filterTimeForUsers);
            if (entriesCatAndUser != null) {
                setEntriesByCatAndUser(entriesCatAndUser);
            } else {
                let entriesCatAndUserEmpty = [];
                setEntriesByCatAndUser(entriesCatAndUserEmpty);
            }
        }
        getEntriesByCatAndUser();
    }, [filterTimeForUsers]);

    // const optionsEntriesDate = {
    //     title: {text: 'Expenses in time'},
    //     xAxis: {
    //         categories: entriesByDate.map(entry => new Date(entry.entry_date).toLocaleString('en-GB', {
    //             day: 'numeric', // numeric, 2-digit
    //             year: '2-digit', // numeric, 2-digit
    //             month: 'short', // numeric, 2-digit, long, short, narrow
    //         }))
    //     },
    //     plotOptions: {
    //         line: {
    //             dataLabels: {
    //                 enabled: true
    //             },
    //             enableMouseTracking: false
    //         }
    //     },
    //     series: [{data: entriesByDate.map(entry => entry.amount)}]
    // };

    const optionsEntriesByWeek = {
        chart: {
            type: 'bar'
        },
        title: {text: 'Expenses by week'},
        xAxis: {
            categories: entriesByWeek.map(entry => entry.entry_name)
        },
        plotOptions: {
            line: {
                dataLabels: {
                    enabled: true
                },
                enableMouseTracking: false
            }
        },
        series: [{
            data: entriesByWeek.map(entry => entry.amount)}]

    };


// this is not yet working
//     function genSeries(entriesByTime) {
//         let newSeries = []
//         let newCategories = []
//         let oldCategories = []
//         for (let i = 0; i < entriesByTime.length; i++) {
//             if (!newCategories.includes(entriesByTime[i].entry_name)) {
//                 newCategories.push(entriesByTime[i].entry_name)
//             }
//             if (!oldCategories.includes(entriesByTime[i].category)) {
//                 oldCategories.push(entriesByTime[i].category)
//             }
//         }
//
//         for (let i = 0; i < oldCategories.length; i++) {
//                 newSeries.push({
//                     "name": oldCategories[i],
//                     "data": []
//                 })
//         }
//         console.log("entriesByTime", entriesByTime);
//         console.log("newSeries", newSeries);
//
//         for (let i = 0; i < newSeries.length; i++) {
//             for (let j = 0; j < entriesByTime.length; j++) {
//                 if (entriesByTime[j].category === newSeries[i].name) {
//                     newSeries[i].data.push({
//                         "cat": entriesByTime[j].entry_name,
//                         "amount": entriesByTime[j].amount
//                     })
//                 }
//             }
//         }
//
//         return {"series": newSeries, "categories": newCategories}
//     }

    const monthNames = ["January", "February", "March", "April", "May", "June",
        "July", "August", "September", "October", "November", "December"
    ];

    const optionsEntriesByMonth = {
        chart: {
            type: 'bar'
        },
        title: {text: 'Expenses by month'},
        xAxis: {
            categories: entriesByMonth.map(entry => monthNames[new Date(entry.entry_date).getMonth()] + ' ' + new Date(entry.entry_date).getFullYear())
        },
        plotOptions: {
            line: {
                dataLabels: {
                    enabled: true
                },
                enableMouseTracking: false
            }
        },
        series: [{
            data: entriesByMonth.map(entry => entry.amount)}]
    };

    function getCategoryNamesForOptions(entries) {
        let categoryNames = []
        if (entries == null) {
            return categoryNames
        }
        for (let i = 0; i < entries.length; i++) {
            for (let j = 0; j < categories.length; j++) {
                if (categories[j].id === entries[i].category) {
                    categoryNames.push({
                        "name": categories[j].category_name,
                        "y": entries[i].amount
                    })
                }
            }
        }
        return categoryNames
    }

    function getCategoryName(entry) {
        if (entry === 0) {
            return "get all"
        }
        for (let i = 0; i < categories.length; i++) {
            if (categories[i].id === entry) {
                return categories[i].category_name
            }
        }
    }

    const optionsEntriesByCat = {
        chart: {
            type: 'column'
        },
        title: {text: ''},
        xAxis: {
            categories: getCategoryNamesForOptions(entriesByCat).map(cat => cat.name)
        },
        plotOptions: {
            line: {
                dataLabels: {
                    enabled: true
                },
                enableMouseTracking: false
            }
        },
        series: [{data: getCategoryNamesForOptions(entriesByCat).map(cat => cat.y),
            name: 'categories',
            lineWidth: 0.5,
        }]
    };

    const optionsEntriesPieByCat = {
        chart: {
            plotBackgroundColor: null,
            plotBorderWidth: null,
            plotShadow: false,
            type: 'pie'
        },
        title: {text: ''},
        tooltip: {
            pointFormat: '<b>{point.percentage:.1f}%</b>'
        },
        accessibility: {
            point: {
                valueSuffix: '%'
            }
        },
        plotOptions: {
            pie: {
                allowPointSelect: true,
                cursor: 'pointer',
                dataLabels: {
                    enabled: true,
                    format: '<b>{point.name}</b>: {point.percentage:.1f} %'
                },
                showInLegend: true
            }
        },
        series: [{
            colorByPoint: true,
            data: getCategoryNamesForOptions(entriesPieByCat)}]
    };

    // function getUserNamesForOptions(entries) {
    //     let categoryNames = []
    //     if (entries == null) {
    //         return categoryNames
    //     }
    //     for (let i = 0; i < entries.length; i++) {
    //         for (let j = 0; j < categories.length; j++) {
    //             if (categories[j].id === entries[i].category) {
    //                 categoryNames.push({
    //                     "name": categories[j].category_name,
    //                     "y": entries[i].amount
    //                 })
    //             }
    //         }
    //     }
    //     return categoryNames
    // }

    function isEntriesByCatAndUserZero() {
        let isZero = true;
        if (entriesByCatAndUser[0] === undefined) {
            return true;
        }
        for (let i = 0; i < entriesByCatAndUser.length; i++) {
            for (let j = 0; j < entriesByCatAndUser[i].series.categories.length; j++) {
                if (entriesByCatAndUser[i].series.data[j] !== 0) {
                    isZero = false;
                    break;
                }
            }
        }
        return isZero;
    }
    let users = entriesByCatAndUser;
    let user_categories = "";
    if ( users !== null) {
        if ( users[0] !== undefined) {
            user_categories = users[0].series.categories
        }
    }
    const optionsEntriesByUsers = {
        chart: {
            type: 'column'
        },
        title: {text: ''},
        xAxis: {
            categories: user_categories
        },
        plotOptions: {
            line: {
                dataLabels: {
                    enabled: true
                },
                enableMouseTracking: false
            }
        },
        series: users.map(user => user.series)
    };

    let chartFilterCategory;
    if (filterCategory === 0) {
        chartFilterCategory = (
            <h4 className="title is-5 center-text chart-title">Expenses by time</h4>
        )
    } else {
        chartFilterCategory = (
            <h4 className="title is-5 center-text chart-title">Expenses by time for {getCategoryName(filterCategory)}</h4>
        )
    }

    let categoriesCharts, chartFilterTime;
    if (entriesByCat === null) {
        chartFilterTime = (
            <div className="chart-filter">
                <h4 className="title is-5 center-text chart-title">Expenses by categories for {filterTime}</h4>
                <h4 className="title is-5 center-text chart-title">There are none - choose another time frame!</h4>
                <FilterTime currentTimeFilter={filterTime}
                            onChange={filter => setFilterTime(filter)}
                />
            </div>
        )
    } else if (filterTime === "get all") {
        chartFilterTime = (
            <div className="chart-filter">
                <h4 className="title is-5 center-text chart-title">Expenses by categories</h4>
                <FilterTime currentTimeFilter={filterTime}
                            onChange={filter => setFilterTime(filter)}
                />
            </div>
        )
        categoriesCharts = (
            <div className="chart-category">
                <div className="chart-item-right">
                    <HighchartsReact highcharts={Highcharts}
                                     options={optionsEntriesPieByCat} />
                </div>
                <div className="chart-item-left">
                    <HighchartsReact highcharts={Highcharts}
                                     options={optionsEntriesByCat} />
                </div>
            </div>
        )
    } else {
        chartFilterTime = (
            <div className="chart-filter">
                <h4 className="title is-5 center-text chart-title">Expenses by categories for {filterTime}</h4>
                <FilterTime currentTimeFilter={filterTime}
                            onChange={filter => setFilterTime(filter)}
                />
            </div>
        )
        categoriesCharts = (
            <div className="chart-category">
                <div className="chart-item-right">
                    <HighchartsReact highcharts={Highcharts}
                                     options={optionsEntriesPieByCat} />
                </div>
                <div className="chart-item-left">
                    <HighchartsReact highcharts={Highcharts}
                                     options={optionsEntriesByCat} />
                </div>
            </div>
        )
    }

    let categoriesByUserChart, chartFilterTimeForUsers;
    if (isEntriesByCatAndUserZero()) {
        chartFilterTimeForUsers = (
            <div className="chart-filter">
                <h4 className="title is-5 center-text chart-title">Expenses by users for { filterTimeForUsers }</h4>
                <h4 className="title is-5 center-text chart-title">There are none - choose another time frame!</h4>
                <FilterTime currentTimeFilter={filterTimeForUsers}
                            onChange={filter => setFilterTimeForUsers(filter)}
                />
            </div>
            )
    } else if (filterTimeForUsers === "get all") {
        chartFilterTimeForUsers = (
            <div className="chart-filter">
                <h4 className="title is-5 center-text chart-title">Expenses by users</h4>
                <FilterTime currentTimeFilter={filterTimeForUsers}
                            onChange={filter => setFilterTimeForUsers(filter)}
                />
            </div>
        )
        categoriesByUserChart = (
            <div className="chart-category-and-users">
                <div className="chart-item-center">
                    <HighchartsReact highcharts={Highcharts}
                                     options={optionsEntriesByUsers} />
                </div>
            </div>
        )
    } else {
        chartFilterTimeForUsers = (
            <div className="chart-filter">
                <h4 className="title is-5 center-text chart-title">Expenses by users for { filterTimeForUsers }</h4>
                <FilterTime currentTimeFilter={filterTimeForUsers}
                            onChange={filter => setFilterTimeForUsers(filter)}
                />
            </div>
        )
        categoriesByUserChart = (
            <div className="chart-category-and-users">
                <div className="chart-item-center">
                    <HighchartsReact highcharts={Highcharts}
                                     options={optionsEntriesByUsers} />
                </div>
            </div>
        )
    }

    return (
        <div className="chart-container">
                {/*<div className="chart-item">*/}
                {/*    <HighchartsReact highcharts={Highcharts}*/}
                {/*                     options={optionsEntriesDate} />*/}
                {/*</div>*/}
            <div className="chart-filter">
                {chartFilterCategory}
                <FilterCategory
                    items={categories}
                    selectedItem={filterCategory}
                    onItemSelect={filter => setFilterCategory(filter)}
                />
            </div>
            <div className="chart-time">
                <div className="chart-item-right">
                    <HighchartsReact highcharts={Highcharts}
                                     options={optionsEntriesByWeek} />
                </div>
                <div className="chart-item-left">
                    <HighchartsReact highcharts={Highcharts}
                                     options={optionsEntriesByMonth} />
                </div>
            </div>
            <div>
                { chartFilterTime }
                { categoriesCharts }
            </div>
            <div>
                { chartFilterTimeForUsers }
                { categoriesByUserChart }
            </div>

        </div>
    );
};

export default Overview;