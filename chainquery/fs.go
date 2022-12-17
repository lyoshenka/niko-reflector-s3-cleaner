package chainquery

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/nikooo777/reflector-s3-cleaner/shared"

	"github.com/lbryio/lbry.go/v2/extras/errors"

	"github.com/sirupsen/logrus"
)

func SaveHashes(streamData []shared.StreamData, path string) error {
	logrus.Printf("saving %d stream data objects to %s", len(streamData), path)

	buff := bytes.Buffer{}
	enc := json.NewEncoder(&buff)

	for _, item := range streamData {
		err := enc.Encode(item)
		if err != nil {
			return errors.Err(err)
		}
	}

	err := ioutil.WriteFile(path, buff.Bytes(), 0644)
	if err != nil {
		return errors.Err(err)
	}
	return nil
}

func LoadResolvedHashes(path string) ([]shared.StreamData, error) {
	logrus.Printf("loading stream data objects from %s", path)
	streamData := make([]shared.StreamData, 0)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Err(err)
	}

	dec := json.NewDecoder(bytes.NewReader(content))
	for {
		var item shared.StreamData
		err := dec.Decode(&item)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.Err(err)
		}
		streamData = append(streamData, item)
	}

	logrus.Printf("loaded %d stream data objects from %s", len(streamData), path)
	return streamData, nil
}
