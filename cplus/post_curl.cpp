#include <stdio.h>
#include <curl/curl.h>
bool getUrl(char *filename)
{
    CURL *curl;
    CURLcode res;
    FILE *fp;
    if ((fp = fopen(filename, "w")) == NULL)  // 返回结果用文件存储
        return false;
    struct curl_slist *headers = NULL;
    headers = curl_slist_append(headers, "Accept: Agent-007");
    curl = curl_easy_init();    // 初始化
    if (curl)
    {
        //curl_easy_setopt(curl, CURLOPT_PROXY, "10.99.60.201:8080");// 代理
        curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);// 改协议头
        curl_easy_setopt(curl, CURLOPT_URL,"http://www.baidu.com");
        curl_easy_setopt(curl, CURLOPT_WRITEDATA, fp); //将返回的http头输出到fp指向的文件
        curl_easy_setopt(curl, CURLOPT_HEADERDATA, fp); //将返回的html主体数据输出到fp指向的文件
        res = curl_easy_perform(curl);   // 执行
        if (res != 0) {

            curl_slist_free_all(headers);
            curl_easy_cleanup(curl);
        }
        fclose(fp);
        return true;
    }
}
bool postUrl(char *filename)
{
    CURL *curl;
    CURLcode res;
    FILE *fp;
    if ((fp = fopen(filename, "w")) == NULL)
        return false;
    curl = curl_easy_init();
    if (curl)
    {
        //curl_easy_setopt(curl, CURLOPT_COOKIEFILE, "/tmp/cookie.txt"); // 指定cookie文件
        curl_easy_setopt(curl, CURLOPT_POSTFIELDS, "&logintype=uid&u=xieyan&psw=xxx86");    // 指定post内容
        //curl_easy_setopt(curl, CURLOPT_PROXY, "10.99.60.201:8080");
        curl_easy_setopt(curl, CURLOPT_URL, " http://127.0.0.1:8080/");   // 指定url
        curl_easy_setopt(curl, CURLOPT_WRITEDATA, fp);
        res = curl_easy_perform(curl);
        curl_easy_cleanup(curl);
    }
    fclose(fp);
    return true;
}

bool postUrl1() {
    CURL *curl = curl_easy_init();

    curl_easy_setopt(curl, CURLOPT_URL, "http://127.0.0.1:8080/");
    curl_easy_setopt(curl, CURLOPT_POSTFIELDS, "{\"hi\" : \"there\"}");

    curl_easy_perform(curl);
}
int main(void)
{
    //getUrl("/tmp/get.html");
    //postUrl("./1.txt");
    postUrl1();
}
