import http from 'k6/http';
import { check } from 'k6';


export function TestObjectAPI(serverUrl, headers){
    const objectUrl = `${serverUrl}/object`
    let res;
    let payload;

    payload = {
        object: {
            namespace: "test_file",
            name: "1",
            relation: "read",
        }
    };
    res = http.post(`${objectUrl}/get-user-relations`, JSON.stringify(payload), {headers:headers});
    check(res, { 'GetUserRelations': (r) => r.status == 200 });


    payload = {
        object: {
            namespace: "test_file",
            name: "1",
            relation: "read",
        }
    };
    res = http.post(`${objectUrl}/get-role-relations`, JSON.stringify(payload), {headers:headers});
    check(res, { 'GetRoleRelations': (r) => r.status == 200 });
}