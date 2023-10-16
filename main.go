package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"time"

	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

const name = "nostr-undelete"

const version = "0.0.6"

var revision = "HEAD"

var relays = []string{
	"wss://paid.nostrified.org",
	"wss://1.noztr.com",
	"wss://btc-italia.online",
	"wss://at.nostrworks.com",
	"wss://bitcoiner.social",
	"wss://bitcoinmaximalists.online",
	"wss://brb.io",
	"wss://btc.klendazu.com",
	"wss://e.nos.lol",
	"wss://eden.nostr.land",
	"wss://eden.nostr.land",
	"wss://freespeech.casa",
	"wss://global.relay.red",
	"wss://knostr.neutrine.com",
	"wss://knostr.neutrine.com:8880",
	"wss://lbrygen.xyz",
	"wss://lightningrelay.com",
	"wss://no-str.org",
	"wss://no.str.cr",
	"wss://nos.lol",
	"wss://noster.online",
	"wss://nostr-01.bolt.observer",
	"wss://nostr-2.afarazit.eu",
	"wss://nostr-au.coinfundit.com",
	"wss://nostr-eu.coinfundit.com",
	"wss://nostr-pub.wellorder.net",
	"wss://nostr-pub1.southflorida.ninja",
	"wss://nostr-relay.alekberg.net",
	"wss://nostr-relay.bitcoin.ninja",
	"wss://nostr-relay.derekross.me",
	"wss://nostr-relay.freedomnode.com",
	"wss://nostr-relay.lnmarkets.com",
	"wss://nostr-relay.pcdkd.fyi",
	"wss://nostr-relay.schnitzel.world",
	"wss://nostr-relay.texashedge.xyz",
	"wss://nostr-us.coinfundit.com",
	"wss://nostr-verif.slothy.win",
	"wss://nostr.1729.cloud",
	"wss://nostr.1f52b.xyz",
	"wss://nostr.21m.fr",
	"wss://nostr.21sats.net",
	"wss://nostr.600.wtf",
	"wss://nostr.8e23.net",
	"wss://nostr.actn.io",
	"wss://nostr.app.runonflux.io",
	"wss://nostr.arguflow.gg",
	"wss://nostr.arguflow.gg",
	"wss://nostr.bch.ninja",
	"wss://nostr.beta3.dev",
	"wss://nostr.bitcoin-21.org",
	"wss://nostr.bitcoin.sex",
	"wss://nostr.bitcoinplebs.de",
	"wss://nostr.blockpower.capital",
	"wss://nostr.bongbong.com",
	"wss://nostr.bostonbtc.com",
	"wss://nostr.chainofimmortals.net",
	"wss://nostr.cizmar.net",
	"wss://nostr.coinsamba.com.br",
	"wss://nostr.com.de",
	"wss://nostr.coollamer.com",
	"wss://nostr.corebreach.com",
	"wss://nostr.drss.io",
	"wss://nostr.easydns.ca",
	"wss://nostr.einundzwanzig.space",
	"wss://nostr.ethtozero.fr",
	"wss://nostr.fluidtrack.in",
	"wss://nostr.fmt.wiz.biz",
	"wss://nostr.fractalized.ovh",
	"wss://nostr.globals.fans",
	"wss://nostr.gromeul.eu",
	"wss://nostr.hackerman.pro",
	"wss://nostr.handyjunky.com",
	"wss://nostr.herci.one",
	"wss://nostr.hugo.md",
	"wss://nostr.inosta.cc",
	"wss://nostr.island.network",
	"wss://nostr.itas.li",
	"wss://nostr.klabo.blog",
	"wss://nostr.localhost.re",
	"wss://nostr.lu.ke",
	"wss://nostr.massmux.com",
	"wss://nostr.middling.mydns.jp",
	"wss://nostr.mikedilger.com",
	"wss://nostr.milou.lol",
	"wss://nostr.milou.lol",
	"wss://nostr.mom",
	"wss://nostr.mouton.dev",
	"wss://nostr.mustardnodes.com",
	"wss://nostr.nodeofsven.com",
	"wss://nostr.noones.com",
	"wss://nostr.ownscale.org",
	"wss://nostr.oxtr.dev",
	"wss://nostr.pleb.network",
	"wss://nostr.pleb.network",
	"wss://nostr.plebchain.org",
	"wss://nostr.portemonero.com",
	"wss://nostr.randomdevelopment.biz",
	"wss://nostr.rdfriedl.com",
	"wss://nostr.rocket-tech.net",
	"wss://nostr.rocks",
	"wss://nostr.roundrockbitcoiners.com",
	"wss://nostr.satsophone.tk",
	"wss://nostr.screaminglife.io",
	"wss://nostr.sebastix.dev",
	"wss://nostr.shawnyeager.net",
	"wss://nostr.sidnlabs.nl",
	"wss://nostr.slothy.win",
	"wss://nostr.supremestack.xyz",
	"wss://nostr.up.railway.app",
	"wss://nostr.uselessshit.co",
	"wss://nostr.uthark.com",
	"wss://nostr.vulpem.com",
	"wss://nostr.w3ird.tech",
	"wss://nostr.wine",
	"wss://nostr.wine",
	"wss://nostr.yuv.al",
	"wss://nostr.yuv.al",
	"wss://nostr.zaprite.io",
	"wss://nostr.zebedee.cloud",
	"wss://nostr.zenon.info",
	"wss://nostr.zkid.social",
	"wss://nostr01.opencult.com",
	"wss://nostr01.vida.dev",
	"wss://nostr2.actn.io",
	"wss://nostr3.actn.io",
	"wss://nostrafrica.pcdkd.fyi",
	"wss://nostre.cc",
	"wss://nostream.nostrly.io",
	"wss://nostrex.fly.dev",
	"wss://nostrical.com",
	"wss://nostrich.friendship.tw",
	"wss://nostring.deno.dev",
	"wss://nostro.cc",
	"wss://nostrsatva.net",
	"wss://offchain.pub",
	"wss://paid.no.str.cr",
	"wss://paid.spore.ws",
	"wss://pow.nostrati.com",
	"wss://private.red.gb.net",
	"wss://private.red.gb.net",
	"wss://private.red.gb.net",
	"wss://puravida.nostr.land",
	"wss://relay.cryptocculture.com",
	"wss://relay.damus.io",
	"wss://relay.honk.pw",
	"wss://relay.lexingtonbitcoin.org",
	"wss://relay.minds.com/nostr/v1/ws",
	"wss://relay.n057r.club",
	"wss://relay.nostr.africa",
	"wss://relay.nostr.band",
	"wss://relay.nostr.bg",
	"wss://relay.nostr.ch",
	"wss://relay.nostr.express",
	"wss://relay.nostr.info",
	"wss://relay.nostr.moe",
	"wss://relay.nostr.net.in",
	"wss://relay.nostr.nu",
	"wss://relay.nostr.ro",
	"wss://relay.nostr.wirednet.jp",
	"wss://relay.nostrati.com",
	"wss://relay.nostrich.de",
	"wss://relay.nostriches.org",
	"wss://relay.nostrid.com",
	"wss://relay.nostrology.org",
	"wss://relay.nostrview.com",
	"wss://relay.orangepill.dev",
	"wss://relay.plebstr.com",
	"wss://relay.ryzizub.com",
	"wss://relay.sendstr.com",
	"wss://relay.snort.social",
	"wss://relay.stoner.com",
	"wss://relay.taxi",
	"wss://relay.zeh.app",
	"wss://relay1.gems.xyz",
	"wss://rsslay.nostr.moe",
	"wss://spore.ws",
	"wss://tmp-relay.cesc.trade",
	"wss://relay.punkhub.me",
	"wss://relay.nostrdocs.com",
	"wss://nostr.rikmeijer.nl",
	"wss://nostr.xmr.rocks",
	"wss://nerostr.xmr.rocks",
	"wss://nostr.azte.co",
	"wss://nostr.rocketstyle.com.au",
	"wss://xmr.usenostr.org",
	"wss://africa.nostr.joburg",
	"wss://relay.pineapple.pizza",
	"wss://nostrelay.yeghro.site",
	"wss://nostr.zerofiat.world",
	"wss://rs1.abaiba.top",
	"wss://rs2.abaiba.top",
	"wss://yabu.me",
}

func sendUndelete(r string, ev nostr.Event) {
	relay, err := nostr.RelayConnect(context.Background(), r)
	if err != nil {
		log.Println(err, relay.URL)
		return
	}
	defer relay.Close()

	_, err = relay.Publish(context.Background(), ev)
	if err != nil {
		log.Println(err, relay.URL)
	}
}

func main() {
	var checkRelay string
	flag.StringVar(&checkRelay, "relay", "wss://nostr-relay.nokotaro.com", "relay for checking metadata")
	flag.Parse()

	priv := flag.Arg(0)
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

	relay, err := nostr.RelayConnect(context.Background(), checkRelay)
	if err != nil {
		log.Fatal(err)
	}

	evs, err := relay.QuerySync(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	if len(evs) == 0 {
		log.Fatalf("metadata not found on %s", relay)
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

	for _, r := range relays {
		log.Println(r)
		sendUndelete(r, ev)
	}
}
