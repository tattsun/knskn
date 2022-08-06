package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

//go:embed static/*
var staticRoot embed.FS

type Data struct {
	Timestamp int64
	Temp      float64
	Press     float64
	Hum       float64
}

func main() {
	dbFilePath := os.Getenv("KNSKN_DB_FILEPATH")
	if dbFilePath == "" {
		dbFilePath = "./knskn.db"
	}

	bme, err := NewBME()
	if err != nil {
		log.Fatalf("failed to initialize BME: %+v", err)
	}

	db, err := leveldb.OpenFile(dbFilePath, nil)
	if err != nil {
		log.Fatalf("failed to open db: %+v", err)
	}
	defer db.Close()

	go func() {
		for {
			temp, press, hum, err := bme.GetEnv()
			if err != nil {
				log.Printf("failed to get environments: %+v", err)
			}

			key := []byte(strconv.FormatInt(time.Now().UnixMilli(), 10))
			value := []byte(fmt.Sprintf("%.2f,%.2f,%.2f", temp, press, hum))
			if err := db.Put(key, value, nil); err != nil {
				log.Printf("failed to put data to DB: %+v", err)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		from := r.URL.Query().Get("from")
		if from == "" {
			from = strconv.FormatInt(time.Now().Add(-24*time.Hour).UnixMilli(), 10)
		}

		to := r.URL.Query().Get("to")
		if to == "" {
			to = strconv.FormatInt(time.Now().UnixMilli(), 10)
		}

		datum := []Data{}

		iter := db.NewIterator(&util.Range{Start: []byte(from), Limit: []byte(to)}, nil)
		defer iter.Release()
		for iter.Next() {
			timestamp, err := strconv.ParseInt(string(iter.Key()), 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "unexpected error")
				return
			}

			parts := strings.Split(string(iter.Value()), ",")
			temp, err := strconv.ParseFloat(parts[0], 64)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "unexpected error")
				return
			}

			press, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "unexpected error")
				return
			}

			hum, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "unexpected error")
				return
			}

			datum = append(datum, Data{
				Timestamp: timestamp,
				Temp:      temp,
				Press:     press,
				Hum:       hum,
			})
		}
		if err := iter.Error(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "unexpected error")
			return
		}

		w.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(w)
		enc.Encode(datum)
	})

	static, err := fs.Sub(staticRoot, "static")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.FS(static)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
