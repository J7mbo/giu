package main

import (
	"fmt"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
)

var (
	name         string
	items        []string
	itemSelected int32
	checked      bool
	checked2     bool
	dragInt      int32
	multiline    string
	radioOp      int
)

func btnClickMeClicked() {
	fmt.Println("Click me is clicked")
}

func comboChanged() {
	fmt.Println(items[itemSelected])
}

func contextMenu1Clicked() {
	fmt.Println("Context menu 1 is clicked")
}

func contextMenu2Clicked() {
	fmt.Println("Context menu 2 is clicked")
}

func btnPopupCLicked() {
	imgui.OpenPopup("Confirm")
}

func loop() {
	// Create main menu bar for master window.
	g.MainMenuBar(
		g.Layout{
			g.Menu("File", g.Layout{
				g.MenuItem("Open"),
				g.MenuItem("Save"),
				// You could add any kind of widget here, not just menu item.
				g.Menu("Save as ...", g.Layout{
					g.MenuItem("Excel file"),
					g.MenuItem("CSV file"),
					g.Button("Button inside menu", nil),
				},
				),
			},
			),
		},
	).Build()

	// Build a new window
	size := g.Context.GetPlatform().DisplaySize()
	g.Window("Overview", 0, 20, size[0], size[1], g.Layout{
		g.Label("One line label"),
		g.Line(
			g.InputText("##name", 0, &name),
			g.Button("Click Me", btnClickMeClicked),
			g.Tooltip("I'm a tooltip"),
		),

		g.Line(
			g.Checkbox("Checkbox", &checked, func() {
				fmt.Println(checked)
			}),
			g.Checkbox("Checkbox 2", &checked2, func() {
				fmt.Println(checked2)
			}),
			g.Dummy(30, 0),
			g.RadioButton("Radio 1", radioOp == 0, func() { radioOp = 0 }),
			g.RadioButton("Radio 2", radioOp == 1, func() { radioOp = 1 }),
			g.RadioButton("Radio 3", radioOp == 2, func() { radioOp = 2 }),
		),

		g.ProgressBar(0.8, -1, 0, "Progress"),
		g.DragInt("DragInt", &dragInt),
		g.SliderInt("Slider", &dragInt, 0, 100, ""),

		g.Combo("Combo", items[itemSelected], items, &itemSelected, 0, comboChanged),

		g.Line(
			g.Button("Popup Modal", btnPopupCLicked),
			g.Popup("Confirm", g.Layout{
				g.Label("Confirm to close me?"),
				g.Line(
					g.Button("Yes", func() { imgui.CloseCurrentPopup() }),
					g.Button("No", nil),
				),
			}),
			g.Label("Right click me to see the context menu"),
			g.ContextMenu(g.Layout{
				g.Selectable("Context menu 1", contextMenu1Clicked),
				g.Selectable("Context menu 2", contextMenu2Clicked),
			}),
		),

		g.TabBar("Tabbar Input", g.Layout{
			g.TabItem("Multiline Input", g.Layout{
				g.Label("This is first tab with a multiline input text field"),
				g.InputTextMultiline("##multiline", &multiline, 0, 0, 0, nil, nil),
			}),
			g.TabItem("Tree", g.Layout{
				g.TreeNode("TreeNode1", imgui.TreeNodeFlagsCollapsingHeader|imgui.TreeNodeFlagsDefaultOpen, g.Layout{
					g.Label("Tree node 1"),
					g.Label("Tree node 1"),
					g.Label("Tree node 1"),
					g.Button("Button inside tree", nil),
				}),
				g.TreeNode("TreeNode2", 0, g.Layout{
					g.Label("Tree node 2"),
					g.Label("Tree node 2"),
					g.Label("Tree node 2"),
					g.Button("Button inside tree", nil),
				}),
			}),
			g.TabItem("Table", g.Layout{
				g.Table("Table", true, g.Rows{
					g.Row(g.Label("Name"), g.Label("Age"), g.Label("Location")),
					g.Row(g.Label("Allen"), g.Label("33"), g.Label("Shanghai/China")),
					g.Row(g.Checkbox("check me", &checked, nil), g.Button("click me", nil), g.Label("Anything")),
				}),
			}),
			g.TabItem("Group", g.Layout{
				g.Line(
					g.Group(g.Layout{
						g.Label("I'm inside group 1"),
					}),
					g.Group(g.Layout{
						g.Label("I'm inside group 2"),
					}),
				),
			}),
		}),
	})
}

func main() {
	items = make([]string, 100)
	for i := range items {
		items[i] = fmt.Sprintf("Item %d", i)
	}

	w := g.NewMasterWindow("Overview", 800, 600, true, nil)
	w.Main(loop)
}
