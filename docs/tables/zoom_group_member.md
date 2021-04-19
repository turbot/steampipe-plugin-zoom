# Table: zoom_group_member

Query information about members of a given group in the Zoom account.

Note: The `group_id` field must be set in the `where` clause.

## Examples

### List all members of the Sales group

```sql
select
  *
from
  zoom_group_member
where
  group_id in (select id from zoom_group where name = 'Sales')
```

### List all members of all groups

```sql
select
  g.id,
  g.name,
  m.user_id,
  m.first_name,
  m.last_name
from
  zoom_group as g,
  zoom_group_member as m
where
  g.id = m.group_id
order by
  g.name,
  m.first_name,
  m.last_name
```

### List all groups that have dwight@dundermifflin.com as a member

```sql
select
  g.id,
  g.name
from
  zoom_group as g,
  zoom_group_member as m
where
  g.id = m.group_id
  and m.email = 'dwight@dundermifflin.com'
order by
  g.name
```

### List all users that are not a member of any group

```sql
select
  u.email
from
  zoom_user as u
where
  u.id not in
  (
    select distinct
      m.user_id
    from
      zoom_group as g,
      zoom_group_member as m
    where
      g.id = m.group_id
  )
order by
  u.email
```
