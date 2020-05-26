package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rodzy/flash/middlewares"
	"github.com/rodzy/flash/routers"
	"github.com/rs/cors"
)

//DirectDrivers func allows to connect API w/ routing using mux
func DirectDrivers() {
	//Instancing a new router w mux
	router := mux.NewRouter()
	//Register
	router.HandleFunc("/register", middlewares.CheckDataBase(routers.Register)).Methods("POST")
	//Login
	router.HandleFunc("/login", middlewares.CheckDataBase(routers.Login)).Methods("POST")
	//Profile
	router.HandleFunc("/profile", middlewares.CheckDataBase(middlewares.CheckJWT(routers.ViewProfile))).Methods("GET")
	//Updating profile
	router.HandleFunc("/updateProfile", middlewares.CheckDataBase(middlewares.CheckJWT(routers.ModifyUserInfo))).Methods("PUT")
	//Publish pubs
	router.HandleFunc("/pubpub", middlewares.CheckDataBase(middlewares.CheckJWT(routers.PublishPub))).Methods("POST")
	//View pubs
	router.HandleFunc("/pubRead", middlewares.CheckDataBase(middlewares.CheckJWT(routers.PubReader))).Methods("GET")
	//Deleting pubs
	router.HandleFunc("/pubDelete", middlewares.CheckDataBase(middlewares.CheckJWT(routers.DeletePub))).Methods("DELETE")
	//Avatar upload
	router.HandleFunc("/uploadAvatar", middlewares.CheckDataBase(middlewares.CheckJWT(routers.UploadProfilePic))).Methods("POST")
	//Avatar request
	router.HandleFunc("/avatar", middlewares.CheckDataBase(routers.ShowAvatar)).Methods("GET")
	//Banner upload
	router.HandleFunc("/uploadBanner", middlewares.CheckDataBase(middlewares.CheckJWT(routers.UploadBannerPic))).Methods("POST")
	//Banner request
	router.HandleFunc("/banner", middlewares.CheckDataBase(routers.ShowBanner)).Methods("GET")
	//Follower
	router.HandleFunc("/follow", middlewares.CheckDataBase(middlewares.CheckJWT(routers.FollowUser))).Methods("POST")
	//Unfollow
	router.HandleFunc("/unfollow", middlewares.CheckDataBase(middlewares.CheckJWT(routers.UnfollowUser))).Methods("DELETE")
	//My followers
	router.HandleFunc("/myfollowers", middlewares.CheckDataBase(middlewares.CheckJWT(routers.AskFollower))).Methods("GET")
	
	//Setting port env var
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	//Setting up the cors for the cloud API
	handler := cors.AllowAll().Handler(router)
	//Just in case for http error listening
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
