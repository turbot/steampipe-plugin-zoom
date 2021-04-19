# Table: zoom_meeting

Query information about meetings in the Zoom account.

Note: A `user_id` must be provided in all queries to this table.

## Examples

### List all meetings for a given host

```sql
select
  *
from
  zoom_meeting
where
  user_id = 'RCKlotFLRpe-Hbnv-VK3CA'
```

### List all meetings for a given host by email

```sql
select
  *
from
  zoom_meeting
where
  user_id in (select id from zoom_user where email = 'dwight@dundermifflin.com')
```

### Get a meeting by ID

```sql
select
  *
from
  zoom_meeting
where
  id = 912357281124
```
