package environment

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type EnvVariable struct {
  Key string
  Value string
}

func envVariable(chk string) EnvVariable {
	data, err := os.Open(".env")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

  result := EnvVariable{
    Key: "",
    Value: "",
  }
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
      key := strings.TrimSpace(parts[0])
			if key == chk {
				result.Key = key
				result.Value = strings.TrimSpace(parts[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func YoutubeApi() EnvVariable {
  api := envVariable("YOUTUBE_API")
	return api 
}
