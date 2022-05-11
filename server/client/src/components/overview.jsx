import React, { useState, useEffect } from "react";
import Highcharts from 'highcharts';
import HighchartsReact from 'highcharts-react-official';
import {getCategories} from "../services/categoryService";
import {getEntriesByDate, getEntriesByWeek, getEntriesByMonth, getEntriesByCategory, getEntriesPieByCategory} from "../services/chartsService";
import FilterTime from "./common/filterTime";
import FilterCategory from "./common/filterCategories";


const Overview = (props) => {
    const [entriesByDate, setEntriesByDate] = useState([]);
    const [entriesByWeek, setEntriesByWeek] = useState([]);
    const [entriesByMonth, setEntriesByMonth] = useState([]);
    const [categories, setCategories] = useState([]);
    const [entriesByCat, setEntriesByCat] = useState([]);
    const [entriesPieByCat, setEntriesPieByCat] = useState([]);
    const [filterTime, setFilterTime] = useState("Current week")
    const [filterCategory, setFilterCategory] = useState(0)

    useEffect( () => {
        async function getEntriesByCat() {
            const { data: categories } = await getCategories();
            setCategories(categories);
            const { data: entriesCat } = await getEntriesByCategory(filterTime);
            setEntriesByCat(entriesCat);
            const { data: entriesPieByCat } = await getEntriesPieByCategory(filterTime);
            setEntriesPieByCat(entriesPieByCat);
        }
        getEntriesByCat();
    }, [filterTime]);

    useEffect( () => {
        async function getEntriesByTime() {
            const { data: entriesDate } = await getEntriesByDate();
            setEntriesByDate(entriesDate);
            const { data: entriesWeek } = await getEntriesByWeek(filterCategory);
            if (entriesWeek != null) {
                setEntriesByWeek(entriesWeek);
            }
            console.log("entriesWeek", entriesWeek);
            const { data: entriesMonth } = await getEntriesByMonth(filterCategory);
            if (entriesMonth != null) {
                setEntriesByMonth(entriesMonth);
            }
        }
        getEntriesByTime();
    }, [filterCategory]);

    const optionsEntriesDate = {
        title: {text: 'Expenses in time'},
        xAxis: {
            categories: entriesByDate.map(entry => new Date(entry.entry_date).toLocaleString('en-GB', {
                day: 'numeric', // numeric, 2-digit
                year: '2-digit', // numeric, 2-digit
                month: 'short', // numeric, 2-digit, long, short, narrow
            }))
        },
        plotOptions: {
            line: {
                dataLabels: {
                    enabled: true
                },
                enableMouseTracking: false
            }
        },
        series: [{data: entriesByDate.map(entry => entry.amount)}]
    };

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

    function getCategoryNames(entries) {
        let categoryNames = []
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

    const optionsEntriesByCat = {
        chart: {
            type: 'column'
        },
        title: {text: ''},
        xAxis: {
            categories: getCategoryNames(entriesByCat).map(cat => cat.name)
        },
        plotOptions: {
            line: {
                dataLabels: {
                    enabled: true
                },
                enableMouseTracking: false
            }
        },
        series: [{data: getCategoryNames(entriesByCat).map(cat => cat.y),
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
            data: getCategoryNames(entriesPieByCat)}]
    };
    console.log("entriesWeek2", entriesByWeek);
    return (
        <div className="chart-container">
                {/*<div className="chart-item">*/}
                {/*    <HighchartsReact highcharts={Highcharts}*/}
                {/*                     options={optionsEntriesDate} />*/}
                {/*</div>*/}
            <div className="chart-filter">
                <h4 className="title is-5 center-text chart-title">Expenses by time</h4>
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
            <div className="chart-filter">
                <h4 className="title is-5 center-text chart-title">Expenses by categories</h4>
                <FilterTime currentTimeFilter={filterTime}
                            onChange={filter => setFilterTime(filter)}
                />
            </div>
            <div>
                <div className="chart-item-right">
                    <HighchartsReact highcharts={Highcharts}
                                     options={optionsEntriesPieByCat} />
                </div>
                <div className="chart-item-left">
                    <HighchartsReact highcharts={Highcharts}
                                     options={optionsEntriesByCat} />
                </div>
            </div>

        </div>
    );
};

export default Overview;