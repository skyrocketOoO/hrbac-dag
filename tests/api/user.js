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
    check(res, { 
        'GetAllUsers: status == 200': (r) => r.status == 200,
        'GetAllUsers: User "Jimmy" exists in response': (r) => {
            return r.json().data[0] == "Jimmy";
        }
    });

    res = http.get(`${userUrl}/${userName}`, null, {headers:headers});
    check(res, { 'GetUser: status == 200': (r) => r.status == 200 });


    res = http.del(`${userUrl}/${userName}`, null, {headers:headers});
    check(res, { 'DeleteUser: status == 200': (r) => r.status == 200 });
    
    payload = {
        user_name: userName,
        role_name: roleName,
    };
    http.post(`${userUrl}/add-role`, JSON.stringify(payload), {headers:headers});
    payload = {
        user_name: "Jimmy",
        role_name: "rd-director",
    };
    res = http.post(`${userUrl}/remove-role`, JSON.stringify(payload), {headers:headers});
    check(res, { 'RemoveRole: status == 200': (r) => r.status == 200 });
    
    payload = {
        name: "Jimmy",
    };
    res = http.post(`${userUrl}/find-all-object-relations`, JSON.stringify(payload), {headers:headers});
    check(res, { 'FindAllObjectRelations: status == 200': (r) => r.status == 200 });
    
    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "*",
        user_name: "Jimmy",
    };
    res = http.post(`${userUrl}/add-relation`, JSON.stringify(payload), {headers:headers});
    check(res, { 'AddRelation: status == 200': (r) => r.status == 200 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "*",
        user_name: "Jimmy",
    };
    res = http.post(`${userUrl}/check`, JSON.stringify(payload), {headers:headers});
    check(res, { 'Check: status == 200': (r) => r.status == 200 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "*",
        user_name: "Jimmy",
    };
    res = http.post(`${userUrl}/remove-relation`, JSON.stringify(payload), {headers:headers});
    check(res, { 'RemoveRelation: status == 200': (r) => r.status == 200 });
}
