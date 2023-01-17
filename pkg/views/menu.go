package views 

func MainMenu(m TUIMain) string {
  menu := ""
  menu += "1. Print website header layout\n"
  
  return Layout(
    "Press number to continue",
    menu,
    m,
  )
}

