package vars

var TianGan = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
var DiZhi = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

var TianGanWuXing = map[string]string{
	"甲": "木",
	"乙": "木",
	"丙": "火",
	"戊": "土",
	"已": "土",
	"庚": "金",
	"辛": "金",
	"壬": "水",
	"癸": "水",
}

// Ke 起局的四课
// todo 天干在旬中名字
type Ke struct {
	KeNum            int    //第几课
	Shang            string //日上或者辰上
	Xia              string //干或者支
	ShangXiaRelation string //上克下，下贼上 无克
}

type SiKe struct {
	Ke1 Ke //干阳
	Ke2 Ke //干阳
	Ke3 Ke //支阳
	Ke4 Ke //支阴

}
type BiYong struct {
	Shang    string //那个元素
	IsTongRi string //是否同日
}
type SheHai struct {
	Shang           string //那个元素
	SheHaiNum       int    //涉害次数
	SheHaiWeiZhi    string //涉害位置
	YinYangIsTongRi string //是否同日
}

// JiuZongMem 九宗门
type JiuZongMem struct {
	KeNeiShengKeList []string //一课下贼上 二课下贼上 四课下贼上 四课下贼上 一课上克下 二课上克下 三课上克下 四课上克下
	KeDesc           string   //一克 多克 无克
	BiYongList       []BiYong
	IsBiYong         bool              //是否必用
	SheHaiMap        map[string]SheHai //涉害情况
	KeWaiShengKeList []string          //二课遥克日 三课遥克日 四课遥克日 日遥克二课 日遥克三课 日遥克四课
	KeWaiIsKe        string            //四课是否遥克，默认无
	IsSheHai         string            //是否涉害
	SheHaiFa         string            //涉害用啥法
	YaoKe            string            //日遥克他人 他人遥克日
	YaoKeDesc        string            //日遥克他人 他人遥克日
	IsMaoXing        string            //是否昴星
	MaoXingDesc      string
	IsBieZe          string //是否别责
	BieZeDesc        string //
	IsBaZhuan        string //是否八专
	BaZhuanDesc      string
	IsFuYin          string //默认否
	FuYinDesc        string
	IsFanYin         string //默认否
	FaYinDesc        string
	Fa               string //最终使用的九宗门的方法
	FaDesc           string //最终使用的九宗门的方法
}
