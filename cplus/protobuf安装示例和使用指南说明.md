# protobuf 3.1.0 安装示例和使用指南说明

说明: 本文例子见： https://github.com/githublefantian/protobuf-3.1.0example.git

1、protobuf是google公司提出的数据存储格式，详细介绍可以参考：https://code.google.com/p/protobuf/

2、下载最新的protobuf源代码，下载地址：

https://github.com/google/protobuf/archive/v3.1.0.tar.gz

3、下载protobuf-3.1.0版本，protobuf-3.1.0.tar.gz解压并进行安装。
解压：tar xvzf protobuf-3.1.0.tar.gz
安装步骤：
（0）./autogen.sh
（1）./configure? --prefix=/usr/local/protobuf
（2）make
（3）make check
（4）make install

注意：
安装成功后，将它的bin和lib目录分别加入到PATH和LD_LIBRARY_PATH环境变量，以方便直接调用。

设置环境变量过程：编辑/etc/profile，在文件末尾添加：
export PATH=$PATH:/usr/local/protobuf/bin
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/protobuf/lib
此时最好重启一次。

4、测试例子：
1）创建一个.proto文件，例子中文件为：lm.helloworld.proto

```
syntax = "proto3";

package lm;
message helloworld
{
    required int32   id = 1;  // ID
    required string  str = 2; // str
    optional int32   opt = 3; //optional field
}
```
编译成目标语言： protoc -I=./ --cpp_out=./?? lm.helloworld.proto

2）编写 c++文件：reader.cc? writer.cc

```
/*reader.cc*/读文件：
#include "lm.helloworld.pb.h"
#include <iostream>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <string>
#include <pthread.h>
#include <fstream>

using namespace std;

void ListMsg(const lm::helloworld & msg) {
	cout << msg.id() << endl;
	cout << msg.str() << endl;
}

int main(int argc, char* argv[]) {
lm::helloworld msg1;
{
 fstream input("./log", ios::in | ios::binary);
 if (!msg1.ParseFromIstream(&input)) {
     cerr << "Failed to parse address book." << endl;
	 return -1;
 }
}

ListMsg(msg1);
}


/*writer.cc*/ 写文件：
#include "lm.helloworld.pb.h"
#include <iostream>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <string>
#include <pthread.h>
#include <fstream>

using namespace std;


int main(void)
{
 lm::helloworld msg1;
 msg1.set_id(101);
 msg1.set_str("hello");

 // Write the new address book back to disk.
 fstream output("./log", ios::out | ios::trunc | ios::binary);

 if (!msg1.SerializeToOstream(&output)) {
     cerr << "Failed to write msg." << endl;
	 return -1;
 }
 return 0;
}
```


5、编译程序：C++语言编译命令添加后缀如下：

注意：编译 是要指定protobuf头文件位置和protobuf的库文件位置? 例如 -I/usr/local/protobuf/include ?-L/usr/local/protobuf/lib -lprotobuf -pthread

g++ reader.cc lm.helloworld.pb.cc -o reader -I/usr/local/protobuf/include -L/usr/local/protobuf/lib -lprotobuf
g++ writer.cc lm.helloworld.pb.cc -o writer -I/usr/local/protobuf/include -L/usr/local/protobuf/lib -lprotobuf


6、运行结果

运行 Writer 和 Reader 的结果如下：
>writer
>reader
101
Hello


