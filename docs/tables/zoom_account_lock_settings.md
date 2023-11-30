---
title: "Steampipe Table: zoom_account_lock_settings - Query Zoom Account Lock Settings using SQL"
description: "Allows users to query Zoom Account Lock Settings, specifically the account settings that are locked at the account level and cannot be modified by the user."
---

# Table: zoom_account_lock_settings - Query Zoom Account Lock Settings using SQL

Zoom is a cloud-based video communications app that allows you to set up virtual video and audio conferencing, webinars, live chats, screen-sharing, and other collaborative capabilities. With the Account Lock Settings, you can lock certain settings at the account level, preventing them from being modified by the user. This includes settings related to meetings, recordings, telephony, and more.

## Table Usage Guide

The `zoom_account_lock_settings` table provides insights into locked settings within Zoom's Account Lock Settings. As a system administrator, explore these locked settings through this table, including details on meetings, recordings, telephony, and more. Utilize it to manage and maintain control over the account settings, ensuring that they cannot be altered by the user.

## Examples

### Query all lock settings for the account
Explore the security settings of your account to understand how your meetings, recordings, and notifications are configured. This can help you identify areas for potential improvement or changes in your security preferences.
If true, then the setting is locked to the account level setting.


```sql
select
  jsonb_pretty(email_notification) as email_notification,
  jsonb_pretty(in_meeting) as in_meeting,
  jsonb_pretty(meeting_security) as meeting_security,
  jsonb_pretty(recording) as recording,
  jsonb_pretty(schedule_meeting) as schedule_meeting,
  jsonb_pretty(telephony) as telephony,
  jsonb_pretty(tsp) as tsp
from
  zoom_account_lock_settings
```

### Ensure join before host is set to disabled (CIS v1.1.2.4)
Explore which Zoom meeting settings allow participants to join before the host, to ensure compliance with the CIS v1.1.2.4 standard. This can help improve meeting security by preventing unauthorized access.
Ensure the setting is locked at account level:

```sql
select
  schedule_meeting ->> 'join_before_host'
from
  zoom_account_lock_settings
```