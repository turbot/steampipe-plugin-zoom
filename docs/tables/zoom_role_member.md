---
title: "Steampipe Table: zoom_role_member - Query Zoom Role Members using SQL"
description: "Allows users to query Zoom Role Members, providing detailed information about each member's role within the Zoom platform."
---

# Table: zoom_role_member - Query Zoom Role Members using SQL

Zoom is a cloud-based video communications app that allows you to set up virtual video and audio conferencing, webinars, live chats, screen-sharing, and other collaborative capabilities. You can have one-on-one meetings or host meetings with up to hundreds of participants. Zoom is used by businesses of all sizes, and both individuals and businesses can schedule meetings and invite guests.

## Table Usage Guide

The `zoom_role_member` table provides insights into the roles assigned to members within the Zoom platform. As a system administrator or team lead, you can use this table to understand the permissions and capabilities assigned to each member, aiding in the management and organization of your team's Zoom usage. Use it to monitor role assignments, verify member permissions, and ensure the appropriate distribution of access rights.

**Important Notes**
- You must specify the `role_id` in the `where` clause to query this table.

## Examples

### List all Owners in the account
Discover the segments that contain all account owners. This is particularly useful when you need to understand the distribution of responsibilities and roles within your account.

```sql+postgres
select
  *
from
  zoom_role_member
where
  role_id = '0';
```

```sql+sqlite
select
  *
from
  zoom_role_member
where
  role_id = '0';
```

### List all members of all roles
Explore the organizational structure of your Zoom account by understanding the allocation of roles to users. This query is useful for auditing user access and ensuring appropriate permissions are assigned.

```sql+postgres
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
  m.last_name;
```

```sql+sqlite
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
  m.last_name;
```

### List all roles that have dwight@dundermifflin.com as a member
Uncover the details of all roles associated with a specific email address to better understand user permissions and group associations. This can be particularly useful in auditing user access and managing role-based access control.

```sql+postgres
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
  r.name;
```

```sql+sqlite
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
  r.name;
```