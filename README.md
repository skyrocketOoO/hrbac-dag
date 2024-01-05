# HRBAC-DAG

HRBAC implement using DAG

## Feature

- [x] Admin(all permission, only implement for check, no path or list)
- [x] User Permission
- [x] RBAC
- [x] HRBAC
- [x] Access link (same object namespace, role, no user)
- [x] Access Inheritance
- [x] Children role access control
- [x] Zero trust
- [x] Fine grained
- [x] Multiple role
- [ ] Object *(all object | all operations) support (unsupport distinct object namespace link)
- [ ] List who has access to object
- [ ] Regex
- [ ] Temporal Constraints

## Restriction

### Reserved words

- All: @ # ( ) : %
- Relatoion: "member" "parent" "modify-permission"
- Namespace: "role", "user"
- Name: "admin"

## Development benchmark

[Link](https://docs.google.com/spreadsheets/d/1RLyWh62_trEEWyLYD34sX4jUrBSGOLbnxi7ZRWubi5s/edit?usp=sharing)
  