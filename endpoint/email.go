package endpoint

import (
	"fmt"

	"../extras"
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
		extras.SendEmail(fmt.Sprintf("%s <%s>", mds.Members[m].Name, mds.Members[m].Email), "Moderation "+msg.MessageID, `Moderation
			👍 Approve .../moderation/`+msg.MessageID+`/approved/`+extras.GenerateModeratorLink(msg.MessageID, msg.CreatedAt, mds.Members[m].Email)+`

			👎 Disapproved .../moderation/`+msg.MessageID+`/disapproved/`+extras.GenerateModeratorLink(msg.MessageID, msg.CreatedAt, mds.Members[m].Email)+``)
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

	if append >= ep.minModApproves {
		gms, err := ep.db.GetGroup(msg.GroupID)
		if err != nil {
			fmt.Println(err)
			return
		}
		for m := range gms.Members {
			extras.SendEmail(fmt.Sprintf("%s <%s>", gms.Members[m].Name, gms.Members[m].Email), "Email de "+msg.Email, msg.Message+"\nTracking: "+"\nResponder: ")
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
