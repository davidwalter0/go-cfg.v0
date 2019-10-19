package cfg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// func SaveCfg(sptr interface{}, filename string) (err error) {
// 	var sti *StructInfo = &StructInfo{
// 		StructPtr: sptr,
// 	}

// 	ioutil.WriteFile(filename, []byte(sti.DumpStructInfo()), 0664)

// 	return err
// }

// // Dump info from struct
// func (app *App) Dump() {
// 	if app.Debug || app.DumpCfg {
// 		var j = []byte{}
// 		if j, err = json.MarshalIndent(app, "", "  "); err != nil {
// 			log.Println(err)
// 			os.Exit(1)
// 		}
// 		fmt.Println(string(j))
// 	}
// 	if app.DumpCfg {
// 		os.Exit(0)
// 	}
// }

// Jsonify an object
func Jsonify(data interface{}) string {
	s, err := json.MarshalIndent(data, "", "  ") // spaces)
	if err != nil {
		return fmt.Sprintf("%+v", err)
	}
	return string(s)
}

// Load a file to a byte array
func Load(path string) (content []byte, err error) {
	if content, err = ioutil.ReadFile(path); err == nil {
		if err != nil {
			content = nil
		}
	}
	return
}

// **TODO** need to decide on standard method for passing in the
// configuration file name

// LoadCfg from a file to initialize the certificate definition
func LoadCfg(o interface{}, filename string) error {

	text, err := Load(filename)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(text, o); err != nil {
		return err
	}

	fmt.Printf("%+v\n", o)
	return nil
}

// LoadCfg from a file to initialize the certificate definition
func (sti *StructInfo) LoadCfg() error {
	if sti.LoadCfgSet() {
		text, err := Load("config.test.yaml")
		if err != nil {
			return err
		}
		if err = yaml.Unmarshal(text, sti.StructPtr); err != nil {
			return err
		}
	}
	fmt.Printf("%+v\n", sti.StructPtr)
	return nil
}

// SaveCfg from a file to initialize the certificate definition
func SaveCfg(o interface{}, filename string) error {
	var text []byte
	var err error
	if text, err = yaml.Marshal(o); err != nil {
		return err
	}
	Save(filename, text)

	return nil
}

// SaveCfg from a file to initialize the certificate definition
func (sti *StructInfo) SaveCfg() error {
	var text []byte
	var err error
	if sti.SaveCfgSet() {
		if text, err = yaml.Marshal(sti.StructPtr); err != nil {
			return err
		}
		ApplTmpFile(text)
	}
	return nil
}

// ApplTmpFile writes to a generated tmpfile name using the
// application name from os.Arg[0] as the prefix
func Save(filename string, content []byte) error {
	var err error
	w, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer func() {
		_ = w.Close()
	}()
	_, err = w.Write(content)

	return err
}

// ApplTmpFile writes to a generated tmpfile name using the
// application name from os.Arg[0] as the prefix
func ApplTmpFile(content []byte) error {
	var err error
	filename := filepath.Base(os.Args[0])

	tmpfile, err := ioutil.TempFile("", filename+"-*.yaml")

	if err != nil {
		return err
	}

	defer tmpfile.Close()

	_, err = tmpfile.Write(content)

	return err
}
