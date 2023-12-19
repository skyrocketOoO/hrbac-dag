import http from 'k6/http';
import { check } from 'k6';

export function checkAPI(serverUrl, headers){

    // User
    const userUrl = `${serverUrl}/user`

    let payload = {};
    let res = http.get(`${userUrl}`, JSON.stringify(payload), {headers:headers});
    check(res, { 'list users': (r) => r.status == 200 });

    const userName = "Jimmy"
    // payload = {};
    // res = http.get(`${userUrl}/${userName}`, JSON.stringify(payload), {headers:headers});
    // check(res, { 'get user': (r) => r.status == 200 });

    payload = {};
    res = http.del(`${userUrl}/${userName}`, JSON.stringify(payload), {headers:headers});
    check(res, { 'delete user': (r) => r.status == 200 });

    payload = {
        user_name: "Jimmy",
        role_name: "rd-director",
    };
    res = http.post(`${userUrl}/add-role`, JSON.stringify(payload), {headers:headers});
    check(res, { 'user add role': (r) => r.status == 200 });
    
    payload = {
        user_name: "Jimmy",
        role_name: "rd-director",
    };
    res = http.post(`${userUrl}/remove-role`, JSON.stringify(payload), {headers:headers});
    check(res, { 'user remove role': (r) => r.status == 200 });
    
    payload = {
        user_name: "Jimmy",
    };
    res = http.post(`${userUrl}/list-relation`, JSON.stringify(payload), {headers:headers});
    check(res, { 'user list relation': (r) => r.status == 200 });
    
    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "*",
        user_name: "Jimmy",
    };
    res = http.post(`${userUrl}/add-relation`, JSON.stringify(payload), {headers:headers});
    check(res, { 'user add relation': (r) => r.status == 200 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "*",
        user_name: "Jimmy",
    };
    res = http.post(`${userUrl}/check`, JSON.stringify(payload), {headers:headers});
    check(res, { 'user check': (r) => r.status == 200 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "*",
        user_name: "Jimmy",
    };
    res = http.post(`${userUrl}/remove-relation`, JSON.stringify(payload), {headers:headers});
    check(res, { 'user remove relation': (r) => r.status == 200 });
    
    // Role
    const roleUrl = `${serverUrl}/role`

    payload = {};
    res = http.get(`${roleUrl}`, JSON.stringify(payload), {headers:headers});
    check(res, { 'list role ': (r) => r.status == 200 });

    const roleName = "rd";
    // payload = {};
    // res = http.get(`${roleUrl}/${roleName}`, JSON.stringify(payload), {headers:headers});
    // check(res, { 'get role ': (r) => r.status == 200 });

    payload = {};
    res = http.del(`${roleUrl}/${roleName}`, JSON.stringify(payload), {headers:headers});
    check(res, { 'delete role ': (r) => r.status == 200 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "*",
        role_name: "rd",
    };
    res = http.post(`${roleUrl}/add-relation`, JSON.stringify(payload), {headers:headers});
    check(res, { 'add relation ': (r) => r.status == 200 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "*",
        role_name: "rd",
    };
    res = http.post(`${roleUrl}/check`, JSON.stringify(payload), {headers:headers});
    check(res, { 'check': (r) => r.status == 200 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "*",
        role_name: "rd",
    };
    res = http.post(`${roleUrl}/remove-relation`, JSON.stringify(payload), {headers:headers});
    check(res, { 'Remove relation ': (r) => r.status == 200 });

    payload = {
        child_role_name: "rd",
        parent_role_name: "rd-director",
    };
    res = http.post(`${roleUrl}/add-parent`, JSON.stringify(payload), {headers:headers});
    check(res, { 'add-parent ': (r) => r.status == 200 });

    payload = {
        child_role_name: "rd",
        parent_role_name: "rd-director",
    };
    res = http.post(`${roleUrl}/remove-parent`, JSON.stringify(payload), {headers:headers});
    check(res, { 'remove-parent': (r) => r.status == 200 });

    payload = {};
    res = http.post(`${roleUrl}/list-relation`, JSON.stringify(payload), {headers:headers});
    check(res, { 'list-relation': (r) => r.status == 200 });

    payload = {};
    res = http.post(`${roleUrl}/get-members`, JSON.stringify(payload), {headers:headers});
    check(res, { 'get-role-members': (r) => r.status == 200 });

    // Object
    const objectUrl = `${serverUrl}/object`
    // payload = {};
    // res = http.post(`${objectUrl}/list-user-has-relation`, JSON.stringify(payload), {headers:headers});
    // check(res, { 'list-user-has-relation': (r) => r.status == 200 });

    // payload = {};
    // res = http.post(`${objectUrl}/list-role-has-relation`, JSON.stringify(payload), {headers:headers});
    // check(res, { 'list-role-has-relation': (r) => r.status == 200 });

    // payload = {};
    // res = http.post(`${objectUrl}/list-user-or-role-has-relation`, JSON.stringify(payload), {headers:headers});
    // check(res, { 'list-user-or-role-has-relation': (r) => r.status == 200 });

    // payload = {};
    // res = http.post(`${objectUrl}/list-all-relations`, JSON.stringify(payload), {headers:headers});
    // check(res, { 'list-all-relations': (r) => r.status == 200 });

    // Relation
    const relationUrl = `${serverUrl}/relation`

    res = http.get(`${relationUrl}`, null, {headers:headers});
    check(res, { 'list all relations': (r) => r.status == 200 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "read",
        subject_namespace: "test_file",
        subject_name: "1",
        subject_relation: "write",
    };
    res = http.post(`${relationUrl}/link`, JSON.stringify(payload), {headers:headers});
    check(res, { 'link relation': (r) => r.status == 200 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "read",
        subject_namespace: "test_file",
        subject_name: "1",
        subject_relation: "write", 
    };
    res = http.post(`${relationUrl}/check`, JSON.stringify(payload), {headers:headers});
    check(res, { 'relation check': (r) => r.status ==  403 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "read",
        subject_namespace: "test_file",
        subject_name: "1",
        subject_relation: "write", 
    };
    res = http.post(`${relationUrl}/path`, JSON.stringify(payload), {headers:headers});
    check(res, { 'relation path': (r) => r.status == 500 });

}