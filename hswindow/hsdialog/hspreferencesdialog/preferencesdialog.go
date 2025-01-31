// Package hspreferencesdialog contains preferences dialog data
package hspreferencesdialog

import (
	"image/color"

	"github.com/OpenDiablo2/dialog"
	g "github.com/ianling/giu"

	"github.com/OpenDiablo2/HellSpawner/hscommon/hsutil"
	"github.com/OpenDiablo2/HellSpawner/hsconfig"
	"github.com/OpenDiablo2/HellSpawner/hswindow/hsdialog"
)

const (
	mainWindowW, mainWindowH = 300, 200
	textboxSize              = 245
	btnW, btnH               = 30, 0
)

// PreferencesDialog represents preferences dialog
type PreferencesDialog struct {
	*hsdialog.Dialog

	config             *hsconfig.Config
	onConfigChanged    func(config *hsconfig.Config)
	windowColorChanger func(c color.RGBA)
}

// Create creates a new preferences dialog
func Create(onConfigChanged func(config *hsconfig.Config), windowColorChanger func(c color.RGBA)) *PreferencesDialog {
	result := &PreferencesDialog{
		Dialog:             hsdialog.New("Preferences"),
		onConfigChanged:    onConfigChanged,
		windowColorChanger: windowColorChanger,
	}
	result.Visible = false

	return result
}

// Build builds a preferences dialog
func (p *PreferencesDialog) Build() {
	p.IsOpen(&p.Visible).Layout(
		g.Child("PreferencesLayout").Size(mainWindowW, mainWindowH).Layout(
			g.Label("Auxiliary MPQ Path"),
			g.Line(
				g.InputText("##AppPreferencesAuxMPQPath", &p.config.AuxiliaryMpqPath).Size(textboxSize).Flags(g.InputTextFlags_ReadOnly),
				g.Button("...##AppPreferencesAuxMPQPathBrowse").Size(btnW, btnH).OnClick(p.onBrowseAuxMpqPathClicked),
			),
			g.Separator(),
			g.Label("External MPQ listfile Path"),
			g.Line(
				g.InputText("##AppPreferencesListfilePath", &p.config.ExternalListFile).Size(textboxSize).Flags(g.InputTextFlags_ReadOnly),
				g.Button("...##AppPreferencesListfilePathBrowse").Size(btnW, btnH).OnClick(p.onBrowseExternalListfileClicked),
			),
			g.Separator(),
			g.Label("Abyss Engine Path"),
			g.Line(
				g.InputText("##AppPreferencesAbyssEnginePath", &p.config.AbyssEnginePath).Size(textboxSize).Flags(g.InputTextFlags_ReadOnly),
				g.Button("...##AppPreferencesAbyssEnginePathBrowse").Size(btnW, btnH).OnClick(p.onBrowseAbyssEngineClicked),
			),
			g.Separator(),
			g.Checkbox("Open most recent project on start-up", &p.config.OpenMostRecentOnStartup),
			g.Separator(),
			g.Label("Background color:"),
			g.Line(
				g.ColorEdit("##BackgroundColor", &p.config.BGColor).
					Flags(g.ColorEditFlagsNoAlpha).OnChange(func() {
					p.windowColorChanger(p.config.BGColor)
				}),
				g.Button("Default##BackgroundColorDefault").OnClick(func() {
					p.config.BGColor = hsutil.Color(hsconfig.DefaultBGColor)
					p.windowColorChanger(p.config.BGColor)
				}),
			),
		),
		g.Line(
			g.Button("Save##AppPreferencesSave").OnClick(p.onSaveClicked),
			g.Button("Cancel##AppPreferencesCancel").OnClick(p.onCancelClicked),
		),
	).Build()
}

// Show switch preferences dialog to visible state
func (p *PreferencesDialog) Show(config *hsconfig.Config) {
	p.Dialog.Show()

	p.config = config
}

func (p *PreferencesDialog) onBrowseAuxMpqPathClicked() {
	path, err := dialog.Directory().Browse()
	if err != nil || path == "" {
		return
	}

	p.config.AuxiliaryMpqPath = path
}

func (p *PreferencesDialog) onBrowseExternalListfileClicked() {
	path := dialog.File()
	path.Filter("Text file", "txt")
	filePath, err := path.Load()

	if err != nil || filePath == "" {
		return
	}

	p.config.ExternalListFile = filePath
}

func (p *PreferencesDialog) onSaveClicked() {
	p.onConfigChanged(p.config)
	p.Visible = false
}

func (p *PreferencesDialog) onCancelClicked() {
	p.Visible = false
}

func (p *PreferencesDialog) onBrowseAbyssEngineClicked() {
	path := dialog.File()

	filePath, err := path.Load()

	if err != nil || filePath == "" {
		return
	}

	p.config.AbyssEnginePath = filePath
}
