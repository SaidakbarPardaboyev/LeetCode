#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char * mergeAlternately(char * word1, char * word2){
    char* merchedString = malloc(strlen(word1)+strlen(word2)+1*sizeof(char));
    int SmallLength;
    if (strlen(word1) > strlen(word2)) {
        SmallLength = strlen(word2);
    } else {
        SmallLength = strlen(word1);
    }
    int i = 0;
    while (i < SmallLength) {
        merchedString[i*2] = word1[i];
        merchedString[i*2+1] = word2[i];
        i++;
    }

    if (i == strlen(word1) && i < strlen(word2)) {
        int k = i*2;
        for (int j = i; j < strlen(word2); j++) {
            merchedString[k++] = word2[j]; 
        }
    } else if (i == strlen(word2) && i < strlen(word1)) {
        int k = i*2;
        for (int j = i; j < strlen(word1); j++) {
            merchedString[k] = word1[j];
            k++; 
        }
    }
    return merchedString;
}

int main() {
    char word1[] = "abc";
    char word2[] = "12";
    char *result = mergeAlternately(word1, word2);
    if (result != NULL) {
        printf("Merged string: %s\n", result);
    }
    return 0;
}