#include <iostream>
#include <stdio.h>
#include <time.h>
#include <chrono>

static uint64_t getTick() {
  std::chrono::time_point<std::chrono::system_clock, std::chrono::milliseconds> tp =
  std::chrono::time_point_cast<std::chrono::milliseconds>(std::chrono::system_clock::now());
  return tp.time_since_epoch().count();
}


int main() 
{
#ifdef __linux
	// linux only
	std::cout << "=== linux only time analysis ===" << std::endl;

	struct timespec timestamp = { 0, 0 };

	clock_gettime(CLOCK_REALTIME, &timestamp);
	uint64_t now_epoch_time = timestamp.tv_sec * 1000000000 + timestamp.tv_nsec;
	std::cout << "current epoch time: " << now_epoch_time << " ns" << std::endl;

	clock_gettime(CLOCK_MONOTONIC, &timestamp);
	uint64_t machine_running_duration = timestamp.tv_sec * 1000000000 + timestamp.tv_nsec;
	std::cout << "machine running to now duration: " << machine_running_duration << " ns" << std::endl;

	uint64_t launch_epoch_time = now_epoch_time - machine_running_duration;
	std::cout << "machine launch epoch time: " << launch_epoch_time << " ns" << std::endl;
#endif

	// cross platform
	std::cout << "=== cross platform time analysis ===" << std::endl;

	auto now = std::chrono::system_clock::now(); // system_clock can get the current timestamp, accuracy to 1 ns in linux and 100 ns in win
	long long now_timestamp = std::chrono::duration_cast<std::chrono::nanoseconds>(now.time_since_epoch()).count(); // milliseconds, microseconds, nanoseconds, all are ok
	std::cout << "current epoch time: " << now_timestamp << " ns" << std::endl;

	long long duration = std::chrono::steady_clock::now().time_since_epoch().count(); // steady_clock can get maching running to now duration, accuracy to 1 ns
	std::cout << "machine running to now duration: " << duration << " ns" << std::endl;

	uint64_t launch_timestamp = now_timestamp - duration;
	std::cout << "machine launch epoch time: " << launch_timestamp << " ns" << std::endl;

	return 0;
}
