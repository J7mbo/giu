package giu

import (
	"image/color"

	"github.com/AllenDang/giu/imgui"
)

type LineWidget struct {
	widgets []Widget
}

func Line(widgets ...Widget) *LineWidget {
	return &LineWidget{
		widgets: widgets,
	}
}

func (l *LineWidget) Build() {
	for i, w := range l.widgets {
		_, isTooltip := w.(*TooltipWidget)
		_, isContextMenu := w.(*ContextMenuWidget)
		_, isPopup := w.(*PopupWidget)
		_, isTabItem := w.(*TabItemWidget)

		if i > 0 && !isTooltip && !isContextMenu && !isPopup && !isTabItem {
			imgui.SameLine()
		}

		w.Build()
	}
}

type InputTextMultilineWidget struct {
	label         string
	text          *string
	width, height float32
	flags         InputTextFlags
	cb            imgui.InputTextCallback
	changed       func()
}

func (i *InputTextMultilineWidget) Build() {
	if imgui.InputTextMultilineV(i.label, i.text, imgui.Vec2{X: i.width, Y: i.height}, int(i.flags), i.cb) && i.changed != nil {
		i.changed()
	}
}

func InputTextMultiline(label string, text *string, width, height float32, flags InputTextFlags, cb imgui.InputTextCallback, changed func()) *InputTextMultilineWidget {
	return &InputTextMultilineWidget{
		label:   label,
		text:    text,
		width:   width,
		height:  height,
		flags:   flags,
		cb:      cb,
		changed: changed,
	}
}

type ButtonWidget struct {
	id      string
	width   float32
	height  float32
	clicked func()
}

func (b *ButtonWidget) Build() {
	if imgui.ButtonV(b.id, imgui.Vec2{X: b.width, Y: b.height}) && b.clicked != nil {
		b.clicked()
	}
}

func Button(id string, clicked func()) *ButtonWidget {
	return ButtonV(id, 0, 0, clicked)
}

func ButtonV(id string, width, height float32, clicked func()) *ButtonWidget {
	return &ButtonWidget{
		id:      id,
		width:   width,
		height:  height,
		clicked: clicked,
	}
}

type InvisibleButtonWidget struct {
	id      string
	width   float32
	height  float32
	clicked func()
}

func InvisibleButton(id string, width, height float32, clicked func()) *InvisibleButtonWidget {
	return &InvisibleButtonWidget{
		id:      id,
		width:   width,
		height:  height,
		clicked: clicked,
	}
}

func (ib *InvisibleButtonWidget) Build() {
	if imgui.InvisibleButton(ib.id, imgui.Vec2{X: ib.width, Y: ib.height}) && ib.clicked != nil {
		ib.clicked()
	}
}

type ImageButtonWidget struct {
	texture *Texture
	width   float32
	height  float32
	clicked func()
}

func (i *ImageButtonWidget) Build() {
	if imgui.ImageButton(i.texture.id, imgui.Vec2{}) && i.clicked != nil {
		i.clicked()
	}
}

func ImageButton(texture *Texture, width, height float32, clicked func()) *ImageButtonWidget {
	return &ImageButtonWidget{
		texture: texture,
		width:   width,
		height:  height,
		clicked: clicked,
	}
}

type CheckboxWidget struct {
	text     string
	selected *bool
	changed  func()
}

func (c *CheckboxWidget) Build() {
	if imgui.Checkbox(c.text, c.selected) && c.changed != nil {
		c.changed()
	}
}

func Checkbox(text string, selected *bool, changed func()) *CheckboxWidget {
	return &CheckboxWidget{
		text:     text,
		selected: selected,
		changed:  changed,
	}
}

type RadioButtonWidget struct {
	text    string
	active  bool
	changed func()
}

func (r *RadioButtonWidget) Build() {
	if imgui.RadioButton(r.text, r.active) && r.changed != nil {
		r.changed()
	}
}

func RadioButton(text string, active bool, changed func()) *RadioButtonWidget {
	return &RadioButtonWidget{
		text:    text,
		active:  active,
		changed: changed,
	}
}

type ChildWidget struct {
	id     string
	width  float32
	height float32
	border bool
	flags  int
	layout Layout
}

func (c *ChildWidget) Build() {
	imgui.BeginChildV(c.id, imgui.Vec2{X: c.width, Y: c.height}, c.border, c.flags)
	if c.layout != nil {
		c.layout.Build()
	}
	imgui.EndChild()
}

func Child(id string, border bool, width, height float32, flags int, layout Layout) *ChildWidget {
	return &ChildWidget{
		id:     id,
		width:  width,
		height: height,
		border: border,
		flags:  flags,
		layout: layout,
	}
}

type ComboWidget struct {
	label        string
	previewValue string
	items        []string
	selected     *int32
	flags        int
	changed      func()
}

func (c *ComboWidget) Build() {
	if imgui.BeginComboV(c.label, c.previewValue, c.flags) {
		for i, item := range c.items {
			if imgui.Selectable(item) {
				*c.selected = int32(i)
				if c.changed != nil {
					c.changed()
				}
			}
		}

		imgui.EndCombo()
	}
}

func Combo(label, previewValue string, items []string, selected *int32, flags int, changed func()) *ComboWidget {
	return &ComboWidget{
		label:        label,
		previewValue: previewValue,
		items:        items,
		selected:     selected,
		flags:        flags,
		changed:      changed,
	}
}

type ContextMenuWidget struct {
	label       string
	mouseButton int
	layout      Layout
}

func (c *ContextMenuWidget) Build() {
	if imgui.BeginPopupContextItemV(c.label, c.mouseButton) {
		if c.layout != nil {
			c.layout.Build()
		}
		imgui.EndPopup()
	}
}

func ContextMenu(layout Layout) *ContextMenuWidget {
	return ContextMenuV("", 1, layout)
}

func ContextMenuV(label string, mouseButton int, layout Layout) *ContextMenuWidget {
	return &ContextMenuWidget{
		label:       label,
		mouseButton: mouseButton,
		layout:      layout,
	}
}

type DragIntWidget struct {
	label  string
	value  *int32
	speed  float32
	min    int32
	max    int32
	format string
}

func (d *DragIntWidget) Build() {
	imgui.DragIntV(d.label, d.value, d.speed, d.min, d.max, d.format)
}

func DragInt(label string, value *int32) *DragIntWidget {
	return DragIntV(label, value, 1.0, 0, 0, "%d")
}

func DragIntV(label string, value *int32, speed float32, min, max int32, format string) *DragIntWidget {
	return &DragIntWidget{
		label:  label,
		value:  value,
		speed:  speed,
		min:    min,
		max:    max,
		format: format,
	}
}

type GroupWidget struct {
	layout Layout
}

func (g *GroupWidget) Build() {
	imgui.BeginGroup()
	if g.layout != nil {
		g.layout.Build()
	}
	imgui.EndGroup()
}

func Group(layout Layout) *GroupWidget {
	return &GroupWidget{
		layout: layout,
	}
}

type ImageWidget struct {
	texture *Texture
	width   float32
	height  float32
}

func (i *ImageWidget) Build() {
	size := imgui.Vec2{X: i.width, Y: i.height}
	if i.texture != nil && i.texture.id != 0 {
		rect := imgui.ContentRegionAvail()
		if size.X == -1 {
			size.X = rect.X
		}
		if size.Y == -1 {
			size.Y = rect.Y
		}
		imgui.Image(i.texture.id, size)
	}
}

func Image(texture *Texture, width, height float32) *ImageWidget {
	return &ImageWidget{
		texture: texture,
		width:   width,
		height:  height,
	}
}

type InputTextWidget struct {
	label   string
	value   *string
	width   float32
	flags   InputTextFlags
	cb      imgui.InputTextCallback
	changed func()
}

func (i *InputTextWidget) Build() {
	if i.width != 0 {
		PushItemWidth(i.width)
	}
	if imgui.InputTextV(i.label, i.value, int(i.flags), i.cb) && i.changed != nil {
		i.changed()
	}
}

func InputText(label string, width float32, value *string) *InputTextWidget {
	return InputTextV(label, width, value, 0, nil, nil)
}

func InputTextV(label string, width float32, value *string, flags InputTextFlags, cb imgui.InputTextCallback, changed func()) *InputTextWidget {
	return &InputTextWidget{
		label:   label,
		value:   value,
		width:   width,
		flags:   flags,
		cb:      cb,
		changed: changed,
	}
}

type LabelWidget struct {
	label string
	color *color.RGBA
	font  *imgui.Font
}

func (l *LabelWidget) Build() {
	if l.color != nil {
		PushColorText(*l.color)
	}

	if l.font != nil {
		PushFont(*l.font)
	}

	imgui.Text(l.label)

	if l.font != nil {
		PopFont()
	}

	if l.color != nil {
		PopStyleColor()
	}
}

func Label(label string) *LabelWidget {
	return LabelV(label, nil, nil)
}

func LabelV(label string, color *color.RGBA, font *imgui.Font) *LabelWidget {
	return &LabelWidget{
		label: label,
		color: color,
		font:  font,
	}
}

type MainMenuBarWidget struct {
	layout Layout
}

func (m *MainMenuBarWidget) Build() {
	if imgui.BeginMainMenuBar() {
		if m.layout != nil {
			m.layout.Build()
		}
		imgui.EndMainMenuBar()
	}
}

func MainMenuBar(layout Layout) *MainMenuBarWidget {
	return &MainMenuBarWidget{
		layout: layout,
	}
}

type MenuBarWidget struct {
	layout Layout
}

func (m *MenuBarWidget) Build() {
	if imgui.BeginMenuBar() {
		if m.layout != nil {
			m.layout.Build()
		}
		imgui.EndMenuBar()
	}
}

func MenuBar(layout Layout) *MenuBarWidget {
	return &MenuBarWidget{
		layout: layout,
	}
}

type MenuItemWidget struct {
	label    string
	selected bool
	enabled  bool
	clicked  func()
}

func (m *MenuItemWidget) Build() {
	if imgui.MenuItemV(m.label, "", m.selected, m.enabled) && m.clicked != nil {
		m.clicked()
	}
}

func MenuItem(label string) *MenuItemWidget {
	return MenuItemV(label, false, true, nil)
}

func MenuItemV(label string, selected, enabled bool, clicked func()) *MenuItemWidget {
	return &MenuItemWidget{
		label:    label,
		selected: selected,
		enabled:  enabled,
		clicked:  clicked,
	}
}

type MenuWidget struct {
	label   string
	enabled bool
	layout  Layout
}

func (m *MenuWidget) Build() {
	if imgui.BeginMenuV(m.label, m.enabled) {
		if m.layout != nil {
			m.layout.Build()
		}
		imgui.EndMenu()
	}
}

func Menu(label string, layout Layout) *MenuWidget {
	return MenuV(label, true, layout)
}

func MenuV(label string, enabled bool, layout Layout) *MenuWidget {
	return &MenuWidget{
		label:   label,
		enabled: enabled,
		layout:  layout,
	}
}

type PopupWidget struct {
	name   string
	open   *bool
	flags  int
	layout Layout
}

func (p *PopupWidget) Build() {
	if imgui.BeginPopupModalV(p.name, p.open, p.flags) {
		if p.layout != nil {
			p.layout.Build()
		}
		imgui.EndPopup()
	}
}

func Popup(name string, layout Layout) *PopupWidget {
	return PopupV(name, nil, 0, layout)
}

func PopupV(name string, open *bool, flags int, layout Layout) *PopupWidget {
	return &PopupWidget{
		name:   name,
		open:   open,
		flags:  flags,
		layout: layout,
	}
}

type ProgressBarWidget struct {
	fraction float32
	width    float32
	height   float32
	overlay  string
}

func (p *ProgressBarWidget) Build() {
	imgui.ProgressBarV(p.fraction, imgui.Vec2{X: p.width, Y: p.height}, p.overlay)
}

func ProgressBar(fraction float32, width, height float32, overlay string) *ProgressBarWidget {
	return &ProgressBarWidget{
		fraction: fraction,
		width:    width,
		height:   height,
		overlay:  overlay,
	}
}

type SelectableWidget struct {
	label    string
	selected bool
	flags    int
	width    float32
	height   float32
	clicked  func()
}

func (s *SelectableWidget) Build() {
	if imgui.SelectableV(s.label, s.selected, s.flags, imgui.Vec2{X: s.width, Y: s.height}) && s.clicked != nil {
		s.clicked()
	}
}

func Selectable(label string, clicked func()) *SelectableWidget {
	return SelectableV(label, false, 0, 0, 0, clicked)
}

type SelectableFlags int

const (
	// SelectableFlagsNone default = 0
	SelectableFlagsNone SelectableFlags = 0
	// SelectableFlagsDontClosePopups makes clicking the selectable not close any parent popup windows.
	SelectableFlagsDontClosePopups SelectableFlags = 1 << 0
	// SelectableFlagsSpanAllColumns allows the selectable frame to span all columns (text will still fit in current column).
	SelectableFlagsSpanAllColumns SelectableFlags = 1 << 1
	// SelectableFlagsAllowDoubleClick generates press events on double clicks too.
	SelectableFlagsAllowDoubleClick SelectableFlags = 1 << 2
	// SelectableFlagsDisabled disallows selection and displays text in a greyed out color.
	SelectableFlagsDisabled SelectableFlags = 1 << 3
)

func SelectableV(label string, selected bool, flags SelectableFlags, width, height float32, clicked func()) *SelectableWidget {
	return &SelectableWidget{
		label:    label,
		selected: selected,
		flags:    int(flags),
		width:    width,
		height:   height,
		clicked:  clicked,
	}
}

type SeparatorWidget struct{}

func (s *SeparatorWidget) Build() {
	imgui.Separator()
}

func Separator() *SeparatorWidget {
	return &SeparatorWidget{}
}

type SliderIntWidget struct {
	label  string
	value  *int32
	min    int32
	max    int32
	format string
}

func (s *SliderIntWidget) Build() {
	imgui.SliderIntV(s.label, s.value, s.min, s.max, s.format)
}

func SliderInt(label string, value *int32, min, max int32, format string) *SliderIntWidget {
	return &SliderIntWidget{
		label:  label,
		value:  value,
		min:    min,
		max:    max,
		format: format,
	}
}

type DummyWidget struct {
	width  float32
	height float32
}

func (d *DummyWidget) Build() {
	imgui.Dummy(imgui.Vec2{X: d.width, Y: d.height})
}

func Dummy(width, height float32) *DummyWidget {
	return &DummyWidget{
		width:  width,
		height: height,
	}
}

type HSplitterWidget struct {
	id     string
	width  float32
	height float32
	delta  *float32
}

func (h *HSplitterWidget) Build() {
	imgui.InvisibleButton(h.id, imgui.Vec2{X: h.width, Y: h.height})
	if imgui.IsItemActive() {
		*(h.delta) = imgui.CurrentIO().GetMouseDelta().Y
	} else {
		*(h.delta) = 0
	}
	if imgui.IsItemHovered() {
		imgui.SetMouseCursor(imgui.MouseCursorResizeNS)
	}
}

func HSplitter(id string, width, height float32, delta *float32) *HSplitterWidget {
	return &HSplitterWidget{
		id:     id,
		width:  width,
		height: height,
		delta:  delta,
	}
}

type VSplitterWidget struct {
	id     string
	width  float32
	height float32
	delta  *float32
}

func (v *VSplitterWidget) Build() {
	imgui.InvisibleButton(v.id, imgui.Vec2{X: v.width, Y: v.height})
	if imgui.IsItemActive() {
		*(v.delta) = imgui.CurrentIO().GetMouseDelta().X
	} else {
		*(v.delta) = 0
	}
	if imgui.IsItemHovered() {
		imgui.SetMouseCursor(imgui.MouseCursorResizeEW)
	}
}

func VSplitter(id string, width, height float32, delta *float32) *VSplitterWidget {
	return &VSplitterWidget{
		id:     id,
		width:  width,
		height: height,
		delta:  delta,
	}
}

type TabItemWidget struct {
	label  string
	open   *bool
	flags  int
	layout Layout
}

func (t *TabItemWidget) Build() {
	if imgui.BeginTabItemV(t.label, t.open, t.flags) {
		if t.layout != nil {
			t.layout.Build()
		}
		imgui.EndTabItem()
	}
}

func TabItem(label string, layout Layout) *TabItemWidget {
	return TabItemV(label, nil, 0, layout)
}

func TabItemV(label string, open *bool, flags int, layout Layout) *TabItemWidget {
	return &TabItemWidget{
		label:  label,
		open:   open,
		flags:  flags,
		layout: layout,
	}
}

type TabBarWidget struct {
	id     string
	flags  int
	layout Layout
}

func (t *TabBarWidget) Build() {
	if imgui.BeginTabBarV(t.id, t.flags) {
		if t.layout != nil {
			t.layout.Build()
		}
		imgui.EndTabBar()
	}
}

func TabBar(id string, layout Layout) *TabBarWidget {
	return TabBarV(id, 0, layout)
}

func TabBarV(id string, flags int, layout Layout) *TabBarWidget {
	return &TabBarWidget{
		id:     id,
		flags:  flags,
		layout: layout,
	}
}

type RowWidget struct {
	layout Layout
}

func (r *RowWidget) Build() {
	for i, w := range r.layout {
		if i > 0 {
			imgui.NextColumn()
		}
		w.Build()
	}
}

func Row(widgets ...Widget) *RowWidget {
	return &RowWidget{
		layout: widgets,
	}
}

type Rows []*RowWidget

type TabelWidget struct {
	label  string
	border bool
	rows   Rows
}

func (t *TabelWidget) Build() {
	if len(t.rows) > 0 && len(t.rows[0].layout) > 0 {
		imgui.ColumnsV(len(t.rows[0].layout), t.label, t.border)

		for i, r := range t.rows {
			if t.border {
				imgui.Separator()
			}

			if i > 0 {
				imgui.NextColumn()
			}

			r.Build()
		}

		imgui.Columns()
		if t.border {
			imgui.Separator()
		}
	}
}

func Table(label string, border bool, rows Rows) *TabelWidget {
	return &TabelWidget{
		label:  label,
		border: border,
		rows:   rows,
	}
}

type TooltipWidget struct {
	tip string
}

func (t *TooltipWidget) Build() {
	if imgui.IsItemHovered() {
		imgui.SetTooltip(t.tip)
	}
}

func Tooltip(tip string) *TooltipWidget {
	return &TooltipWidget{
		tip: tip,
	}
}

type TreeNodeWidget struct {
	label  string
	flags  int
	layout Layout
}

func (t *TreeNodeWidget) Build() {
	if imgui.TreeNodeV(t.label, t.flags) {
		if t.layout != nil {
			t.layout.Build()
		}
		if (t.flags & imgui.TreeNodeFlagsNoTreePushOnOpen) == 0 {
			imgui.TreePop()
		}
	}
}

func TreeNode(label string, flags int, layout Layout) *TreeNodeWidget {
	return &TreeNodeWidget{
		label:  label,
		flags:  flags,
		layout: layout,
	}
}

type SpacingWidget struct{}

func (s *SpacingWidget) Build() {
	imgui.Spacing()
}

func Spacing() *SpacingWidget {
	return &SpacingWidget{}
}

type CustomWidget struct {
	builder func()
}

func (c *CustomWidget) Build() {
	if c.builder != nil {
		c.builder()
	}
}

func Custom(builder func()) *CustomWidget {
	return &CustomWidget{
		builder: builder,
	}
}

type ConditionWidget struct {
	cond   bool
	layout Layout
}

func (c *ConditionWidget) Build() {
	if c.cond {
		c.layout.Build()
	}
}

func Condition(cond bool, layout Layout) *ConditionWidget {
	return &ConditionWidget{
		cond:   cond,
		layout: layout,
	}
}
