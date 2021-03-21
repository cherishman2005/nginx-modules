#include <stdlib.h>
#include <iostream>
#include <vector>
#include <string>
#include <set>
#include <iterator>
using namespace std;

typedef unsigned int uint32_t;

static void remove_space(string& str){
	string buff(str); 
	char space = ' '; 
	str.assign(buff.begin() + buff.find_first_not_of(space), 
		buff.begin() + buff.find_last_not_of(space) + 1); 
} 

static vector<string> tokenize_str_trim(const string & str,
                            const string & delims)
{
    // Skip delims at beginning, find start of first token
    string::size_type lastPos = str.find_first_not_of(delims, 0);
    // Find next delimiter @ end of token
    string::size_type pos     = str.find_first_of(delims, lastPos);

    // output vector
    vector<string> tokens;

    while (string::npos != pos || string::npos != lastPos)
    {
        // Found a token, add it to the vector.
        string partStr = str.substr(lastPos, pos - lastPos);
        remove_space( partStr );
        tokens.push_back( partStr );
        // Skip delims.  Note the "not_of". this is beginning of token
        lastPos = str.find_first_not_of(delims, pos);
        // Find next delimiter at end of token.
        pos     = str.find_first_of(delims, lastPos);
    }

    return tokens;
}

class GraySidByTail
{
public:
    //
    // function
    //
    GraySidByTail() : m_mod(10){};
    ~GraySidByTail() {};
    // 以逗号为分隔符的尾号串，做初始化
    void Init(const std::string& strSids, const std::string& strTails, const std::string& strMod);
    bool Match(uint32_t sid);
    //std::string ShowTails();

private:
    // 灰度发布
    //uint32_t m_grayEnable;
    std::set<uint32_t> m_sids;
    std::set<uint32_t> m_sidTails;
    uint32_t m_mod;
};

void GraySidByTail::Init(const std::string& strSids, const std::string& strTails, const std::string& strMod)
{
    if (strSids.size() > 0) {
        const std::vector<std::string> sids = tokenize_str_trim(strSids, ",");
        for (std::vector<std::string>::const_iterator it = sids.begin(); it != sids.end(); it++) {
            uint32_t id = strtoul((*it).c_str(), 0, 0);
            m_sids.insert(id);
        }
    }
    
    if (strTails.size() > 0) {
        const std::vector<std::string> tails = tokenize_str_trim(strTails, ",");
        for (std::vector<std::string>::const_iterator it = tails.begin(); it != tails.end(); it++) {
            uint32_t id = strtoul((*it).c_str(), 0, 0);
            m_sidTails.insert(id);
        }
    }

    if (strMod.size() > 0) {
        m_mod = strtoul(strMod.c_str(), 0, 0);
    }
    
    cout << "m_mod=" << m_mod << endl;
    copy(m_sids.begin(), m_sids.end(), ostream_iterator<uint32_t>(cout, ","));
    cout << endl;
    
    copy(m_sidTails.begin(), m_sidTails.end(), ostream_iterator<uint32_t>(cout, ","));
    cout << endl;
}

bool GraySidByTail::Match(uint32_t sid)
{
    if (m_sids.find(sid) != m_sids.end()) {
        return true;
    }
    
    if (m_sidTails.find(sid%m_mod) != m_sidTails.end()) {
        return true;
    }
    
    return false;
}

int main() 
{
    GraySidByTail *gray = new GraySidByTail();
   
    std::string strSids = "111111, 222333, 777888";
    std::string strTails = "6,     0";
    std::string strMod = "10";
    gray->Init(strSids, strTails, strMod);
    
    cout << gray->Match(777888) << endl;
    cout << gray->Match(7778881) << endl;
    cout << gray->Match(55556) << endl;
    
    return 0;
}