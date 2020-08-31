package endpoint

import (
	"fmt"

	"github.com/jonathanhecl/public-feedback-api/extras"
)

func EmailConfirmation(MessageID string) {

	msg, err := ep.db.GetMessage(MessageID)
	if err != nil {
		fmt.Println(err)
		return
	}

	extras.SendEmail(fmt.Sprintf("%s <%s>", msg.Name, msg.Email), "Confirmation "+msg.MessageID, "🔑 Confirmation Code: "+msg.ConfirmationCode)

	return

}

func EmailWaitModeration(MessageID string) {

	msg, err := ep.db.GetMessage(MessageID)
	if err != nil {
		fmt.Println(err)
		return
	}

	mds, err := ep.db.GetGroup("MOD")
	if err != nil {
		fmt.Println(err)
		return
	}

	for m := range mds.Members {
		extras.SendEmail(fmt.Sprintf("%s <%s>", mds.Members[m].Name, mds.Members[m].Email), "Moderation "+msg.MessageID, `Moderation
			👍 Approve .../moderation/`+msg.MessageID+`/approved/`+extras.GenerateModeratorLink(msg.MessageID, msg.CreatedAt, mds.Members[m].Email)+`

			👎 Disapproved .../moderation/`+msg.MessageID+`/disapproved/`+extras.GenerateModeratorLink(msg.MessageID, msg.CreatedAt, mds.Members[m].Email)+``)
	}

	return

}
