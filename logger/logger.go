package logger

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"runtime"
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

		client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false))
		if err != nil {
			log.Panic(err)
		}

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
