import http from 'k6/http';
import { check, group, sleep } from 'k6';
import { checkAPI } from './api/api_test.js';
import { checkScenario } from './scenario/scenario_test.js';
import { TestUserAPI } from './api/user.js';
import { TestRepeatTuple } from './feature/repetition_test.js';
import { TestRoleAPI } from './api/role.js';
import { TestRelationAPI } from './api/relation.js';


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
    // checkAPI(SERVER_URL, Headers)
  });

  group("feature", () => {
    // TestRepeatTuple(SERVER_URL, Headers);
  });

  // group("scenario", () => {
  //   checkScenario(SERVER_URL, Headers)
  // })
}
