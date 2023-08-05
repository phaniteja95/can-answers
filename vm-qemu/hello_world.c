#include <stdio.h>
#include <unistd.h>

int main(int argc, char** argv) {
  char *args[] = {"bin/ash", NULL};
  
  printf("********************\n");
  printf("Hello, world!\n");
  printf("********************\n");
  
  execv("/bin/ash", args);
  
  while (1) sleep(10);
  return 0;
}
