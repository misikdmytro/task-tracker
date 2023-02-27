import http from 'k6/http';
import { check } from 'k6';

export const options = {
    stages: [{ target: 20, duration: '10s' }],
    stages: [{ target: 50, duration: '20s' }],
    stages: [{ target: 75, duration: '30s' }],
    stages: [{ target: 200, duration: '1m' }],
    stages: [{ target: 0, duration: '30s' }],
    thresholds: {
        http_req_failed: ['rate<0.01'],
        http_req_duration: ['p(95)<200'],
    },
};

let num = 0;
export default function () {
    const body = {
        name: `load_test${num++}`,
    };

    const res = http.put('http://localhost:8000/lists/', JSON.stringify(body), {
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
