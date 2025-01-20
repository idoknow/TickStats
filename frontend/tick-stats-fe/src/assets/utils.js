import { useGlobalStore } from '@/stores/global';
import moment from 'moment';
import { da } from 'vuetify/locale';

export const fetchWrapper = async (url, options = {}) => {
    if (!/^http/.test(url)) {
        const globalStore = useGlobalStore();
        url = globalStore.baseUrl + url;
    }

    return fetch(url, {
        headers: {
            'Authorization': localStorage.getItem('token'),
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
export const fillingTimeData = (items, interval_mins = 30) => {
    // 将 items 转换为 [timestamp_ms, value] 的数组
    let arr = items.map((item) => {
        return [item.k, item.v]; // [timestamp_ms, value]
    });

    if (arr.length === 0) {
        return arr;
    }
    arr.sort((a, b) => a[0] - b[0]);

    let data = [];
    let currentTime = arr[0][0];
    const intervalMs = interval_mins * 60 * 1000;
    const lastTime = arr[arr.length - 1][0];

    // 遍历所有时间点，填充缺失的数据
    while (currentTime <= lastTime) {
        if (arr.length > 0 && currentTime === arr[0][0]) {
            data.push(arr.shift());
        } else {
            data.push([currentTime, 0]);
        }
        currentTime += intervalMs;
    }

    return data;
};

export const chartsPresetConfigs = [
    {
        title: 'Simple Line Chart / Plain Number',
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
        extra_config: [
            {
                name: "method",
                type: "selectable",
                default: "sum",
                options: ["sum", "count", "accumulate"],
                description: "The method to aggregate the data points."
            },
            {
                name: "distinct_ip",
                type: "bool",
                default: false,
                description: "Only available for `count` method."
            },
            {
                name: "only_represent_number",
                type: "bool",
                default: false,
                description: "Only display the number of the latest data point."
            },
            {
                name: "bucket_mins",
                type: "number",
                default: 30,
                description: "The time interval in minutes for each data point. Default is 30 mins."
            }
        ],
        multiple_keys: false
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
        extra_config: [],
        multiple_keys: false
    },
    {
        title: 'Table',
        chart_type: 'table',
        option_model: {_chart_type: 'table'},
        demo_data: [],
        description: 'A simple table to display data.',
        metric_example: `{
    "metrics_data": {
        "os_name": "windows",
        "user_count": 100,
        ...
    }
}`,
        extra_config: [],
        multiple_keys: true
    },
]