# Table: zoom_account_settings

Query settings for the Zoom account.

## Examples

### Get all settings for the current account

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

Check the setting at account level:
```sql
select
  security -> 'password_requirement' ->> 'minimum_password_length' as minimum_password_length
from
  zoom_account_settings
```

### Ensure join before host is set to disabled (CIS v1.1.2.4)

Check the setting at account level:
```sql
select
  schedule_meeting ->> 'join_before_host' as join_before_host
from
  zoom_account_settings
```
