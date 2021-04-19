# Table: zoom_group

Query information about groups in the Zoom account.

## Examples

### List all groups

```sql
select
  *
from
  zoom_group
order by
  name
```

### Get a group by ID

```sql
select
  *
from
  zoom_group
where
  id = 'siPuH6LvQfKSkeaVyxZCRQ'
```

### Groups with no members (to clean up?)

```sql
select
  *
from
  zoom_group
where
  total_members = 0
```
