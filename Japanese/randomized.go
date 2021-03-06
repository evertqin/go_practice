package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Gojuon struct {
	Hiragana string
	Katakana string
	Roman    string
	Tip      string
}

var lst []Gojuon

func init() {
	lst = append(lst, Gojuon{Hiragana: "あ", Katakana: "ア", Roman: "a", Tip: "平假名あ很像汉字的“女”字，像女孩子有个大肚子。片假名ア很像汉字的“了”字，读音类似“啊”。一句话记忆：“女”孩子胖“了”一圈，惨叫了一声“啊”！"})
	lst = append(lst, Gojuon{Hiragana: "い", Katakana: "イ", Roman: "i", Tip: "平假名い和片假名イ拼在一起就会组成一个汉字“以”，读音也是“以”。"})
	lst = append(lst, Gojuon{Hiragana: "う", Katakana: "ウ", Roman: "u", Tip: "平假名う很像字母W横过来，而W在拼音中读音类似“屋”。片假名ウ很像一个偏旁：“家”字上面的宝盖头，有家就会有房“屋”。读音也类似“屋”。"})
	lst = append(lst, Gojuon{Hiragana: "え", Katakana: "エ", Roman: "e", Tip: "平假名え很像汉字的“元”字。片假名エ很像汉字的“工”字。读音是e，但为了记忆 有点点像“哀”。一句话记忆：“元”朝的“工”人很悲“哀（e）”。"})
	lst = append(lst, Gojuon{Hiragana: "お", Katakana: "オ", Roman: "o", Tip: "平假名お有些像写的很潦草的汉字的“术”字。片假名オ根本就是汉字的“才”字。读音是o但是为了方便记 想想“傲”。一句话记忆：有学“术”的“才”子都很“傲（o）”气。"})

	lst = append(lst, Gojuon{Hiragana: "か", Katakana: "カ", Roman: "ka", Tip: "平假名か和片假名カ都和汉字“力”很像，区别只是平假名多了右上角一点。读音类似“卡”。一句话记忆：因为被“卡”住了，所以要多用“一点”“力”哦！"})
	lst = append(lst, Gojuon{Hiragana: "き", Katakana: "キ", Roman: "ki", Tip: "琵琶的简化是不是很像这个符号？（好了我知道画的太渣）平假名き就可以想象成一个琵琶。片假名キ可以想象成琵琶上面的部分，而这个部分是干嘛用的呢？调音啊～读音有点像英文里面的“KEY”一句话记忆：琵琶（き）的上半部分（キ）是用来调“KEY”的～"})
	lst = append(lst, Gojuon{Hiragana: "く", Katakana: "ク", Roman: "ku", Tip: "平假名く很像平时用的“＜”（小于）符号。片假名ク有点像我们用的“＞”（大于）符号，但是左上角多了一点。读音类似“哭”。一句话记忆：小时候“小于”别人，长大后终于“大于”别人“一点”，感动的要“哭”了！（对就是一个励志故事哈哈哈）"})
	lst = append(lst, Gojuon{Hiragana: "け", Katakana: "ケ", Roman: "ke", Tip: "平假名け类似汉字“汁”字。片假名ケ类似汉字“个”字。读音是ke，但为了记忆 有点点像“开”。一句话记忆：“开”了一“个”豆“汁”店～"})
	lst = append(lst, Gojuon{Hiragana: "こ", Katakana: "コ", Roman: "ko", Tip: "平假名こ可以想象成上、下两条鱼摆摆，片假名コ可以想象成一个鱼篓，倾斜在水里的样子（考验想象力了，毕竟我不想画图了实在太渣），读音是ko，但为了记忆 有点点像“烤”， 一句话记忆：“两条鱼”被“鱼篓”抓住了，只能被做成“烤”鱼啦！好想吃！"})

	lst = append(lst, Gojuon{Hiragana: "さ", Katakana: "サ", Roman: "sa", Tip: "平假名さ可以想象成一个人跪着（又一次考验想象力）。片假名サ你们一定认识！像不像操！哦不，是“艹”！读音类似“傻”一句话记忆：那个人“跪”在“艹”上面，你说他是不是“傻”？"})
	lst = append(lst, Gojuon{Hiragana: "し", Katakana: "シ", Roman: "shi", Tip: "平假名し可以想象成一根弯弯的“吸管”。片假名シ看起来就是“三点水”。读音类似“吸”。一句话记忆：用“吸管”喝水，“吸”出来“三点水”。"})
	lst = append(lst, Gojuon{Hiragana: "す", Katakana: "ス", Roman: "su", Tip: "平假名す可以想象成一个人在上吊（又一次考验想象力）。片假名ス类似汉字“又”，但字是比“又”字少了一点。读音类似“死”。一句话记忆：那个人“上吊”，“差一点”“又”“死”了。"})
	lst = append(lst, Gojuon{Hiragana: "せ", Katakana: "セ", Roman: "se", Tip: "平假名せ类似汉字“世”但是少一竖，就好像“世”字是大马路而せ就是单行道了。片假名セ和平假名长得很像基本记住一个就能记住另外一个。读音是se，但为了记忆 有点点像“塞”。一句话记忆：单行道的“世”界被“塞”的满满的。"})
	lst = append(lst, Gojuon{Hiragana: "そ", Katakana: "ソ", Roman: "so", Tip: "平假名そ有点像写的比较潦草的“艺”字。片假名ソ有点像拍照时候的剪刀手哦。读音是so，但为了记忆 有点点像“少”。一句话记忆：现在拍“艺”术照，还比“V”的人真的太“少”了啦!"})

	lst = append(lst, Gojuon{Hiragana: "た", Katakana: "タ", Roman: "ta", Tip: "平假名た 拆分成汉字的“十”和“二”。片假名タ类似汉字“夕”。读音是ta，但为了记忆，引申到 “谭”。一句话记忆：半夜“十二”点还有“夕”阳，简直是天方夜“谭”。"})
	lst = append(lst, Gojuon{Hiragana: "ち", Katakana: "チ", Roman: "chi", Tip: "平假名ち类似数字“5”。片假名チ 请想到汉字“我”，像不像这个我字刚刚写了个开头（一定写写看啊你！）读音类似“气”。一句话记忆：这小孩儿！都“5”岁了还写不会一个完整的“我”字，你说“气”人不气人！"})
	lst = append(lst, Gojuon{Hiragana: "つ", Katakana: "ツ", Roman: "tsu", Tip: "平假名つ想象成一个倒了的瓶子。片假名ツ想象成溅起的三点水花。读音类似“呲”。一句话记忆：瓶子倒了，溅起三点水花，发出“呲”的声音。（这里和しシ很像，别混了，可以对比记忆。）"})
	lst = append(lst, Gojuon{Hiragana: "て", Katakana: "テ", Roman: "te", Tip: "平假名て类似汉字“乙”。片假名テ 类似汉字“街”的偏旁“彳”。读音是te，但为了记忆 引申到 “太”。一句话记忆：路人“乙”在“街”边上撞到了一个老“太太”。"})
	lst = append(lst, Gojuon{Hiragana: "と", Katakana: "ト", Roman: "to", Tip: "平假名と很像一个人正伸向衣服兜儿的手。片假名ト类似汉字“卜”。读音是to，但为了记忆 引申到 “掏”。一句话记忆：小偷在街上“掏”兜儿，结果掏到一根萝“卜”。"})

	lst = append(lst, Gojuon{Hiragana: "な", Katakana: "ナ", Roman: "na", Tip: "平假名な 拆分成汉字的“十”和比较潦草的数字“3”。片假名ナ类似汉字“十”。读音类似“那”。一句话记忆：把“十三”算成了“十”，也太木“讷”了（木讷读音其实是木ne，但是这里就当做木na吧）。"})
	lst = append(lst, Gojuon{Hiragana: "に", Katakana: "ニ", Roman: "ni", Tip: "平假名に类似汉字“仁”。片假名ニ类似汉字“二”，二的读音引申到“爱”上。读音类似“尼”。一句话记忆：“仁”者“二（爱）”人，是孟子观点，说到孟子怎能忘记孔仲“尼”（孔子）。"})
	lst = append(lst, Gojuon{Hiragana: "ぬ", Katakana: "ヌ", Roman: "nu", Tip: "平假名ぬ类似汉字“奴”。片假名ヌ类似汉字“又”。读音类似“奴”。一句话记忆：“奴”隶制度“又”出现了！（这里ヌ和前面的ス长得很像，可以对比记忆)"})
	lst = append(lst, Gojuon{Hiragana: "ね", Katakana: "ネ", Roman: "ne", Tip: "平假名ね类似汉字“权”。片假名ネ类似汉字“福”的偏旁“礻”。读音是ne，但为了记忆 引申到 “捺”。一句话记忆：有了“权”利就是“福”利，所以很多人都按“捺”不住啊。"})
	lst = append(lst, Gojuon{Hiragana: "の", Katakana: "ノ", Roman: "no", Tip: "平假名の类似数字符号“①”。片假名ノ类似数字“1”。读音是no，但为了记忆 引申到 “闹”。一句话：一个“①”，一个“1”，都争着抢当第一，“闹”个不停。"})

	lst = append(lst, Gojuon{Hiragana: "は", Katakana: "ハ", Roman: "ha", Tip: "平假名は有点像潦草的汉字“仗”。片假名ハ类似汉字“八”。读音类似“哈”。一句话记忆：居然想和“八”路军打“仗”哈哈哈哈哈。"})
	lst = append(lst, Gojuon{Hiragana: "ひ", Katakana: "ヒ", Roman: "hi", Tip: "平假名ひ请把它想象成一个装匕首的套子。片假名ヒ类似汉字“匕”。读音是hi，但为了记忆 引申到 “嘿”。一句话记忆：从剑套ひ里拿出“匕”首，“嘿”的一声！(我觉得这里把读音想的和另外一个假名容易记串，如果有更好的办法告诉我昂！）"})
	lst = append(lst, Gojuon{Hiragana: "ふ", Katakana: "フ", Roman: "hu", Tip: "平假名ふ有点像潦草的汉字“小”。片假名フ像汉字“不”的开头。读音类似“服”。一句话记忆：“小”时候和长大以后都坚持说“不”，真是“服”了他了。"})
	lst = append(lst, Gojuon{Hiragana: "へ", Katakana: "ヘ", Roman: "he", Tip: "平假名和片假名很像，就是感觉一大一小而已，像一对弯弯的眼睛。读音类似“黑”。一句话记忆：太“happy”了，笑的眼睛都弯弯的～"})
	lst = append(lst, Gojuon{Hiragana: "ほ", Katakana: "ホ", Roman: "ho", Tip: "平假名ほ 拆分成数字的“1”和汉字“天”。片假名类似汉字“木”。读音是ho，但为了记忆 引申到 “嚎”。一句话记忆：“嚎”了“1”整“天”，人都变“木”了！"})

	lst = append(lst, Gojuon{Hiragana: "ま", Katakana: "マ", Roman: "ma", Tip: "平假名ま像是汉字“天”但是上面出头了一笔。片假名类似汉字“捅”右上角的部分。读音类似“妈”。一句话记忆：“妈”呀！“天”被“捅”破了！ "})
	lst = append(lst, Gojuon{Hiragana: "み", Katakana: "ミ", Roman: "mi", Tip: "平假名み请把它想象成一个在射箭的动作， 片假名ミ就想象成射出来的三只箭好了～ 读音类似“眯”。一句话记忆：丘比特“眯”起眼睛射了三箭～ （背到这里已经有三个假名是三点这样，很类似的 还有印象不） "})
	lst = append(lst, Gojuon{Hiragana: "む", Katakana: "ム", Roman: "mu", Tip: "平假名む有点点像潦草的汉字“生”。片假名ム像是汉字“牟”上面的部分。读音类似“牟”（很多人把释迦mu尼读成释迦mou尼，不对哦，这里读音是MU！） 一句话记忆：错“生”了释迦“牟”尼。"})
	lst = append(lst, Gojuon{Hiragana: "め", Katakana: "メ", Roman: "me", Tip: "平假名め类似汉字“女”。片假名メ类似汉字“丫”。读音是me，但为了记忆 引申到 “卖”。一句话记忆：好好的“女”孩子被当成“丫”头“卖”了…… "})
	lst = append(lst, Gojuon{Hiragana: "も", Katakana: "モ", Roman: "mo", Tip: "平假名和片假名都很像汉字“毛”。读音是mo但是也引申想到“毛”。一句话记忆：像“毛”。（这个假名的记忆方法给的太简单，但是我觉得这个假名蛮好记的，就是不用什么办法都容易记住的那种。如果大家有更好的办法告诉我，我补充上来。）"})

	lst = append(lst, Gojuon{Hiragana: "や", Katakana: "ヤ", Roman: "ya", Tip: "平假名和片假名都很像汉字“也”。读音是ya但是也引申想到“也”。一句话记忆：像“也”。"})
	lst = append(lst, Gojuon{Hiragana: "ゆ", Katakana: "ユ", Roman: "yu", Tip: "这个太考想象力。平假名ゆ请想象成太极的云手（你们按笔画顺序打打试试看，还是挺像的）。片假名ユ请想象成跳芭蕾舞的倒下了～（难以想象的两条腿……） 读音类似“有”。 一句话记忆：太极大战芭蕾，芭蕾被打倒。太“有”看头了！"})
	lst = append(lst, Gojuon{Hiragana: "よ", Katakana: "ヨ", Roman: "yo", Tip: "平假名よ很像一把钥匙的形状啊。片假名ヨ很像梳子的梳齿部分哦。读音是yo但是也引申想到“妖”。一句话记忆：把“钥匙”当做“梳子”使，这实在是太“妖”气了啦！"})

	lst = append(lst, Gojuon{Hiragana: "ら", Katakana: "ラ", Roman: "ra", Tip: "平假名ら想象成一个人在拉臭臭坐在马桶上。片假名ラ想象成马桶（带盖子的那种哦）。读音类似“拉”。一句话记忆：“拉”完就盖上马桶盖。"})
	lst = append(lst, Gojuon{Hiragana: "り", Katakana: "リ", Roman: "ri", Tip: "平假名和片假名都很类似汉字的立刀旁“刂”。读音类似“立”。一句话记忆：像立刀旁“刂”。"})
	lst = append(lst, Gojuon{Hiragana: "る", Katakana: "ル", Roman: "ru", Tip: "平假名る类似汉字“歹”。平假名ル类似汉字“儿”。读音类似“路”。一句话记忆：“路”上遇见“歹”徒，怒斥歹徒是“儿”子！"})
	lst = append(lst, Gojuon{Hiragana: "れ", Katakana: "レ", Roman: "re", Tip: "平假名れ类似潦草的汉字“水”。片假名レ有点像英文字母“V”。读音是le但是为了记忆引申想到“来”。一句话记忆：“水”终于“来”了，比出了象征胜利的“V”手势！"})
	lst = append(lst, Gojuon{Hiragana: "ろ", Katakana: "ロ", Roman: "ro", Tip: "平假名ろ类似汉字“万”。平假名ロ类似汉字“口”。读音是lo但是为了记忆引申想到“老”。一句话记忆：生了一“万”“口”，人都“老”了。"})

	lst = append(lst, Gojuon{Hiragana: "わ", Katakana: "ワ", Roman: "wa", Tip: "平假名わ类似潦草的汉字“水”（和れ写起来有点像哦）。片假名ワ想象成一个水箱，但是底儿没啦～读音类似“哇”。一句话记忆：“水”箱的底儿没了，水“哇”的全漏光了！"})
	lst = append(lst, Gojuon{Hiragana: "を", Katakana: "ヲ", Roman: "wo", Tip: "又一次考验想象力。平假名を可以拆成上下两部分，有点像动车头“c”上趴着一个人“大”。片假名ヲ说是现代日语基本不用，不用刻意记。读音是o但是为了记忆引申想到“嗷”。一句话记忆：动车头上趴着一个人，被带着跑吓得“嗷”嗷直叫。"})

	lst = append(lst, Gojuon{Hiragana: "ん", Katakana: "ン", Roman: "n", Tip: "平假名ん很像那种手摇的铃铛，不知道啥样的去看爸哪第三季夏克立家那个亲吻铃铛。片假名ン像是字母“V”倒过来。读音类似“摁”。一句话记忆：手摇铃“倒”着“摁”一下就胜利“V”了！"})
}

func genRandom() Gojuon {
	rn := rand.Intn(len(lst))
	return lst[rn]
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	for {
		gojuon := genRandom()

		var ran int
		var display, backup string

		if ran, display, backup = rand.Intn(2), gojuon.Katakana, gojuon.Hiragana; ran > 0 {
			display = gojuon.Hiragana
			backup = gojuon.Katakana
		}

		fmt.Println(display)
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
		fmt.Println(backup, gojuon.Roman, gojuon.Tip)
	}

}
