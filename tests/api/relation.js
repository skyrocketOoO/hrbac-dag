import http from 'k6/http';
import { check } from 'k6';


export function TestRelationAPI(serverUrl, headers){
    const relationUrl = `${serverUrl}/relation`
    let res;
    let payload;

    res = http.get(`${relationUrl}`, null, {headers:headers});
    check(res, { 'GetAllRelations': (r) => r.status == 200 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "read",
        subject_namespace: "test_file",
        subject_name: "1",
        subject_relation: "write",
    };
    res = http.post(`${relationUrl}/add-link`, JSON.stringify(payload), {headers:headers});
    check(res, { 'AddLink': (r) => r.status == 200 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "read",
        subject_namespace: "test_file",
        subject_name: "1",
        subject_relation: "write",
    };
    res = http.post(`${relationUrl}/check`, JSON.stringify(payload), {headers:headers});
    check(res, { 'Check': (r) => r.status ==  200 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "read",
        subject_namespace: "test_file",
        subject_name: "1",
        subject_relation: "write",
    };
    res = http.post(`${relationUrl}/remove-link`, JSON.stringify(payload), {headers:headers});
    check(res, { 'RemoveLink': (r) => r.status == 200 });

    // payload = {
    //     object_namespace: "test_file",
    //     object_name: "1",
    //     relation: "read",
    //     subject_namespace: "test_file",
    //     subject_name: "1",
    //     subject_relation: "write", 
    // };
    // res = http.post(`${relationUrl}/path`, JSON.stringify(payload), {headers:headers});
    // check(res, { 'relation path': (r) => r.status == 500 });

}