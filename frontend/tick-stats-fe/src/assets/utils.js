import { useGlobalStore } from '@/stores/global';
import moment from 'moment';

export const fetchWrapper = async (url, options = {}) => {
    if (!/^http/.test(url)) {
        const globalStore = useGlobalStore();
        url = globalStore.baseUrl + url;
    }

    console.log(url)

    return fetch(url, {
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
            ...options.headers,
        },
        ...options,
    }).then(async response => {
        if (!response.ok) {
            return Promise.reject(response);
        }

        let resp = await response.json();
        return Promise.resolve(resp.data);
    })
};

export const simpleLineChartOptionModel = {
    tooltip: {
        trigger: 'axis'
    },
    xAxis: {
        type: 'time',
        axisLabel: {
            formatter: function (value) {
                return moment(value).format('MM-DD HH:mm');
            }
        },
    },
    yAxis: {
        type: 'value'
    },
    color: ['#4247CB'],
    dataZoom: [
        {
            show: true,
            realtime: true,
            start: 90,
            end: 100,
        }
    ],
    _chart_type: 'simple_line'
};

export const simplePieChartOptionModel = {
    tooltip: {
        trigger: 'item',
        formatter: '{a} <br/>{b} : {c} ({d}%)'
    },
    padAngle: 5,
    itemStyle: {
        borderRadius: 5
    },
    legend: {
        orient: 'vertical',
        left: 'left',
        data: []
    },
    _chart_type: 'simple_pie'
};

export const generateTimeDemoData = (count = 1496) => {
    let now = Date.now();
    let data = [];
    for (let i = 0; i < count; i++) {
        data.push([
            now + i * 2000,
            Math.random() * 100
        ]);
    }
    console.log(data);
    return data;
}

export const fillingTimeData = (items, count = 1496, interval_mins = 30) => {
    console.log(items);
    let arr = items.map((item) => {
        return [item.k, item.v];
    })

    let data = [];
    let nearest = new Date();
    nearest.setSeconds(0);
    nearest.setMilliseconds(0);
    let minutes = nearest.getMinutes();
    if (minutes < 30) {
        nearest.setMinutes(0);
    } else {
        nearest.setMinutes(30);
    }
    nearest = nearest.getTime();

    console.log(nearest);

    let interval = interval_mins * 60 * 1000;
    for (let i = 0; i < count; i++) {
        let time = nearest - i * interval
        let found = arr.find((item) => {
            return item[0] === time;
        });
        if (found) {
            data.push(found);
        } else {
            data.push([time, 0]);
        }
    }

    console.log(data);

    return data;
}



