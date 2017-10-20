package com.row8.manifest

object Settings {
//    def jwt_secret = "Frogs"
    def jwt_secret = sys.env("JWT_SECRET")
}