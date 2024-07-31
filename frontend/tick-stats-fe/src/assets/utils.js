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
            const error = await response.json();
            throw new Error(error.message || 'Something went wrong');
        }

        return await response.json();
    }).catch(error => {
        console.error('Fetch error:', error);
        throw error;
    });
};

export const simpleLineChartOptionModel = {
    tooltip: {
        trigger: 'axis'
    },
    xAxis: {
        type: 'category',
        axisLabel: {
            formatter: function (value) {
                return moment(value).format('MM-DD HH:mm');
            }
        }
    },
    yAxis: {
        type: 'value'
    },
    color: ['#21CBF3'],
    dataZoom: [
        {
            show: true,
            realtime: true,
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