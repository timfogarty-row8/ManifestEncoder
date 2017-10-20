package com.row8.manifest

import org.scalatra.test.scalatest._
import org.scalatest.FunSuiteLike

class ManifestEncoderServletTests extends ScalatraSuite with FunSuiteLike {

  addServlet(classOf[ManifestEncoderServlet], "/*")

  test("GET / on ManifestEncoderServlet should return status 200"){
    get("/"){
      status should equal (200)
    }
  }

}
