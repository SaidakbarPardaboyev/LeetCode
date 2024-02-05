int firstUniqChar(char* s) {
    for(int i = 0; i < strlen(s); i++){
        int count = 0;
        for(int j = 0; j < strlen(s); j++){
            if(s[i] == s[j]){
                count++;
                if(count > 1){
                    break;
                }
            }
        }
        if(count <= 1){
            return i;
        }
    }
    return -1;
}
