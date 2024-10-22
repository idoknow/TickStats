import { useGlobalStore } from '@/stores/global';
import moment from 'moment';

export const fetchWrapper = async (url, options = {}) => {
    if (!/^http/.test(url)) {
        const globalStore = useGlobalStore();
        url = globalStore.baseUrl + url;
    }

    return fetch(url, {
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
            ...options.headers,
        },
        ...options,
    }).then(async response => {
        if (!response.ok) {
            let resp = await response.json();
            response.data = resp;
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
    return data;
}

export const generatePieDemoData = () => {
    let mock_os = ['Windows', 'Linux', 'MacOS']
    return Array.from({ length: 3 }, (_, i) => {
        return {
            name: mock_os[i],
            value: Math.floor(Math.random() * 100)
        }
    })
}

export const fillingTimeData = (items, count = 1496, interval_mins = 30) => {
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
    return data;
}

export const chartsPresetConfigs = [
    {
        title: 'Simple Line Chart',
        chart_type: 'simple_line',
        option_model: simpleLineChartOptionModel,
        demo_data: generateTimeDemoData(),
        description: 'A simple line chart with time x-axis and value y-axis.',
        metric_example: `{
    "metrics_data": {
        "used_count": 2,
        ...
    }
}`,
        extra_configs: [
            {
                name: "distinct_ip",
                type: "bool",
                default: false,
                description: "Count distinct IP addresses."
            },
            {
                name: "method",
                type: "selectable",
                default: "sum",
                options: ["sum", "count", "accumulate"],
                description: "The method to aggregate the data points."
            }
        ]
    },
    {
        title: 'Simple Pie Chart',
        chart_type: 'simple_pie',
        option_model: simplePieChartOptionModel,
        demo_data: generatePieDemoData(),
        description: 'A simple pie chart with data in name-value format.',
        metric_example: `{
    "metrics_data": {
        "os_name": "windows",
        ...
    }
}`,
        extra_configs: []
    },

]


