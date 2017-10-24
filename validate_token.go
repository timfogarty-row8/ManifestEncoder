package manifest-encoder

import (
	"os"
	"net/http"
)

type User struct {
	user_id int
	other_stuff
}

func validateToken( token ) {
	url := os.Getenv("CTS_TOKEN_VALIDATE")+token;
	ctsClient := http.Client{
        Timeout: time.Second * 5, // Maximum of 5 secs
	}
	
	req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    req.Header.Set("User-Agent", "row8-manifest-encoder")
    req.SetBasicAuth( os.Getenv("CTS_USER"), os.Getenv("CTS_PASS") )

    res, getErr := ctsClient.Do(req)
    if getErr != nil {
        log.Fatal(getErr)
    }
    
    body, readErr := ioutil.ReadAll(res.Body)
    if readErr != nil {
        log.Fatal(readErr)
    }
    
    user := User{}
    jsonErr := json.Unmarshal(body, &user)
    if jsonErr != nil {
        log.Fatal(jsonErr)
    }
    
    return user
}
