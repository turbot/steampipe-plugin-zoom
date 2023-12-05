---
title: "Steampipe Table: zoom_user - Query Zoom Users using SQL"
description: "Allows users to query Zoom Users, specifically the information about each user including email, first name, last name, and status. This table provides insights into the user management within Zoom."
---

# Table: zoom_user - Query Zoom Users using SQL

Zoom is a cloud-based video communications app that allows you to set up virtual video and audio conferencing, webinars, live chats, screen-sharing, and other collaborative capabilities. Users are the individuals who have registered by providing their email, name, and other personal information to use this platform. Zoom provides a comprehensive set of user management features, including user roles, statuses, settings, and more.

## Table Usage Guide

The `zoom_user` table provides insights into user management within Zoom. As an IT admin, explore user-specific details through this table, including email, first name, last name, and status. Utilize it to uncover information about users, such as their roles, statuses, settings, and the verification of their profiles.

## Examples

### List all users
Explore all registered users within your Zoom account to ensure proper management and oversight. This can be particularly useful for administrators seeking a comprehensive overview of all account users.

```sql+postgres
select
  *
from
  zoom_user;
```

```sql+sqlite
select
  *
from
  zoom_user;
```

### Get a user by ID
Explore the details of a specific user in the Zoom platform. This is particularly useful when you need to assess user-related issues or understand their activity within the system.

```sql+postgres
select
  *
from
  zoom_user
where
  id = 'RCKlotFLRpe-Hbnv-VK3CA';
```

```sql+sqlite
select
  *
from
  zoom_user
where
  id = 'RCKlotFLRpe-Hbnv-VK3CA';
```

### Get a user by email
Discover the details of a specific user in your Zoom account by using their email address. This can be particularly useful for administrators who need to manage or track user activities.

```sql+postgres
select
  *
from
  zoom_user
where
  email = 'dwight@dundermifflin.com';
```

```sql+sqlite
select
  *
from
  zoom_user
where
  email = 'dwight@dundermifflin.com';
```

### Most recently created users
Discover the most recent additions to your Zoom user base. This query helps you keep track of new users, allowing for timely onboarding and account management.

```sql+postgres
select
  *
from
  zoom_user
order by
  created_at desc
limit 5;
```

```sql+sqlite
select
  *
from
  zoom_user
order by
  created_at desc
limit 5;
```

### Users by license type
Discover the segments that are most common among Zoom users based on their license type. This can help prioritize resources and tailor services for the most prevalent user categories.

```sql+postgres
select
  type,
  count(*)
from
  zoom_user
group by
  type
order by
  count desc;
```

```sql+sqlite
select
  type,
  count(*)
from
  zoom_user
group by
  type
order by
  count(*) desc;
```

### Users with Security:Edit permission
Explore which Zoom users have the ability to edit security settings, a useful query for managing access control and maintaining security protocols within your organization.

```sql+postgres
select
  *
from
  zoom_user
where
  permissions ? 'Security:Edit';
```

```sql+sqlite
select
  *
from
  zoom_user
where
  json_extract(permissions, '$.Security:Edit') is not null;
```