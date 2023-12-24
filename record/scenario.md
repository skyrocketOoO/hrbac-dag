# Scenario

## Domain

- User: Jimmy, Tasha, Ivy, Heidi
- Role: RD-Director, RD, Sales
- Object:
  - source-code: 1, 2
  - sales-data: 1, 2
  - profile: Ivy, Heidi

## Relatio

- Jimmy is a member of RD-Director
- Tasha is a member of RD
- Ivy is a member of Sales

- RD-Director is a parent of RD

## Permission

- Ivy has read access to her profile
- Heidi has read access to her profile
- RD-Director has full access to all source code
- RD has write access to all source code
- RD-Director has read access to all sales data
- Sales has full access to all sales data
- RD has write access to sales-data:1

## Link

- RD-Director can modify permissions of RD
- Source Code read access implies write access
