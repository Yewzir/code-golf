//-------------------------------------------------------------------------
// [Includes]
//-------------------------------------------------------------------------
#include <dirent.h>
#include <errno.h>
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/stat.h>
#include <sys/wait.h>
#include <unistd.h>

//-------------------------------------------------------------------------
// [Defines]
//-------------------------------------------------------------------------
#define BUFFER_SIZE 4096
#define ERR_AND_EXIT(msg) do { perror(msg); exit(EXIT_FAILURE); } while (0)

//-------------------------------------------------------------------------
// [Prototypes]
//-------------------------------------------------------------------------
void copy_file(const char* src, const char* dst);
void copy_folder(const char* src, const char* dst);

//-------------------------------------------------------------------------
// [Variables]
//-------------------------------------------------------------------------
const char* bin;
const char* src;
FILE* fp;
char buffer[BUFFER_SIZE];
ssize_t nbytes;
pid_t pid;
int status;

//-------------------------------------------------------------------------
// Lorem ipsum...
//
// <Parameters>
// src: Lorem ipsum...
// dst: Lorem ipsum...
//-------------------------------------------------------------------------
void copy_file(const char* src, const char* dst) {
    FILE* src_fp, *dst_fp;

    if (!(src_fp = fopen(src, "r")))
        ERR_AND_EXIT("fopen");

    if (!(dst_fp = fopen(dst, "w"))) {
        if (fclose(src_fp))
            ERR_AND_EXIT("fclose");

        ERR_AND_EXIT("fopen");
    }

    char buffer[sizeof(short)];
    ssize_t nbytes;

    while ((nbytes = fread(buffer, sizeof(char), sizeof(buffer), src_fp)))
        if (fwrite(buffer, sizeof(char), nbytes, dst_fp) != (size_t) nbytes)
            ERR_AND_EXIT("fwrite");

    if (fclose(src_fp) || fclose(dst_fp))
        ERR_AND_EXIT("fclose");
}

//-------------------------------------------------------------------------
// Lorem ipsum...
//
// <Parameters>
// src: Lorem ipsum...
// dst: Lorem ipsum...
//-------------------------------------------------------------------------
void copy_folder(const char* src, const char* dst) {
    DIR* dp;

    if (!(dp = opendir(src)))
        ERR_AND_EXIT("opendir");

    struct stat st;

    if (stat(dst, &st))
        if (mkdir(dst, 0755)) {
            if (closedir(dp))
                ERR_AND_EXIT("closedir");

            ERR_AND_EXIT("mkdir");
        }

    struct dirent* dep;
    char src_buf[BUFFER_SIZE], dst_buf[BUFFER_SIZE];

    while ((dep = readdir(dp))) {
        if (!strcmp(dep->d_name, ".") || !strcmp(dep->d_name, ".."))
            continue;

        snprintf(src_buf, sizeof(src_buf), "%s/%s", src, dep->d_name);
        snprintf(dst_buf, sizeof(dst_buf), "%s/%s", dst, dep->d_name);

        if (!stat(src_buf, &st) && S_ISDIR(st.st_mode))
            copy_folder(src_buf, dst_buf);
        else if (!stat(src_buf, &st) && S_ISREG(st.st_mode))
            copy_file(src_buf, dst_buf);
    }

    if (closedir(dp))
        ERR_AND_EXIT("closedir");
}
