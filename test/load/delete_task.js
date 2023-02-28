import http from 'k6/http';
import { check } from 'k6';

export const options = {
    stages: [
        { target: 50, duration: '30s' },
        { target: 50, duration: '1m' },
        { target: 0, duration: '30s' },
    ],
    thresholds: {
        http_req_failed: ['rate<0.01'],
        http_req_duration: ['p(95)<300'],
    },
};

let minTaskID = 1;
let maxTaskID = 42397;

const randomTaskID = () => {
    return Math.floor(Math.random() * (maxTaskID - minTaskID + 1)) + minTaskID;
};

export default function () {
    const res = http.del(`http://localhost:8000/tasks/${randomTaskID()}`, null, {
        headers: { 'Content-Type': 'application/json' },
    });

    check(res, {
        'status is 204': (r) => r.status === 204
    });
}
