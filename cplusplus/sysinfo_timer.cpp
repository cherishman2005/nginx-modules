#include <iostream>
#include <stdio.h>
#include <time.h>
#include <chrono>

static uint64_t uptime() {
	// cross platform
	std::cout << "=== cross platform time analysis ===" << std::endl;

	auto now = std::chrono::system_clock::now(); // system_clock can get the current timestamp, accuracy to 1 ns in linux and 100 ns in win
	uint64_t now_timestamp = std::chrono::duration_cast<std::chrono::nanoseconds>(now.time_since_epoch()).count(); // milliseconds, microseconds, nanoseconds, all are ok
	std::cout << "current epoch time: " << now_timestamp << " ns" << std::endl;

	uint64_t duration = std::chrono::steady_clock::now().time_since_epoch().count(); // steady_clock can get maching running to now duration, accuracy to 1 ns
	std::cout << "machine running to now duration: " << duration << " ns" << std::endl;

	uint64_t launch_timestamp = now_timestamp - duration;
	std::cout << "machine launch epoch time: " << launch_timestamp << " ns" << std::endl;
    return launch_timestamp/1000000;
}

int main() 
{
    std::cout << "uptime: " << uptime() << " ms" << std::endl;
	return 0;
}
