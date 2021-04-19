# Table: zoom_role

Query information about roles in the Zoom account.

## Examples

### List all roles

```sql
select
  *
from
  zoom_role
order by
  name
```

### Get a role by ID

```sql
select
  *
from
  zoom_role
where
  id = '1' -- Owner role
```

### Get privileges for each role

```sql
select
  r.id,
  r.name,
  p
from
  zoom_role as r,
  jsonb_array_elements(r.privileges) as p
order by
  r.name,
  p
```

### Find all roles with permission to edit account settings

```sql
select
  id,
  name,
  total_members
from
  zoom_role
where
  privileges ? 'AccountSetting:Edit'
order by
  total_members
```
