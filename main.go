package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

var name string
var ip string
var ttl int64
var weight = int64(1)
var zoneID string

func init() {
	flag.StringVar(&name, "d", "", "domain name")
	flag.StringVar(&ip, "ip", getIP(), "ip for a record (default is your current wan ip)")
	flag.StringVar(&zoneID, "z", "", "AWS Zone Id for domain")
	flag.Int64Var(&ttl, "ttl", int64(60), "ttl for DNS Cache")

}

func main() {
	flag.Parse()
	if name == "" || zoneID == "" {
		log.Fatalf("Incomplete arguments: d: %s, ip: %s, z: %s", name, ip, zoneID)
		flag.PrintDefaults()
		return
	}

	// check if ip is same as record, so no updated is needed.
	// else create aws session, and update record
	if getIP() != lookupRecord(name) {
		sess, err := session.NewSession()
		if err != nil {
			log.Fatalln("failed to create session,", err)
			return
		}

		svc := route53.New(sess)
		createARecord(svc)
		log.Print("updating record " + name + " to new value " + ip)

	} else {
		log.Println("No updated needed, record is up to date!", getIP())
	}
}

func createARecord(svc *route53.Route53) {

	params := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{ // Required
			Changes: []*route53.Change{ // Required
				{ // Required
					Action: aws.String("UPSERT"), // Required
					ResourceRecordSet: &route53.ResourceRecordSet{ // Required
						Name: aws.String(name), // Required
						Type: aws.String("A"),  // Required
						ResourceRecords: []*route53.ResourceRecord{
							{ // Required
								Value: aws.String(ip), // Required
							},
						},
						TTL:           aws.Int64(ttl),
						Weight:        aws.Int64(weight),
						SetIdentifier: aws.String("Arbitrary Id describing this change set"),
					},
				},
			},
			Comment: aws.String("updating record " + name + " to new value " + ip),
		},
		HostedZoneId: aws.String(zoneID), // Required
	}

	// Change to resp, err := svc.ChangeResourceRecordSets(params)
	// fmt.Println("Change Response:")
	// fmt.Println(resp)
	_, err := svc.ChangeResourceRecordSets(params)
	if err != nil {
		log.Fatalln("Error updating record:", err)
		return
	}
}

func getIP() string {
	resp, err := http.Get("https://diagnostic.opendns.com/myip")
	if err != nil {
		log.Fatalln("Error getting ip:", err)
		return ""
	}
	defer resp.Body.Close()

	bs, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln("Error parsing ip:", err)
		return ""
	}
	return string(bs)
}

func lookupRecord(n string) string {
	l, err := net.LookupIP(n)
	if err != nil {
		log.Fatalln("Error parsing ip:", err)
		return ""
	}
	return l[0].String()
}
