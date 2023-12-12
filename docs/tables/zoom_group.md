---
title: "Steampipe Table: zoom_group - Query Zoom Groups using SQL"
description: "Allows users to query Zoom Groups, specifically details about each group in a Zoom account, providing insights into group management and structure."
---

# Table: zoom_group - Query Zoom Groups using SQL

Zoom Groups is a feature within the Zoom video communications platform that allows account administrators to manage and organize users based on group. It provides a centralized way to apply settings and permissions to a group of users, facilitating efficient user management. Zoom Groups help administrators streamline the process of managing large user bases by grouping users based on department, role, or any other criteria.

## Table Usage Guide

The `zoom_group` table provides insights into Zoom Groups within the Zoom video communications platform. As an account administrator, explore group-specific details through this table, including group members, settings, and associated metadata. Utilize it to uncover information about groups, such as those with specific settings, the members within each group, and the overall structure of user organization within a Zoom account.

## Examples

### List all groups
Identify instances where you need to analyze all groups within your Zoom account, sorted by their names. This can provide a comprehensive overview of all groups, beneficial for administrative tasks and group management.

```sql+postgres
select
  *
from
  zoom_group
order by
  name
```

```sql+sqlite
select
  *
from
  zoom_group
order by
  name
```

### Get a group by ID
Explore the specific group details in the Zoom application by using a unique identifier. This can be particularly useful in managing and organizing large numbers of groups.

```sql+postgres
select
  *
from
  zoom_group
where
  id = 'siPuH6LvQfKSkeaVyxZCRQ'
```

```sql+sqlite
select
  *
from
  zoom_group
where
  id = 'siPuH6LvQfKSkeaVyxZCRQ'
```

### Groups with no members (to clean up?)
Discover the segments that consist of Zoom groups with no members, to determine areas that may require cleanup or removal. This helps streamline group management and optimise resource allocation.

```sql+postgres
select
  *
from
  zoom_group
where
  total_members = 0
```

```sql+sqlite
select
  *
from
  zoom_group
where
  total_members = 0
```