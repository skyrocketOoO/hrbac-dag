import http from 'k6/http';
import { check, group, sleep } from 'k6';
import { TestUserAPI } from './api/user.js';
import { TestRepeatTuple } from './feature/repetition_test.js';
import { TestRoleAPI } from './api/role.js';
import { TestRelationAPI } from './api/relation.js';
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

  // healthy
  let res = http.get(`${SERVER_URL}/healthy`);
  check(res, { 'Server is healthy': (r) => r.status == 200 });

  group("api", () => {
    group("user", () => {
      TestUserAPI(SERVER_URL, Headers);
    });
    group("role", () => {
      TestRoleAPI(SERVER_URL, Headers);
    });
    group("relation", () => {
      TestRelationAPI(SERVER_URL, Headers);
    });
  });

  group("clear", () => {
    res = http.del(`${SERVER_URL}/relation/`, null, {headers:Headers});
    check(res, { 'ClearAllRelations: status == 200': (r) => r.status == 200 });
  });

  group("feature", () => {
    group("Repeat tuple", () => {
      TestRepeatTuple(SERVER_URL, Headers);
    })
    group("Access inheritance", () => {
      TestAccessInheritance(SERVER_URL, Headers);
    });
    group("HRBAC", () => {
      TestHRBAC(SERVER_URL, Headers);
    });
    group("* syntax", () => {
      TestUniversalSyntax(SERVER_URL, Headers);
    });
  });

  group("clear2", () => {
    res = http.del(`${SERVER_URL}/relation/`, null, {headers:Headers});
    check(res, { 'ClearAllRelations: status == 200': (r) => r.status == 200 });
  });

  group("scenario", () => {
    // checkScenario(SERVER_URL, Headers)
    TestCycle(SERVER_URL, Headers);
  })
}
