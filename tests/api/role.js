import http from 'k6/http';
import { check } from 'k6';

export function TestRoleAPI(serverUrl, headers) {
    const roleUrl = `${serverUrl}/role`
    const userUrl = `${serverUrl}/user`
    const userName = "Jimmy"
    const roleName = "rd-director"
    const parentRoleName = roleName
    const childRoleName = "rd"
    let payload;
    let res;

    payload = {
        child_role_name: childRoleName,
        parent_role_name: parentRoleName,
    };
    res = http.post(`${roleUrl}/add-parent`, JSON.stringify(payload), {headers:headers});
    check(res, { 'AddParent: status == 200': (r) => r.status == 200 });

    payload = {
        child_role_name: childRoleName,
        parent_role_name: parentRoleName,
    };
    res = http.post(`${roleUrl}/remove-parent`, JSON.stringify(payload), {headers:headers});
    check(res, { 'RemoveParent: status == 200': (r) => r.status == 200 });

    payload = {
        user_name: userName,
        role_name: roleName,
    };
    res = http.post(`${userUrl}/add-role`, JSON.stringify(payload), {headers:headers});
    payload = {
        name: roleName,
    };
    res = http.post(`${roleUrl}/get-members`, JSON.stringify(payload), {headers:headers});
    check(res, { 
        'GetMembers: status == 200': (r) => r.status == 200,
        'GetMembers: Jimmy in members': (r) => {
            return r.json().data[0] == "Jimmy";
        }
    });

    res = http.get(`${roleUrl}`, null, {headers:headers});
    check(res, { 
        'GetAllRoles: status == 200': (r) => r.status == 200,
        'GetAllRoles: rd-director exists in response': (r) => {
            return r.json().data.includes(roleName);
        }
    });

    res = http.get(`${roleUrl}/${roleName}`, null, {headers:headers});
    check(res, { 'GetRole: status == 200': (r) => r.status == 200 });


    res = http.del(`${roleUrl}/${roleName}`, null, {headers:headers});
    check(res, { 'DeleteRole: status == 200': (r) => r.status == 200 });
    
    payload = {
        name: roleName,
    };
    res = http.post(`${roleUrl}/find-all-object-relations`, JSON.stringify(payload), {headers:headers});
    check(res, { 'FindAllObjectRelations: status == 200': (r) => r.status == 200 });
    
    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "*",
        roleName: roleName,
    };
    res = http.post(`${roleUrl}/add-relation`, JSON.stringify(payload), {headers:headers});
    check(res, { 'AddRelation: status == 200': (r) => r.status == 200 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "*",
        roleName: roleName,
    };
    res = http.post(`${roleUrl}/check`, JSON.stringify(payload), {headers:headers});
    check(res, { 'Check: status == 200': (r) => r.status == 200 });

    payload = {
        object_namespace: "test_file",
        object_name: "1",
        relation: "*",
        roleName: roleName,
    };
    res = http.post(`${roleUrl}/remove-relation`, JSON.stringify(payload), {headers:headers});
    check(res, { 'RemoveRelation: status == 200': (r) => r.status == 200 });
}
