---
title: "Steampipe Table: zoom_cloud_recording - Query Zoom Cloud Recordings using SQL"
description: "Allows users to query Zoom Cloud Recordings, specifically the details of each recording such as the meeting ID, recording start and end times, file type, and download URL."
---

# Table: zoom_cloud_recording - Query Zoom Cloud Recordings using SQL

Zoom Cloud Recording is a feature within Zoom Meetings that allows you to record your meetings and webinars, and securely store them in the cloud. It provides the ability to access, share, and download the recorded meetings and webinars. Zoom Cloud Recording helps you keep track of your meetings and webinars, ensuring that important information is not lost.

## Table Usage Guide

The `zoom_cloud_recording` table provides insights into Zoom Cloud Recordings within Zoom Meetings. As a meeting organizer or participant, explore recording-specific details through this table, including meeting IDs, recording start and end times, file types, and download URLs. Utilize it to manage and keep track of your recorded meetings and webinars, ensuring that important information is easily accessible and securely stored.

**Important Notes**
- You must specify the `user_id` in the `where` clause to query this table.
- You can use `start_time` to specify the range of meetings to return (max 30 days).
- Default time range is the last 30 days.

## Examples

### List all cloud recordings for a given host
Explore all cloud recordings associated with a specific host to better manage storage or review past meetings. This can be particularly useful for tracking meeting history, auditing content, or retrieving specific recordings for reference.

```sql+postgres
select
  *
from
  zoom_cloud_recording
where
  user_id = 'RCKlotFLRpe-Hbnv-VK3CA';
```

```sql+sqlite
select
  *
from
  zoom_cloud_recording
where
  user_id = 'RCKlotFLRpe-Hbnv-VK3CA';
```

### List all cloud recordings for a given host by email
Explore all cloud recordings associated with a specific host's email. This is useful for auditing purposes or for gathering insights into a particular user's activity on the platform.

```sql+postgres
select
  *
from
  zoom_cloud_recording
where
  user_id in (select id from zoom_user where email = 'dwight@dundermifflin.com');
```

```sql+sqlite
select
  *
from
  zoom_cloud_recording
where
  user_id in (select id from zoom_user where email = 'dwight@dundermifflin.com');
```

### List all cloud recordings for Sep 2019
Explore all cloud recordings from September 2019, specifically for a user with a particular email address. This is useful for auditing purposes or to find a specific recording from that period.

```sql+postgres
select
  *
from
  zoom_cloud_recording
where
  user_id in (select id from zoom_user where email = 'dwight@dundermifflin.com')
  and start_time >= '2019-09-01'
  and start_time <= '2021-09-30';
```

```sql+sqlite
select
  *
from
  zoom_cloud_recording
where
  user_id in (select id from zoom_user where email = 'dwight@dundermifflin.com')
  and start_time >= date('2019-09-01')
  and start_time <= date('2021-09-30');
```

### Get a cloud recording by ID
Discover the details of a specific Zoom cloud recording by using its unique identifier. This can be beneficial for auditing purposes, or to gain insights into a particular meeting or webinar.
To get a recording by ID it must be within the time range of the list. The
underlying API does not support getting a historical recording by ID.

```sql+postgres
select
  *
from
  zoom_cloud_recording
where
  id = 912357281124;
```

```sql+sqlite
select
  *
from
  zoom_cloud_recording
where
  id = 912357281124;
```