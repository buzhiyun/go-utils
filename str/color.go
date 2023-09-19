package str

import "fmt"

// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色
//
// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见

type ColorFont int
type ColorBG int

const (
	// 字体本身
	BlackFont  ColorFont = 30 + iota // 黑色
	RedFont                          // 红色
	GreenFont                        // 绿色
	YellowFont                       // 黄色
	BlueFont                         // 蓝色
	PurpleFont                       // 紫色
	CyanFont                         // 青色
	WhiteFont                        // 白色

	// 背景色
	BlackBG  ColorBG = 40 + iota
	RedBG            // 红色
	GreenBG          // 绿色
	YellowBG         // 黄色
	BlueBG           // 蓝色
	PurpleBG         // 紫色
	CyanBG           // 青色
	WhiteBG          // 白色
)

/*
*
获取带有颜色的字符串 ，windows 不能用
*/
func ColorfulString(str string, fontColor ColorFont, backgroudColor ColorBG) string {
	return fmt.Sprintf("\033[1;%v;%vm%s\033[0m", fontColor, backgroudColor, str)
}

/*
*
显示带有颜色的字符串 ，windows 不能用
*/
func ColorfulPrint(str string, fontColor ColorFont, backgroudColor ColorBG) {
	fmt.Printf("\033[1;%v;%vm%s\033[0m", fontColor, backgroudColor, str)
}
