---
title: "Steampipe Table: zoom_meeting - Query Zoom Meetings using SQL"
description: "Allows users to query Zoom Meetings, specifically the details of each meeting including the ID, topic, type, start time, duration, timezone, and status."
---

# Table: zoom_meeting - Query Zoom Meetings using SQL

Zoom Meetings is a service offered by Zoom that allows users to host or join virtual meetings with high-quality video, audio, and screen sharing options. It supports a variety of meeting types, including one-on-one meetings, group video conferences, and webinars. Zoom Meetings allows for easy scheduling, recording, and sharing of meetings.

## Table Usage Guide

The `zoom_meeting` table provides insights into Zoom Meetings within the Zoom platform. As a system administrator or a team manager, explore meeting-specific details through this table, including meeting ID, topic, type, start time, duration, timezone, and status. Utilize it to monitor and manage the meetings occurring within your organization, ensuring efficient communication and collaboration.

**Important Notes**
- You must specify the `user_id` in the `where` clause to query this table.

## Examples

### List all meetings for a given host
Explore all meetings organized by a specific host. This is useful for auditing purposes, allowing you to understand the host's meeting patterns and frequency.

```sql+postgres
select
  *
from
  zoom_meeting
where
  user_id = 'RCKlotFLRpe-Hbnv-VK3CA';
```

```sql+sqlite
select
  *
from
  zoom_meeting
where
  user_id = 'RCKlotFLRpe-Hbnv-VK3CA';
```

### List all meetings for a given host by email
Determine the meetings conducted by a specific host using their email address. This is useful for tracking a host's meeting history and understanding their engagement levels.

```sql+postgres
select
  *
from
  zoom_meeting
where
  user_id in (select id from zoom_user where email = 'dwight@dundermifflin.com');
```

```sql+sqlite
select
  *
from
  zoom_meeting
where
  user_id in (select id from zoom_user where email = 'dwight@dundermifflin.com');
```

### Get a meeting by ID
Discover the details of a specific meeting using its unique identifier, which can be useful for auditing purposes or for gathering information about a past meeting.

```sql+postgres
select
  *
from
  zoom_meeting
where
  id = 912357281124;
```

```sql+sqlite
select
  *
from
  zoom_meeting
where
  id = 912357281124;
```