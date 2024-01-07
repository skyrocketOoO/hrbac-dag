import { group, check } from "k6"
import { BuildGraph } from "./benchmark/build_graph.js"
import http from 'k6/http';


export const options = {
  vus: 1,
}

export default function() {
  const SERVER_URL = "http://localhost:3000"
  const Headers = {
    'Content-Type': 'application/json',
  }
  const layer = 6, base = 6;

  let res = http.post(`${SERVER_URL}/relation/clear-all-relations`, null, {headers:Headers});
  check(res, { 'ClearAllRelations: status == 200': (r) => r.status == 200 });

  group("build graph", () => {
    BuildGraph(SERVER_URL, Headers, layer, base);
  })
}
