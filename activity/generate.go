package activity

import (
	"time"
)

func (activity *ASActivity) GenerateAnnounceActivity(id string, actor APActor, object ASObject) {
	published := time.Now()

	if !contains(activity.Context, "https://www.w3.org/ns/activitystreams") {
		activity.Context = append(activity.Context, "https://www.w3.org/ns/activitystreams")
	}
	activity.ID = id
	activity.Type = "Announce"
	activity.Actor = actor
	activity.Object = object
	activity.Published = &published
}
