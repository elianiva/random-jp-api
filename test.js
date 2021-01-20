const a = {
  abc: "1",
  abc: "1",
  abc: "1",
  abc: "1",
  abc: "1",
}

const func = obj => {
  console.log(obj)
}

const c = (...obj) => console.log([...obj])

func({
    b: "thing",
    c: "thing",
    d: ["stuff"],
  })

c(1, 2, 3, 4)

const arr = [1, 2, 3]

try {

} catch(err) {
  console.log(err)
} finally {

}

arr
.map(x => x)
.filter(x => x != 1)
.join("\n")
.split("\n")
.pop()
.unshift()
