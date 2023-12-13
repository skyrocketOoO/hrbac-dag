# rbac (in processing...)


## How to use?
### Reserved words
@ # () :

### Add user to a group
group:B#members@(user:1)

### Init a group
no need to do anything because this is a tuple based approach

### Give permission to a group
file:a#view@(group:A#member)
file:a#view@(group:A#parent)

### Level the group
group:A#parent@(group:B#member)
group:A#parent@(group:B#parent)

----------------------------------------------------------------
group:B#parent@(group:C#member)

### Link permission
file:a#edit@(file:a#view)

### Admin
directly return true without zanzibar



