package df

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func FilterByActionAndLifespan(folders []string, action string, lifespan int) (res []Intent) {
	var intents []string

	for _, folder := range folders {
		files, err := ioutil.ReadDir(folder)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			if !strings.Contains(file.Name(), "usersays") {
				intents = append(intents, fmt.Sprintf("%v/%v", folder, file.Name()))
			}
		}

	}

	var reported []string
	for _, intent := range intents {
		var intentData Intent
		content, err := ioutil.ReadFile(intent)

		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(content, &intentData)
		if err != nil {
			fmt.Println("Error " + err.Error())
		} else {
			if intentData.Responses[0].Action == action {
				for _, ctx := range intentData.Responses[0].AffectedContexts {
					if ctx.Lifespan == lifespan {
						reported = append(reported, fmt.Sprintf("intent name: %v and lifespan: %v\n", intentData.Name, ctx.Lifespan))
						res = append(res, intentData)
						break
					}
				}
			}
		}
	}

	for _, r := range reported {
		fmt.Println(r)
	}
	return
}
