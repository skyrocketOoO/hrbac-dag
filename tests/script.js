import http from 'k6/http';
import { check, group, sleep } from 'k6';
import { checkAPI } from './api/api_test.js';
import { checkScenario } from './scenario/scenario_test.js';

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
    checkAPI(SERVER_URL, Headers)
  })

  group("scenario", () => {
    checkScenario(SERVER_URL, Headers)
  })
}
