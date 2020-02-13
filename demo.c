#include <stdio.h>
#include "iplookuplib.h"


int main() {
    int val = IsValidEnvironment();
    printf("%d", val);
    if (val > 0) {
        printf("Valid Environment!");
    }
    else {
        printf("Invalid Environment");
    }
}
// See client1.c on GitHub