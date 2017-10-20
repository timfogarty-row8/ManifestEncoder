package com.row8.manifest

import org.json4s.{DefaultFormats, Formats}
import org.scalatra._
import org.scalatra.json._

import com.row8.manifest.models._
//import com.row8.auth._

class ManifestEncoderServlet extends ScalatraServlet with CorsSupport with JacksonJsonSupport  {
	protected implicit lazy val jsonFormats: Formats = DefaultFormats
	
	options("/*"){
    	response.setHeader("Access-Control-Allow-Headers", request.getHeader("Access-Control-Request-Headers"));
  }
  
  before() {
    	contentType = formats("json")
  }
  
  get("/") {
    	"Hi there " + Settings.jwt_secret
  }
  
  get("/version") {
    	"Row8 Manifest Generator version 0.1.0"
  }
  
  get("/hello/:name") {
    	<p>Hello, {params("name")}</p>
  }
  
  get("/flowers") {
//    val user: User = auth.get
  	  Flowers.all
  }

}

