---
title: "Steampipe Table: zoom_role - Query Zoom Roles using SQL"
description: "Allows users to query Roles in Zoom, specifically the details of each role including its id, name, total members, and privileges, providing insights into role-based access control within the Zoom platform."
---

# Table: zoom_role - Query Zoom Roles using SQL

Zoom Roles is a feature within Zoom that allows administrators to assign permissions to users based on their role. It provides a way to manage access control and permissions for different users within the organization. Zoom Roles help to ensure that users have the appropriate access rights for their job function and responsibilities.

## Table Usage Guide

The `zoom_role` table provides insights into roles within Zoom. As an administrator, explore role-specific details through this table, including role id, name, total members, and privileges. Utilize it to uncover information about roles, such as those with specific permissions, the number of users assigned to each role, and the specific privileges associated with each role.

## Examples

### List all roles
Explore the various roles within your Zoom account, allowing you to manage and organize your team more effectively. This query helps in maintaining a clear hierarchy and understanding the permissions associated with each role.

```sql+postgres
select
  *
from
  zoom_role
order by
  name;
```

```sql+sqlite
select
  *
from
  zoom_role
order by
  name;
```

### Get a role by ID
Discover the segments that are associated with a specific role ID in Zoom to better manage permissions and responsibilities within your team. This can be particularly useful for administrators who need to understand the scope of a role for delegation or auditing purposes.

```sql+postgres
select
  *
from
  zoom_role
where
  id = '1'; -- Owner role
```

```sql+sqlite
select
  *
from
  zoom_role
where
  id = '1'; -- Owner role
```

### Get privileges for each role
Explore which privileges are assigned to each role within the Zoom platform to understand user permissions and rights. This can help in managing user access and ensuring appropriate security measures are in place.

```sql+postgres
select
  r.id,
  r.name,
  p
from
  zoom_role as r,
  jsonb_array_elements(r.privileges) as p
order by
  r.name,
  p;
```

```sql+sqlite
select
  r.id,
  r.name,
  p.value as p
from
  zoom_role as r,
  json_each(r.privileges) as p
order by
  r.name,
  p.value;
```

### Find all roles with permission to edit account settings
Identify roles that have the authority to modify account settings, with the intent of understanding which roles have this level of access and how many members each role contains.

```sql+postgres
select
  id,
  name,
  total_members
from
  zoom_role
where
  privileges ? 'AccountSetting:Edit'
order by
  total_members;
```

```sql+sqlite
select
  id,
  name,
  total_members
from
  zoom_role
where
  json_extract(privileges, '$."AccountSetting:Edit"') is not null
order by
  total_members;
```