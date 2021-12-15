#include<iostream>

using namespace std;

int main()
{
    //保存输入图像文件名和输出图像文件名
    char InImgName[10];
    char OutImgName[10];
    //图像数据长度
    int length;
    //文件指针
    FILE* fp;
    //输入要读取的图像名
    cout << "Enter Image name:";
    cin >> InImgName;
    //以二进制方式打开图像
    if ((fp = fopen(InImgName, "rb"))==NULL )
    {
        cout<<"Open image failed!"<<endl;
        exit(0);
    }
    //获取图像数据总长度
    fseek(fp, 0, SEEK_END);
    length = ftell(fp);
    rewind(fp);
    //根据图像数据长度分配内存buffer
    char* ImgBuffer=(char*)malloc(length* sizeof(char));
    //将图像数据读入buffer
    fread(ImgBuffer, length, 1, fp);
    fclose(fp);
    //输入要保存的文件名
    cout << "Enter the name you want to save:";
    cin >> OutImgName;
    //以二进制写入方式
    if ((fp=fopen(OutImgName, "wb"))==NULL)
    {
        cout << "Open File failed!" <<endl;
        exit(0);
    }
    //从buffer中写数据到fp指向的文件中
    fwrite(ImgBuffer, length, 1, fp);
    cout << "Done!" <<endl;
    //关闭文件指针，释放buffer内存
    fclose(fp);
    free(ImgBuffer);
    
    return 0;
}