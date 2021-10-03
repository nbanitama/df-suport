package auth

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"golang.org/x/oauth2/google"
)

func GetAccessTokenFromServiceAccount(filePath string) (res string) {
	jsonByte, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	conf, err := google.JWTConfigFromJSON(jsonByte, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		log.Fatal(fmt.Sprintf("JWTConfigFromJSON() get error %v", err))
	}

	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Duration(3*time.Second))
	defer ctxCancel()
	tokSrc := conf.TokenSource(ctx)
	gToken, err := tokSrc.Token()
	if err != nil {
		log.Fatal(fmt.Sprintf("tokSrc.Token() get err %v", err))
	}

	if gToken == nil {
		log.Fatal("gToken not generated")
	}

	// fmt.Printf("Success generating token %v\n", gToken.AccessToken)
	res = gToken.AccessToken
	return
}
