import requests
import logging
import random
import schedule
import time

# 配置日志
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

# 定义发送 POST 请求的函数
def send_post_request():
    url = 'http://localhost:8080/api/metric/1abacd02'
    data = {
        "metrics_data": {
            "post_num": random.randint(0, 10),
            "tick": 1
        }
    }
    try:
        response = requests.post(url, json=data)
        response.raise_for_status()
        logging.info(f"POST 请求成功: {response.json()}")
    except requests.exceptions.RequestException as e:
        logging.error(f"POST 请求失败: {e}")

# 每 30 分钟执行一次任务
schedule.every(30).minutes.do(send_post_request)

# 立即执行一次任务
send_post_request()

# 保持脚本运行
while True:
    schedule.run_pending()
    time.sleep(1)