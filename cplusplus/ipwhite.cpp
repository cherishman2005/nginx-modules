#include<iostream>
#include<string>
#include<fstream>
using namespace std;

int main()
{
    // 从文件读入多行数据 方法一 正确读取
    cout<<"使用 << 读取数据"<<endl;
    ifstream infile("ipwhite.txt", ios::in);
    if (!infile.fail())
    {
        while (!infile.eof())
        {
            string str5;
            infile >> str5;
            cout << str5 <<endl;
        }
    }
    infile.close();
    //从文件中读入多行数据 方法二 错误读取
    cout<<"使用 getline 读取"<<endl;
    ifstream infile2("ipwhite.txt", ios::in);
    if (!infile2.fail())
    {
        while (!infile2.eof())
        {
            string str6;
            getline(infile2, str6);
            cout << str6 << endl;
        }
    }
    infile2.close();
    
    return 0;
}