#include <stdio.h>
#include <string.h>
#include "libbackend.h"

int main() {
	GoString name = {"Vedant", strlen("Vedant")};
	GoString pword = {"pword", strlen("pword")};
	printf("%s\n", GetUser(1));
}
