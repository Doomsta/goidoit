package main

import (
	"context"
	"github.com/Doomsta/goidoit"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	client := goidoit.NewClient("https://demo.i-doit.com/src/jsonrpc.php", "",
		goidoit.WithUserCredencials("admin", "admin"),
		goidoit.WithHTTPClient(http.DefaultClient),
		goidoit.WithInsecure(),
	)
	//log.Println(client.Search(ctx, "Bremen"))
	addons, _ := client.Idoit.Addons(ctx)
	for i, addon := range addons {
		log.Println(i, addon.Title, addon.Version, addon.Author.Name, addon.Key)
	}

	l, _ := client.Idoit.License(ctx)
	for s, s2 := range l.Addons {
		log.Println(s, s2.Label, s2.Licensed)
	}

	ll, _ := client.CMDB.StatusRead(ctx)
	for i, s := range ll {
		log.Println(i, s.Title, s.Constant, s.Editable)
	}

	//log.Println("=== constants ===")
	//data, err := client.Idoit.Constants(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//for _, s := range data.Categories {
	//	for s2, s3 := range s {
	//		log.Println(s2, s3)
	//	}
	//}

	//log.Println(client.CMDB.GetObjects(ctx, goidoit.GetObjectsFilter{
	//	IDs:   []int{1, 5242, 2942},
	//	Title: "Root Location",
	//}))
}
