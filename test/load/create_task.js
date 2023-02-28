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

let minListID = 1;
let maxListID = 40188;

const randomListID = () => {
    return Math.floor(Math.random() * (maxListID - minListID + 1)) + minListID;
};

let num = 0;
export default function () {
    const body = {
        name: `load_test${num++}`,
    };

    const res = http.put(`http://localhost:8000/lists/${randomListID()}/tasks`, JSON.stringify(body), {
        headers: { 'Content-Type': 'application/json' },
    });

    check(res, {
        'status is 201': (r) => r.status === 201,
        'response body': (r) => {
            const body = JSON.parse(r.body);
            return body.id > 0;
        },
    });
}
