---
title: "Steampipe Table: zoom_account_settings - Query Zoom Account Settings using SQL"
description: "Allows users to query Zoom Account Settings, specifically to retrieve and analyze the configuration and preferences of a Zoom account."
---

# Table: zoom_account_settings - Query Zoom Account Settings using SQL

Zoom Account Settings is a resource in Zoom that allows users to manage and customize their account's configuration and preferences. This includes settings related to scheduling, recording, telephony, and more. Zoom Account Settings provide a centralized way to manage and understand the behavior and controls of your Zoom account.

## Table Usage Guide

The `zoom_account_settings` table provides insights into Zoom Account Settings. As a Zoom administrator or IT professional, explore account-specific details through this table, including scheduling preferences, recording settings, telephony options, and other configuration details. Utilize it to understand and manage the behavior and controls of your Zoom account, ensuring it aligns with organizational policies and preferences.

## Examples

### Get all settings for the current account
Explore the various settings associated with your account to better understand its configuration and features. This can help in identifying areas for improvement, ensuring optimal security and functionality.

```sql
select
  jsonb_pretty(email_notification) as email_notification,
  jsonb_pretty(feature) as feature,
  jsonb_pretty(in_meeting) as in_meeting,
  jsonb_pretty(integration) as integration,
  jsonb_pretty(managed_domains) as managed_domains,
  jsonb_pretty(meeting_authentication) as meeting_authentication,
  jsonb_pretty(meeting_security) as meeting_security,
  jsonb_pretty(recording) as recording,
  jsonb_pretty(recording_authentication) as recording_authentication,
  jsonb_pretty(schedule_meeting) as schedule_meeting,
  jsonb_pretty(security) as security,
  jsonb_pretty(telephony) as telephony,
  jsonb_pretty(trusted_domains) as trusted_domains,
  jsonb_pretty(tsp) as tsp
from
  zoom_account_settings
```

### Ensure a Personal Meeting ID (PMI) is not used for meetings by default
This example helps you verify whether your Zoom account settings are configured to avoid using Personal Meeting IDs (PMIs) for instant and scheduled meetings by default. It's beneficial for maintaining meeting security, as PMIs, if reused, can be exploited by unauthorized individuals to gain access to meetings.
It's more secure to use a random meeting ID rather than using your same PMI
for every meeting, particularly for public participants.

Ensure these options are false to confirm the PMI is not being used for
meetings by default:


```sql
select
  schedule_meeting->>'use_pmi_for_instant_meetings' as use_pmi_for_instant_meetings,
  schedule_meeting->>'use_pmi_for_scheduled_meetings' as use_pmi_for_scheduled_meetings
from
  zoom_account_settings
```

Check that these settings are locked for all users in the account:

```sql
select
  schedule_meeting->>'personal_meeting' as personal_meeting
from
  zoom_account_lock_settings
```

### Ensure minimum passcode length is set to at least 6 characters (CIS v1.1.1.1.1)
Determine if your Zoom account settings comply with CIS v1.1.1.1.1 by checking if the minimum passcode length is set to at least 6 characters. This is useful for maintaining security standards and preventing unauthorized access.
Check the setting at account level:

```sql
select
  security -> 'password_requirement' ->> 'minimum_password_length' as minimum_password_length
from
  zoom_account_settings
```

### Ensure join before host is set to disabled (CIS v1.1.2.4)
Determine if the Zoom meeting settings allow participants to join before the host. This is critical in ensuring the security and privacy of the meeting by preventing unauthorized access before the host arrives.
Check the setting at account level:

```sql
select
  schedule_meeting ->> 'join_before_host' as join_before_host
from
  zoom_account_settings
```