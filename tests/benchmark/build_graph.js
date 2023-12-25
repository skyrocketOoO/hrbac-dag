import http from 'k6/http';
import { check } from 'k6';

export function BuildGraph(serverUrl, headers){
    let payload, res;
    const layer = 10;
    const exp_base = 4;

    payload = {
        object_namespace: "",
        object_name: "",
        relation: "",
        subject_namespace: "",
        subject_name: "",
        subject_relation: "",
    }

    
    res = http.post(`${serverUrl}/user/add-role`, JSON.stringify(payload), {headers:headers});
    check(res, { '(Add Relation) Jimmy is a member of RD-Director': (r) => r.status == 200 } );
}