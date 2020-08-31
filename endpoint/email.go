package endpoint

import (
	"fmt"

	"github.com/jonathanhecl/public-feedback-api/extras"
)

func EmailUserConfirmation(MessageID string) {

	msg, err := ep.db.GetMessage(MessageID)
	if err != nil {
		fmt.Println(err)
		return
	}

	extras.SendEmail(fmt.Sprintf("%s <%s>", msg.Name, msg.Email), "Confirmation "+msg.MessageID, "🔑 Confirmation Code: "+msg.ConfirmationCode)

	return

}

func EmailModerationWait(MessageID string) {

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
		code := extras.GenerateModeratorLink(msg.MessageID, msg.CreatedAt, mds.Members[m].Email)
		extras.SendEmail(fmt.Sprintf("%s <%s>", mds.Members[m].Name, mds.Members[m].Email), "Moderation "+msg.MessageID, `Moderation
			👍 Approve .../moderation/`+msg.MessageID+`/approved/`+code+`

			👎 Disapproved .../moderation/`+msg.MessageID+`/disapproved/`+code+``)
	}

	return

}

func EmailModerationConfirm(MessageID string) {

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

	vmsg, err := ep.db.GetModerationVote(MessageID)
	if err != nil {
		fmt.Println(err)
		return
	}

	votes := 0
	approve := 0
	for v := range vmsg.Votes {
		votes++
		if vmsg.Votes[v].IsApprove {
			approve++
		}
	}

	if approve >= ep.minModApproves {
		gms, err := ep.db.GetGroup(msg.ToGroup)
		if err != nil {
			fmt.Println(err)
			return
		}
		for m := range gms.Members {
			code := extras.GenerateMemberLink(msg.MessageID, msg.CreatedAt, gms.Members[m].Email)
			extras.SendEmail(fmt.Sprintf("%s <%s>", gms.Members[m].Name, gms.Members[m].Email), "Email de "+msg.Email, msg.Message+`
			Tracking .../tracking/`+msg.MessageID+`/`+code+`/pixel.gif

			Reply .../feedback/`+msg.MessageID+`/`+code+`/`)
		}
		err = ep.db.SetMessageSended(msg.MessageID)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	if votes >= len(mds.Members) {
		err = ep.db.SetMessageClosed(msg.MessageID)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
