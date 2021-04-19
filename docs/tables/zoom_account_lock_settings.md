# Table: zoom_account_lock_settings

Accounts may lock settings to prevent them being changed at the user, group or meeting level.
This table represents those lock settings.

## Examples

### Query all lock settings for the account

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

Ensure the setting is locked at account level:
```sql
select
  schedule_meeting ->> 'join_before_host'
from
  zoom_account_lock_settings
```
