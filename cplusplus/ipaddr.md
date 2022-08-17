# ipaddr

```
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>

inline bool   valid_addr(u_long ip)             { return ip != INADDR_NONE; }
inline u_long aton_addr(const char * ip)        { return ::inet_addr(ip); }
inline u_long aton_addr(const std::string & ip) { return aton_addr(ip.c_str()); }

inline std::string addr_ntoa(u_long ip)
{ 
	struct in_addr addr;
	memcpy(&addr, &ip, 4);
	return std::string(::inet_ntoa(addr)); 
}
```

## ipsegment

![image](https://user-images.githubusercontent.com/17688273/185039788-93cdb7c8-489e-4e61-885e-30abe7bf53c9.png)


# 参考链接

ip segment计算
- [https://jodies.de/ipcalc](https://jodies.de/ipcalc)
