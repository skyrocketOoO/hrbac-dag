# Feature record

- avoid infinite recursion
  - if has cycle in DAG, it will cause the issue.
  - resolution:
    - detect cycle on create
      - pros: the original method
      - cons: each create need to check entire graph
    - limited the search length
      - pros: easy to implement and use
      - cons: may loss the permission out of the length, hard to define the limited length
  