# public-feedback-api

## Environment required

* PORT
* MONGODB
* SECRET
* GOOGLECERT
* GOOGLEGROUP
* MINAPPROVED
* MAILDOMAIN
* MAILAPIKEY

## Endpoints

* POST /message - New message
* POST /message/confirm - Confirm message
* GET /message/resend - Resend confirmation
* GET /message/[ID]/ - Get a message
* GET /groups - List groups
* GET /status - Get status API
* GET /tracking/[ID]/[CODE]/pixel.gif - Pixel Tracking

## Feedback

* POST /feedback/[ID]/[CODE] - Send a reply message to the author

## Administration

_Emails and moderators are controlled with Google Spreadsheets_

* GET /moderation/[ID]/approved/[CODE] - Approve a message
* GET /moderation/[ID]/disapproved/[CODE] - Disapprove a message