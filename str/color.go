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

type ColorFont string
type ColorBG string
type FontStyle int

const (
	FontDefault      FontStyle = 0 + iota // 默认
	FontBold                              // 高亮
	FontFuzzy                             // **模糊
	FontItalics                           // **斜体
	FontUnderline                         // 下划线
	FontSlowBlinking                      //慢闪烁
	FontFastBlinking                      // **快闪烁
	FontReverseColor                           // 反白显示
	FontHiding                            // 不可见

	// 字体本身
	BlackFont  ColorFont = ";30" // 黑色
	RedFont    ColorFont = ";31" // 红色
	GreenFont  ColorFont = ";32" // 绿色
	YellowFont ColorFont = ";33" // 黄色
	BlueFont   ColorFont = ";34" // 蓝色
	PurpleFont ColorFont = ";35" // 紫色
	CyanFont   ColorFont = ";36" // 青色
	WhiteFont  ColorFont = ";37" // 白色

	// 背景色
	BlackBG  ColorBG = ";40"
	RedBG    ColorBG = ";41" // 红色
	GreenBG  ColorBG = ";42" // 绿色
	YellowBG ColorBG = ";43" // 黄色
	BlueBG   ColorBG = ";44" // 蓝色
	PurpleBG ColorBG = ";45" // 紫色
	CyanBG   ColorBG = ";46" // 青色
	WhiteBG  ColorBG = ";47" // 白色
)

/*
*
获取带有颜色的字符串 ，windows 不能用
*/
func ColorfulString(str string, fontStyle FontStyle, fontColor ColorFont, backgroudColor ...ColorBG) string {
	if len(backgroudColor) > 0 {
		return fmt.Sprintf("\033[%v%s%sm%s\033[0m", fontStyle, fontColor, backgroudColor[0], str)
	}
	return fmt.Sprintf("\033[%v%sm%s\033[0m", fontStyle, fontColor, str)
}

/*
*
显示带有颜色的字符串 ，windows 不能用
*/
func ColorfulPrint(str string, fontStyle FontStyle, fontColor ColorFont, backgroudColor ...ColorBG) {
	if len(backgroudColor) > 0 {
		fmt.Printf("\033[%v%s%sm%s\033[0m", fontStyle, fontColor, backgroudColor[0], str)
		return
	}
	fmt.Printf("\033[%v%sm%s\033[0m", fontStyle, fontColor, str)
}
