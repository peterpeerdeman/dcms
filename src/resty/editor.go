package resty

import (
	"encoding/json"
	"fmt"
	"github.com/sergi/go-diff/diff"
	"log"
	"storage"
	"strings"
)

type Patch struct {
	Diffs   [][2]interface{} `json:"diffs`
	Start1  int              `json:"start1"`
	Start2  int              `json:"start2"`
	Length1 int              `json:"length1"`
	Length2 int              `json:"length2"`
}

func (this *Patch) ToDMP() []diffmatchpatch.Patch {
	var patch diffmatchpatch.Patch
	patch.Start1 = this.Start1
	patch.Start2 = this.Start2
	patch.Length1 = this.Length1
	patch.Length2 = this.Length2
	patch.Diffs = make([]diffmatchpatch.Diff, len(this.Diffs))
	for diffIndex, diff := range this.Diffs {
		patch.Diffs[diffIndex] = diffmatchpatch.Diff{int8(diff[0].(float64)), diff[1].(string)}
	}
	return []diffmatchpatch.Patch{patch}
}

type Update struct {
	Hash  int     `json:"hash"`
	Patch []Patch `json:"patch"`
}

func handleNotification(notification Notification) {
	if strings.HasPrefix(notification.Type, "document.") {
		id := strings.Replace(notification.Type, "document.", "", -1)
		content, err := storage.Repo.Get(fmt.Sprintf("/documents/%s", id))
		if err != nil {
			log.Panicf("Unable to retrieve document %v", err)
			return
		}
		document := string(content)
		data := notification.Data.(string)
		var update Update
		json.Unmarshal([]byte(data), &update)
		patches := update.Patch[0].ToDMP()
		dmp := diffmatchpatch.New()
		updatedDocument, status := dmp.PatchApply(patches, document)
		if status[0] {
			storage.Repo.Add(fmt.Sprintf("/documents/%s", id), []byte(updatedDocument))
		}
	}
}
