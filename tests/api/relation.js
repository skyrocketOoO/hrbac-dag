import http from 'k6/http';
import { check } from 'k6';


export function TestRelationAPI(serverUrl, headers){
    const relationUrl = `${serverUrl}/relation`
    let res;
    let payload;

    res = http.get(`${relationUrl}`, null, {headers:headers});
    check(res, { 'GetAll': (r) => r.status == 200 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "read",
        subject_namespace: "test_file",
        subject_name: "1",
        subject_relation: "write",
    };
    res = http.post(`${relationUrl}/query`, JSON.stringify(payload), {headers:headers});
    check(res, { 'Query': (r) => r.status == 200 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "read",
        subject_namespace: "test_file",
        subject_name: "1",
        subject_relation: "write",
    };
    res = http.post(`${relationUrl}`, JSON.stringify(payload), {headers:headers});
    check(res, { 'Create': (r) => r.status == 200 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "read",
        subject_namespace: "test_file",
        subject_name: "1",
        subject_relation: "write",
    };
    res = http.del(`${relationUrl}`, JSON.stringify(payload), {headers:headers});
    check(res, { 'Delete': (r) => r.status == 200 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "read",
        subject_namespace: "test_file",
        subject_name: "1",
        subject_relation: "write",
    };
    res = http.post(`${relationUrl}/check`, JSON.stringify(payload), {headers:headers});
    check(res, { 'Check': (r) => r.status == 200 });

    res = http.post(`${relationUrl}/clear-all-relations`, null, {headers:headers});
    check(res, { 'ClearAllRelations': (r) => r.status == 200 });
}