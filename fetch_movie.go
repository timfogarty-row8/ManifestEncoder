package manifest-encoder

import (
	"errors",
	"log"
	"strconv"
	
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/reactivex/rxgo"
	"github.com/reactivex/rxgo/observer"
	"github.com/reactivex/rxgo/observable"

)

type Manifest struct {
	media_guid: string
	file_a: string
	file_b: string
}

func fetchMovie(mediaGuid, format) {
	key := "media_guid:"+ mediaGuid+format
	
	conn, err := Redis.Get()
    if err != nil {
        return nil, err
    }
	defer Redis.Put(conn)

	exists, err := conn.Cmd("EXISTS", key).Int()
	if err != nil {
	    return err
    } else if exists == 0 {
	    	watcher := observer.New(onNext, onDone)
	    	sub := observable.Just( requestMovieAndStore(mediaGuid,format) ).Subscribe( watcher )
	    	<-sub
    } else {
		manifest, err := conn.Cmd("HGETALL", key).Map()		// need to map to an object
		if err != nil {
		}
		return manifest
    }
}

func requestMovieAndStore(mediaGuid,format) {
	manifest := requestMovie(mediaGuid,format);
    err := conn.Cmd("HMSET", manifest ).Err
    if err != nil {
	    	// TODO
    }
	return manifest    
}

function requestMovie(kemediaGuid,formaty) {
	url := os.Getenv("CTS_MANIFEST_URL")+mediaGuid+format;			// TODO: not correct
	ctsClient := http.Client{
        Timeout: time.Second * 5, // Maximum of 5 secs
	}
	
	req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
	    	// TODO
    }
    
    req.Header.Set("User-Agent", "row8-manifest-encoder")
    req.SetBasicAuth( os.Getenv("CTS_USER"), os.Getenv("CTS_PASS") )

    res, getErr := ctsClient.Do(req)
    if getErr != nil {
	    	// TODO:
    }
    
    body, readErr := ioutil.ReadAll(res.Body)
    if readErr != nil {
	    	// TODO:
    }
    
    manifest := Manifest{}
    jsonErr := json.Unmarshal(body, &manifest)
    if jsonErr != nil {
	    	// TODO:
    }
    
    return manifest
}