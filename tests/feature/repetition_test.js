import http from 'k6/http';
import { check } from 'k6';

export function TestRepeatTuple(serverUrl, headers) {
    const userUrl = `${serverUrl}/user`
    const userName = "Jimmy"
    const roleName = "rd-director"
    let payload;
    let res;

    payload = {
        user_name: userName,
        role_name: roleName,
    };
    res = http.post(`${userUrl}/add-role`, JSON.stringify(payload), {headers:headers});
    check(res, { 'add role: status == 200': (r) => r.status == 200 });

    payload = {
        user_name: userName,
        role_name: roleName,
    };
    res = http.post(`${userUrl}/add-role`, JSON.stringify(payload), {headers:headers});
    check(res, { 
        'add role != 200': (r) => r.status != 200,
        'add role err = tuple exist': (r) => {
            return r.json().error = "tuple exist";
        },
    });
};
