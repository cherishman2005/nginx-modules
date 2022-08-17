#include <iostream>
#include <string>
#include <set>
#include <fstream>
#include <sstream>
#include <vector>
#include <stdio.h>
 #include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>

using namespace std;

inline bool   valid_addr(u_long ip)             { return ip != INADDR_NONE; }
inline u_long aton_addr(const char * ip)        { return ::inet_addr(ip); }
inline u_long aton_addr(const std::string & ip) { return aton_addr(ip.c_str()); }

inline std::string addr_ntoa(u_long ip)
{
    struct in_addr addr;
    memcpy(&addr, &ip, 4);
    return std::string(::inet_ntoa(addr)); 
}

std::vector<std::string> split_string(std::string src, std::string sign)
{
    size_t posBeg = 0;
    size_t posEnd = 0;
    vector<string> vecRes;
    posEnd = src.find(sign, posBeg);
    while (posEnd != string::npos)
    {
        vecRes.push_back(src.substr(posBeg, posEnd - posBeg));
        posBeg = posEnd + sign.size();
        posEnd = src.find(sign, posBeg);
    }
    vecRes.push_back(src.substr(posBeg));
    return vecRes;
}

vector<std::string> split(std::string s, std::string delimiter)
{
    size_t pos_start = 0, pos_end, delim_len = delimiter.length();
    string token;
    vector<string> res;

    while ((pos_end = s.find (delimiter, pos_start)) != string::npos) {
        token = s.substr(pos_start, pos_end - pos_start);
        pos_start = pos_end + delim_len;
        res.push_back(token);
    }

    res.push_back(s.substr(pos_start));
    return res;
}

static string trim(const string &str)
{
    int first = 0;
    int last = 0;

    if (str.size() == 0)
        return str;

    int i;
    for (i = 0; i < (int)(str.size()); ++i)
    {
        if ((unsigned char)str[i] > 0x20)
            break;
    }

    first = i;
    for (i = str.size() - 1; i >= 0; --i)
    {
        if ((unsigned char)str[i] > 0x20)
            break;
    }

    last = i;

    if (last < first)
        return string("");
    if (first == 0 && last == (int)(str.size() - 1))
        return str;

    return str.substr(first, last - first + 1);
}

class IpSegment {
public:

    IpSegment(): m_address(0), m_type(0), m_mask(0xFFFFFFFF)
    {
    }
    
    IpSegment(uint32_t addr, uint32_t type)
    {
        m_address = addr;
        m_type = type;
        m_mask = 0xFFFFFFFF << (32 - type);
    }
    IpSegment(const string & addr, uint32_t type)
    {
        uint32_t ip = aton_addr(addr);
        //m_address = htonl(ip);
        m_address = ip;
        m_type = type;
        m_mask = 0xFFFFFFFF << (32 - type);
    }
    
    bool isInRange(const string & ip) const {
        uint32_t _ip = aton_addr(ip);
        _ip = htonl(_ip);
        uint32_t addr = htonl(m_address);
        printf("mask=0x%08lx, _ip=0x%08lx, segment ip=0x%08lx\n", m_mask, _ip, addr);
        return (_ip & m_mask) == (addr & m_mask);
    }
    
    
    bool operator < (const IpSegment& right) const
    {
        return (m_address < right.m_address) ||
            ((m_address == right.m_address) && (m_mask < right.m_mask));
    }

    bool operator == (const IpSegment& right) const
    {
        return (m_address == right.m_address) && (m_mask == right.m_mask);
    }
    
    bool operator != (const IpSegment& right) const
    {
        return (m_address != right.m_address) || (m_mask != right.m_mask);
    }
    
    string dump() const {
        //uint32_t addr = ntohl(m_address);
        string ip = addr_ntoa(m_address);
        ostringstream os;
        os << ip << "/" << m_type;
        return os.str();
    }

private:
    //ip地址（host little-endian）
    uint32_t m_address;
    
    uint32_t m_type;
    //子网掩码
    uint32_t m_mask;
};

set<uint32_t> m_lips;
set<IpSegment> m_ipsegments;

bool read_whitelist()
{
    const char *path = "ipwhite.txt";
    FILE *pfile = fopen(path, "r");
    if (pfile == NULL)
    {
        cout << "read fail, path: " << path << endl;
        return false;
    }

    char buff[1024] = {0};
    while (fgets(buff, 256, pfile) != NULL)
    {
        string tmp = trim(buff);
        if (0 == tmp.size())
        {
            continue;
        }

        vector<std::string> vec = split(tmp, "/");
        if (vec.size() == 2) {
            cout << "segment=" << tmp << endl;
            IpSegment ipsegment(vec[0], uint32_t(::atoi(vec[1].c_str())));
            m_ipsegments.insert(ipsegment);
        } else if (vec.size() == 1) {
            cout << "ip=" << tmp << endl;
            uint32_t ip = aton_addr(tmp);
            //m_lips.insert(htonl(ip));
            m_lips.insert(ip);
        }
    }
    fclose(pfile);

    std::ostringstream os;
    for (std::set<uint32_t>::const_iterator it = m_lips.begin(); it != m_lips.end(); it++) {
        os << ""<< *it << " ";
        printf("0x%08lx\n", *it);
    }

    cout << "IPLIST load from file success, total: " << m_lips.size() << " ip:[" << os.str() << "]" << endl;

    return true;
}


int main()
{
    string s = "192.168.158.1/24";
    vector<string> tmp = split(s, "/");
    size_t size = tmp.size();
    if (size != 2) {
        cout << "not segment" << endl;
        return 0;
    }
    
    IpSegment ipsegment(tmp[0], uint32_t(::atoi(tmp[1].c_str())));
    
    std::string ip = "192.168.158.49";
    bool f = ipsegment.isInRange(ip);
    cout << "isInRange=" << f << endl;
    
    read_whitelist();
    
    for (std::set<IpSegment>::const_iterator it = m_ipsegments.begin(); it != m_ipsegments.end(); it++) {
        cout << "ip=" << ip << ", ip segment=" << it->dump() << ", isInRange=" << it->isInRange(ip) << endl;
    }
    
    return 0;
}
