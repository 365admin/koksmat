package officegraph

import (
	"context"
	"fmt"
	"io"
	"log"
	"testing"
	"time"
)

func TestAuth(t *testing.T) {
	_, err, _ := GetClient()
	if err != nil {
		t.Fatalf("Should not return error")
	}
}
func TestList(t *testing.T) {

	items, err := SubscriptionList()
	if err != nil {
		t.Fatalf("Should not return error")
	}
	log.Println(len(items), "Items")
	for _, item := range items {
		log.Println(*item.Resource)
	}
}

func TestRemove(t *testing.T) {
	items, err := SubscriptionList()
	if err != nil {
		t.Fatalf("Should not return error")
	}
	for _, item := range items {
		log.Println("Removing", *item.Resource)
		_, err = RemoveSubscription(*item.Id)
		if err != nil {
			t.Fatalf("Should not return error")
		}
	}

}

func s(text string) *string {
	return &text
}
func TestAdd(t *testing.T) {
	c, err, _ := GetClient() //NewClient("https://graph.microsoft.com/v1.0/")
	if err != nil {
		t.Fatalf("Should not return error")
	}
	ctx := context.Background()

	time := time.Now().Add(time.Hour * 24 * 1)

	sub := &MicrosoftGraphSubscription{
		//Id:                 s("1"),
		ChangeType: s("created,updated,deleted"),
		//Resource:           s("https://christianiabpos.sharepoint.com/sites/Cava3/_api/Web/GetList('/sites/Cava3/Lists/Test Changes"),
		Resource:           s("/users/niels.johansen@nexigroup.com/events"),
		ExpirationDateTime: &time,
		NotificationUrl:    s("https://niels-mac.nets-intranets.com/api/v1/subscription/notify"),
		//NotificationUrl: s("https://magicbox.nexi-intra.com/api/v1/subscription/notify"),
		ClientState: s("zz"),
	}

	response, err := c.SubscriptionsSubscriptionCreateSubscription(ctx, *sub)

	if err != nil {
		t.Fatalf("Should not return error")
	}
	body, err := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
