// combine 3 arrays to an array of object
// useful to convert scraped data

const fs = require("fs")

const kanji = fs.readFileSync("./thing").toString().split("\n")
const kana = fs.readFileSync("./thing2").toString().split("\n")
const en = fs.readFileSync("./thing3").toString().split("\n")

const thing = kanji
  .map((x, i) => ({
    kanji: x.trim(),
    kana: kana[i].trim(),
    en: en[i].trim(),
  }))
  .filter(x => x.kana != "-------------------")
  .filter(x => x.kana != "" && x.kanji != "" && x.kana != "")

fs.writeFileSync(
  "./result",
  JSON.stringify(
    {
      data: thing,
    },
    null,
    2
  )
)
