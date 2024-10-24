import * as echarts from 'echarts';

class Chart {
    constructor(divId) {
        this.chart = null;
        this.divId = divId;
    }

    initChart(option) {
        this.option = option;
        if (option._chart_type === 'simple_line' && option._display_type === 'plain_number') {
            let data = option.series[0].data[1];
            let value = data[data.length - 1];
            let title = option.title.text;
            let ele = document.getElementById(this.divId);
            ele.innerHTML = `
                <h3 color="#666">${title}</h3>
                <div style="font-size: 72px; font-weight: bold; display: flex; align-items: center; justify-content: center; height: 100%"><span>${value}</span></div>
            `

        } else {
            this.chart = echarts.init(document.getElementById(this.divId));
            this.chart.setOption(option);
        }
    }

    dispose() {
        if (this.option._chart_type === 'simple_line' && this.option._display_type === 'plain_number') {
            document.getElementById(this.divId).innerHTML = '';
        } else {
            this.chart.dispose();
        }
    }

    resize() {
        if (this.option._chart_type === 'simple_line' && this.option._display_type === 'plain_number') {
            // do nothing
        } else {
            this.chart.resize();
        }
    }
}

export { Chart }