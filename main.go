package main

import(
	"net/http"
	"log"
)

fucn main(){
	http.HandleFunc("/",handler)
	log.Fetal(gttp.ListenAndServe(":8080",nill))
}

func handler( w http.ResponseWriter,

)