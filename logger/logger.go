package logger

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/olivere/elastic"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"sort"
	"time"
)

type Logger struct {
	client   *elastic.Client
	testMode bool
}
type Log struct {
	Type       string    `json:"type"`
	Message    string    `json:"message"`
	Error      string    `json:"error"`
	Stacktrace string    `json:"stack"`
	Created    time.Time `json:"created,omitempty"`
}

type LogInfo struct {
	Type    string    `json:"type"`
	Message string    `json:"message"`
	Created time.Time `json:"created,omitempty"`
}

const (
	indexName = "logs-adalpha"
	docType   = "log"
	mapping   = `
		{
			"settings":{
				"number_of_shards": 1,
				"number_of_replicas": 0
		},
			"mappings":{
				"log":{
					"properties":{
						"type":{
							"type":"keyword"
						},
						"message":{
							"type":"text",
							"store": true,
							"fielddata": true
						},
						"error":{
							"type":"text",
							"store": true,
							"fielddata": true
						},
						"stack":{
							"type":"text",
							"store": true,
							"fielddata": true
						},
						"created":{
							"type":"date"
						}
					}
				}
			}
		}`
)

//Creates logger instance, if running in tests testMode set to true, if not ES client created
func GetInstance() Logger {

	l := Logger{}
	if flag.Lookup("test.v") != nil {
		l.testMode = true
	}
	if !l.testMode {

		esUrl, err := url.Parse("http://elastic:9200")
		if err != nil {
			log.Fatalf("invalid -url flag: %v", err)
		}

		if len(os.Args) > 1 {
			switch os.Args[1] {
			case "showenv":
				log.Println("Environment")
				env := os.Environ()
				sort.Strings(env)
				for _, e := range env {
					log.Printf("- %s", e)
				}
				os.Exit(0)
			}
		}

		log.SetFlags(log.LstdFlags | log.Lshortfile)

		log.Printf("Running %s\n", runtime.Version())
		log.Printf("Version of github.com/olivere/elastic: %s\n", elastic.Version)

		log.Printf("Looking up hostname %q", esUrl.Hostname())
		ips, err := net.LookupIP(esUrl.Hostname())
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Lookup for hostname %q returns the following IPs:", esUrl.Hostname())
		for _, ip := range ips {
			log.Printf("%v", ip)
		}

		// Check ES version and status
		{
			log.Printf("Retrieving %s:", "http://elastic:9200")
			res, err := http.Get("http://elastic:9200")
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%v", string(body))
		}

		// Check ES nodes configuration
		{
			log.Printf("Retrieving %s:", "http://elastic:9200/_nodes/http?pretty=true")
			res, err := http.Get("http://elastic:9200/_nodes/http?pretty=true")
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%v", string(body))
		}

		log.Printf("Connecting to %s", "http://elastic:9200")
		client, err := elastic.NewClient(elastic.SetURL("http://elastic:9200"))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Connected to %s", "http://elastic:9200")

		err = createIndexWithLogsIfDoesNotExist(client)
		if err != nil {
			panic(err)
		}

		l.client = client

	}

	return l

}

//Logs error to ES, if test mode logs to stdout
func (l *Logger) LogError(eslog Log) {
	if l.testMode {
		log.Println(eslog.Message)
	} else {
		ctx := context.Background()
		_, err := l.client.Index().
			Index(indexName).
			Type(docType).
			BodyJson(eslog).
			Do(ctx)

		if err != nil {
			log.Println("Error logging to ES")
		}
	}
}

//Logs info log to ES, if test mode logs to stdout
func (l *Logger) LogInfo(info LogInfo) {
	if l.testMode {
		log.Println(info.Message)
	} else {
		ctx := context.Background()
		_, err := l.client.Index().
			Index(indexName).
			Type(docType).
			BodyJson(info).
			Do(ctx)

		if err != nil {
			log.Println("Error logging to ES")
		}
	}

}

//Initialises a Log
func CreateLog(logType string, msg string, err string, stack string, created time.Time) Log {
	l := Log{}
	l.Type = logType
	l.Message = msg
	l.Error = err
	l.Stacktrace = stack
	l.Created = created

	return l
}

//Initialises a InfoLog
func CreateInfoLog(logType string, msg string, created time.Time) LogInfo {
	l := LogInfo{}
	l.Type = logType
	l.Message = msg
	l.Created = created

	return l
}

//creates logs index in ES if does not exist
func createIndexWithLogsIfDoesNotExist(client *elastic.Client) error {
	ctx := context.Background()
	exists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	res, err := client.CreateIndex(indexName).Body(mapping).Do(ctx)

	if err != nil {
		return err
	}
	if !res.Acknowledged {
		return errors.New("CreateIndex was not acknowledged. Check that timeout value is correct.")
	}
	return err
}

//Returns stacktrace information
func Trace() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("%s,:%d %s\n", frame.File, frame.Line, frame.Function)
}
