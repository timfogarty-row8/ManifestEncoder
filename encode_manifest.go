package manifest-encoder

import (
	"github.com/reactivex/rxgo"
	"github.com/reactivex/rxgo/observer"
	"github.com/reactivex/rxgo/observable"
)

func encodeManifest(c *gin.Context) {
	format := c.Params.ByName("format")
	mediaGuid := c.Params.ByName("mediaGuid")
	token := c.Params.ByName("token")
	
	var user_id = 0
	var manifest
	
	tokenIsGood := observable.Just( validateToken(token) )
	sub_token := tokenIsGood.Subscribe( observer.Observer {
			NextHandler: func( item interface{} ) {
				user_id = item.user_id
			},
			
			ErrHandler: func(err error) {
				
			},
			
			DoneHandler: func() {
				if( manifest != nil ) doEncoding( user_id, format, manifest )
			}
		})
	<-sub_token
	
	hasManifest := observable.Just( fetchMovie(mediaGuid, format) )
	sub_manifest := hasManifest.Subscribe( observer.Observer {
			NextHandler: func( item interface{} ) {
				manifest = item
			},
			
			ErrHandler: func(err error) {
				
			},
			
			DoneHandler: func() {
				if( user_id != nil ) doEncoding( user_id, format, manifest )
			}

		})
	<-sub_manifest
}


func doEncoding( user_id, format, manifest ) {
	
}