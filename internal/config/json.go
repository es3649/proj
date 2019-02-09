package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	fp "path/filepath"

	"github.com/es3649/proj/pkg/log"
	"github.com/sirupsen/logrus"
)

// LoadProjData loads the configuration files for the projects
// they are marshalled into a ProjData object
func LoadProjData(projRoot string) (*ProjData, error) {

	// open projRoot/.proj/projData.json

	projDataPath := fp.Join(projRoot, ".proj/projData.json")
	r, err := os.Open(projDataPath)
	if err != nil {
		log.Log.WithFields(logrus.Fields{"where": "config.LoadProjData", "error": err}).Info(fmt.Sprintf("Failed to open `%s'", projDataPath))
		return nil, fmt.Errorf("Failed to open `%s'", projDataPath)
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		log.Log.WithFields(logrus.Fields{"where": "config.LoadProjData", "error": err}).Info("Failed to read data from `projData.json'")
		return nil, fmt.Errorf("Failed to read data from `projData.json'")
	}

	var projData *ProjData

	err = json.Unmarshal(data, projData)
	if err != nil {
		log.Log.WithFields(logrus.Fields{"where": "config.LoadProjData", "error": err}).Info("Failed to unmarshal json data")
		return nil, fmt.Errorf("Failed to unmarshal json data")
	}

	return projData, nil
}

// SaveProjData writes a projData object to disk
func SaveProjData(projData *ProjData) error {
	log.Log.WithField("projData", projData).Debug("Saving project information")

	// do projData to JSON
	data, err := json.Marshal(*projData)
	if err != nil {
		log.Log.WithFields(logrus.Fields{"where": "config.SaveProjData", "error": err}).Info("Failed to Marshal to json")
		log.Log.WithField("ProjData", fmt.Sprintf("%#v", *projData)).Debug("Failed to Marshal to json")
		return fmt.Errorf("Failed to write to json")
	}

	// write the JSON to the .proj folder
	w, err := os.OpenFile(fp.Join(projData.Root, ".proj/projData.json"), os.O_WRONLY, 0666)
	defer w.Close()
	if err != nil {
		log.Log.WithFields(logrus.Fields{"where": "config.SaveProjData", "error": err}).Info("Failed to open `.proj/projData.json'")
		return fmt.Errorf("Failed to open `.proj/projData.json'")
	}
	fmt.Println(fp.Join(projData.Root, ".proj/projData.json.IS_OPEN"))

	_, err = w.Write(data)
	if err != nil {
		log.Log.WithFields(logrus.Fields{"where": "config.SaveProjData", "error": err}).Info("Failed to write the data")
		return fmt.Errorf("Failed to write the data")
	}
	w.Sync()

	return nil
}
