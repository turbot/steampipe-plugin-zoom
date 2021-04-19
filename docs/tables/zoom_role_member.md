# Table: zoom_role_member

Query information about members of a given role in the Zoom account.

Note: The `role_id` field must be set in the `where` clause.

## Examples

### List all Owners in the account

```sql
select
  *
from
  zoom_role_member
where
  role_id = '0'
```

### List all members of all roles

```sql
select
  r.id,
  r.name,
  m.user_id,
  m.first_name,
  m.last_name
from
  zoom_role as r,
  zoom_role_member as m
where
  r.id = m.role_id
order by
  r.name,
  m.first_name,
  m.last_name
```

### List all roles that have dwight@dundermifflin.com as a member

```sql
select
  r.id,
  r.name
from
  zoom_role as r,
  zoom_role_member as m
where
  r.id = m.role_id
  and m.email = 'dwight@dundermifflin.com'
order by
  r.name
```
