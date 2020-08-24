# public-feedback-api

## Endpoints

* POST /message - New message
* POST /message/confirm - Confirm message
* GET /message/resend - Resend confirmation
* GET /groups - List groups

```
    **Admin**
```
Emails and moderators are controlled by Google Spreadsheets
* GET /moderation/[ID]/approved/[CODE] - Approve a message
* GET /moderation/[ID]/disapproved/[CODE] - Disapprove a message