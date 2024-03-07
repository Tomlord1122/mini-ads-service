import http from 'k6/http';
import {check, sleep} from 'k6';


export let options = {
    stages:[
       
        {duration: '10s', target: 3000}, // 3000 per second
        {duration: '10s', target: 10000}, // 10000 per second
        {duration: '10s', target: 3000}, // 3000 per second
    ],
};


export default function () {
    let res = http.get('http://localhost:8080/ads/random');
    check(res, {'status was 200': (r) => r.status == 200});
    sleep(1);
}