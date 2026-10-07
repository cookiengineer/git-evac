//go:build wasm

package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/components/interfaces"
import "git-evac/structs"
import "strconv"
import "strings"

type SettingsArticle struct {
	Name      string
	Component *components.Component
	Backup    *dom.Element
	Folder    *dom.Element
	Port      *dom.Element
	scheduler interfaces.Scheduler
}

func ToSettingsArticle(element *dom.Element) *SettingsArticle {

	var article SettingsArticle

	component := components.NewComponent(element)

	article.Component = &component
	article.Name = strings.TrimSpace(element.GetAttribute("data-name"))

	return &article

}

func (article *SettingsArticle) Disable() bool {
	return false
}

func (article *SettingsArticle) Enable() bool {
	return false
}

func (article *SettingsArticle) SetScheduler(scheduler interfaces.Scheduler) {

	article.scheduler = scheduler

	if article.Component != nil {
		article.Component.SetScheduler(scheduler)
	}

}

func (article *SettingsArticle) Mount() bool {

	if article.Component.Element != nil {

		article.Backup = article.Component.Element.QuerySelector("#settings-backup")
		article.Folder = article.Component.Element.QuerySelector("#settings-folder")
		article.Port = article.Component.Element.QuerySelector("#settings-port")

		return true

	}

	return false

}

func (article *SettingsArticle) Query(query string) interfaces.Component {

	selectors := utils.SplitQuery(query)

	if len(selectors) == 1 {

		if article.Component.Element != nil {

			if utils.MatchesQuery(article.Component.Element, selectors[0]) == true {
				return article
			}

		}

	}

	return nil

}

func (article *SettingsArticle) Render() *dom.Element {
	return article.Component.Element
}

func (article *SettingsArticle) String() string {

	if article.Component != nil {
		return article.Component.String()
	}

	return ""

}

func (article *SettingsArticle) Unmount() bool {
	return true
}

func (article *SettingsArticle) SetSchema(settings *structs.Settings) {

	if settings == nil {
		return
	}

	if article.Backup != nil {
		setValue(article.Backup, settings.GetBackup())
	}

	if article.Folder != nil {
		setValue(article.Folder, settings.GetFolder())
	}

	if article.Port != nil {
		setValue(article.Port, strconv.FormatUint(uint64(settings.GetPort()), 10))
	}

}

func (article *SettingsArticle) GetBackup() string {
	return readValue(article.Backup)
}

func (article *SettingsArticle) GetFolder() string {
	return readValue(article.Folder)
}

func (article *SettingsArticle) GetPort() uint16 {

	port, err := strconv.ParseUint(readValue(article.Port), 10, 16)

	if err == nil {
		return uint16(port)
	}

	return 0

}
