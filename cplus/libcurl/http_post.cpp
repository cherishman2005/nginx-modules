#include<iostream>
#include <curl/curl.h>
using namespace std;

int main()
{
    //保存输入图像文件名和输出图像文件名
    const char *InImgName = "2314776959_1639567355890.jpeg";
    //图像数据长度
    int length;
    //文件指针
    FILE* fp;
    //输入要读取的图像名
    //cout << "Enter Image name:";
    //cin >> InImgName;
    //以二进制方式打开图像
    if (NULL == (fp = fopen(InImgName, "rb")))
    {
        cout << "Open image failed!" << endl;
        exit(0);
    }
    //获取图像数据总长度
    fseek(fp, 0, SEEK_END);
    length = ftell(fp);
    rewind(fp);

	cout << "length=" << length << endl;
	//根据图像数据长度分配内存buffer
    char* ImgBuffer = new char[length* sizeof(char)];
    //将图像数据读入buffer
    fread(ImgBuffer, length, 1, fp);
    fclose(fp);
    
    /*
    //输入要保存的文件名
    cout << "Enter the name you want to save:";
    cin >> OutImgName;
    //以二进制写入方式
    if (NULL == (fp=fopen(OutImgName, "wb")))
    {
        cout << "Open File failed!" <<endl;
        exit(0);
    }
    //从buffer中写数据到fp指向的文件中
    fwrite(ImgBuffer, length, 1, fp);
    cout << "Done!" <<endl;
    //关闭文件指针，释放buffer内存
    */
    
    
    CURL *curl = curl_easy_init();

    curl_easy_setopt(curl, CURLOPT_URL, "http://127.0.0.1:8080/");
    
    
    struct curl_slist *headers = NULL;
    headers = curl_slist_append(headers, "Content-Type: application/binary");
    curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);
    
    //curl_easy_setopt(curl, CURLOPT_POSTFIELDS, "{\"hi\" : \"there\"}");
    curl_easy_setopt(curl, CURLOPT_POSTFIELDS, ImgBuffer);
    //curl_easy_setopt(curl, CURLOPT_POSTFIELDSIZE_LARGE, length_of_data);
    curl_easy_setopt(curl, CURLOPT_POSTFIELDSIZE, length);

    curl_easy_perform(curl);
    
    fclose(fp);
    delete [] ImgBuffer;
    
    return 0;
}
