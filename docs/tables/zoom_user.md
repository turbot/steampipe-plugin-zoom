# Table: zoom_user

Query information about users in the Zoom account.

## Examples

### List all users

```sql
select
  *
from
  zoom_user
```

### Get a user by ID

```sql
select
  *
from
  zoom_user
where
  id = 'RCKlotFLRpe-Hbnv-VK3CA'
```

### Get a user by email

```sql
select
  *
from
  zoom_user
where
  email = 'dwight@dundermifflin.com'
```

### Most recently created users

```sql
select
  *
from
  zoom_user
order by
  created_at desc
limit 5
```

### Users by license type

```sql
select
  type,
  count(*)
from
  zoom_user
group by
  type
order by
  count desc
```

### Users with Security:Edit permission

```sql
select
  *
from
  zoom_user
where
  permissions ? 'Security:Edit'
```
