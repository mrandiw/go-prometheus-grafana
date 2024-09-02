import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    stages: [
        { duration: '1m', target: 10 }, // Ramp-up to 10 users over 1 minute
        { duration: '2m', target: 10 }, // Stay at 10 users for 2 minutes
        { duration: '1m', target: 0 },  // Ramp-down to 0 users over 1 minute
    ],
};

export default function () {
    // success
    let res2 = http.get('http://host.docker.internal:8080/get-user');
    check(res2, {
        'status is 200': (r) => r.status === 200,
        'response is Success Get Users': (r) => r.body === 'Success Get Users',
    });

    let res3 = http.get('http://host.docker.internal:8080/get-role');
    check(res3, {
        'status is 200': (r) => r.status === 200,
        'response is Success Get Roles': (r) => r.body === 'Success Get Roles',
    });

    let res4 = http.get('http://host.docker.internal:8080/get-level');
    check(res4, {
        'status is 200': (r) => r.status === 200,
        'response is Success Get Levels': (r) => r.body === 'Success Get Levels',
    });


    // error
    let res5 = http.get('http://host.docker.internal:8080/get-user?param=error');
    check(res5, {
        'status is 500': (r) => r.status === 500,
        'response contains Internal Server Error': (r) => r.body.includes('Internal Server Error'),
    });

    let res6 = http.get('http://host.docker.internal:8080/get-user?param=not-found');
    check(res6, {
        'status is 500': (r) => r.status === 500,
        'response contains Not Found': (r) => r.body.includes('Not Found'),
    });

    let res7 = http.get('http://host.docker.internal:8080/get-role?param=error');
    check(res7, {
        'status is 500': (r) => r.status === 500,
        'response contains Internal Server Error': (r) => r.body.includes('Internal Server Error'),
    });

    let res8 = http.get('http://host.docker.internal:8080/get-role?param=not-found');
    check(res8, {
        'status is 500': (r) => r.status === 500,
        'response contains Not Found': (r) => r.body.includes('Not Found'),
    });

    // Optional: Add a sleep to simulate user think time
    sleep(1);
}
