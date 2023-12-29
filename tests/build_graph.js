import { group } from "k6"
import { BuildGraph } from "./benchmark/build_graph.js"


export const options = {
  vus: 1,
}

export default function() {
  const SERVER_URL = "http://localhost:3000"
  const Headers = {
    'Content-Type': 'application/json',
  }
  const layer = 7, base = 6;

  group("build graph", () => {
    BuildGraph(SERVER_URL, Headers, layer, base);
  })

}
