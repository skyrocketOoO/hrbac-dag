import http from 'k6/http';
import { check, group, sleep } from 'k6';
import { TestUserAPI } from './api/user.js';
import { TestRepeatTuple } from './feature/repetition_test.js';
import { TestRoleAPI } from './api/role.js';
import { TestRelationAPI } from './api/relation.js';
import { TestObjectAPI } from './api/object.js';
import { TestAccessInheritance } from './feature/access_inheritance.js';
import { TestHRBAC } from './feature/hrbac.js';


export const options = {
  vus: 1,
}

export default function() {
  const SERVER_URL = "http://localhost:3000"
  const Headers = {
    'Content-Type': 'application/json',
  }

  ClearAllRelations(SERVER_URL)
  // healthy
  let res = http.get(`${SERVER_URL}/healthy`);
  check(res, { 'Server is healthy': (r) => r.status == 200 });

  group("api", () => {
    group("user", () => {
      TestUserAPI(SERVER_URL, Headers);
    });
    ClearAllRelations(SERVER_URL)
    group("role", () => {
      TestRoleAPI(SERVER_URL, Headers);
    });
    ClearAllRelations(SERVER_URL)
    group("relation", () => {
      TestRelationAPI(SERVER_URL, Headers);
    });
    ClearAllRelations(SERVER_URL)
    group("object", () => {
      TestObjectAPI(SERVER_URL, Headers);
    });
    ClearAllRelations(SERVER_URL)
  });

  group("feature", () => {
    group("Repeat tuple", () => {
      TestRepeatTuple(SERVER_URL, Headers);
    })
    ClearAllRelations(SERVER_URL)
    group("Access inheritance", () => {
      TestAccessInheritance(SERVER_URL, Headers);
    });
    ClearAllRelations(SERVER_URL)
    group("HRBAC", () => {
      TestHRBAC(SERVER_URL, Headers);
    });
    ClearAllRelations(SERVER_URL)
    // group("* syntax", () => {
    //   TestUniversalSyntax(SERVER_URL, Headers);
    // });
    // ClearAllRelations(SERVER_URL)
  });

  group("scenario", () => {
    // checkScenario(SERVER_URL, Headers)
  })
}

function ClearAllRelations(serverUrl){
  http.post(`${serverUrl}/relation/clear-all-relations`, null, null);
}
