import http from 'k6/http';
import {check, sleep} from 'k6';


export let options = {
    stages:[
       
        {duration: '10s', target: 5000}, 
        {duration: '10s', target: 10000}, 
        {duration: '10s', target: 15000}, 
        {duration: '10s', target: 15000}, 
        {duration: '10s', target: 15000}, 
        {duration: '10s', target: 10000}, 

    ],
};


export default function () {
    let res = http.get('http://localhost:8080/ads?age_start=18&age_end=80&country[]=US&platform[]=ios&gender[]=F&limit=10&offset=0');
    check(res, {'status was 200': (r) => r.status == 200});
    sleep(1);
}


// Path: k6/loadtest.js