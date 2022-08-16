# ipaddr

```
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
