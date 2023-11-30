---
title: "Steampipe Table: zoom_group_member - Query Zoom Group Members using SQL"
description: "Allows users to query Group Members in Zoom, providing insights into group composition, member roles, and member statuses."
---

# Table: zoom_group_member - Query Zoom Group Members using SQL

Zoom is a cloud-based video communications platform that offers video conferencing, audio conferencing, chat, and webinar features. It allows users to schedule, join, and host meetings across various devices. A Zoom Group Member is an individual who is part of a specific group in Zoom, and their details can include their role in the group, their status, and other related information.

## Table Usage Guide

The `zoom_group_member` table provides insights into group members within Zoom's cloud-based video communications platform. As a systems administrator, explore group member-specific details through this table, including their roles, statuses, and other associated metadata. Utilize it to uncover information about group composition, such as the roles of each member, their statuses, and the verification of their details.

## Examples

### List all members of the Sales group
Discover the segments that make up the Sales group in your organization. This can help you gain insights into the team composition and effectively manage resources.

```sql
select
  *
from
  zoom_group_member
where
  group_id in (select id from zoom_group where name = 'Sales')
```

### List all members of all groups
Explore which members belong to which groups in order to understand the distribution of users across different teams or categories. This could be useful in managing team resources, identifying collaboration opportunities, or assessing group-based access controls.

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
Explore which groups within a communication platform include a specific user. This is useful in understanding the user's team involvement and collaboration channels.

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
Explore which Zoom users are not attached to any group. This can be useful for identifying individuals who may be missing out on group-specific communications or functionalities.

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