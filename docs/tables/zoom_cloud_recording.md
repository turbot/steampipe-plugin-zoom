# Table: zoom_recording

Query information about cloud recordings of meetings in the Zoom account.

Notes:
* A `user_id` must be provided in all queries to this table.
* Use `start_time` to specify the range of meetings to return (max 30 days).
* Default time range is the last 30 days.

## Examples

### List all cloud recordings for a given host

```sql
select
  *
from
  zoom_cloud_recording
where
  user_id = 'RCKlotFLRpe-Hbnv-VK3CA'
```

### List all cloud recordings for a given host by email

```sql
select
  *
from
  zoom_cloud_recording
where
  user_id in (select id from zoom_user where email = 'dwight@dundermifflin.com')
```

### List all cloud recordings for a Sep 2019

```sql
select
  *
from
  zoom_cloud_recording
where
  user_id in (select id from zoom_user where email = 'dwight@dundermifflin.com')
  and start_time >= '2019-09-01'
  and start_time <= '2021-09-30'
```

### Get a cloud recording by ID

To get a recording by ID it must be within the time range of the list. The
underlying API does not support getting a historical recording by ID.

```sql
select
  *
from
  zoom_cloud_recording
where
  id = 912357281124
```
