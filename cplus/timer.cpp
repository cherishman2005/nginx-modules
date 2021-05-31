// https://currentmillis.com/

#include <iostream>
#include <chrono>
using namespace std;


static uint64_t getTickCount() {
    std::chrono::time_point<std::chrono::system_clock, std::chrono::milliseconds> tp =
    std::chrono::time_point_cast<std::chrono::milliseconds>(std::chrono::system_clock::now());
    return tp.time_since_epoch().count();
}

int main()
{
	//定义毫秒级别的时钟类型
	typedef chrono::time_point<chrono::system_clock, chrono::milliseconds> microClock_type;
	//获取当前时间点，windows system_clock是100纳秒级别的(不同系统不一样，自己按照介绍的方法测试)，所以要转换
	microClock_type tp = chrono::time_point_cast<chrono::milliseconds>(chrono::system_clock::now());
	//计算距离1970-1-1,00:00的时间长度，因为当前时间点定义的精度为毫秒，所以输出的是毫秒
	cout << "to 1970-1-1,00:00  " << tp.time_since_epoch().count() << "ms" << endl;
    cout << "tick=" << getTickCount() << endl;
    
	return 0;
}