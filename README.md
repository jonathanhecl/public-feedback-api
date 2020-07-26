# public-feedback-api

## Endpoints

* POST /message - New message
* POST /message/confirm - Confirm message
* GET /message/resend - Resend confirmation
* GET /groups - List groups

```
    **Admin**
    Authorization by header required
```
* GET /admin/groups - List groups with details
* POST /admin/groups - New or edit a group
* DELETE /admin/groups - Delete a group