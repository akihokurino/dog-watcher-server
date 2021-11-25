package agora

import (
	"live-server/graph/model"
	"live-server/infra/firebase"
	"log"
	"time"

	"github.com/pkg/errors"

	rtctokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/RtcTokenBuilder"
)

type Client interface {
	GetRTCToken(uid firebase.UID, channelName string, role model.AgoraRole) (string, error)
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

func (c *client) GetRTCToken(uid firebase.UID, channelName string, role model.AgoraRole) (string, error) {
	expireTimeInSeconds := uint32(60 * 30)
	currentTimestamp := uint32(time.Now().UTC().Unix())
	expireTimestamp := currentTimestamp + expireTimeInSeconds

	var ar rtctokenbuilder.Role
	if role == model.AgoraRolePublisher {
		ar = rtctokenbuilder.RolePublisher
	}
	if role == model.AgoraRoleSubscriber {
		ar = rtctokenbuilder.RoleSubscriber
	}

	result, err := rtctokenbuilder.BuildTokenWithUID(c.appID, c.certID, channelName, 0, ar, expireTimestamp)
	if err != nil {
		log.Printf("agora error: %+v", err)
		return "", errors.WithStack(err)
	}

	return result, nil
}
