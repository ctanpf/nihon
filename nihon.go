package nihon

const UpperCaseStart = 0x41
const UpperCaseEnd = 0x5A
const HiraganaStart = 0x3041
const HiraganaEnd = 0x3096
const KatakanaStart = 0x30A1
const KatakanaEnd = 0x30FA

var RomajiToJapaneseMap = map[string]string{
	"あ": "a",
	"い": "i",
	"う": "u",
	"え": "e",
	"お": "o",

	"か": "ka",
	"き": "ki",
	"く": "ku",
	"け": "ke",
	"こ": "ko",

	"さ": "sa",
	"し": "shi",
	"す": "su",
	"せ": "se",
	"そ": "so",

	"た": "ta",
	"ち": "chi",
	"つ": "tsu",
	"て": "te",
	"と": "to",

	"な": "na",
	"に": "ni",
	"ぬ": "nu",
	"ね": "ne",
	"の": "no",

	"は": "ha",
	"ひ": "hi",
	"ふ": "fu",
	"へ": "he",
	"ほ": "ho",

	"ま": "ma",
	"み": "mi",
	"む": "mu",
	"め": "me",
	"も": "mo",

	"ら": "ra",
	"り": "ri",
	"る": "ru",
	"れ": "re",
	"ろ": "ro",

	"が": "ga",
	"ぎ": "gi",
	"ぐ": "gu",
	"げ": "ge",
	"ご": "go",

	"ざ": "za",
	"じ": "ji",
	"ず": "zu",
	"ぜ": "ze",
	"ぞ": "zo",

	"だ": "da",
	"ぢ": "ji",
	"づ": "zu",
	"で": "de",
	"ど": "do",

	"ば": "ba",
	"び": "bi",
	"ぶ": "bu",
	"べ": "be",
	"ぼ": "bo",

	"ぱ": "pa",
	"ぴ": "pi",
	"ぷ": "pu",
	"ぺ": "pe",
	"ぽ": "po",

	"や": "ya",
	"ゆ": "yu",
	"よ": "yo",

	"わ": "wa",
	"を": "wo",
	"ん": "n",

	"きゃ": "kya",
	"きゅ": "kyu",
	"きょ": "kyo",

	"ぎゃ": "gya",
	"ぎゅ": "gyu",
	"ぎょ": "gyo",

	"しゃ": "sha",
	"しゅ": "shu",
	"しょ": "sho",

	"じゃ": "ja",
	"じゅ": "ju",
	"じょ": "jo",

	"ちゃ": "cha",
	"ちゅ": "chu",
	"ちょ": "cho",

	"にゃ": "nya",
	"にゅ": "nyu",
	"にょ": "nyo",

	"ひゃ": "hya",
	"ひゅ": "hyu",
	"ひょ": "hyo",

	"びゃ": "bya",
	"びゅ": "byu",
	"びょ": "byo",

	"ぴゃ": "pya",
	"ぴゅ": "pyu",
	"ぴょ": "pyo",

	"みゃ": "mya",
	"みゅ": "myu",
	"みょ": "myo",

	"りゃ": "rya",
	"りゅ": "ryu",
	"りょ": "ryo",
}

var KatakanaToJapaneseMap = map[string]string{
	"ア": "a",
	"イ": "i",
	"ウ": "u",
	"エ": "e",
	"オ": "o",

	"カ": "ka",
	"キ": "ki",
	"ク": "ku",
	"ケ": "ke",
	"コ": "ko",

	"サ": "sa",
	"シ": "shi",
	"ス": "su",
	"セ": "se",
	"ソ": "so",

	"タ": "ta",
	"チ": "chi",
	"ツ": "tsu",
	"テ": "te",
	"ト": "to",

	"ナ": "na",
	"ニ": "ni",
	"ヌ": "nu",
	"ネ": "ne",
	"ノ": "no",

	"ハ": "ha",
	"ヒ": "hi",
	"フ": "fu",
	"へ": "he",
	"ホ": "ho",

	"マ": "ma",
	"ミ": "mi",
	"ム": "mu",
	"メ": "me",
	"モ": "mo",

	"ラ": "ra",
	"リ": "ri",
	"ル": "ru",
	"レ": "re",
	"ロ": "ro",

	"ガ": "ga",
	"ギ": "gi",
	"グ": "gu",
	"ゲ": "ge",
	"ゴ": "go",

	"ザ": "za",
	"ジ": "ji",
	"ズ": "zu",
	"ゼ": "ze",
	"ゾ": "zo",

	"ダ": "da",
	"ヂ": "ji",
	"ヅ": "zu",
	"デ": "de",
	"ド": "do",

	"バ": "ba",
	"ビ": "bi",
	"ブ": "bu",
	"ベ": "be",
	"ボ": "bo",

	"パ": "pa",
	"ピ": "pi",
	"プ": "pu",
	"ペ": "pe",
	"ポ": "po",

	"ヤ": "ya",
	"ユ": "yu",
	"ヨ": "yo",

	"ワ": "wa",
	"ヲ": "wo",
	"ン": "n",

	"キャ": "kya",
	"キュ": "kyu",
	"キョ": "kyo",

	"ギャ": "gya",
	"ギュ": "gyu",
	"ギョ": "gyo",

	"シャ": "sha",
	"シュ": "shu",
	"ショ": "sho",

	"ジャ": "ja",
	"ジュ": "ju",
	"ジョ": "jo",

	"チャ": "cha",
	"チュ": "chu",
	"チョ": "cho",

	"ニャ": "nya",
	"ニュ": "nyu",
	"ニョ": "nyo",

	"ヒャ": "hya",
	"ヒュ": "hyu",
	"ヒョ": "hyo",

	"ビャ": "bya",
	"ビュ": "byu",
	"ビョ": "byo",

	"ピャ": "pya",
	"ピュ": "pyu",
	"ピョ": "pyo",

	"ミャ": "mya",
	"ミュ": "myu",
	"ミョ": "myo",

	"リャ": "rya",
	"リュ": "ryu",
	"リョ": "ryo",
}

func AllHiragana() map[string]string {
	return merge(RomajiToJapaneseMap, KatakanaToJapaneseMap)
}

func merge(a map[string]string, b map[string]string) map[string]string {
	res := make(map[string]string)
	for k, v := range a {
		res[k] = v
	}
	for k, v := range b {
		res[k] = v
	}
	return res
}
