import http from 'k6/http';
import { check, group, sleep } from 'k6';
import { TestUserAPI } from './api/user.js';
import { TestRepeatTuple } from './feature/repetition_test.js';
import { TestRoleAPI } from './api/role.js';
import { TestRelationAPI } from './api/relation.js';
import { TestObjectAPI } from './api/object.js';
import { TestAccessInheritance } from './feature/access_inheritance.js';
import { TestHRBAC } from './feature/hrbac.js';
import { TestUniversalSyntax } from './feature/regex_*_test.js';
import { TestCycle } from './scenario/cycle_test.js';

export const options = {
  vus: 1,
}

export default function() {
  const SERVER_URL = "http://localhost:3000"
  const Headers = {
    'Content-Type': 'application/json',
  }

  http.post(`${SERVER_URL}/relation/clear-all-relations`, null, {headers:Headers});
  // healthy
  let res = http.get(`${SERVER_URL}/healthy`);
  check(res, { 'Server is healthy': (r) => r.status == 200 });

  group("api", () => {
    group("user", () => {
      TestUserAPI(SERVER_URL, Headers);
    });
    http.post(`${SERVER_URL}/relation/clear-all-relations`, null, {headers:Headers});
    group("role", () => {
      TestRoleAPI(SERVER_URL, Headers);
    });
    http.post(`${SERVER_URL}/relation/clear-all-relations`, null, {headers:Headers});
    group("relation", () => {
      TestRelationAPI(SERVER_URL, Headers);
    });
    http.post(`${SERVER_URL}/relation/clear-all-relations`, null, {headers:Headers});
    group("object", () => {
      TestObjectAPI(SERVER_URL, Headers);
    });
    http.post(`${SERVER_URL}/relation/clear-all-relations`, null, {headers:Headers});
  });

  group("feature", () => {
    group("Repeat tuple", () => {
      TestRepeatTuple(SERVER_URL, Headers);
    })
    http.post(`${SERVER_URL}/relation/clear-all-relations`, null, {headers:Headers});
    group("Access inheritance", () => {
      TestAccessInheritance(SERVER_URL, Headers);
    });
    http.post(`${SERVER_URL}/relation/clear-all-relations`, null, {headers:Headers});
    group("HRBAC", () => {
      TestHRBAC(SERVER_URL, Headers);
    });
    http.post(`${SERVER_URL}/relation/clear-all-relations`, null, {headers:Headers});
    // group("* syntax", () => {
    //   TestUniversalSyntax(SERVER_URL, Headers);
    // });
    // http.post(`${SERVER_URL}/relation/clear-all-relations`, null, {headers:Headers});
  });

  group("scenario", () => {
    // checkScenario(SERVER_URL, Headers)
  })
}
