package directory

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

type FileEntry struct {
	Path    string `json:"path"`
	Content string `json:"content,omitempty"`
	Error   string `json:"error,omitempty"`
}

/*
Directory implements the sort.Interface, making it sortable using
sort.Sort(*Directory).
This is particularly useful to reduce client side sorting code (which is
clumsy, thanks to Javascript).
*/
type Directory struct {
	Exists bool         `json:"exists"`
	Files  []*FileEntry `json:"files"`
}

func (dl *Directory) Len() int {
	if dl.Files == nil {
		return 0
	}
	return len(dl.Files)
}

func (dl *Directory) Less(i, j int) bool {
	return dl.Files[i].Path < dl.Files[j].Path
}

func (dl *Directory) Swap(i, j int) {
	dl.Files[j], dl.Files[i] = dl.Files[i], dl.Files[j]
}

/*
Create a Directory and return its pointer.
Additionally automatically sorts the listing using sort.Sort(*Directory)
*/
func New(rootpath string) (*Directory, error) {
	var (
		err error
	)
	if _, err = os.Stat(rootpath); err == nil {
		files := []*FileEntry{}
		filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
			if os.IsNotExist(err) || info.IsDir() {
				// does not exist anymore or is a directory
				return nil
			}
			key := path[len(rootpath):]
			val, err := ioutil.ReadFile(path)
			if err == nil {
				files = append(files, &FileEntry{key, string(val), ""})
			} else {
				files = append(files, &FileEntry{key, "", err.Error()})
			}
			return nil
		})
		dl := &Directory{true, files}
		sort.Sort(dl)
		return dl, nil
	} else {
		if os.IsNotExist(err) {
			return &Directory{false, nil}, nil
		}
	}
	return nil, err
}
