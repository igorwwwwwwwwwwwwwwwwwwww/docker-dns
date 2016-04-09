package main

import (
	"net"
	"regexp"
	"sync"
	"time"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/miekg/dns"
	"golang.org/x/net/context"
)

var mu sync.Mutex
var active = map[string]types.Container{}

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	// fmt.Println(">")
	// fmt.Println(r)

	m := new(dns.Msg)
	m.SetReply(r)

	for _, question := range r.Question {
		re := regexp.MustCompile("(.*)\\.docker.")
		matches := re.FindStringSubmatch(question.Name)
		if len(matches) < 2 {
			continue
		}

		id := matches[1]

		mu.Lock()
		c, ok := active[id]
		mu.Unlock()
		if !ok {
			continue
		}

		network, ok := c.NetworkSettings.Networks["bridge"]
		if !ok {
			continue
		}

		m.Answer = append(m.Answer, &dns.A{
			Hdr: dns.RR_Header{
				Name:   question.Name,
				Rrtype: question.Qtype,
				Class:  question.Qclass,
				Ttl:    10,
			},
			A: net.ParseIP(network.IPAddress),
		})
	}

	w.WriteMsg(m)

	// fmt.Println("<")
	// fmt.Println(m)
}

func manageContainers(cli *client.Client) {
	for {
		options := types.ContainerListOptions{}
		containers, err := cli.ContainerList(context.Background(), options)
		if err != nil {
			panic(err)
		}

		mu.Lock()
		for _, c := range containers {
			active[c.ID] = c
			active[c.ID[:12]] = c
			for _, name := range c.Names {
				active[name[1:]] = c
			}
		}
		mu.Unlock()

		time.Sleep(3 * time.Second)
	}
}

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	go manageContainers(cli)

	server := &dns.Server{Addr: ":5300", Net: "udp"}
	dns.HandleFunc("docker.", handleRequest)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
