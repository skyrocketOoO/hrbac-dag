import http from 'k6/http';
import { check } from 'k6';

export function TestHRBAC(serverUrl, headers) {
    const roleUrl = `${serverUrl}/role`
    const parentRoleName = "rd-director"
    const childRoleName = "rd"
    const relation = "modify-permission"
    const objNs = "test_obj"
    const objName = "1"
    let payload;
    let res;

    payload = {
        child_role_name: childRoleName,
        parent_role_name: parentRoleName,
    };
    res = http.post(`${roleUrl}/add-parent`, JSON.stringify(payload), {headers:headers});
    check(res, { 'AddParent: status == 200': (r) => r.status == 200 });

    payload = {
        object_namespace: objNs,
        object_name: objName,
        relation: relation,
        roleName: parentRoleName,
    };
    res = http.post(`${roleUrl}/check`, JSON.stringify(payload), {headers:headers});
    check(res, { 'Check: status == 200': (r) => r.status == 200 });

    payload = {
        child_role_name: childRoleName,
        parent_role_name: parentRoleName,
    };
    res = http.post(`${roleUrl}/remove-parent`, JSON.stringify(payload), {headers:headers});
    check(res, { 'AddParent: status == 200': (r) => r.status == 200 });
};
