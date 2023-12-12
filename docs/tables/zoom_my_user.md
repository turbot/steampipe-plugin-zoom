---
title: "Steampipe Table: zoom_my_user - Query Zoom Users using SQL"
description: "Allows users to query Zoom Users, specifically their profile details and settings, providing insights into user management and account configurations."
---

# Table: zoom_my_user - Query Zoom Users using SQL

Zoom is a cloud-based video communications app that allows you to set up virtual video and audio conferencing, webinars, live chats, screen-sharing, and other collaborative capabilities. You can have up to 1,000 video participants and 10,000 viewers. Zoom is used by businesses for remote team meetings, webinars, demonstrations, virtual training, and more.

## Table Usage Guide

The `zoom_my_user` table provides insights into user profiles within Zoom. As a system administrator, explore user-specific details through this table, including profile settings, account configurations, and associated metadata. Utilize it to uncover information about users, such as their role in the organization, their account status, and the configuration of their Zoom settings.

## Examples

### List user information
Explore the details of all user profiles within your Zoom account. This can help in managing users and understanding the overall utilization of your Zoom resources.

```sql+postgres
select
  *
from
  zoom_my_user;
```

```sql+sqlite
select
  *
from
  zoom_my_user;
```