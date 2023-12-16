import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  vus: 1,
}


export default function() {
  const SERVER_URL = "http://localhost:3000"
  const Headers = {
    'Content-Type': 'application/json',
  }

  // healthy
  let res = http.get(`${SERVER_URL}/healthy`);
  check(res, { 'Server is healthy': (r) => r.status == 200 });

  // Add Relation
  let payload = {
    user_name: "Jimmy",
    role_name: "rd-director",
  };
  res = http.post(`${SERVER_URL}/user/add-role`, JSON.stringify(payload), {headers:Headers});
  check(res, { '(Add Relation) Jimmy is a member of RD-Director': (r) => r.status == 200 });

  payload = {
    user_name: "Tasha",
    role_name: "rd",
  };
  res = http.post(`${SERVER_URL}/user/add-role`, JSON.stringify(payload), {headers:Headers});
  check(res, { '(Add Relation) Tasha is a member of RD': (r) => r.status == 200 });

  payload = {
    user_name: "Ivy",
    role_name: "sales",
  };
  res = http.post(`${SERVER_URL}/user/add-role`, JSON.stringify(payload), {headers:Headers});
  check(res, { '(Add Relation) Ivy is a member of Sales': (r) => r.status == 200 });

  payload = {
    child_rolename: "rd",
    parent_rolename: "rd-director",
  };
  res = http.post(`${SERVER_URL}/role/add-parent-role`, JSON.stringify(payload), {headers:Headers});
  check(res, { '(Add Relation) RD-Director is a parent of RD': (r) => r.status == 200 });

  // Add permission
  payload = {
    user_name: "Ivy",
    relation: "read",
    object_namespace: "profile",
    object_name: "Ivy",
  };
  res = http.post(`${SERVER_URL}/user/add-permission`, JSON.stringify(payload), {headers:Headers});
  check(res, { '(Add Permission) Ivy has read access to her profile': (r) => r.status == 200 });

  payload = {
    user_name: "Heidi",
    relation: "read",
    object_namespace: "profile",
    object_name: "Heidi",
  };
  res = http.post(`${SERVER_URL}/user/add-permission`, JSON.stringify(payload), {headers:Headers});
  check(res, { '(Add Permission) Heidi has read access to her profile': (r) => r.status == 200 });

  payload = {
    role_name: "rd-director",
    relation: "*",
    object_namespace: "source-code",
    object_name: "*",
  };
  res = http.post(`${SERVER_URL}/role/add-permission`, JSON.stringify(payload), {headers:Headers});
  check(res, { '(Add Permission) RD-Director has full access to all source code': (r) => r.status == 200 });

  payload = {
    role_name: "rd",
    relation: "write",
    object_namespace: "source-code",
    object_name: "*",
  };
  res = http.post(`${SERVER_URL}/role/add-permission`, JSON.stringify(payload), {headers:Headers});
  check(res, { '(Add Permission) RD has write access to all source code': (r) => r.status == 200 });

  payload = {
    role_name: "rd-director",
    relation: "read",
    object_namespace: "sales-data",
    object_name: "*",
  };
  res = http.post(`${SERVER_URL}/role/add-permission`, JSON.stringify(payload), {headers:Headers});
  check(res, { '(Add Permission) RD-Director has read access to all sales data': (r) => r.status == 200 });

  payload = {
    role_name: "sales",
    relation: "*",
    object_namespace: "sales-data",
    object_name: "*",
  };
  res = http.post(`${SERVER_URL}/role/add-permission`, JSON.stringify(payload), {headers:Headers});
  check(res, { '(Add Permission) Sales has full access to all sales data': (r) => r.status == 200 });

  payload = {
    role_name: "rd",
    relation: "write",
    object_namespace: "sales-data",
    object_name: "1",
  };
  res = http.post(`${SERVER_URL}/role/add-permission`, JSON.stringify(payload), {headers:Headers});
  check(res, { '(Add Permission) RD has write access to sales-data:1': (r) => r.status == 200 });

  // Link permission
  payload = {
    object_namespace: "role",
    object_name: "rd",
    relation: "modify-permission",
    subjectset_namespace: "role",
    subjectset_name: "rd",
    subjectset_relation: "parent",
  };
  res = http.post(`${SERVER_URL}/permission/link`, JSON.stringify(payload), {headers:Headers});
  check(res, { '(Add Link Permission) RD-Director can modify permissions of RD': (r) => r.status == 200 });

  payload = {
    object_namespace: "source-code",
    object_name: "*",
    relation: "read",
    subjectset_namespace: "source-code",
    subjectset_name: "*",
    subjectset_relation: "write",
  };
  res = http.post(`${SERVER_URL}/permission/link`, JSON.stringify(payload), {headers:Headers});
  check(res, { '(Add Link Permission) Source Code read access implies write access': (r) => r.status == 200 });

}
