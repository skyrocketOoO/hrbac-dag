import http from 'k6/http';
import { check } from 'k6';

export function TestUserAPI(serverUrl, headers) {
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
    check(res, { 'AddRole: status == 200': (r) => r.status == 200 });

    res = http.get(`${userUrl}`, null, {headers:headers});
    check(res, { 'GetAll: status == 200': (r) => r.status == 200 });

    res = http.del(`${userUrl}/${userName}`, null, {headers:headers});
    check(res, { 'Delete: status == 200': (r) => r.status == 200 });
    
    payload = {
        user_name: userName,
        role_name: roleName,
    };
    http.post(`${userUrl}/add-role`, JSON.stringify(payload), {headers:headers});
    payload = {
        user_name: userName,
        role_name: roleName,
    };
    res = http.post(`${userUrl}/remove-role`, JSON.stringify(payload), {headers:headers});
    check(res, { 'RemoveRole: status == 200': (r) => r.status == 200 });
    
    payload = {
        name: userName,
    };
    res = http.post(`${userUrl}/get-all-object-relations`, JSON.stringify(payload), {headers:headers});
    check(res, { 'GetAllObjectRelations: status == 200': (r) => r.status == 200 });
    
    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "*",
        user_name: userName,
    };
    res = http.post(`${userUrl}/add-relation`, JSON.stringify(payload), {headers:headers});
    check(res, { 'AddRelation: status == 200': (r) => r.status == 200 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "*",
        user_name: userName,
    };
    res = http.post(`${userUrl}/check`, JSON.stringify(payload), {headers:headers});
    check(res, { 'Check: status == 200': (r) => r.status == 200 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "*",
        user_name: userName,
    };
    res = http.post(`${userUrl}/remove-relation`, JSON.stringify(payload), {headers:headers});
    check(res, { 'RemoveRelation: status == 200': (r) => r.status == 200 });
}
