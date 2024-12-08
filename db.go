package main

import (
	"database/sql"
	"encoding/json"
	"math"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Db struct {
	db *sql.DB
}

func NewDb(path string) (*Db, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `requests` (`method` TEXT NOT NULL, `protocol` TEXT NOT NULL, `host` TEXT NOT NULL, `path` TEXT NOT NULL, `headers` LONGBLOB, `cookies` LONGBLOB, `remoteaddr` TEXT NOT NULL, `multipartform` TEXT, `requesttime` TIMESTAMP NOT NULL);")
	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

func (db *Db) InsertRequest(request *http.Request) error {
	stmt, err := db.db.Prepare("insert into requests (method, protocol, host, path, headers, cookies, remoteaddr, multipartform, requesttime) values (?, ?, ?, ?, ?, ?, ?, ?, unixepoch());")
	if err != nil {
		return err
	}

	headers, err := json.Marshal(request.Header)
	if err != nil {
		return err
	}

	cookies, err := json.Marshal(request.Cookies())
	if err != nil {
		return err
	}

	var multipart []byte
	if request.Header.Get("Content-Type") == "multipart/form-data" {
		err = request.ParseMultipartForm(math.MaxInt32)
		if err != nil {
			return err
		}
		multipart, err = json.Marshal(request.MultipartForm)
		if err != nil {
			return err
		}
	}

	_, err = stmt.Exec(request.Method, request.Proto, request.Host, request.RequestURI, string(headers), string(cookies), request.RemoteAddr, string(multipart))
	if err != nil {
		return err
	}

	return nil
}
