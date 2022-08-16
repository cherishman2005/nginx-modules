#include<iostream>
#include<string>
#include<fstream>
#include<sstream>
using namespace std;

set<std::string> m_lips;

int main()
{

    ifstream infile("ipwhite.txt", ios::in);
    if (!infile.fail())
    {
        while (!infile.eof())
        {
            string str;
            infile >> str;
            cout << str <<endl;
            m_lips.insert(str);
        }
    }
    infile.close();
    
    ostringstream os;
    for (std::set<std::string>::const_iterator it = m_lips.begin(); it != m_lips.end(); it++) {
        os << "\"" << *it << "\" ";
    }
    cout << os.str() << endl;
    
    return 0;
}