package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

const name = "nostr-undelete"

const version = "0.0.2"

var revision = "HEAD"

func main() {
	priv := os.Args[1]
	var pub string
	if _, s, err := nip19.Decode(priv); err != nil {
		log.Fatal(err)
	} else {
		if pub, err = nostr.GetPublicKey(s.(string)); err != nil {
			log.Fatal(err)
		}
	}
	filter := nostr.Filter{
		Kinds:   []int{nostr.KindSetMetadata},
		Authors: []string{pub},
		Limit:   1,
	}

	relay, err := nostr.RelayConnect(context.Background(), "wss://nostr-relay.nokotaro.com")
	if err != nil {
		log.Fatal(err)
	}

	evs, err := relay.QuerySync(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var m map[string]any
	err = json.Unmarshal([]byte(evs[0].Content), &m)
	if err != nil {
		log.Fatal(err)
	}
	delete(m, "deleted")

	var sk string
	if _, s, err := nip19.Decode(priv); err != nil {
		log.Fatal(err)
	} else {
		sk = s.(string)
	}

	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	var ev nostr.Event
	ev.PubKey = pub
	ev.CreatedAt = time.Now()
	ev.Content = string(b)
	ev.Kind = nostr.KindSetMetadata
	if err := ev.Sign(sk); err != nil {
		log.Fatal(err)
	}
	relay.Publish(context.Background(), ev)
}
