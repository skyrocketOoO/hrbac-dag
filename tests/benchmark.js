import { group } from "k6"
import { BuildGraph } from "./benchmark/build_graph.js"
import { Check } from "./benchmark/check.js"


export const options = {
  vus: 1,
}

export default function() {
  const SERVER_URL = "http://localhost:3000"
  const Headers = {
    'Content-Type': 'application/json',
  }
  const layer = 6, base = 6;

  group("check", () => {
    Check(SERVER_URL, Headers, layer, base);
  })
}
