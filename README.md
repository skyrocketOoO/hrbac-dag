# HRBAC-zanzibar
HRBAC implement using Google Zanzibar algorithm

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
- [x] Object *(all object | all operations) support (unsupport distinct object namespace link)
- [ ] List who has access to object
- [ ] Regex

## Restriction

### Reserved words

- All: @ # ( ) :
- Relatoion: "member" "parent"
- Namespace: "role", "user"
- Name: "admin"
  