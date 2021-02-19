package endpoint

import (
	"fmt"

	"github.com/jonathanhecl/public-feedback-api/extras"
)

var subjectEmail = map[string]string{
	"confirm":    "Confirma el envío de tu carta al {{.Group}}",
	"moderation": "Moderar carta de {{.Name}}",
	"success":    "Tu carta al {{.Group}} se ha enviado con éxito.",
	"rejected":   "Tu carta al {{.Group}} ha sido rechazada.",
	"message":    "Carta de {{.Name}} al {{.Group}}",
	"feedback":   "Tu carta ha sido respondida por {{.NameMember}} del {{.Group}}",
}

func EmailUserConfirmation(MessageID string) {

	msg, err := ep.db.GetMessage(MessageID)
	if err != nil {
		fmt.Println(err)
		return
	}

	gms, err := ep.db.GetGroup(msg.ToGroup)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]string)
	data["WebDomain"] = extras.GetWebDomain()
	data["Brand"] = ep.brandTitle
	data["Name"] = msg.Name
	data["Message"] = msg.Message
	data["URL"] = "https://" + extras.GetWebDomain() + "/confirm/" + msg.MessageID + "/" + msg.ConfirmationCode
	data["Group"] = gms.Label

	t := ParseTemplate("confirm", data)
	if len(t) != 0 {
		subject := ParseTemplateText("confirm", subjectEmail["confirm"], data)
		extras.SendEmail(msg.Email, subject, t)
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
		data["WebDomain"] = extras.GetWebDomain()
		data["Brand"] = ep.brandTitle
		data["Name"] = msg.Name
		data["Email"] = msg.Email
		data["Message"] = msg.Message
		data["Moderator"] = mds.Members[m].Name
		data["URLApprove"] = "https://" + extras.GetWebDomain() + "/moderation/" + msg.MessageID + "/approved/" + code
		data["URLDisapprove"] = "https://" + extras.GetWebDomain() + "/moderation/" + msg.MessageID + "/disapproved/" + code

		t := ParseTemplate("moderation", data)
		if len(t) != 0 {
			subject := ParseTemplateText("moderation", subjectEmail["moderation"], data)
			extras.SendEmail(mds.Members[m].Email, subject, t)
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

	gms, err := ep.db.GetGroup(msg.ToGroup)
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

	data := make(map[string]string)
	data["WebDomain"] = extras.GetWebDomain()
	data["Brand"] = ep.brandTitle
	data["Name"] = msg.Name
	data["Message"] = msg.Message
	data["Group"] = gms.Label

	sended := false
	if approve >= ep.minModApproves {
		sended = true
		t := ParseTemplate("success", data)
		if len(t) != 0 {
			subject := ParseTemplateText("success", subjectEmail["success"], data)
			extras.SendEmail(msg.Email, subject, t)
		}
		for m := range gms.Members {
			code := extras.GenerateMemberLink(msg.MessageID, msg.CreatedAt, gms.Members[m].Email)
			// fmt.Sprintf("%s <%s>", gms.Members[m].Name, gms.Members[m].Email)

			data["Name"] = msg.Name                  // user name
			data["NameMember"] = gms.Members[m].Name // member name
			data["URLTracking"] = "https://" + extras.GetAPIDomain() + "/tracking/" + msg.MessageID + "/" + code + "/pixel.gif"
			data["URLFeedback"] = "https://" + extras.GetWebDomain() + "/feedback/" + msg.MessageID + "/authorization/" + code

			t := ParseTemplate("message", data)
			if len(t) != 0 {
				subject := ParseTemplateText("message", subjectEmail["message"], data)
				extras.SendEmail(gms.Members[m].Email, subject, t)
			}
		}
		err = ep.db.SetMessageSended(msg.MessageID)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	} else if disapprove >= ep.minModApproves {
		sended = true
		t := ParseTemplate("rejected", data)
		if len(t) != 0 {
			subject := ParseTemplateText("rejected", subjectEmail["rejected"], data)
			extras.SendEmail(msg.Email, subject, t)
		}
	}

	if votes >= ep.minModApproves && sended {
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
	data["WebDomain"] = extras.GetWebDomain()
	data["Brand"] = ep.brandTitle
	data["Name"] = msg.Name
	data["NameMember"] = name
	data["Group"] = gms.Label
	data["Message"] = msg.Message
	data["Feedback"] = fbk.Message

	t := ParseTemplate("feedback", data)
	if len(t) != 0 {
		subject := ParseTemplateText("feedback", subjectEmail["feedback"], data)
		extras.SendEmail(msg.Email, subject, t)
	}

}
