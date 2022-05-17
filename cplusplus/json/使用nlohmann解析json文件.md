# [C++]使用nlohmann解析json文件

nlohmann 是德国工程师，以其名字为工程名的 nlohmann/json 项目又被成为 JSON for Modern C++。
网上常见如何使用 nlohmann 生成 JSON 文件的中英文资料。但如何使用 nlohmann 解析 JSON 文件的 资料 不多，或者，不够清楚直接。
其实，工程的 README.md 写得也算清楚。但是对于从未接触过 JSON 文件的新手来说，还是不太友好。此篇主要向未接触过 JSON 文件的新手介绍如何快速使用 nlohmann 解析 JSON 文件。

工程引用
nlohmann 使用了 hpp，这种形式被称为 header only。在自己的工程中使用 nlohmann 时，仅需 #include “json.hpp” 即可。在所有 #include 之后，再添加一句 using json = nlohmann::json 会更加便捷。

## 简单文件解析

如果 JSON 文件够简单，如下，
```
{
	"pi":3.1415,
	"happy":true
}
```

不包括任何结构体或数组，则解析代码如下：
```
#include "json.hpp"
#include <fstream>
#include <iostream>
using namespace std;
using json = nlohmann::json;
int main() {
	json j;			// 创建 json 对象
	ifstream jfile("test.json");
	jfile >> j;		// 以文件流形式读取 json 文件
	float pi = j.at("pi");
	bool happy = j.at("happy");
	return 0;
}
```

## 较复杂的 JSON 文件解析

一般来说，JSON 文件包含很多的层级，转换成 C++ 则是结构体。以下 json 文件为例：
```
{
  "output": {
    "width": 720,
    "height": 1080,
    "frameRate": 20,
    "crf": 31
  },
  "tracks": [
    {
      "name": "t1",
      "pieces": [
        {
          "file": "x.mp4",
          "startTime": 2,
          "endTime": 6
        },
        {
          "file": "y.mp4",
          "startTime": 9,
          "endTime": 13
        }
      ]
    },
    {
      "name": "t2",
      "pieces": [
        {
          "file": "z.mp4",
          "startTime": 0,
          "endTime": 10
        }
      ]
    }
  ]
}
```

从这个 json 文件，人脑解析的话，我们可以看到，它包括两大部分 "output" 和 "tracks"。其中，output 可以用结构体来表示，其代码如下：
```
struct video_info {
	int width;
	int height;
	int frame_rate;
	int crf;
}
```

而另一部分可看作是有两个元素的结构体 "tracks" 的数组，其结构体包括 string 的 name，和另一个结构体 "pieces" 的数组。用 C++ 代码可表示为如下：

```
    struct pieceinfo {
        string  pathname;
        int     startTime;
        int     endTime;
    };
    struct trackinfo {
        string      name;
        pieceinfo   pieces[10];
        int         size;   // piceces 大小
    };
```

为了解析结构体类内容，需要定义对结构体的解析方法。因此，整个解析过程如下：

```
#include "json.hpp"
#include <iostream>
#include <fstream>

using namespace std;
using json = nlohmann::json;

namespace jsonns {
    struct videoinfo {
        int width;
        int height;
        int frameRate;
        int crf;
    };

    void from_json(const json& j, videoinfo& v) {
        j.at("width").get_to(v.width);
        j.at("height").get_to(v.height);
        j.at("frameRate").get_to(v.frameRate);
        j.at("crf").get_to(v.crf);
    }

    struct pieceinfo {
        string  pathname;
        int     startTime;
        int     endTime;
    };
    
    void from_json(const json&j, pieceinfo &p) {
        j.at("file").get_to(p.pathname);
        j.at("startTime").get_to(p.startTime);
        j.at("endTime").get_to(p.endTime);
    }

    struct trackinfo {
        string      name;
        pieceinfo   pieces[10];
        int         size;
    };

    void from_json(const json&j, trackinfo &t) {
        j.at("name").get_to(t.name);
        for(int i = 0; i < j["pieces"].size(); i++) {
            t.pieces[i] = j["pieces"][i];
        }
        t.size = j["pieces"].size();
    }
}

int main()
{
    json j;
    ifstream jfile("test.json");
    jfile >> j;
    jsonns::videoinfo vi = j.at("output");
    int tilength = j["tracks"].size();
    jsonns::trackinfo ti[tilength];
    for (int i = 0; i < tilength; i++) {
        ti[i] = j["tracks"][i];
    }
    return 0;
}
```
