package util

import (
	"bytes"
	"fmt"
	glua "github.com/yuin/gopher-lua"
	"html/template"
	"io/ioutil"
	"sync"
)

var (
	// WidgetTemplate holds all the widget templates
	WidgetTemplate Tmpl

	// Widgets holds all the widget results
	Widgets = WidgetList{
		rw: &sync.RWMutex{},
	}
)

// WidgetList the list of widget application
type WidgetList struct {
	rw   *sync.RWMutex
	List []*Widget
}

// Widget is an application sidebar content
type Widget struct {
	rw     *sync.RWMutex
	Name   string
	Result template.HTML
}

// Load loads all the widgets from the given directory
func (w *WidgetList) Load(path string) error {
	// Lock widget list
	w.rw.Lock()

	// Unlock widget list
	defer w.rw.Unlock()

	// Load all directories of the widget folder
	widgets, err := ioutil.ReadDir(path)

	if err != nil {
		return err
	}

	// Set slice
	w.List = []*Widget{}

	// Loop folder items
	for _, file := range widgets {

		// Check if file is a directory
		if file.IsDir() {

			// Append widget
			w.List = append(w.List, &Widget{
				Name: file.Name(),
				rw:   &sync.RWMutex{},
			})
		}
	}

	return nil
}

// IsCached checks if the given widget is cached
func (w *Widget) IsCached() (template.HTML, bool) {
	// Lock mutex
	w.rw.RLock()
	defer w.rw.RUnlock()

	// Get value from cache
	buff, found := Cache.Get(
		fmt.Sprintf("widget_%v", w.Name),
	)

	if !found {
		return "", false
	}

	return buff.(template.HTML), true
}

// Execute gets the result of the given widget
func (w *Widget) Execute(luaState *glua.LState) (*glua.LTable, bool, error) {
	// Lock mutex
	w.rw.RLock()
	defer w.rw.RUnlock()

	// Execute widget function
	if err := luaState.CallByParam(glua.P{
		Fn:      luaState.GetGlobal("widget"),
		NRet:    2,
		Protect: true,
	}); err != nil {
		return nil, false, err
	}

	// Get value of the top of the stack
	v := luaState.Get(-2)

	// Remove top value
	luaState.Pop(-2)

	// Get cache option
	c := luaState.ToBool(-1)

	// Remove top value
	luaState.Pop(-1)

	// Check for valid returned type
	if v.Type() != glua.LTTable {
		return nil, c, nil
	}

	// Return value as table
	return v.(*glua.LTable), c, nil
}

// SetResult changes a widget result
func (w *Widget) SetResult(buff *bytes.Buffer) {
	// Lock mutex
	w.rw.Lock()
	defer w.rw.Unlock()

	// Set result
	w.Result = template.HTML(buff.String())
}
