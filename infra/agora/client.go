package agora

import (
	"live-server/infra/firebase"
	"log"
	"time"

	"github.com/pkg/errors"

	rtctokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/RtcTokenBuilder"
)

const channelName = "channel-alpha"

type Client interface {
	GetRTCToken(uid firebase.UID) (string, error)
}

func NewClient(appID string, certID string) Client {
	return &client{
		appID:  appID,
		certID: certID,
	}
}

type client struct {
	appID  string
	certID string
}

func (c *client) GetRTCToken(uid firebase.UID) (string, error) {
	expireTimeInSeconds := uint32(60 * 30)
	currentTimestamp := uint32(time.Now().UTC().Unix())
	expireTimestamp := currentTimestamp + expireTimeInSeconds

	log.Printf("call agora token, app_id: %s, uid: %s", c.appID, uid.String())
	result, err := rtctokenbuilder.BuildTokenWithUID(c.appID, c.certID, channelName, 0, 1, expireTimestamp)
	if err != nil {
		log.Printf("agora error: %+v", err)
		return "", errors.WithStack(err)
	}

	return result, nil
}
