import http from 'k6/http';
import {check, sleep} from 'k6';


export let options = {
    stages:[
        // 模擬 100 個使用者在 30 秒內進行登入，巔峰時期有 10000 個使用者在 1 分鐘內進行登入
        {duration: '1s', target: 10000}, // 10000 per second
    ],
};


export default function () {
    let res = http.get('http://localhost:8080/ads/random');
    check(res, {'status was 200': (r) => r.status == 200});
    sleep(1);
}