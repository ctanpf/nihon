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
