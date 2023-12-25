import http from 'k6/http';
import { check } from 'k6';

export function BuildGraph(serverUrl, headers){
    const roleUrl = `${serverUrl}/role`
    let payload, res;
    const roleLayer = 6, roleExpBase = 5;
    
    let curLayer = 1;
    while (curLayer <= roleLayer){
        const count = Math.pow(roleExpBase, curLayer);

        for (let i = 0; i < count; i++){
            payload = {
                child_role_name: (curLayer-1).toString() + "_" + (i / roleExpBase).toString(),
                parent_role_name: curLayer.toString() + "_" + i.toString(),
            };
            res = http.post(`${roleUrl}/add-parent`, JSON.stringify(payload), {headers: headers});
            check(res, { 'AddParent: status == 200': (r) => r.status == 200 });
        };

        curLayer += 1;
    };
};