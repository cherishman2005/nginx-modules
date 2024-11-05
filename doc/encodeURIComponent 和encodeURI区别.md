# encodeURIComponent 和encodeURI区别

## 编码范围不同

* encodeURI：主要用于对整个 URL 进行编码，但它对一些在 URL 中有特殊含义的字符不进行编码，例如冒号（:）、斜杠（/）、问号（?）和井号（#）。因为这些字符是用于分隔 URL 的不同部分，如协议、路径、查询参数和锚点等。如果对这些字符进行编码，URL 可能会失去其原有的结构和语义，导致无法正确解析。
  * 示例：
    encodeURI("http://example.com/path?query=value#anchor")会返回原始的 URL 字符串，因为其中的特殊字符是 URL 正常结构的一部分，不需要编码。

* encodeURIComponent：用于对 URL 的组成部分（如查询参数的值）进行编码，它会对更多的字符进行编码，几乎包括了所有非字母数字字符。这是因为在 URL 的组成部分（如查询字符串的值）中，这些字符可能会引起问题，需要进行转义。
  * 示例：
    encodeURIComponent("query=value&param=other value")会对&、空格等字符进行编码，因为在查询字符串内部，这些字符有特殊的含义，需要正确处理以避免解析错误。

## 应用场景不同

* encodeURI：

适合对完整的 URL 进行编码，特别是当你不确定 URL 是否包含需要特殊处理的字符时。例如，当你需要在 JavaScript 中动态生成一个完整的 URL 链接，并且希望保留 URL 的基本结构，就可以使用encodeURI。不过要注意，它不能用于对包含在 URL 中的用户输入或数据进行编码，因为它不会对所有可能引起问题的字符进行编码。
示例：
假设你有一个函数用于根据用户输入的网站名称和路径来生成 URL，并且希望这个 URL 可以直接在浏览器中使用，就可以使用encodeURI。
```javascript
function generateURL(siteName, path) {
  return encodeURI("http://" + siteName + "/" + path);
}
```

* encodeURIComponent：

常用于对 URL 中的参数值进行编码，尤其是在构建包含用户输入的查询字符串时。例如，当你从用户输入获取搜索关键词，然后将其添加到查询参数中时，必须使用encodeURIComponent来确保关键词中的特殊字符不会破坏查询字符串的语法。
示例：
假设你有一个搜索功能，需要将用户输入的关键词添加到 URL 的查询参数中，就可以使用encodeURIComponent。
```javascript
function buildSearchURL(keyword) {
  return "http://example.com/search?keyword=" + encodeURIComponent(keyword);
}
```

* 编码结果不同

  * encodeURI：
    由于它对部分特殊字符不编码，所以编码后的结果看起来更接近原始的 URL 结构。例如，一个简单的 URL http://example.com/路径（其中 “路径” 是中文）经过encodeURI编码后可能是http://example.com/%E8%B7%AF%E5%BE%84，其中中文部分被编码，但 URL 的基本结构字符未变。
  * encodeURIComponent：
    编码后的结果通常会包含更多的编码字符，因为它对更广泛的字符进行编码。例如，对于同样的中文 “路径”，使用encodeURIComponent编码后可能是%E8%B7%AF%E5%BE%84，并且如果将其用于构建 URL 的查询参数，可能会与其他已经编码的字符一起出现在?之后的部分，如http://example.com/search?param=%E8%B7%AF%E5%BE%84。
