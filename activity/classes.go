package activity

import "time"

// ASObject : activitystreams base struct
type ASObject struct {
	ID           string      `json:"id,omitempty"`
	Object       interface{} `json:"object,omitempty"`
	Type         string      `json:"type,omitempty"`
	Attachment   []*ASObject `json:"attachment,omitempty"`
	AttributedTo []*ASObject `json:"attributedTo,omitempty"`
	Audience     []*ASObject `json:"audience,omitempty"`
	Content      interface{} `json:"content,omitempty"`
	Context      []string    `json:"@context,omitempty"`
	Name         string      `json:"name,omitempty"`
	EndTime      *time.Time  `json:"endTime,omitempty"`
	Generator    []*ASObject `json:"generator,omitempty"`
	Icon         *ASObject   `json:"icon,omitempty"`
	Image        *ASObject   `json:"image,omitempty"`
	InReplyTo    *ASObject   `json:"inReplyTo,omitempty"`
	Location     *ASObject   `json:"location,omitempty"`
	Preview      *ASObject   `json:"preview,omitempty"`
	Published    *time.Time  `json:"published,omitempty"`
	Replies      *ASObject   `json:"replies,omitempty"`
	StartTime    *time.Time  `json:"startTime,omitempty"`
	Summary      string      `json:"summary,omitempty"`
	Tag          *ASObject   `json:"tag,omitempty"`
	Updated      *time.Time  `json:"updated,omitempty"`
	URL          string      `json:"url,omitempty"`
	To           []string    `json:"to,omitempty"`
	Bto          []string    `json:"bto,omitempty"`
	Cc           []string    `json:"cc,omitempty"`
	Bcc          []string    `json:"bcc,omitempty"`
	MediaType    *ASObject   `json:"mediaType,omitempty"`
	Duration     *time.Time  `json:"duration,omitempty"`
}

type APObject struct {
	Inbox                      string    `json:"inbox,omitempty"`
	Outbox                     string    `json:"outbox,omitempty"`
	Following                  string    `json:"following,omitempty"`
	Followers                  string    `json:"followers,omitempty"`
	Liked                      string    `json:"liked,omitempty"`
	Streams                    string    `json:"streams,omitempty"`
	PreferredUsername          string    `json:"preferredUsername,omitempty"`
	Endpoints                  Endpoints `json:"endpoints,omitempty"`
	ProxyURL                   string    `json:"proxyUrl,omitempty"`
	OauthAuthorizationEndpoint string    `json:"oauthAuthorizationEndpoint,omitempty"`
	OauthTokenEndpoint         string    `json:"oauthTokenEndpoint,omitempty"`
	ProvideClientKey           string    `json:"provideClientKey,omitempty"`
	SignClientKey              string    `json:"signClientKey,omitempty"`
	SharedInbox                string    `json:"sharedInbox,omitempty"`
}

//Endpoints : Contains SharedInbox address.
type Endpoints struct {
	SharedInbox string `json:"sharedInbox"`
}

type W3IdSecurityObject struct {
	Context   interface{}           `json:"@context,omitempty"`
	PublicKey W3IdSecurityPublicKey `json:"publicKey"`
}

type W3IdSecurityPublicKey struct {
	ID           string `json:"id"`
	Owner        string `json:"owner"`
	PublicKeyPem string `json:"publicKeyPem"`
}

// ASRootObject : implements ASObject, W3IdSecurityObject
type ASRootObject struct {
	ASObject
	W3IdSecurityObject
}

// ASActivity : implements ASObject
type ASActivity struct {
	ASObject
	Actor      APActor     `json:"actor,omitempty"`
	Target     interface{} `json:"target,omitempty"`
	Result     interface{} `json:"result,omitempty"`
	Origin     interface{} `json:"origin,omitempty"`
	Instrument interface{} `json:"instrument,omitempty"`
}

// APActor : implements ASObject, APObject
type APActor struct {
	ASRootObject
	APObject
}
