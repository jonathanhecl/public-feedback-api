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

	data := make(map[string]string)
	data["Name"] = msg.Name
	data["Message"] = msg.Message
	data["URL"] = "https://" + extras.GetWebDomain() + "/confirm/" + msg.MessageID + "?verify=" + msg.ConfirmationCode

	t := ParseTemplate("confirm", data)
	if len(t) != 0 {
		extras.SendEmail(msg.Email, "Confirma tu mensaje", t)
	}

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
		//fmt.Sprintf("%s <%s>", mds.Members[m].Name, mds.Members[m].Email)

		data := make(map[string]string)
		data["Name"] = msg.Name
		data["Email"] = msg.Email
		data["Message"] = msg.Message
		data["Moderator"] = mds.Members[m].Name
		data["URLApprove"] = "https://" + extras.GetWebDomain() + "/moderation/" + msg.MessageID + "?approved=" + code
		data["URLDisapprove"] = "https://" + extras.GetWebDomain() + "/moderation/" + msg.MessageID + "?disapproved=" + code

		t := ParseTemplate("moderation", data)
		if len(t) != 0 {
			extras.SendEmail(mds.Members[m].Email, "Acción de moderación requerida", t)
		}
	}

	return

}

func EmailModerationConfirm(MessageID string) {

	msg, err := ep.db.GetMessage(MessageID)
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
	disapprove := 0
	for v := range vmsg.Votes {
		votes++
		if vmsg.Votes[v].IsApprove {
			approve++
		} else if !vmsg.Votes[v].IsApprove {
			disapprove++
		}
	}

	if approve >= ep.minModApproves {
		data := make(map[string]string)
		data["Name"] = msg.Name
		data["Message"] = msg.Message
		t := ParseTemplate("success", data)
		if len(t) != 0 {
			extras.SendEmail(msg.Email, "Tu mensaje ha sido enviado con exito!", t)
		}
		gms, err := ep.db.GetGroup(msg.ToGroup)
		if err != nil {
			fmt.Println(err)
			return
		}
		for m := range gms.Members {
			code := extras.GenerateMemberLink(msg.MessageID, msg.CreatedAt, gms.Members[m].Email)
			// fmt.Sprintf("%s <%s>", gms.Members[m].Name, gms.Members[m].Email)

			data := make(map[string]string)
			data["Name"] = gms.Members[m].Name
			data["Message"] = msg.Message
			data["URLTracking"] = "https://" + extras.GetAPIDomain() + "/tracking/" + msg.MessageID + "/" + code + "/pixel.gif"
			data["URLFeedback"] = "https://" + extras.GetWebDomain() + "/feedback/" + msg.MessageID + "?authorization=" + code

			t := ParseTemplate("message", data)
			if len(t) != 0 {
				extras.SendEmail(gms.Members[m].Email, "Nuevo mensaje de "+msg.Name, t)
			}
		}
		err = ep.db.SetMessageSended(msg.MessageID)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	} else if disapprove >= ep.minModApproves {
		data := make(map[string]string)
		data["Name"] = msg.Name
		data["Message"] = msg.Message
		t := ParseTemplate("rejected", data)
		if len(t) != 0 {
			extras.SendEmail(msg.Email, "Tu mensaje ha sido rechazado.", t)
		}
	}

	if votes >= ep.minModApproves {
		err = ep.db.SetMessageClosed(msg.MessageID)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

func EmailFeedbackUser(FeedbackID string) {

	fbk, err := ep.db.GetFeedback(FeedbackID)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg, err := ep.db.GetMessage(fbk.MessageID)
	if err != nil {
		fmt.Println(err)
		return
	}

	name := ""
	gms, err := ep.db.GetGroup(fbk.ToGroup)
	if err != nil {
		fmt.Println(err)
		return
	}
	for m := range gms.Members {
		if gms.Members[m].Email == fbk.Email {
			name = gms.Members[m].Name
			break
		}
	}

	// fmt.Sprintf("%s <%s>", msg.Name, msg.Email)
	data := make(map[string]string)
	data["Name"] = msg.Name
	data["NameFeedback"] = name
	data["Message"] = msg.Message
	data["Feedback"] = fbk.Message

	t := ParseTemplate("feedback", data)
	if len(t) != 0 {
		extras.SendEmail(msg.Email, name+" ha respondido tu mensaje!", t)
	}

}
