#include <stdlib.h>
#include <iostream>
#include <vector>
#include <string>
#include <set>
using namespace std;

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

int main() 
{
    std::set<int> graySidTails;
    std::string strGraySidTails = "1, 2 ,     40, 7";
    if (strGraySidTails.size() > 0) {
        const std::vector<std::string> tails = tokenize_str_trim(strGraySidTails, ",");
        for (std::vector<std::string>::iterator it = tails.begin(); it != tails.end(); it++) {
            int id = strtoul((*it).c_str(), 0, 0);
            graySidTails.insert(id);
        }
   }
   
   
   for (std::set<int>::iterator it = graySidTails.begin(); it != graySidTails.end(); it++) {
         cout << *it << endl;
    }
    
    return 0;
}