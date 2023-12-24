import http from 'k6/http';
import { check } from 'k6';

export function TestUniversalSyntax(serverUrl, headers) {
    const userUrl = `${serverUrl}/user`
    const objNs = "test_obj";
    const objName = "1";
    const relation = "write";
    const userName = "JImmy";
    let payload;
    let res;

    payload = {
        object_namespace: objNs,
        obj_name: "*",
        relation: "*",
        user_name: userName,
    };
    res = http.post(`${userUrl}/add-relation`, JSON.stringify(payload), {headers:headers});
    check(res, { 'add-relation: status == 200': (r) => r.status == 200 });

    payload = {
        object_namespace: objNs,
        object_name: objName,
        relation: relation,
        userName: userName,
    };
    res = http.post(`${userUrl}/check`, JSON.stringify(payload), {headers:headers});
    check(res, { 'Check: status == 200': (r) => r.status == 200 });

    payload = {
        object_namespace: objNs,
        obj_name: "*",
        relation: "*",
        user_name: userName,
    };
    res = http.post(`${userUrl}/remove-relation`, JSON.stringify(payload), {headers:headers});
    check(res, { 'remove-relation: status == 200': (r) => r.status == 200 });


    payload = {
        object_namespace: objNs,
        obj_name: "*",
        relation: "write",
        user_name: userName,
    };
    res = http.post(`${userUrl}/add-relation`, JSON.stringify(payload), {headers:headers});
    check(res, { 'add-relation: status == 200': (r) => r.status == 200 });

    payload = {
        object_namespace: objNs,
        object_name: objName,
        relation: relation,
        userName: userName,
    };
    res = http.post(`${userUrl}/check`, JSON.stringify(payload), {headers:headers});
    check(res, { 'Check: status == 200': (r) => r.status == 200 });

    payload = {
        object_namespace: objNs,
        obj_name: "*",
        relation: "write",
        user_name: userName,
    };
    res = http.post(`${userUrl}/remove-relation`, JSON.stringify(payload), {headers:headers});
    check(res, { 'remove-relation: status == 200': (r) => r.status == 200 });


    payload = {
        object_namespace: objNs,
        obj_name: "1",
        relation: "*",
        user_name: userName,
    };
    res = http.post(`${userUrl}/add-relation`, JSON.stringify(payload), {headers:headers});
    check(res, { 'add-relation: status == 200': (r) => r.status == 200 });

    payload = {
        object_namespace: objNs,
        object_name: objName,
        relation: relation,
        userName: userName,
    };
    res = http.post(`${userUrl}/check`, JSON.stringify(payload), {headers:headers});
    check(res, { 'Check: status == 200': (r) => r.status == 200 });

    payload = {
        object_namespace: objNs,
        obj_name: "1",
        relation: "*",
        user_name: userName,
    };
    res = http.post(`${userUrl}/remove-relation`, JSON.stringify(payload), {headers:headers});
    check(res, { 'remove-relation: status == 200': (r) => r.status == 200 });
};
