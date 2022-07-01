
#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <time.h>
 
typedef struct procstat {
     char processorName[20];
     unsigned int user;
     unsigned int nice;
     unsigned int system;
     unsigned int idle;
     unsigned int iowait;
     unsigned int irq;
     unsigned int softirq;
     unsigned int stealstolen;
     unsigned int guest;
} Procstat;
 
Procstat getCPUStatus() {
    // Get "/proc/stat" info.
    FILE* inputFile = NULL;
    
    chdir("/proc");
    inputFile = fopen("stat", "r");
    if (!inputFile) {
        perror("error: Can not open file.\n");
    }

    char buff[1024];
    fgets(buff, sizeof(buff), inputFile); // Read 1 line.
    printf(buff);
    Procstat ps;
    sscanf(buff, "%s %u %u %u %u %u %u %u %u %u", ps.processorName, &ps.user, &ps.nice, &ps.system, &ps.idle, &ps.iowait, &ps.irq, &ps.softirq, &ps.stealstolen, &ps.guest); // Scan from "buff".
    printf("user: %u\n", ps.user);

    fclose(inputFile);
 
    return ps;
    
}
 
float calculateCPUUse(Procstat ps1, Procstat ps2) {
    unsigned int totalCPUTime = (ps2.user + ps2.nice + ps2.system + ps2.idle + ps2.iowait + ps2.irq + ps2.softirq + ps2.stealstolen + ps2.guest) - (ps1.user + ps1.nice + ps1.system + ps1.idle + ps1.iowait + ps1.irq + ps1.softirq + ps1.stealstolen + ps1.guest);
    unsigned int idleCPUTime = ps2.idle - ps1.idle;
 
    float CPUUse = ((float) totalCPUTime - (float) idleCPUTime) / (float) totalCPUTime;
 
    printf("totalCPUTime: %u\nidleCPUTime: %u\n", totalCPUTime, idleCPUTime);
 
    return CPUUse;
}
 
int main(int argc, char* argv[]) {
    printf("Test CPU\n");
 
    // Get processor num.
    int processorNum = sysconf(_SC_NPROCESSORS_CONF); // "unistd.h" is required.
    printf("Processors: %d\n", processorNum);
    
    while(1) {
        // Test
        Procstat ps1, ps2;
        int i = 0;
        for (i = 0; i <= 100000; i++) {
              
             srand((unsigned) time(NULL));
             int m = rand() % 100000;
             int n = 1 + rand() % 100000;
             int k = m / n;
              
             if (i == 10) {
                  ps1 = getCPUStatus();
             }
     
             if (i == 10000) {
                  ps2 = getCPUStatus();
             }
        }
        float CPUUse = calculateCPUUse(ps1, ps2);
        printf("CPUUse: %.2f%%\n", CPUUse*100);
        sleep(1);
    }
    
    return 0;
}