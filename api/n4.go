package handler

import "net/http"

const n4 = `
{
  "data": [
    {
      "kanji": "会う",
      "kana": "あう",
      "en": "to meet"
    },
    {
      "kanji": "合う",
      "kana": "あう",
      "en": "to fit, to suit"
    },
    {
      "kanji": "上がる",
      "kana": "あがる",
      "en": "to go up"
    },
    {
      "kanji": "開く",
      "kana": "あく",
      "en": "to open"
    },
    {
      "kanji": "空く",
      "kana": "あく",
      "en": "empty, with space"
    },
    {
      "kanji": "開ける",
      "kana": "あける",
      "en": "to open"
    },
    {
      "kanji": "あげる",
      "kana": "あげる",
      "en": "to give"
    },
    {
      "kanji": "上げる",
      "kana": "あげる",
      "en": "to rise"
    },
    {
      "kanji": "遊ぶ",
      "kana": "あそぶ",
      "en": "to play"
    },
    {
      "kanji": "浴びる",
      "kana": "あびる",
      "en": "to take a shower"
    },
    {
      "kanji": "洗う",
      "kana": "あらう",
      "en": "to wash"
    },
    {
      "kanji": "有る",
      "kana": "ある",
      "en": "to be, to exist"
    },
    {
      "kanji": "ある",
      "kana": "ある",
      "en": "to possess"
    },
    {
      "kanji": "歩く",
      "kana": "あるく",
      "en": "to walk"
    },
    {
      "kanji": "言う",
      "kana": "いう",
      "en": "to say, to tell"
    },
    {
      "kanji": "行く",
      "kana": "いく",
      "en": "to go"
    },
    {
      "kanji": "いる",
      "kana": "いる",
      "en": "need, must haven be required"
    },
    {
      "kanji": "いる",
      "kana": "いる",
      "en": "to exist"
    },
    {
      "kanji": "入れる",
      "kana": "いれる",
      "en": "to insert, to put in"
    },
    {
      "kanji": "歌う",
      "kana": "うたう",
      "en": "to sing"
    },
    {
      "kanji": "生まれる",
      "kana": "うまれる",
      "en": "to be born"
    },
    {
      "kanji": "売る",
      "kana": "うる",
      "en": "to sell"
    },
    {
      "kanji": "起きる",
      "kana": "おきる",
      "en": "to get up, to stand up"
    },
    {
      "kanji": "置く",
      "kana": "おく",
      "en": "to put"
    },
    {
      "kanji": "送る",
      "kana": "おくる",
      "en": "to send"
    },
    {
      "kanji": "押す",
      "kana": "おす",
      "en": "to push"
    },
    {
      "kanji": "覚える",
      "kana": "おぼえる",
      "en": "to memorize, to remember"
    },
    {
      "kanji": "泳ぐ",
      "kana": "およぐ",
      "en": "to swim"
    },
    {
      "kanji": "降りる",
      "kana": "おりる",
      "en": "to get off"
    },
    {
      "kanji": "終わる",
      "kana": "おわる",
      "en": "to end"
    },
    {
      "kanji": "買う",
      "kana": "かう",
      "en": "to buy"
    },
    {
      "kanji": "返す",
      "kana": "かえす",
      "en": "to return an object"
    },
    {
      "kanji": "帰る",
      "kana": "かえる",
      "en": "to return home"
    },
    {
      "kanji": "掛かる",
      "kana": "かかる",
      "en": "to take time or money"
    },
    {
      "kanji": "書く",
      "kana": "かく",
      "en": "to write"
    },
    {
      "kanji": "かける",
      "kana": "かける",
      "en": "to wear"
    },
    {
      "kanji": "かける",
      "kana": "かける",
      "en": "to make a phone call"
    },
    {
      "kanji": "貸す",
      "kana": "かす",
      "en": "to lend"
    },
    {
      "kanji": "冠る",
      "kana": "かぶる",
      "en": "to put on a hat"
    },
    {
      "kanji": "借りる",
      "kana": "かりる",
      "en": "to borrow"
    },
    {
      "kanji": "聞こえる",
      "kana": "きこえる",
      "en": "hear, can hear, be audible"
    },
    {
      "kanji": "消える",
      "kana": "きえる",
      "en": "to go out, to vanish"
    },
    {
      "kanji": "決まる",
      "kana": "きまる",
      "en": "to be decided"
    },
    {
      "kanji": "決める",
      "kana": "きめる",
      "en": "to decide"
    },
    {
      "kanji": "聞く",
      "kana": "きく",
      "en": "to listen"
    },
    {
      "kanji": "切る",
      "kana": "きる",
      "en": "to cut"
    },
    {
      "kanji": "着る",
      "kana": "きる",
      "en": "to wear, to put on"
    },
    {
      "kanji": "来る",
      "kana": "くる",
      "en": "to come"
    },
    {
      "kanji": "くださる",
      "kana": "くださる",
      "en": "to give (to me) respectful"
    },
    {
      "kanji": "比べる",
      "kana": "くらべる",
      "en": "to compare"
    },
    {
      "kanji": "暮れる",
      "kana": "くれる",
      "en": "to get dark"
    },
    {
      "kanji": "くれる",
      "kana": "くれる",
      "en": "to give (to me) simple form"
    },
    {
      "kanji": "消す",
      "kana": "けす",
      "en": "to turn off, to switch off"
    },
    {
      "kanji": "答える",
      "kana": "こたえる",
      "en": "to answer"
    },
    {
      "kanji": "困る",
      "kana": "こまる",
      "en": "to be in trouble"
    },
    {
      "kanji": "込む",
      "kana": "こむ",
      "en": "to be crowded"
    },
    {
      "kanji": "壊す",
      "kana": "こわす",
      "en": "to break"
    },
    {
      "kanji": "壊れる",
      "kana": "こわれる",
      "en": "to be broken"
    },
    {
      "kanji": "咲く",
      "kana": "さく",
      "en": "to blossom"
    },
    {
      "kanji": "探す",
      "kana": "さがす",
      "en": "to look for, to search"
    },
    {
      "kanji": "さす",
      "kana": "さす",
      "en": "to open an umbrella"
    },
    {
      "kanji": "下がる",
      "kana": "さがる",
      "en": "to get down"
    },
    {
      "kanji": "下げる",
      "kana": "さげる",
      "en": "to lower"
    },
    {
      "kanji": "差し上げる",
      "kana": "さしあげる",
      "en": "to give (respectful)"
    },
    {
      "kanji": "騒ぐ",
      "kana": "さわぐ",
      "en": "to make noise, to be excited"
    },
    {
      "kanji": "触る",
      "kana": "さわる",
      "en": "to touch"
    },
    {
      "kanji": "叱る",
      "kana": "しかる",
      "en": "to scold"
    },
    {
      "kanji": "知らせる",
      "kana": "しらせる",
      "en": "to notify"
    },
    {
      "kanji": "調べる",
      "kana": "しらべる",
      "en": "to investigate"
    },
    {
      "kanji": "死ぬ",
      "kana": "しぬ",
      "en": "to die, to pass away"
    },
    {
      "kanji": "閉まる",
      "kana": "しまる",
      "en": "to close"
    },
    {
      "kanji": "閉める",
      "kana": "しめる",
      "en": "to close"
    },
    {
      "kanji": "締める",
      "kana": "しめる",
      "en": "to fasten a belt"
    },
    {
      "kanji": "知る",
      "kana": "しる",
      "en": "to know"
    },
    {
      "kanji": "知らせる",
      "kana": "しらせる",
      "en": "to notify"
    },
    {
      "kanji": "調べる",
      "kana": "しらべる",
      "en": "to investigate"
    },
    {
      "kanji": "過ぎる",
      "kana": "すぎる",
      "en": "to exceed"
    },
    {
      "kanji": "進む",
      "kana": "すすむ",
      "en": "to make progress"
    },
    {
      "kanji": "捨てる",
      "kana": "すてる",
      "en": "to throw away"
    },
    {
      "kanji": "滑る",
      "kana": "すべる",
      "en": "to slide, to slip"
    },
    {
      "kanji": "吸う",
      "kana": "すう",
      "en": "to breath, to smoke"
    },
    {
      "kanji": "住む",
      "kana": "すむ",
      "en": "to live, to reside somewhere"
    },
    {
      "kanji": "済む",
      "kana": "すむ",
      "en": "to finish"
    },
    {
      "kanji": "する",
      "kana": "する",
      "en": "to do"
    },
    {
      "kanji": "座る",
      "kana": "すわる",
      "en": "to sit"
    },
    {
      "kanji": "育てる",
      "kana": "そだてる",
      "en": "to rear, to bring up"
    },
    {
      "kanji": "出す",
      "kana": "だす",
      "en": "to take out, hand in"
    },
    {
      "kanji": "立つ",
      "kana": "たつ",
      "en": "to stand"
    },
    {
      "kanji": "倒れる",
      "kana": "たおれる",
      "en": "to break down"
    },
    {
      "kanji": "足す",
      "kana": "たす",
      "en": "to add a number"
    },
    {
      "kanji": "訪ねる",
      "kana": "たずねる",
      "en": "to visit"
    },
    {
      "kanji": "尋ねる",
      "kana": "たずねる",
      "en": "to ask"
    },
    {
      "kanji": "建てる",
      "kana": "たてる",
      "en": "to build"
    },
    {
      "kanji": "立てる",
      "kana": "たてる",
      "en": "to stand up, to put up"
    },
    {
      "kanji": "楽しむ",
      "kana": "たのしむ",
      "en": "to enjoy"
    },
    {
      "kanji": "足りる",
      "kana": "たりる",
      "en": "to be sufficient"
    },
    {
      "kanji": "頼む",
      "kana": "たのむ",
      "en": "to ask, to request"
    },
    {
      "kanji": "食べる",
      "kana": "たべる",
      "en": "to eat"
    },
    {
      "kanji": "違う",
      "kana": "ちがう",
      "en": "to be different"
    },
    {
      "kanji": "使う",
      "kana": "つかう",
      "en": "to use"
    },
    {
      "kanji": "疲れる",
      "kana": "つかれる",
      "en": "to get tired"
    },
    {
      "kanji": "着く",
      "kana": "つく",
      "en": "to arrive"
    },
    {
      "kanji": "作る",
      "kana": "つくる",
      "en": "to make, to produce"
    },
    {
      "kanji": "捕まえる",
      "kana": "つかまえる",
      "en": "to seize"
    },
    {
      "kanji": "付く",
      "kana": "つく",
      "en": "to be attached"
    },
    {
      "kanji": "漬ける",
      "kana": "つける",
      "en": "to soak, to pickle"
    },
    {
      "kanji": "伝える",
      "kana": "つたえる",
      "en": "to report"
    },
    {
      "kanji": "続く",
      "kana": "つづく",
      "en": "to continue, to last"
    },
    {
      "kanji": "続ける",
      "kana": "つづける",
      "en": "to continue, go on"
    },
    {
      "kanji": "包む",
      "kana": "つつむ",
      "en": "to wrap"
    },
    {
      "kanji": "釣る",
      "kana": "つる",
      "en": "to fish"
    },
    {
      "kanji": "連れる",
      "kana": "つれる",
      "en": "to lead"
    },
    {
      "kanji": "点ける",
      "kana": "つける",
      "en": "to turn on"
    },
    {
      "kanji": "勤める",
      "kana": "つとめる",
      "en": "to work for someone"
    },
    {
      "kanji": "出かける",
      "kana": "でかける",
      "en": "to go out"
    },
    {
      "kanji": "出来る",
      "kana": "できる",
      "en": "can do"
    },
    {
      "kanji": "出る",
      "kana": "でる",
      "en": "to leave"
    },
    {
      "kanji": "手伝う",
      "kana": "てつだう",
      "en": "to assist"
    },
    {
      "kanji": "飛ぶ",
      "kana": "とぶ",
      "en": "to fly"
    },
    {
      "kanji": "止まる",
      "kana": "とまる",
      "en": "to stop"
    },
    {
      "kanji": "取る",
      "kana": "とる",
      "en": "to take"
    },
    {
      "kanji": "通る",
      "kana": "とおる",
      "en": "to go through"
    },
    {
      "kanji": "届ける",
      "kana": "とどける",
      "en": "to reach"
    },
    {
      "kanji": "泊まる",
      "kana": "とまる",
      "en": "to lodge at"
    },
    {
      "kanji": "止める",
      "kana": "とめる",
      "en": "to stop"
    },
    {
      "kanji": "取り替える",
      "kana": "とりかえる",
      "en": "to exchange"
    },
    {
      "kanji": "撮る",
      "kana": "とる",
      "en": "to take a photo"
    },
    {
      "kanji": "直す",
      "kana": "なおす",
      "en": "to fix, to repair"
    },
    {
      "kanji": "直る",
      "kana": "なおる",
      "en": "to be fixed, repaired"
    },
    {
      "kanji": "治る",
      "kana": "なおる",
      "en": "to be cured, to heal"
    },
    {
      "kanji": "泣く",
      "kana": "なく",
      "en": "to cry"
    },
    {
      "kanji": "無くなる",
      "kana": "なくなる",
      "en": "to disappear, to get lost"
    },
    {
      "kanji": "亡くなる",
      "kana": "なくなる",
      "en": "to die"
    },
    {
      "kanji": "投げる",
      "kana": "なげる",
      "en": "to throw away"
    },
    {
      "kanji": "なさる",
      "kana": "なさる",
      "en": "to do (respectful)"
    },
    {
      "kanji": "慣れる",
      "kana": "なれる",
      "en": "to grow accustomed to"
    },
    {
      "kanji": "並ぶ",
      "kana": "ならぶ",
      "en": "to form a line"
    },
    {
      "kanji": "並べる",
      "kana": "ならべる",
      "en": "to line up"
    },
    {
      "kanji": "なる",
      "kana": "なる",
      "en": "to become"
    },
    {
      "kanji": "逃げる",
      "kana": "にげる",
      "en": "to escape"
    },
    {
      "kanji": "似る",
      "kana": "にる",
      "en": "to be similar"
    },
    {
      "kanji": "脱ぐ",
      "kana": "ぬぐ",
      "en": "to take off clothes"
    },
    {
      "kanji": "盗む",
      "kana": "ぬすむ",
      "en": "to steal"
    },
    {
      "kanji": "塗る",
      "kana": "ぬる",
      "en": "to paint, lacquer"
    },
    {
      "kanji": "濡れる",
      "kana": "ぬれる",
      "en": "to get wet"
    },
    {
      "kanji": "寝る",
      "kana": "ねる",
      "en": "to sleep"
    },
    {
      "kanji": "眠る",
      "kana": "ねむる",
      "en": "to sleep"
    },
    {
      "kanji": "登る",
      "kana": "のぼる",
      "en": "to climb up"
    },
    {
      "kanji": "飲む",
      "kana": "のむ",
      "en": "to drink"
    },
    {
      "kanji": "残る",
      "kana": "のこる",
      "en": "to remain"
    },
    {
      "kanji": "乗り換える",
      "kana": "のりかえる",
      "en": "to take off clothes"
    },
    {
      "kanji": "乗る",
      "kana": "のる",
      "en": "to take, to ride"
    },
    {
      "kanji": "入る",
      "kana": "はいる",
      "en": "to enter"
    },
    {
      "kanji": "履く",
      "kana": "はく",
      "en": "to put on shoes"
    },
    {
      "kanji": "運ぶ",
      "kana": "はこぶ",
      "en": "to transport"
    },
    {
      "kanji": "始まる",
      "kana": "はじまる",
      "en": "to begin, to start"
    },
    {
      "kanji": "始める",
      "kana": "はじめる",
      "en": "to begin"
    },
    {
      "kanji": "走る",
      "kana": "はしる",
      "en": "to run"
    },
    {
      "kanji": "働く",
      "kana": "はたらく",
      "en": "to work"
    },
    {
      "kanji": "話す",
      "kana": "はなす",
      "en": "to talk, to speak, to tell"
    },
    {
      "kanji": "張る",
      "kana": "はる",
      "en": "to put something on, to stick"
    },
    {
      "kanji": "払う",
      "kana": "はらう",
      "en": "to pay"
    },
    {
      "kanji": "晴れる",
      "kana": "はれる",
      "en": "to clear up"
    },
    {
      "kanji": "引く",
      "kana": "ひく",
      "en": "to pull"
    },
    {
      "kanji": "冷える",
      "kana": "ひえる",
      "en": "get cold, feel chilly"
    },
    {
      "kanji": "光る",
      "kana": "ひかる",
      "en": "to shine"
    },
    {
      "kanji": "引き出す",
      "kana": "ひきだす",
      "en": "to draw out"
    },
    {
      "kanji": "弾く",
      "kana": "ひく",
      "en": "to play an instrument"
    },
    {
      "kanji": "引っ越す",
      "kana": "ひっこす",
      "en": "to move (house)"
    },
    {
      "kanji": "開く",
      "kana": "ひらく",
      "en": "to open an event"
    },
    {
      "kanji": "拾う",
      "kana": "ひろう",
      "en": "topic up, to gather"
    },
    {
      "kanji": "吹く",
      "kana": "ふく",
      "en": "to blow (wind)"
    },
    {
      "kanji": "増える",
      "kana": "ふえる",
      "en": "to　increase"
    },
    {
      "kanji": "太る",
      "kana": "ふとる",
      "en": "to grow fat"
    },
    {
      "kanji": "踏む",
      "kana": "ふむ",
      "en": "to step on"
    },
    {
      "kanji": "降り出す",
      "kana": "ふりだす",
      "en": "start to rain"
    },
    {
      "kanji": "降る",
      "kana": "ふる",
      "en": "to fall (rain, snow)"
    },
    {
      "kanji": "褒める",
      "kana": "ほめる",
      "en": "to praise"
    },
    {
      "kanji": "曲がる",
      "kana": "まがる",
      "en": "to turn"
    },
    {
      "kanji": "参る",
      "kana": "まいる",
      "en": "to go, to come (humble form)"
    },
    {
      "kanji": "負ける",
      "kana": "まける",
      "en": "to lose"
    },
    {
      "kanji": "間違える",
      "kana": "まちがえる",
      "en": "to make a mistake"
    },
    {
      "kanji": "間に合う",
      "kana": "まにあう",
      "en": "to be in time for"
    },
    {
      "kanji": "回る",
      "kana": "まわる",
      "en": "to go around"
    },
    {
      "kanji": "待つ",
      "kana": "まつ",
      "en": "to wait"
    },
    {
      "kanji": "見える",
      "kana": "みえる",
      "en": "to be visible, to be seen"
    },
    {
      "kanji": "磨く",
      "kana": "みがく",
      "en": "to polish, to brush"
    },
    {
      "kanji": "見つかる",
      "kana": "みつかる",
      "en": "to be discovered"
    },
    {
      "kanji": "見つける",
      "kana": "みつける",
      "en": "to discover"
    },
    {
      "kanji": "見せる",
      "kana": "みせる",
      "en": "to show"
    },
    {
      "kanji": "見る",
      "kana": "みる",
      "en": "to see, to watch"
    },
    {
      "kanji": "向かう",
      "kana": "むかう",
      "en": "to face"
    },
    {
      "kanji": "迎える",
      "kana": "むかえる",
      "en": "to go to meet"
    },
    {
      "kanji": "召し上がる",
      "kana": "めしあがる",
      "en": "to eat (respectful)"
    },
    {
      "kanji": "持つ",
      "kana": "もつ",
      "en": "to have, to own"
    },
    {
      "kanji": "申し上げる",
      "kana": "もうしあげる",
      "en": "to say, to tell (respectful)"
    },
    {
      "kanji": "申す",
      "kana": "もうす",
      "en": "to be called, to say (respectful)"
    },
    {
      "kanji": "戻る",
      "kana": "もどる",
      "en": "to return, to go back"
    },
    {
      "kanji": "貰う",
      "kana": "もらう",
      "en": "to receive"
    },
    {
      "kanji": "休む",
      "kana": "やすむ",
      "en": "to rest"
    },
    {
      "kanji": "焼く",
      "kana": "やく",
      "en": "to bake, to grill"
    },
    {
      "kanji": "役に立つ",
      "kana": "やくにたつ",
      "en": "to be helpful"
    },
    {
      "kanji": "痩せる",
      "kana": "やせる",
      "en": "to become thin"
    },
    {
      "kanji": "止む",
      "kana": "やむ",
      "en": "to stop"
    },
    {
      "kanji": "止める",
      "kana": "やめる",
      "en": "to stop"
    },
    {
      "kanji": "やる",
      "kana": "やる",
      "en": "to do"
    },
    {
      "kanji": "ゆれる",
      "kana": "ゆれる",
      "en": "to shake, to sway"
    },
    {
      "kanji": "呼ぶ",
      "kana": "よぶ",
      "en": "to call"
    },
    {
      "kanji": "読む",
      "kana": "よむ",
      "en": "to read"
    },
    {
      "kanji": "汚れる",
      "kana": "よごれる",
      "en": "to get dirty"
    },
    {
      "kanji": "寄る",
      "kana": "よる",
      "en": "to visit"
    },
    {
      "kanji": "喜ぶ",
      "kana": "よろこぶ",
      "en": "to be delighted"
    },
    {
      "kanji": "沸かす",
      "kana": "わかす",
      "en": "to boil, to heat"
    },
    {
      "kanji": "別れる",
      "kana": "わかれる",
      "en": "to separate"
    },
    {
      "kanji": "分かる",
      "kana": "わかる",
      "en": "to know, to understand"
    },
    {
      "kanji": "沸く",
      "kana": "わく",
      "en": "to boil, to grow hot"
    },
    {
      "kanji": "笑う",
      "kana": "わらう",
      "en": "to laugh, to smile"
    },
    {
      "kanji": "割れる",
      "kana": "われる",
      "en": "to break"
    },
    {
      "kanji": "忘る",
      "kana": "わすれる",
      "en": "to forget"
    },
    {
      "kanji": "渡す",
      "kana": "わたす",
      "en": "to hand over"
    },
    {
      "kanji": "渡る",
      "kana": "わたる",
      "en": "to cross"
    },
    {
      "kanji": "好き",
      "kana": "すき",
      "en": "Like, love"
    },
    {
      "kanji": "大切",
      "kana": "たいせつ",
      "en": "important, great"
    },
    {
      "kanji": "結構",
      "kana": "けっこう",
      "en": "wonderful,enviable"
    },
    {
      "kanji": "大勢",
      "kana": "おおぜい",
      "en": "lots of people, crowded"
    },
    {
      "kanji": "有名",
      "kana": "ゆうめい",
      "en": "Famous"
    },
    {
      "kanji": "きれい",
      "kana": "きれい",
      "en": "Pretty, beautiful"
    },
    {
      "kanji": "簡単",
      "kana": "かんたん",
      "en": "Easy"
    },
    {
      "kanji": "丁寧",
      "kana": "ていねい",
      "en": "polite"
    },
    {
      "kanji": "静か",
      "kana": "しずか",
      "en": "Quiet"
    },
    {
      "kanji": "嫌い",
      "kana": "きらい",
      "en": "to not like"
    },
    {
      "kanji": "静か",
      "kana": "しずか",
      "en": "quiet"
    },
    {
      "kanji": "暇",
      "kana": "ひま",
      "en": "free (time)"
    },
    {
      "kanji": "賑やか",
      "kana": "にぎやか",
      "en": "lively"
    },
    {
      "kanji": "便利",
      "kana": "べんり",
      "en": "useful, convenient"
    },
    {
      "kanji": "元気",
      "kana": "げんき",
      "en": "healthy"
    },
    {
      "kanji": "色々",
      "kana": "いろいろ",
      "en": "various"
    },
    {
      "kanji": "大丈夫",
      "kana": "だいじょうぶ",
      "en": "OK,fine"
    },
    {
      "kanji": "丈夫",
      "kana": "じょうぶ",
      "en": "Healthy, robust"
    },
    {
      "kanji": "大変",
      "kana": "たいへん",
      "en": "Terrible"
    },
    {
      "kanji": "楽",
      "kana": "らく",
      "en": "comfortable, easy"
    },
    {
      "kanji": "嫌",
      "kana": "いや",
      "en": "unpleasant"
    },
    {
      "kanji": "大切",
      "kana": "たいせつ",
      "en": "important,precious"
    },
    {
      "kanji": "上手",
      "kana": "じょうず",
      "en": "good, skilled"
    },
    {
      "kanji": "下手",
      "kana": "へた",
      "en": "bad at, unskilled"
    },
    {
      "kanji": "一生懸命",
      "kana": "いっしょうけんめい",
      "en": "to one’s fullest possibility"
    },
    {
      "kanji": "危険",
      "kana": "きけん",
      "en": "Dangerous"
    },
    {
      "kanji": "残念",
      "kana": "ざんねん",
      "en": "regrettable, dissapointing"
    },
    {
      "kanji": "心配",
      "kana": "しんぱい",
      "en": "care, worry"
    },
    {
      "kanji": "自由",
      "kana": "じゆう",
      "en": "free, unrestrained"
    },
    {
      "kanji": "十分",
      "kana": "じゅうぶん",
      "en": "sufficient"
    },
    {
      "kanji": "大好き",
      "kana": "だいすき",
      "en": "Passionate"
    },
    {
      "kanji": "適当",
      "kana": "てきとう",
      "en": "proper, suitable"
    },
    {
      "kanji": "特別",
      "kana": "とくべつ",
      "en": "special"
    },
    {
      "kanji": "熱心",
      "kana": "ねっしん",
      "en": "eager, enthusiastic"
    },
    {
      "kanji": "必要",
      "kana": "ひつよう",
      "en": "necessary"
    },
    {
      "kanji": "真面目",
      "kana": "まじめ",
      "en": "serious, earnest"
    },
    {
      "kanji": "真直ぐ",
      "kana": "まっすぐ",
      "en": "straight"
    },
    {
      "kanji": "無理",
      "kana": "むり",
      "en": "unreasonnable"
    },
    {
      "kanji": "立派",
      "kana": "りっぱ",
      "en": "Excellent, splendid,fine"
    },
    {
      "kanji": "不便",
      "kana": "ふべん",
      "en": "Inconvenient"
    },
    {
      "kanji": "青い",
      "kana": "あおい",
      "en": "blue"
    },
    {
      "kanji": "赤い",
      "kana": "あかい",
      "en": "red"
    },
    {
      "kanji": "明るい",
      "kana": "あかるい",
      "en": "light, bright"
    },
    {
      "kanji": "温かい",
      "kana": "あたたかい",
      "en": "warm"
    },
    {
      "kanji": "新しい",
      "kana": "あたらしい",
      "en": "new"
    },
    {
      "kanji": "暑い",
      "kana": "あつい",
      "en": "hot (air)"
    },
    {
      "kanji": "厚い",
      "kana": "あつい",
      "en": "thick"
    },
    {
      "kanji": "危ない",
      "kana": "あぶない",
      "en": "dangerous"
    },
    {
      "kanji": "甘い",
      "kana": "あまい",
      "en": "sweet"
    },
    {
      "kanji": "良い",
      "kana": "よい / いい",
      "en": "good"
    },
    {
      "kanji": "忙しい",
      "kana": "いそがしい",
      "en": "to be busy"
    },
    {
      "kanji": "痛い",
      "kana": "いたい",
      "en": "to be painful"
    },
    {
      "kanji": "薄い",
      "kana": "うすい",
      "en": "thin"
    },
    {
      "kanji": "美味しい",
      "kana": "おいしい",
      "en": "tasty, delicious"
    },
    {
      "kanji": "大きい",
      "kana": "おおきい",
      "en": "big"
    },
    {
      "kanji": "遅い",
      "kana": "おそい",
      "en": "late, slow"
    },
    {
      "kanji": "重い",
      "kana": "おもい",
      "en": "heavy"
    },
    {
      "kanji": "面白い",
      "kana": "おもしろい",
      "en": "intersting, funny"
    },
    {
      "kanji": "辛い",
      "kana": "からい",
      "en": "hot, spicy"
    },
    {
      "kanji": "軽い",
      "kana": "かるい",
      "en": "light (not heavy)"
    },
    {
      "kanji": "可愛い",
      "kana": "かわいい",
      "en": "cute, pretty"
    },
    {
      "kanji": "厳しい",
      "kana": "きびしい",
      "en": "strict"
    },
    {
      "kanji": "黄色い",
      "kana": "きいろい",
      "en": "yellow"
    },
    {
      "kanji": "汚い",
      "kana": "きたない",
      "en": "dirty"
    },
    {
      "kanji": "暗い",
      "kana": "くらい",
      "en": "dark"
    },
    {
      "kanji": "黒い",
      "kana": "くろい",
      "en": "black"
    },
    {
      "kanji": "細かい",
      "kana": "こまかい",
      "en": "small, fin"
    },
    {
      "kanji": "怖い",
      "kana": "こわい",
      "en": "Frightening"
    },
    {
      "kanji": "寂しい",
      "kana": "さびしい",
      "en": "lonely"
    },
    {
      "kanji": "寒い",
      "kana": "さむい",
      "en": "cold"
    },
    {
      "kanji": "白い",
      "kana": "しろい",
      "en": "white"
    },
    {
      "kanji": "涼しい",
      "kana": "すずしい",
      "en": "cool"
    },
    {
      "kanji": "凄い",
      "kana": "すごい",
      "en": "terrific"
    },
    {
      "kanji": "素晴らしい",
      "kana": "すばらしい",
      "en": "wonderful"
    },
    {
      "kanji": "狭い",
      "kana": "せまい",
      "en": "narrow"
    },
    {
      "kanji": "正しい",
      "kana": "ただしい",
      "en": "correct"
    },
    {
      "kanji": "高い",
      "kana": "たかい",
      "en": "high, expensive"
    },
    {
      "kanji": "楽しい",
      "kana": "たのしい",
      "en": "pleasant, enjoyable"
    },
    {
      "kanji": "小さい",
      "kana": "ちいさい",
      "en": "small"
    },
    {
      "kanji": "近い",
      "kana": "ちかい",
      "en": "near, close"
    },
    {
      "kanji": "詰らない",
      "kana": "つまらない",
      "en": "uninteresting"
    },
    {
      "kanji": "冷たい",
      "kana": "つめたい",
      "en": "cold"
    },
    {
      "kanji": "強い",
      "kana": "つよい",
      "en": "strong"
    },
    {
      "kanji": "遠い",
      "kana": "とおい",
      "en": "far"
    },
    {
      "kanji": "長い",
      "kana": "ながい",
      "en": "long"
    },
    {
      "kanji": "苦い",
      "kana": "にがい",
      "en": "bitter"
    },
    {
      "kanji": "眠い",
      "kana": "ねむい",
      "en": "to be sleepy"
    },
    {
      "kanji": "早い",
      "kana": "はやい",
      "en": "early"
    },
    {
      "kanji": "速い",
      "kana": "はやい",
      "en": "fast, quick"
    },
    {
      "kanji": "恥ずかしい",
      "kana": "はずかしい",
      "en": "to be embarrassed"
    },
    {
      "kanji": "低い",
      "kana": "ひくい",
      "en": "low"
    },
    {
      "kanji": "広い",
      "kana": "ひろい",
      "en": "wide, spacious"
    },
    {
      "kanji": "酷い",
      "kana": "ひどい",
      "en": "awful"
    },
    {
      "kanji": "太い",
      "kana": "ふとい",
      "en": "thick, fat"
    },
    {
      "kanji": "古い",
      "kana": "ふるい",
      "en": "old"
    },
    {
      "kanji": "深い",
      "kana": "ふかい",
      "en": "deep"
    },
    {
      "kanji": "欲しい",
      "kana": "ほしい",
      "en": "to want something"
    },
    {
      "kanji": "細い",
      "kana": "ほそい",
      "en": "thin, fine"
    },
    {
      "kanji": "不味い",
      "kana": "まずい",
      "en": "bad tasting"
    },
    {
      "kanji": "丸い",
      "kana": "まるい",
      "en": "round"
    },
    {
      "kanji": "短い",
      "kana": "みじかい",
      "en": "short"
    },
    {
      "kanji": "難しい",
      "kana": "むずかしい",
      "en": "difficult"
    },
    {
      "kanji": "珍しい",
      "kana": "めずらしい",
      "en": "rare"
    },
    {
      "kanji": "優しい",
      "kana": "やさしい",
      "en": "gentle"
    },
    {
      "kanji": "安い",
      "kana": "やすい",
      "en": "cheap"
    },
    {
      "kanji": "柔らかい",
      "kana": "やわらかい",
      "en": "tender, soft"
    },
    {
      "kanji": "宜しい",
      "kana": "よろしい",
      "en": "ok, all right (respectful)"
    },
    {
      "kanji": "若い",
      "kana": "わかい",
      "en": "young"
    },
    {
      "kanji": "あ",
      "kana": "あ",
      "en": "Ah!, Oh!"
    },
    {
      "kanji": "ああ",
      "kana": "ああ",
      "en": "Ah!"
    },
    {
      "kanji": "挨拶",
      "kana": "あいさつ",
      "en": "greeting"
    },
    {
      "kanji": "間",
      "kana": "あいだ",
      "en": "between"
    },
    {
      "kanji": "合う",
      "kana": "あう",
      "en": "fit, suit"
    },
    {
      "kanji": "会う",
      "kana": "あう",
      "en": "to meet"
    },
    {
      "kanji": "青い",
      "kana": "あおい",
      "en": "blue"
    },
    {
      "kanji": "赤い",
      "kana": "あかい",
      "en": "red"
    },
    {
      "kanji": "赤ちゃん",
      "kana": "あかちゃん",
      "en": "baby, newborn"
    },
    {
      "kanji": "上がる",
      "kana": "あがる",
      "en": "to go up"
    },
    {
      "kanji": "明るい",
      "kana": "あかるい",
      "en": "light, bright"
    },
    {
      "kanji": "赤ん坊",
      "kana": "あかんぼう",
      "en": "baby, newborn"
    },
    {
      "kanji": "秋",
      "kana": "あき",
      "en": "autumn, fall"
    },
    {
      "kanji": "開く",
      "kana": "あく",
      "en": "open"
    },
    {
      "kanji": "空く",
      "kana": "あく",
      "en": "empty, has space"
    },
    {
      "kanji": "開ける",
      "kana": "あける",
      "en": "to open"
    },
    {
      "kanji": "上げる",
      "kana": "あげる",
      "en": "to rise"
    },
    {
      "kanji": "あげる",
      "kana": "あげる",
      "en": "to give"
    },
    {
      "kanji": "朝",
      "kana": "あさ",
      "en": "morning"
    },
    {
      "kanji": "朝ご飯",
      "kana": "あさごはん",
      "en": "breakfast"
    },
    {
      "kanji": "あさって",
      "kana": "あさって",
      "en": "the day after tomorrow"
    },
    {
      "kanji": "朝寝坊",
      "kana": "あさねぼう",
      "en": "a late riser"
    },
    {
      "kanji": "足",
      "kana": "あし",
      "en": "leg, foot"
    },
    {
      "kanji": "味",
      "kana": "あじ",
      "en": "taste, flavor"
    },
    {
      "kanji": "明日",
      "kana": "あした",
      "en": "tomorrow"
    },
    {
      "kanji": "明日",
      "kana": "あす",
      "en": "tomorrow (polite)"
    },
    {
      "kanji": "あそこ",
      "kana": "あそこ",
      "en": "over there"
    },
    {
      "kanji": "遊ぶ",
      "kana": "あそぶ",
      "en": "to play"
    },
    {
      "kanji": "温かい",
      "kana": "あたたかい",
      "en": "warm"
    },
    {
      "kanji": "頭",
      "kana": "あたま",
      "en": "head"
    },
    {
      "kanji": "新しい",
      "kana": "あたらしい",
      "en": "new"
    },
    {
      "kanji": "あちら",
      "kana": "あちら",
      "en": "over there (polite)"
    },
    {
      "kanji": "暑い",
      "kana": "あつい",
      "en": "hot (air)"
    },
    {
      "kanji": "厚い",
      "kana": "あつい",
      "en": "thick"
    },
    {
      "kanji": "後",
      "kana": "あと",
      "en": "later, after"
    },
    {
      "kanji": "貴方",
      "kana": "あなた",
      "en": "you"
    },
    {
      "kanji": "兄",
      "kana": "あに",
      "en": "older brother"
    },
    {
      "kanji": "姉",
      "kana": "あね",
      "en": "older sister"
    },
    {
      "kanji": "あの",
      "kana": "あの",
      "en": "that (over there)"
    },
    {
      "kanji": "あの",
      "kana": "あの",
      "en": "well, then"
    },
    {
      "kanji": "アパート",
      "kana": "アパート",
      "en": "apartment"
    },
    {
      "kanji": "浴びる",
      "kana": "あびる",
      "en": "to take a shower"
    },
    {
      "kanji": "危ない",
      "kana": "あぶない",
      "en": "dangerous"
    },
    {
      "kanji": "甘い",
      "kana": "あまい",
      "en": "sweet"
    },
    {
      "kanji": "あまり",
      "kana": "あまり",
      "en": "not so"
    },
    {
      "kanji": "雨",
      "kana": "あめ",
      "en": "rain"
    },
    {
      "kanji": "洗う",
      "kana": "あらう",
      "en": "to wash"
    },
    {
      "kanji": "有る",
      "kana": "ある",
      "en": "to be, to exist"
    },
    {
      "kanji": "ある",
      "kana": "ある",
      "en": "to possess"
    },
    {
      "kanji": "歩く",
      "kana": "ありく",
      "en": "to walk"
    },
    {
      "kanji": "あれ",
      "kana": "あれ",
      "en": "that one"
    },
    {
      "kanji": "良い",
      "kana": "いい / よい",
      "en": "good"
    },
    {
      "kanji": "いいえ",
      "kana": "いいえ",
      "en": "no"
    },
    {
      "kanji": "言う",
      "kana": "いう",
      "en": "to say, to tell"
    },
    {
      "kanji": "家",
      "kana": "いえ",
      "en": "house, home"
    },
    {
      "kanji": "行く",
      "kana": "いく",
      "en": "to go"
    },
    {
      "kanji": "いくつ",
      "kana": "いくつ",
      "en": "how many, how old"
    },
    {
      "kanji": "いくら",
      "kana": "いくら",
      "en": "how much"
    },
    {
      "kanji": "池",
      "kana": "いけ",
      "en": "pond"
    },
    {
      "kanji": "医者",
      "kana": "いしゃ",
      "en": "doctor"
    },
    {
      "kanji": "椅子",
      "kana": "いす",
      "en": "chair"
    },
    {
      "kanji": "忙しい",
      "kana": "いそがしい",
      "en": "to be busy"
    },
    {
      "kanji": "痛い",
      "kana": "いたい",
      "en": "to be painful"
    },
    {
      "kanji": "一",
      "kana": "いち",
      "en": "one"
    },
    {
      "kanji": "一日",
      "kana": "いちにち",
      "en": "one day"
    },
    {
      "kanji": "一番",
      "kana": "いちばん",
      "en": "No. 1, the best, the first"
    },
    {
      "kanji": "いつ",
      "kana": "いつ",
      "en": "when"
    },
    {
      "kanji": "五日",
      "kana": "いつか",
      "en": "the 5th day of the month, 5 days"
    },
    {
      "kanji": "一緒",
      "kana": "いっしょ",
      "en": "together"
    },
    {
      "kanji": "五つ",
      "kana": "いつつ",
      "en": "five"
    },
    {
      "kanji": "いつも",
      "kana": "いつも",
      "en": "always"
    },
    {
      "kanji": "今",
      "kana": "いま",
      "en": "now"
    },
    {
      "kanji": "意味",
      "kana": "いみ",
      "en": "meaning"
    },
    {
      "kanji": "妹",
      "kana": "いもうと",
      "en": "someone’s younger sister"
    },
    {
      "kanji": "いや",
      "kana": "いや",
      "en": "not likable, unpleasant"
    },
    {
      "kanji": "入口",
      "kana": "いりぐち",
      "en": "entrance"
    },
    {
      "kanji": "いる",
      "kana": "いる",
      "en": "need, must have, be required"
    },
    {
      "kanji": "いる",
      "kana": "いる",
      "en": "to exist"
    },
    {
      "kanji": "入れる",
      "kana": "いれる",
      "en": "to insert, to put in"
    },
    {
      "kanji": "色",
      "kana": "いろ",
      "en": "color"
    },
    {
      "kanji": "色々",
      "kana": "いろいろ",
      "en": "various"
    },
    {
      "kanji": "上",
      "kana": "うえ",
      "en": "top, on, above"
    },
    {
      "kanji": "後ろ",
      "kana": "うしろ",
      "en": "back, rear, behind"
    },
    {
      "kanji": "薄い",
      "kana": "うすい",
      "en": "thin"
    },
    {
      "kanji": "歌",
      "kana": "うた",
      "en": "song"
    },
    {
      "kanji": "歌う",
      "kana": "うたう",
      "en": "to sing"
    },
    {
      "kanji": "内",
      "kana": "うち",
      "en": "home"
    },
    {
      "kanji": "生まれる",
      "kana": "うまれる",
      "en": "to be born"
    },
    {
      "kanji": "海",
      "kana": "うみ",
      "en": "sea"
    },
    {
      "kanji": "売る",
      "kana": "うる",
      "en": "to sell"
    },
    {
      "kanji": "上着",
      "kana": "うわぎ",
      "en": "coat, jacket"
    },
    {
      "kanji": "絵",
      "kana": "え",
      "en": "picture"
    },
    {
      "kanji": "映画",
      "kana": "えいが",
      "en": "movie"
    },
    {
      "kanji": "映画館",
      "kana": "えいがかん",
      "en": "cinema"
    },
    {
      "kanji": "英語",
      "kana": "英語",
      "en": "English language"
    },
    {
      "kanji": "ええ",
      "kana": "ええ",
      "en": "Yes, I see"
    },
    {
      "kanji": "駅",
      "kana": "えき",
      "en": "station"
    },
    {
      "kanji": "エレベータ",
      "kana": "エレベータ",
      "en": "elevator"
    },
    {
      "kanji": "円",
      "kana": "えん",
      "en": "Yen"
    },
    {
      "kanji": "鉛筆",
      "kana": "えんぴつ",
      "en": "pencil"
    },
    {
      "kanji": "御",
      "kana": "お",
      "en": "honorific prefix"
    },
    {
      "kanji": "美味しい",
      "kana": "おいしい",
      "en": "tasty, delicious"
    },
    {
      "kanji": "大きい",
      "kana": "おおきい",
      "en": "big"
    },
    {
      "kanji": "おおぜい",
      "kana": "おおぜい",
      "en": "many people"
    },
    {
      "kanji": "お母さん",
      "kana": "おかあさん",
      "en": "my own mother"
    },
    {
      "kanji": "お菓子",
      "kana": "okashi",
      "en": "confectionary, cake"
    },
    {
      "kanji": "お金",
      "kana": "おかね",
      "en": "money"
    },
    {
      "kanji": "起きる",
      "kana": "おきる",
      "en": "to get up, to stand up"
    },
    {
      "kanji": "おきる",
      "kana": "おきる",
      "en": "to put"
    },
    {
      "kanji": "奥さん",
      "kana": "おくさん",
      "en": "someone’s wife"
    },
    {
      "kanji": "送る",
      "kana": "おくる",
      "en": "to send"
    },
    {
      "kanji": "お酒",
      "kana": "おさけ",
      "en": "alcohol, sake"
    },
    {
      "kanji": "お皿",
      "kana": "おさら",
      "en": "plate"
    },
    {
      "kanji": "伯父さん",
      "kana": "おじさん",
      "en": "uncle"
    },
    {
      "kanji": "おじいさん",
      "kana": "おじいさん",
      "en": "grand father"
    },
    {
      "kanji": "押す",
      "kana": "おす",
      "en": "to push"
    },
    {
      "kanji": "遅い",
      "kana": "おそい",
      "en": "late, slow"
    },
    {
      "kanji": "お茶",
      "kana": "おちゃ",
      "en": "tea"
    },
    {
      "kanji": "お手洗い",
      "kana": "おてあらい",
      "en": "toilet, lavatory"
    },
    {
      "kanji": "お父さん",
      "kana": "おとうさん",
      "en": "father"
    },
    {
      "kanji": "弟",
      "kana": "おとうと",
      "en": "someone’s younger brother"
    },
    {
      "kanji": "男",
      "kana": "おとこ",
      "en": "man"
    },
    {
      "kanji": "男の子",
      "kana": "おとこのこ",
      "en": "boy"
    },
    {
      "kanji": "一昨日",
      "kana": "おととい",
      "en": "the day before yesterday"
    },
    {
      "kanji": "一昨年",
      "kana": "おととし",
      "en": "the year before last"
    },
    {
      "kanji": "大人",
      "kana": "おとな",
      "en": "adult"
    },
    {
      "kanji": "お腹",
      "kana": "おなか",
      "en": "stomach"
    },
    {
      "kanji": "同じ",
      "kana": "おなじ",
      "en": "same"
    },
    {
      "kanji": "お兄さん",
      "kana": "おにいさん",
      "en": "someone’s elder brother"
    },
    {
      "kanji": "お姉さん",
      "kana": "おねえさん",
      "en": "someone’s elder sister"
    },
    {
      "kanji": "伯母さん",
      "kana": "おばさん",
      "en": "aunt"
    },
    {
      "kanji": "おばあさん",
      "kana": "おばあさん",
      "en": "grandmother"
    },
    {
      "kanji": "お弁当",
      "kana": "おべんとう",
      "en": "lunchbox"
    },
    {
      "kanji": "覚える",
      "kana": "おぼえる",
      "en": "to memorize, to remember"
    },
    {
      "kanji": "重い",
      "kana": "おもい",
      "en": "heavy"
    },
    {
      "kanji": "面白い",
      "kana": "おもしろい",
      "en": "interesting, funny"
    },
    {
      "kanji": "泳ぐ",
      "kana": "およぐ",
      "en": "to swim"
    },
    {
      "kanji": "降りる",
      "kana": "おりる",
      "en": "to get off"
    },
    {
      "kanji": "終わる",
      "kana": "おわる",
      "en": "to end"
    },
    {
      "kanji": "音楽",
      "kana": "おんがく",
      "en": "music"
    },
    {
      "kanji": "女",
      "kana": "おんな",
      "en": "woman"
    },
    {
      "kanji": "女の子",
      "kana": "おんなのこ",
      "en": "girl"
    },
    {
      "kanji": "〜回",
      "kana": "〜かい",
      "en": "~times"
    },
    {
      "kanji": "〜階",
      "kana": "〜かい",
      "en": "~floor"
    },
    {
      "kanji": "外国",
      "kana": "がいこく",
      "en": "foreign country"
    },
    {
      "kanji": "外国人",
      "kana": "がいこくじん",
      "en": "foreigner"
    },
    {
      "kanji": "会社",
      "kana": "かいしゃ",
      "en": "company, enterprise"
    },
    {
      "kanji": "階段",
      "kana": "かいだん",
      "en": "stairs"
    },
    {
      "kanji": "買物",
      "kana": "かいもの",
      "en": "shopping"
    },
    {
      "kanji": "買う",
      "kana": "かう",
      "en": "to buy"
    },
    {
      "kanji": "返す",
      "kana": "かえす",
      "en": "to return an object"
    },
    {
      "kanji": "帰る",
      "kana": "かえる",
      "en": "to return home"
    },
    {
      "kanji": "顔",
      "kana": "かお",
      "en": "face"
    },
    {
      "kanji": "かかる",
      "kana": "かかる",
      "en": "to take time, money"
    },
    {
      "kanji": "鍵",
      "kana": "かぎ",
      "en": "key"
    },
    {
      "kanji": "書く",
      "kana": "かく",
      "en": "to write"
    },
    {
      "kanji": "学生",
      "kana": "がくせい",
      "en": "student"
    },
    {
      "kanji": "〜か月",
      "kana": "〜かげつ",
      "en": "~ number of months"
    },
    {
      "kanji": "かける",
      "kana": "かける",
      "en": "to wear"
    },
    {
      "kanji": "かける",
      "kana": "かける",
      "en": "to make a phone call"
    },
    {
      "kanji": "傘",
      "kana": "かさ",
      "en": "umbrella"
    },
    {
      "kanji": "貸す",
      "kana": "かす",
      "en": "to lend"
    },
    {
      "kanji": "風",
      "kana": "かぜ",
      "en": "wind"
    },
    {
      "kanji": "風邪",
      "kana": "かぜ",
      "en": "a cold"
    },
    {
      "kanji": "家族",
      "kana": "かぞく",
      "en": "family"
    },
    {
      "kanji": "方",
      "kana": "かた",
      "en": "person (polite)"
    },
    {
      "kanji": "片仮名",
      "kana": "かたかな",
      "en": "Katakana"
    },
    {
      "kanji": "一月",
      "kana": "いちがつ",
      "en": "January"
    },
    {
      "kanji": "二月",
      "kana": "にがつ",
      "en": "February"
    },
    {
      "kanji": "三月",
      "kana": "さんがつ",
      "en": "March"
    },
    {
      "kanji": "四月",
      "kana": "しがつ",
      "en": "April"
    },
    {
      "kanji": "五月",
      "kana": "ごがつ",
      "en": "May"
    },
    {
      "kanji": "六月",
      "kana": "ろくがつ",
      "en": "June"
    },
    {
      "kanji": "七月",
      "kana": "しちがつ",
      "en": "July"
    },
    {
      "kanji": "八月",
      "kana": "はちがつ",
      "en": "August"
    },
    {
      "kanji": "九月",
      "kana": "くがつ",
      "en": "September"
    },
    {
      "kanji": "十月",
      "kana": "じゅうがつ",
      "en": "October"
    },
    {
      "kanji": "十一月",
      "kana": "じゅういちがつ",
      "en": "November"
    },
    {
      "kanji": "十二月",
      "kana": "じゅうにがつ",
      "en": "December"
    },
    {
      "kanji": "学校",
      "kana": "がっこう",
      "en": "school"
    },
    {
      "kanji": "角",
      "kana": "かど",
      "en": "corner"
    },
    {
      "kanji": "家内",
      "kana": "かない",
      "en": "my wife"
    },
    {
      "kanji": "鞄",
      "kana": "かばん",
      "en": "bag"
    },
    {
      "kanji": "花瓶",
      "kana": "かびん",
      "en": "vase"
    },
    {
      "kanji": "冠る",
      "kana": "かぶる",
      "en": "to put on a hat"
    },
    {
      "kanji": "紙",
      "kana": "かみ",
      "en": "paper"
    },
    {
      "kanji": "カメラ",
      "kana": "かめら",
      "en": "camera"
    },
    {
      "kanji": "火曜日",
      "kana": "かようび",
      "en": "Tuesday"
    },
    {
      "kanji": "辛い",
      "kana": "からい",
      "en": "hot, spicy"
    },
    {
      "kanji": "体",
      "kana": "からだ",
      "en": "body"
    },
    {
      "kanji": "借りる",
      "kana": "かりる",
      "en": "to borrow"
    },
    {
      "kanji": "〜がります",
      "kana": "〜がります",
      "en": "3rd person wants to"
    },
    {
      "kanji": "軽い",
      "kana": "かるい",
      "en": "light (not heavy)"
    },
    {
      "kanji": "カレンダー",
      "kana": "カレンダー",
      "en": "calendar"
    },
    {
      "kanji": "川",
      "kana": "かわ",
      "en": "river"
    },
    {
      "kanji": "〜側",
      "kana": "~がわ",
      "en": "~side"
    },
    {
      "kanji": "可愛い",
      "kana": "かわいい",
      "en": "cute, pretty"
    },
    {
      "kanji": "漢字",
      "kana": "かんじ",
      "en": "Kanji character"
    },
    {
      "kanji": "気",
      "kana": "き",
      "en": "Spirit, mood"
    },
    {
      "kanji": "機会",
      "kana": "きかい",
      "en": "Opportunity"
    },
    {
      "kanji": "危険",
      "kana": "きけん",
      "en": "Danger"
    },
    {
      "kanji": "聞こえる",
      "kana": "きこえる",
      "en": "to be heard"
    },
    {
      "kanji": "汽車",
      "kana": "きしゃ",
      "en": "steam train"
    },
    {
      "kanji": "技術",
      "kana": "ぎじゅつ",
      "en": "Skill, art, technique"
    },
    {
      "kanji": "季節",
      "kana": "きせつ",
      "en": "season"
    },
    {
      "kanji": "規則",
      "kana": "きそく",
      "en": "regulations"
    },
    {
      "kanji": "きっと",
      "kana": "きっと",
      "en": "surely"
    },
    {
      "kanji": "絹",
      "kana": "きぬ",
      "en": "Silk"
    },
    {
      "kanji": "厳しい",
      "kana": "きびしい",
      "en": "strict"
    },
    {
      "kanji": "気分",
      "kana": "きぶん",
      "en": "mood"
    },
    {
      "kanji": "決まる",
      "kana": "きまる",
      "en": "to be decided"
    },
    {
      "kanji": "君",
      "kana": "きみ",
      "en": "You (informal)"
    },
    {
      "kanji": "決める",
      "kana": "きめる",
      "en": "to decide"
    },
    {
      "kanji": "気持ち",
      "kana": "きもち",
      "en": "feeling, mood"
    },
    {
      "kanji": "着物",
      "kana": "きもの",
      "en": "a kimono"
    },
    {
      "kanji": "客",
      "kana": "きゃく",
      "en": "guest, customer"
    },
    {
      "kanji": "急",
      "kana": "きゅう",
      "en": "urgent"
    },
    {
      "kanji": "急行",
      "kana": "きゅうこう",
      "en": "speedy, express"
    },
    {
      "kanji": "教育",
      "kana": "きょういく",
      "en": "education"
    },
    {
      "kanji": "教会",
      "kana": "きょうかい",
      "en": "church"
    },
    {
      "kanji": "競争",
      "kana": "きょうそう",
      "en": "competition"
    },
    {
      "kanji": "興味",
      "kana": "きょうみ",
      "en": "an interest"
    },
    {
      "kanji": "近所",
      "kana": "きんじょ",
      "en": "neighbourhood"
    },
    {
      "kanji": "具合",
      "kana": "ぐあい",
      "en": "condition, health"
    },
    {
      "kanji": "空気",
      "kana": "くうき",
      "en": "air, atmosphere"
    },
    {
      "kanji": "空港",
      "kana": "くうこう",
      "en": "airport"
    },
    {
      "kanji": "草",
      "kana": "くさ",
      "en": "grass"
    },
    {
      "kanji": "くださる",
      "kana": "くださる",
      "en": "to give (to me) respectful"
    },
    {
      "kanji": "首",
      "kana": "くび",
      "en": "neck"
    },
    {
      "kanji": "雲",
      "kana": "くも",
      "en": "cloud"
    },
    {
      "kanji": "比べる",
      "kana": "くらべる",
      "en": "to compare"
    },
    {
      "kanji": "くれる",
      "kana": "くれる",
      "en": "to give (to me) simple form"
    },
    {
      "kanji": "暮れる",
      "kana": "くれる",
      "en": "to get dark, to come to an end"
    },
    {
      "kanji": "君",
      "kana": "くん",
      "en": "suffix for young male (familiar)"
    },
    {
      "kanji": "毛",
      "kana": "け",
      "en": "hair or fur"
    },
    {
      "kanji": "計画",
      "kana": "けいかく",
      "en": "plan"
    },
    {
      "kanji": "経験",
      "kana": "けいけん",
      "en": "experience"
    },
    {
      "kanji": "経済",
      "kana": "けいざい",
      "en": "finance, economy"
    },
    {
      "kanji": "警察",
      "kana": "けいさつ",
      "en": "police"
    },
    {
      "kanji": "ケーキ",
      "kana": "ケーキ",
      "en": "Cake"
    },
    {
      "kanji": "怪我",
      "kana": "けが",
      "en": "an injury"
    },
    {
      "kanji": "景色",
      "kana": "けしき",
      "en": "landscape"
    },
    {
      "kanji": "消しゴム",
      "kana": "けしごむ",
      "en": "an eraser"
    },
    {
      "kanji": "下宿",
      "kana": "げしゅく",
      "en": "lodging"
    },
    {
      "kanji": "決して",
      "kana": "けっして",
      "en": "never"
    },
    {
      "kanji": "けれど/けれども",
      "kana": "けれど/けれども",
      "en": "however"
    },
    {
      "kanji": "原因",
      "kana": "げんいん",
      "en": "source, cause"
    },
    {
      "kanji": "喧嘩",
      "kana": "けんか",
      "en": "a quarrel"
    },
    {
      "kanji": "研究",
      "kana": "けんきゅう",
      "en": "research"
    },
    {
      "kanji": "研究室",
      "kana": "けんきゅうしつ",
      "en": "laboratory"
    },
    {
      "kanji": "見物",
      "kana": "けんぶつ",
      "en": "sightseeing"
    },
    {
      "kanji": "子",
      "kana": "こ",
      "en": "child"
    },
    {
      "kanji": "こう",
      "kana": "こう",
      "en": "this way"
    },
    {
      "kanji": "郊外",
      "kana": "こうがい",
      "en": "outskirts"
    },
    {
      "kanji": "講義",
      "kana": "こうぎ",
      "en": "lecture"
    },
    {
      "kanji": "工業",
      "kana": "こうぎょう",
      "en": "the manufacturing industry"
    },
    {
      "kanji": "高校",
      "kana": "こうこう",
      "en": "high school"
    },
    {
      "kanji": "高校生",
      "kana": "こうこうせい",
      "en": "high school students"
    },
    {
      "kanji": "工場",
      "kana": "こうじょう",
      "en": "factory"
    },
    {
      "kanji": "校長",
      "kana": "こうちょう",
      "en": "headmaster"
    },
    {
      "kanji": "交通",
      "kana": "こうつう",
      "en": "traffic, transportation"
    },
    {
      "kanji": "講堂",
      "kana": "こうどう",
      "en": "auditorium"
    },
    {
      "kanji": "高等学校",
      "kana": "こうとうがっこう",
      "en": "High school (full word)"
    },
    {
      "kanji": "公務員",
      "kana": "こうむいん",
      "en": "Government worker"
    },
    {
      "kanji": "国際",
      "kana": "こくさい",
      "en": "international"
    },
    {
      "kanji": "心",
      "kana": "こころ",
      "en": "core,heart"
    },
    {
      "kanji": "ご主人",
      "kana": "ごしゅじん",
      "en": "Your husband (respectful)"
    },
    {
      "kanji": "故障",
      "kana": "こしょう",
      "en": "trouble, damage, breakdown"
    },
    {
      "kanji": "ご存知",
      "kana": "ごぞんじ",
      "en": "knowing, acquaintance"
    },
    {
      "kanji": "答",
      "kana": "こたえ",
      "en": "response"
    },
    {
      "kanji": "ごちそう",
      "kana": "ごちそう",
      "en": "a feast"
    },
    {
      "kanji": "こと",
      "kana": "こと",
      "en": "thing,matter"
    },
    {
      "kanji": "小鳥",
      "kana": "ことり",
      "en": "small, bird"
    },
    {
      "kanji": "このあいだ",
      "kana": "このあいだ",
      "en": "the other day, recently"
    },
    {
      "kanji": "このごろ",
      "kana": "このごろ",
      "en": "these days, nowadays"
    },
    {
      "kanji": "細かい",
      "kana": "こまかい",
      "en": "smal, fin"
    },
    {
      "kanji": "ごみ",
      "kana": "ごみ",
      "en": "rubbish"
    },
    {
      "kanji": "込む",
      "kana": "こむ",
      "en": "to be crowded"
    },
    {
      "kanji": "米",
      "kana": "こめ",
      "en": "uncooked rice"
    },
    {
      "kanji": "ご覧になる",
      "kana": "ごらんになる",
      "en": "to see(respectful)"
    },
    {
      "kanji": "これから",
      "kana": "これから",
      "en": "after this"
    },
    {
      "kanji": "怖い",
      "kana": "こわい",
      "en": "frightening"
    },
    {
      "kanji": "壊す",
      "kana": "こわす",
      "en": "to break"
    },
    {
      "kanji": "壊れる",
      "kana": "こわれる",
      "en": "to be broken"
    },
    {
      "kanji": "コンサート",
      "kana": "コンサート",
      "en": "Concert"
    },
    {
      "kanji": "今度",
      "kana": "こんど",
      "en": "now, next time"
    },
    {
      "kanji": "コンピューター",
      "kana": "コンピューター",
      "en": "Computer"
    },
    {
      "kanji": "今夜",
      "kana": "こんや",
      "en": "tonight"
    },
    {
      "kanji": "最近",
      "kana": "さいきん",
      "en": "nowadays"
    },
    {
      "kanji": "最後",
      "kana": "さいご",
      "en": "last, end"
    },
    {
      "kanji": "最初",
      "kana": "さいしょ",
      "en": "beginning, first"
    },
    {
      "kanji": "坂",
      "kana": "さか",
      "en": "slope, hill"
    },
    {
      "kanji": "探す",
      "kana": "さがす",
      "en": "to look for"
    },
    {
      "kanji": "下がる",
      "kana": "さがる",
      "en": "to get down"
    },
    {
      "kanji": "盛ん",
      "kana": "さかん",
      "en": "popular, prosperous"
    },
    {
      "kanji": "下げる",
      "kana": "さげる",
      "en": "to lower"
    },
    {
      "kanji": "差し上げる",
      "kana": "さしあげる",
      "en": "to give (respectful)"
    },
    {
      "kanji": "さっき",
      "kana": "さっき",
      "en": "a little while ago"
    },
    {
      "kanji": "寂しい",
      "kana": "さびしい",
      "en": "lonely"
    },
    {
      "kanji": "再来月",
      "kana": "さらいげつ",
      "en": "the month after next"
    },
    {
      "kanji": "再来週",
      "kana": "さらいしゅう",
      "en": "the week after next"
    },
    {
      "kanji": "サラダ",
      "kana": "さらだ",
      "en": "salad"
    },
    {
      "kanji": "騒ぐ",
      "kana": "さわぐ",
      "en": "to make noise, to be excited"
    },
    {
      "kanji": "触る",
      "kana": "さわる",
      "en": "to touch"
    },
    {
      "kanji": "産業",
      "kana": "さんぎょう",
      "en": "industry"
    },
    {
      "kanji": "サンダル",
      "kana": "さんだる",
      "en": "sandal"
    },
    {
      "kanji": "サンドイッチ",
      "kana": "さんどいっち",
      "en": "sandwich"
    },
    {
      "kanji": "残念",
      "kana": "ざんねん",
      "en": "disappointment"
    },
    {
      "kanji": "市",
      "kana": "し",
      "en": "city"
    },
    {
      "kanji": "字",
      "kana": "じ",
      "en": "city"
    },
    {
      "kanji": "試合",
      "kana": "しあい",
      "en": "match, game"
    },
    {
      "kanji": "仕方",
      "kana": "しかた",
      "en": "method"
    },
    {
      "kanji": "叱る",
      "kana": "しかる",
      "en": "to scold"
    },
    {
      "kanji": "試験",
      "kana": "しけん",
      "en": "exam"
    },
    {
      "kanji": "事故",
      "kana": "じこ",
      "en": "accident"
    },
    {
      "kanji": "地震",
      "kana": "じしん",
      "en": "earthquake"
    },
    {
      "kanji": "時代",
      "kana": "じだい",
      "en": "era"
    },
    {
      "kanji": "下着",
      "kana": "したぎ",
      "en": "underwear"
    },
    {
      "kanji": "支度",
      "kana": "したく",
      "en": "preparations"
    },
    {
      "kanji": "しっかり",
      "kana": "しっかり",
      "en": "firmly, steadily"
    },
    {
      "kanji": "失敗",
      "kana": "しっぱい",
      "en": "failure, mistake"
    },
    {
      "kanji": "辞典",
      "kana": "じてん",
      "en": "dictionary"
    },
    {
      "kanji": "品物",
      "kana": "しなもの",
      "en": "goods"
    },
    {
      "kanji": "しばらく",
      "kana": "しばらく",
      "en": "little while"
    },
    {
      "kanji": "島",
      "kana": "しま",
      "en": "island"
    },
    {
      "kanji": "市民",
      "kana": "しみん",
      "en": "citizen"
    },
    {
      "kanji": "事務所",
      "kana": "じむしょ",
      "en": "office"
    },
    {
      "kanji": "社会",
      "kana": "しゃかい",
      "en": "society"
    },
    {
      "kanji": "社長",
      "kana": "しゃちょう",
      "en": "company president"
    },
    {
      "kanji": "じゃま",
      "kana": "じゃま",
      "en": "intrusion, obstacle"
    },
    {
      "kanji": "ジャム",
      "kana": "じゃむ",
      "en": "jam"
    },
    {
      "kanji": "自由",
      "kana": "じゆう",
      "en": "freedom"
    },
    {
      "kanji": "習慣",
      "kana": "しゅうかん",
      "en": "custom, manners"
    },
    {
      "kanji": "住所",
      "kana": "じゅうしょ",
      "en": "an address, a residence"
    },
    {
      "kanji": "柔道",
      "kana": "じゅうどう",
      "en": "judo"
    },
    {
      "kanji": "十分",
      "kana": "じゅうぶん",
      "en": "enough"
    },
    {
      "kanji": "出席",
      "kana": "しゅっせき",
      "en": "attendance"
    },
    {
      "kanji": "出発",
      "kana": "しゅっぱつ",
      "en": "departure"
    },
    {
      "kanji": "趣味",
      "kana": "しゅみ",
      "en": "a hobby"
    },
    {
      "kanji": "準備",
      "kana": "じゅんび",
      "en": "preparation"
    },
    {
      "kanji": "紹介",
      "kana": "しょうかい",
      "en": "introduction"
    },
    {
      "kanji": "小学校",
      "kana": "しょうがっこう",
      "en": "elementary school"
    },
    {
      "kanji": "小説",
      "kana": "しょうせつ",
      "en": "novel"
    },
    {
      "kanji": "招待",
      "kana": "しょうたい",
      "en": "invitation"
    },
    {
      "kanji": "承知",
      "kana": "しょうち",
      "en": "consentment"
    },
    {
      "kanji": "将来",
      "kana": "しょうらい",
      "en": "future, prospects"
    },
    {
      "kanji": "食事",
      "kana": "しょくじ",
      "en": "a meal"
    },
    {
      "kanji": "食料品",
      "kana": "しょくりょうひん",
      "en": "groceries"
    },
    {
      "kanji": "女性",
      "kana": "じょせい",
      "en": "woman"
    },
    {
      "kanji": "知らせる",
      "kana": "しらせる",
      "en": "to notify"
    },
    {
      "kanji": "調べる",
      "kana": "しらべる",
      "en": "to investigate"
    },
    {
      "kanji": "人口",
      "kana": "じんこう",
      "en": "population"
    },
    {
      "kanji": "神社",
      "kana": "じんじゃ",
      "en": "Shinto shrine"
    },
    {
      "kanji": "親切",
      "kana": "しんせつ",
      "en": "kindness"
    },
    {
      "kanji": "心配",
      "kana": "しんぱい",
      "en": "worries"
    },
    {
      "kanji": "新聞社",
      "kana": "しんぶんしゃ",
      "en": "newspaper company"
    },
    {
      "kanji": "水泳",
      "kana": "すいえい",
      "en": "swimming"
    },
    {
      "kanji": "水道",
      "kana": "すいどう",
      "en": "water supply"
    },
    {
      "kanji": "ずいぶん",
      "kana": "ずいぶん",
      "en": "very, extremely"
    },
    {
      "kanji": "数学",
      "kana": "すうがく",
      "en": "mathematic, arithmetic"
    },
    {
      "kanji": "スーツ",
      "kana": "すーつ",
      "en": "suit"
    },
    {
      "kanji": "スーツケース",
      "kana": "すーつけーす",
      "en": "suitcase"
    },
    {
      "kanji": "過ぎる",
      "kana": "すぎる",
      "en": "to exceed"
    },
    {
      "kanji": "空く",
      "kana": "すく",
      "en": "to become empty"
    },
    {
      "kanji": "スクリーン",
      "kana": "すくりーん",
      "en": "screen"
    },
    {
      "kanji": "凄い",
      "kana": "すごい",
      "en": "terrific"
    },
    {
      "kanji": "進む",
      "kana": "すすむ",
      "en": "to make progress"
    },
    {
      "kanji": "すっかり",
      "kana": "すっかり",
      "en": "completely"
    },
    {
      "kanji": "すっと",
      "kana": "すっと",
      "en": "straight, all of a sudden"
    },
    {
      "kanji": "ステーキ",
      "kana": "すてーき",
      "en": "steak"
    },
    {
      "kanji": "捨てる",
      "kana": "すてる",
      "en": "to throw away"
    },
    {
      "kanji": "ステレオ",
      "kana": "すてれお",
      "en": "Stereo"
    },
    {
      "kanji": "砂",
      "kana": "すな",
      "en": "sand"
    },
    {
      "kanji": "素晴らしい",
      "kana": "すばらしい",
      "en": "wonderful"
    },
    {
      "kanji": "滑る",
      "kana": "すべる",
      "en": "to slide, to slip"
    },
    {
      "kanji": "隅",
      "kana": "すみ",
      "en": "corner, nook"
    },
    {
      "kanji": "済む",
      "kana": "すむ",
      "en": "to finish"
    },
    {
      "kanji": "すり",
      "kana": "すり",
      "en": "a pickpocket"
    },
    {
      "kanji": "すると",
      "kana": "すると",
      "en": "then"
    },
    {
      "kanji": "生活",
      "kana": "せいかつ",
      "en": "life, living"
    },
    {
      "kanji": "生産",
      "kana": "せいさん",
      "en": "production"
    },
    {
      "kanji": "政治",
      "kana": "政治",
      "en": "politics, government"
    },
    {
      "kanji": "西洋",
      "kana": "せいよう",
      "en": "western countries"
    },
    {
      "kanji": "世界",
      "kana": "せかい",
      "en": "the world"
    },
    {
      "kanji": "席",
      "kana": "せき",
      "en": "seat"
    },
    {
      "kanji": "説明",
      "kana": "せつめい",
      "en": "explanation"
    },
    {
      "kanji": "背中",
      "kana": "せなか",
      "en": "the back (of the body)"
    },
    {
      "kanji": "是非",
      "kana": "ぜひ",
      "en": "without fail"
    },
    {
      "kanji": "世話",
      "kana": "せわ",
      "en": "to look after"
    },
    {
      "kanji": "線",
      "kana": "せん",
      "en": "line"
    },
    {
      "kanji": "全然",
      "kana": "ぜんぜん",
      "en": "not at all (in negative sentence)"
    },
    {
      "kanji": "戦争",
      "kana": "せんそう",
      "en": "war"
    },
    {
      "kanji": "先輩",
      "kana": "せんぱい",
      "en": "senior"
    },
    {
      "kanji": "そう",
      "kana": "そう",
      "en": "really"
    },
    {
      "kanji": "相談",
      "kana": "そうだん",
      "en": "to discuss"
    },
    {
      "kanji": "育てる",
      "kana": "そだてる",
      "en": "to rear, to bring up"
    },
    {
      "kanji": "卒業",
      "kana": "そつぎょう",
      "en": "graduation"
    },
    {
      "kanji": "祖父",
      "kana": "そふ",
      "en": "grandfather"
    },
    {
      "kanji": "ソフト",
      "kana": "そふと",
      "en": "soft"
    },
    {
      "kanji": "祖母",
      "kana": "そぼ",
      "en": "grandmother"
    },
    {
      "kanji": "それで",
      "kana": "それで",
      "en": "because of that"
    },
    {
      "kanji": "それに",
      "kana": "それに",
      "en": "moreover"
    },
    {
      "kanji": "それほど",
      "kana": "それほど",
      "en": "to that extent"
    },
    {
      "kanji": "そろそろ",
      "kana": "そろそろ",
      "en": "gradually, soon"
    },
    {
      "kanji": "そんな",
      "kana": "そんな",
      "en": "that kind of"
    },
    {
      "kanji": "そんなに",
      "kana": "そんなに",
      "en": "so much, like that"
    },
    {
      "kanji": "退院",
      "kana": "たいいん",
      "en": "leave hospital"
    },
    {
      "kanji": "大学生",
      "kana": "だいがくせい",
      "en": "university student"
    },
    {
      "kanji": "大事",
      "kana": "だいじ",
      "en": "important, valuable"
    },
    {
      "kanji": "大体",
      "kana": "だいたい",
      "en": "generally"
    },
    {
      "kanji": "たいてい",
      "kana": "たいてい",
      "en": "usually"
    },
    {
      "kanji": "タイプ",
      "kana": "たいぷ",
      "en": "type, style"
    },
    {
      "kanji": "大分",
      "kana": "だいぶ",
      "en": "greatly"
    },
    {
      "kanji": "台風",
      "kana": "たいふう",
      "en": "typhoon"
    },
    {
      "kanji": "倒れる",
      "kana": "たおれる",
      "en": "to break down"
    },
    {
      "kanji": "だから",
      "kana": "だから",
      "en": "so, therefore"
    },
    {
      "kanji": "確か",
      "kana": "たしか",
      "en": "definite"
    },
    {
      "kanji": "足す",
      "kana": "たす",
      "en": "to add a number"
    },
    {
      "kanji": "訪ねる",
      "kana": "たずねる",
      "en": "to visit"
    },
    {
      "kanji": "尋ねる",
      "kana": "たずねる",
      "en": "to ask"
    },
    {
      "kanji": "正しい",
      "kana": "ただしい",
      "en": "correct"
    },
    {
      "kanji": "畳",
      "kana": "たたみ",
      "en": "a Tatami"
    },
    {
      "kanji": "立てる",
      "kana": "たてる",
      "en": "stand up, put up"
    },
    {
      "kanji": "建てる",
      "kana": "たてる",
      "en": "to build"
    },
    {
      "kanji": "例えば",
      "kana": "たとえば",
      "en": "for example"
    },
    {
      "kanji": "棚",
      "kana": "たな",
      "en": "shelves"
    },
    {
      "kanji": "楽しみ",
      "kana": "たのしみ",
      "en": "pleasure, amusement"
    },
    {
      "kanji": "楽しむ",
      "kana": "たのしむ",
      "en": "enjoy, enjoy oneself"
    },
    {
      "kanji": "たまに",
      "kana": "たまに",
      "en": "occasionally"
    },
    {
      "kanji": "為",
      "kana": "ため",
      "en": "in order to"
    },
    {
      "kanji": "駄目",
      "kana": "だめ",
      "en": "no good"
    },
    {
      "kanji": "足りる",
      "kana": "たりる",
      "en": "be sufficient"
    },
    {
      "kanji": "男性",
      "kana": "だんせい",
      "en": "man"
    },
    {
      "kanji": "暖房",
      "kana": "だんぼう",
      "en": "heating"
    },
    {
      "kanji": "血",
      "kana": "ち",
      "en": "blood"
    },
    {
      "kanji": "チェック",
      "kana": "ちぇっく",
      "en": "check"
    },
    {
      "kanji": "力",
      "kana": "ちから",
      "en": "strengh, power"
    },
    {
      "kanji": "ちっとも",
      "kana": "ちっとも",
      "en": "not at all (in negative sentence)"
    },
    {
      "kanji": "ちゃん",
      "kana": "ちゃん",
      "en": "familiar suffix usually for girls"
    },
    {
      "kanji": "注意",
      "kana": "ちゅうい",
      "en": "caution"
    },
    {
      "kanji": "中学校",
      "kana": "ちゅうがっこう",
      "en": "junior high school"
    },
    {
      "kanji": "注射",
      "kana": "ちゅうしゃ",
      "en": "injection"
    },
    {
      "kanji": "駐車場",
      "kana": "ちゅうしゃじょう",
      "en": "parking lot"
    },
    {
      "kanji": "地理",
      "kana": "ちり",
      "en": "geography"
    },
    {
      "kanji": "捕まえる",
      "kana": "つかまえる",
      "en": "to seize"
    },
    {
      "kanji": "月",
      "kana": "つき",
      "en": "moon"
    },
    {
      "kanji": "付く",
      "kana": "つく",
      "en": "to be attached"
    },
    {
      "kanji": "漬ける",
      "kana": "つける",
      "en": "to soak, to pickle"
    },
    {
      "kanji": "都合",
      "kana": "つごう",
      "en": "circumstances"
    },
    {
      "kanji": "伝える",
      "kana": "つたえる",
      "en": "to report"
    },
    {
      "kanji": "続く",
      "kana": "つづく",
      "en": "continue, last"
    },
    {
      "kanji": "続ける",
      "kana": "つづける",
      "en": "to continue, go on"
    },
    {
      "kanji": "包む",
      "kana": "つつむ",
      "en": "to wrap"
    },
    {
      "kanji": "妻",
      "kana": "つま",
      "en": "wife"
    },
    {
      "kanji": "つもり",
      "kana": "つもり",
      "en": "intention"
    },
    {
      "kanji": "釣る",
      "kana": "つる",
      "en": "to fish"
    },
    {
      "kanji": "連れる",
      "kana": "つれる",
      "en": "to lead"
    },
    {
      "kanji": "丁寧",
      "kana": "ていねい",
      "en": "polite"
    },
    {
      "kanji": "テキスト",
      "kana": "てきすと",
      "en": "text, text book"
    },
    {
      "kanji": "適当",
      "kana": "てきとう",
      "en": "suitable"
    },
    {
      "kanji": "できるだけ",
      "kana": "できるだけ",
      "en": "as much as possible"
    },
    {
      "kanji": "手伝う",
      "kana": "てつだう",
      "en": "to assist"
    },
    {
      "kanji": "テニス",
      "kana": "てにす",
      "en": "tenis"
    },
    {
      "kanji": "手袋",
      "kana": "てぶくろ",
      "en": "glove"
    },
    {
      "kanji": "寺",
      "kana": "てら",
      "en": "temple"
    },
    {
      "kanji": "点",
      "kana": "てん",
      "en": "point, dot"
    },
    {
      "kanji": "店員",
      "kana": "てんいん",
      "en": "shop assistant"
    },
    {
      "kanji": "天気予報",
      "kana": "てんきよほう",
      "en": "weather forecast"
    },
    {
      "kanji": "電灯",
      "kana": "でんとう",
      "en": "electric light"
    },
    {
      "kanji": "電報",
      "kana": "でんぽう",
      "en": "telegram"
    },
    {
      "kanji": "展覧会",
      "kana": "てんらんかい",
      "en": "exhibition"
    },
    {
      "kanji": "都",
      "kana": "と",
      "en": "metropolitan"
    },
    {
      "kanji": "道具",
      "kana": "どうぐ",
      "en": "tool, means"
    },
    {
      "kanji": "とうとう",
      "kana": "とうとう",
      "en": "finally, after all"
    },
    {
      "kanji": "動物園",
      "kana": "どうぶつえん",
      "en": "zoo"
    },
    {
      "kanji": "遠く",
      "kana": "とおく",
      "en": "distant"
    },
    {
      "kanji": "通る",
      "kana": "とおる",
      "en": "to go through"
    },
    {
      "kanji": "特に",
      "kana": "とくに",
      "en": "particularly, especially"
    },
    {
      "kanji": "特別",
      "kana": "とくべつ",
      "en": "special"
    },
    {
      "kanji": "床屋",
      "kana": "とこや",
      "en": "a barber shop"
    },
    {
      "kanji": "途中",
      "kana": "とちゅう",
      "en": "on the way"
    },
    {
      "kanji": "特急",
      "kana": "とっきゅう",
      "en": "limites express train"
    },
    {
      "kanji": "届ける",
      "kana": "とどける",
      "en": "to reach"
    },
    {
      "kanji": "泊まる",
      "kana": "とまる",
      "en": "to lodge at"
    },
    {
      "kanji": "止める",
      "kana": "とめる",
      "en": "to stop"
    },
    {
      "kanji": "取り替える",
      "kana": "とりかえる",
      "en": "to exchange"
    },
    {
      "kanji": "泥棒",
      "kana": "どろぼう",
      "en": "thief"
    },
    {
      "kanji": "どんどん",
      "kana": "どんどん",
      "en": "more and more"
    },
    {
      "kanji": "直す",
      "kana": "なおす",
      "en": "to fix, to repair"
    },
    {
      "kanji": "直る",
      "kana": "なおる",
      "en": "to be fixed, repaired"
    },
    {
      "kanji": "治る",
      "kana": "なおる",
      "en": "to be cured, to heal"
    },
    {
      "kanji": "中々",
      "kana": "なかなか",
      "en": "considerably"
    },
    {
      "kanji": "泣く",
      "kana": "なく",
      "en": "to cry"
    },
    {
      "kanji": "無くなる",
      "kana": "なくなる",
      "en": "to disappear, to get lost"
    },
    {
      "kanji": "亡くなる",
      "kana": "なくなる",
      "en": "to die"
    },
    {
      "kanji": "投げる",
      "kana": "なげる",
      "en": "to throw away"
    },
    {
      "kanji": "なさる",
      "kana": "なさる",
      "en": "to do (respectful)"
    },
    {
      "kanji": "鳴る",
      "kana": "なる",
      "en": "sound, ring"
    },
    {
      "kanji": "なるべく",
      "kana": "なるべく",
      "en": "as much as possible"
    },
    {
      "kanji": "なるほど",
      "kana": "なるほど",
      "en": "Indeed, really, I see"
    },
    {
      "kanji": "慣れる",
      "kana": "なれる",
      "en": "to grow accustomed to"
    },
    {
      "kanji": "匂い",
      "kana": "におい",
      "en": "a smell"
    },
    {
      "kanji": "苦い",
      "kana": "にがい",
      "en": "bitter"
    },
    {
      "kanji": "二階建て",
      "kana": "にかいだて",
      "en": "two storied"
    },
    {
      "kanji": "逃げる",
      "kana": "にげる",
      "en": "to escape"
    },
    {
      "kanji": "日記",
      "kana": "にっき",
      "en": "diary"
    },
    {
      "kanji": "入院",
      "kana": "にゅういん",
      "en": "be hospitalized"
    },
    {
      "kanji": "入学",
      "kana": "にゅうがく",
      "en": "enter school or university"
    },
    {
      "kanji": "似る",
      "kana": "にる",
      "en": "to be similar"
    },
    {
      "kanji": "人形",
      "kana": "にんぎょう",
      "en": "a doll"
    },
    {
      "kanji": "盗む",
      "kana": "ぬすむ",
      "en": "to steal"
    },
    {
      "kanji": "塗る",
      "kana": "ぬる",
      "en": "to paint, lacquer"
    },
    {
      "kanji": "濡れる",
      "kana": "ぬれる",
      "en": "to get wet"
    },
    {
      "kanji": "値段",
      "kana": "ねだん",
      "en": "price"
    },
    {
      "kanji": "熱",
      "kana": "ねつ",
      "en": "fever"
    },
    {
      "kanji": "熱心",
      "kana": "ねっしん",
      "en": "enthusiasm"
    },
    {
      "kanji": "寝坊",
      "kana": "ねぼう",
      "en": "sleeping in late"
    },
    {
      "kanji": "眠い",
      "kana": "ねむい",
      "en": "to be sleepy"
    },
    {
      "kanji": "眠る",
      "kana": "ねむる",
      "en": "to sleep"
    },
    {
      "kanji": "残る",
      "kana": "のこる",
      "en": "to remain"
    },
    {
      "kanji": "のど",
      "kana": "のど",
      "en": "throat"
    },
    {
      "kanji": "乗り換える",
      "kana": "のりかえる",
      "en": "to tansfer, change"
    },
    {
      "kanji": "乗り物",
      "kana": "のりもの",
      "en": "vehicle"
    },
    {
      "kanji": "葉",
      "kana": "は",
      "en": "leaf"
    },
    {
      "kanji": "場合",
      "kana": "ばあい",
      "en": "situation"
    },
    {
      "kanji": "パート",
      "kana": "ぱーと",
      "en": "part time"
    },
    {
      "kanji": "倍",
      "kana": "ばい",
      "en": "twice, double"
    },
    {
      "kanji": "拝見",
      "kana": "はいけん",
      "en": "to look at, find out (respectful)"
    },
    {
      "kanji": "歯医者",
      "kana": "はいしゃ",
      "en": "dentist"
    },
    {
      "kanji": "運ぶ",
      "kana": "はこぶ",
      "en": "to transport"
    },
    {
      "kanji": "始める",
      "kana": "はじめる",
      "en": "to begin"
    },
    {
      "kanji": "場所",
      "kana": "ばしょ",
      "en": "location"
    },
    {
      "kanji": "はず",
      "kana": "はず",
      "en": "it should be so"
    },
    {
      "kanji": "恥ずかしい",
      "kana": "はずかしい",
      "en": "embarrassed"
    },
    {
      "kanji": "パソコン",
      "kana": "ぱそこん",
      "en": "personal computer"
    },
    {
      "kanji": "発音",
      "kana": "はつおん",
      "en": "pronunciation"
    },
    {
      "kanji": "はっきり",
      "kana": "はっきり",
      "en": "clearly"
    },
    {
      "kanji": "花見",
      "kana": "はなみ",
      "en": "cherry-blossom viewing"
    },
    {
      "kanji": "林",
      "kana": "はやし",
      "en": "woods, forest"
    },
    {
      "kanji": "払う",
      "kana": "はらう",
      "en": "to pay"
    },
    {
      "kanji": "番組",
      "kana": "ばんぐみ",
      "en": "television or radio program"
    },
    {
      "kanji": "反対",
      "kana": "はんたい",
      "en": "Opposition"
    },
    {
      "kanji": "ハンドバッグ",
      "kana": "はんどばっぐ",
      "en": "Handbag"
    },
    {
      "kanji": "日",
      "kana": "ひ",
      "en": "day, sun"
    },
    {
      "kanji": "火",
      "kana": "ひ",
      "en": "fire"
    },
    {
      "kanji": "ピアノ",
      "kana": "ぴあの",
      "en": "piano"
    },
    {
      "kanji": "冷える",
      "kana": "ひえる",
      "en": "get cold, feel chilly"
    },
    {
      "kanji": "光",
      "kana": "ひかり",
      "en": "light"
    },
    {
      "kanji": "光る",
      "kana": "ひかる",
      "en": "to shine"
    },
    {
      "kanji": "引き出し",
      "kana": "ひきだし",
      "en": "drawer"
    },
    {
      "kanji": "引き出す",
      "kana": "ひきだす",
      "en": "to draw out"
    },
    {
      "kanji": "髭",
      "kana": "ひげ",
      "en": "a mustache, a beard"
    },
    {
      "kanji": "飛行場",
      "kana": "ひこうじょう",
      "en": "an airport"
    },
    {
      "kanji": "久しぶり",
      "kana": "ひさしぶり",
      "en": "after a long time"
    },
    {
      "kanji": "美術館",
      "kana": "びじゅつかん",
      "en": "art gallery"
    },
    {
      "kanji": "非常に",
      "kana": "ひじょうに",
      "en": "extremely"
    },
    {
      "kanji": "吃驚",
      "kana": "びっくり",
      "en": "be surprised"
    },
    {
      "kanji": "引っ越す",
      "kana": "ひっこす",
      "en": "to move (house)"
    },
    {
      "kanji": "必要",
      "kana": "ひつよう",
      "en": "necessary"
    },
    {
      "kanji": "酷い",
      "kana": "ひどい",
      "en": "awful"
    },
    {
      "kanji": "開く",
      "kana": "ひらく",
      "en": "to open an event"
    },
    {
      "kanji": "ビル",
      "kana": "びる",
      "en": "building"
    },
    {
      "kanji": "昼間",
      "kana": "ひるま",
      "en": "daytime, during the day"
    },
    {
      "kanji": "昼休み",
      "kana": "ひるやすみ",
      "en": "noon break"
    },
    {
      "kanji": "拾う",
      "kana": "ひろう",
      "en": "to pick up, to gather"
    },
    {
      "kanji": "ファックス",
      "kana": "ふぁっくす",
      "en": "fax"
    },
    {
      "kanji": "増える",
      "kana": "ふえる",
      "en": "to increase"
    },
    {
      "kanji": "深い",
      "kana": "ふかい",
      "en": "deep"
    },
    {
      "kanji": "複雑",
      "kana": "ふくざつ",
      "en": "complexity, complication"
    },
    {
      "kanji": "復習",
      "kana": "ふくしゅう",
      "en": "revision"
    },
    {
      "kanji": "部長",
      "kana": "ぶちょう",
      "en": "head of a section"
    },
    {
      "kanji": "普通",
      "kana": "ふつう",
      "en": "usually"
    },
    {
      "kanji": "葡萄",
      "kana": "ぶどう",
      "en": "a grape"
    },
    {
      "kanji": "太る",
      "kana": "ふとる",
      "en": "to grow fat"
    },
    {
      "kanji": "布団",
      "kana": "ふとん",
      "en": "Japanese bed, futon"
    },
    {
      "kanji": "船",
      "kana": "ふね",
      "en": "ship, boat"
    },
    {
      "kanji": "不便",
      "kana": "ふべん",
      "en": "inconvenience"
    },
    {
      "kanji": "踏む",
      "kana": "ふむ",
      "en": "to step on"
    },
    {
      "kanji": "降り出す",
      "kana": "ふりだす",
      "en": "start to rain"
    },
    {
      "kanji": "プレゼント",
      "kana": "ぷれぜんと",
      "en": "a present, a gift"
    },
    {
      "kanji": "文化",
      "kana": "ぶんか",
      "en": "culture"
    },
    {
      "kanji": "文学",
      "kana": "ぶんがく",
      "en": "litterature"
    },
    {
      "kanji": "文法",
      "kana": "ぶんぽう",
      "en": "grammar"
    },
    {
      "kanji": "別",
      "kana": "べつ",
      "en": "different"
    },
    {
      "kanji": "ベル",
      "kana": "べる",
      "en": "bell"
    },
    {
      "kanji": "変",
      "kana": "へん",
      "en": "strange"
    },
    {
      "kanji": "返事",
      "kana": "へんじ",
      "en": "reply"
    },
    {
      "kanji": "貿易",
      "kana": "ぼうえき",
      "en": "trade"
    },
    {
      "kanji": "放送",
      "kana": "ほうそう",
      "en": "to broadcast"
    },
    {
      "kanji": "法律",
      "kana": "ほうりつ",
      "en": "law"
    },
    {
      "kanji": "僕",
      "kana": "ぼく",
      "en": "I (for male)"
    },
    {
      "kanji": "星",
      "kana": "ほし",
      "en": "star"
    },
    {
      "kanji": "ほど",
      "kana": "ほど",
      "en": "extent"
    },
    {
      "kanji": "ほとんど",
      "kana": "ほとんど",
      "en": "mostly"
    },
    {
      "kanji": "褒める",
      "kana": "ほめる",
      "en": "praise, speak well of"
    },
    {
      "kanji": "翻訳",
      "kana": "ほんやく",
      "en": "translation"
    },
    {
      "kanji": "参る",
      "kana": "まいる",
      "en": "to go, to come (humble form)"
    },
    {
      "kanji": "負ける",
      "kana": "まける",
      "en": "to lose"
    },
    {
      "kanji": "真面目",
      "kana": "まじめ",
      "en": "serious"
    },
    {
      "kanji": "まず",
      "kana": "まず",
      "en": "first of all"
    },
    {
      "kanji": "または",
      "kana": "または",
      "en": "or, otherwise"
    },
    {
      "kanji": "間違える",
      "kana": "まちがえる",
      "en": "to make a mistake"
    },
    {
      "kanji": "間に合う",
      "kana": "まにあう",
      "en": "to be in time for"
    },
    {
      "kanji": "周り",
      "kana": "まわり",
      "en": "surroundings"
    },
    {
      "kanji": "回る",
      "kana": "まわる",
      "en": "to go around"
    },
    {
      "kanji": "漫画",
      "kana": "まんが",
      "en": "japanese comic, manga"
    },
    {
      "kanji": "真ん中",
      "kana": "まんなか",
      "en": "middle"
    },
    {
      "kanji": "見える",
      "kana": "みえる",
      "en": "be seen, be visible"
    },
    {
      "kanji": "湖",
      "kana": "みずうみ",
      "en": "a lake"
    },
    {
      "kanji": "味噌",
      "kana": "みそ",
      "en": "bean paste"
    },
    {
      "kanji": "見つかる",
      "kana": "みつかる",
      "en": "to be discovered"
    },
    {
      "kanji": "見つける",
      "kana": "みつける",
      "en": "to discover"
    },
    {
      "kanji": "皆",
      "kana": "みな",
      "en": "everybody"
    },
    {
      "kanji": "港",
      "kana": "みなと",
      "en": "harbour"
    },
    {
      "kanji": "向かう",
      "kana": "むかう",
      "en": "to face"
    },
    {
      "kanji": "迎える",
      "kana": "むかえる",
      "en": "to go to meet"
    },
    {
      "kanji": "昔",
      "kana": "むかし",
      "en": "olden days, former"
    },
    {
      "kanji": "虫",
      "kana": "むし",
      "en": "insect"
    },
    {
      "kanji": "息子",
      "kana": "むすこ",
      "en": "son"
    },
    {
      "kanji": "娘",
      "kana": "むすめ",
      "en": "daughter"
    },
    {
      "kanji": "無理",
      "kana": "むり",
      "en": "impossible"
    },
    {
      "kanji": "召し上がる",
      "kana": "めしあがる",
      "en": "to eat (respectful)"
    },
    {
      "kanji": "珍しい",
      "kana": "めずらしい",
      "en": "rare"
    },
    {
      "kanji": "申し上げる",
      "kana": "もうしあげる",
      "en": "to say, to tell (respectful)"
    },
    {
      "kanji": "申す",
      "kana": "もうす",
      "en": "to be called, to say (respectful)"
    },
    {
      "kanji": "もうすぐ",
      "kana": "もうすぐ",
      "en": "soon"
    },
    {
      "kanji": "もし",
      "kana": "もし",
      "en": "if"
    },
    {
      "kanji": "もちろん",
      "kana": "もちろん",
      "en": "of course"
    },
    {
      "kanji": "もっとも",
      "kana": "もっとも",
      "en": "extremely"
    },
    {
      "kanji": "戻る",
      "kana": "もどる",
      "en": "to return, to go back to"
    },
    {
      "kanji": "木綿",
      "kana": "もめん",
      "en": "cotton"
    },
    {
      "kanji": "貰う",
      "kana": "もらう",
      "en": "to receive"
    },
    {
      "kanji": "森",
      "kana": "もり",
      "en": "forest"
    },
    {
      "kanji": "焼く",
      "kana": "やく",
      "en": "to bake, to grill"
    },
    {
      "kanji": "約束",
      "kana": "やくそく",
      "en": "promise"
    },
    {
      "kanji": "役に立つ",
      "kana": "やくにたつ",
      "en": "to be helpful"
    },
    {
      "kanji": "焼ける",
      "kana": "やける",
      "en": "to burn, to be roasted"
    },
    {
      "kanji": "優しい",
      "kana": "やさしい",
      "en": "kind"
    },
    {
      "kanji": "痩せる",
      "kana": "やせる",
      "en": "to become thin"
    },
    {
      "kanji": "やっと",
      "kana": "やっと",
      "en": "at last"
    },
    {
      "kanji": "やはり/やっぱり",
      "kana": "やはり/やっぱり",
      "en": "as I thought"
    },
    {
      "kanji": "止む",
      "kana": "やむ",
      "en": "to stop"
    },
    {
      "kanji": "止める",
      "kana": "やめる",
      "en": "to stop"
    },
    {
      "kanji": "柔らかい",
      "kana": "やわらかい",
      "en": "soft, tender"
    },
    {
      "kanji": "湯",
      "kana": "ゆ",
      "en": "hot water"
    },
    {
      "kanji": "輸出",
      "kana": "ゆしゅつ",
      "en": "export"
    },
    {
      "kanji": "輸入",
      "kana": "ゆにゅう",
      "en": "import"
    },
    {
      "kanji": "指",
      "kana": "ゆび",
      "en": "finger"
    },
    {
      "kanji": "指輪",
      "kana": "ゆびわ",
      "en": "a ring"
    },
    {
      "kanji": "夢",
      "kana": "ゆめ",
      "en": "dream"
    },
    {
      "kanji": "揺れる",
      "kana": "ゆれる",
      "en": "to shake, to sway"
    },
    {
      "kanji": "用",
      "kana": "よう",
      "en": "use"
    },
    {
      "kanji": "用意",
      "kana": "ようい",
      "en": "preparation"
    },
    {
      "kanji": "用事",
      "kana": "ようじ",
      "en": "things to do"
    },
    {
      "kanji": "汚れる",
      "kana": "よごれる",
      "en": "to get dirty"
    },
    {
      "kanji": "予習",
      "kana": "よしゅう",
      "en": "preparation of a lesson"
    },
    {
      "kanji": "予定",
      "kana": "よてい",
      "en": "arrangement"
    }
  ]
}
`

func Dumb(w http.ResponseWriter, r *http.Request) {
	return
}
