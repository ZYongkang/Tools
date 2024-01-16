package main

import (
	"fmt"
	"os"

	rtctokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtctokenbuilder2"
)

func main() {
	// Need to set environment variable AGORA_APP_ID
	appID := os.Getenv("AGORA_APP_ID")
	// Need to set environment variable AGORA_APP_CERTIFICATE
	appCertificate := os.Getenv("AGORA_APP_CERTIFICATE")

	channelName := "7d72365eb983485397e3e3f9d460bdda"
	uid := uint32(2882341273)
	uidStr := "2882341273"
	tokenExpirationInSeconds := uint32(3600)
	privilegeExpirationInSeconds := uint32(3600)
	joinChannelPrivilegeExpireInSeconds := uint32(3600)
	pubAudioPrivilegeExpireInSeconds := uint32(3600)
	pubVideoPrivilegeExpireInSeconds := uint32(3600)
	pubDataStreamPrivilegeExpireInSeconds := uint32(3600)

	result, err := rtctokenbuilder.BuildTokenWithUid(appID, appCertificate, channelName, uid, rtctokenbuilder.RoleSubscriber, tokenExpirationInSeconds, privilegeExpirationInSeconds)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Token with int uid: %s\n", result)
	}

	result, err = rtctokenbuilder.BuildTokenWithUserAccount(appID, appCertificate, channelName, uidStr, rtctokenbuilder.RoleSubscriber, tokenExpirationInSeconds, privilegeExpirationInSeconds)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Token with user account: %s\n", result)
	}

	result, err1 := rtctokenbuilder.BuildTokenWithUidAndPrivilege(appID, appCertificate, channelName, uid,
		tokenExpirationInSeconds, joinChannelPrivilegeExpireInSeconds, pubAudioPrivilegeExpireInSeconds, pubVideoPrivilegeExpireInSeconds, pubDataStreamPrivilegeExpireInSeconds)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Printf("Token with int uid and privilege: %s\n", result)
	}

	result, err1 = rtctokenbuilder.BuildTokenWithUserAccountAndPrivilege(appID, appCertificate, channelName, uidStr,
		tokenExpirationInSeconds, joinChannelPrivilegeExpireInSeconds, pubAudioPrivilegeExpireInSeconds, pubVideoPrivilegeExpireInSeconds, pubDataStreamPrivilegeExpireInSeconds)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Printf("Token with user account and privilege: %s\n", result)
	}
}
