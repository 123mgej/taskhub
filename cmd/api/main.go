package main 
import (
	"log"
	"taskhub/internal/router"
)

func main(){
	r :=router.New()
	addr :=":8080"
	log.Printf("taskhub api listening on %s\n",addr)
	if err :=r.Run(addr);err !=nil {
		log.Fatalf("server failed %v",err)
	}
}