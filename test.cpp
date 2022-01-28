
// char* substr(char* arr, int begin, int len)
// {
//     char* res = new char[len + 1];
//     for (int i = 0; i < len; i++)
//         res[i] = *(arr + begin + i);
//     res[len] = 0;
//     return res;
// }

int main (){

    int messageLengthWhateverThatIs = 200;
    char message [messageLengthWhateverThatIs];
    int splitLength  = 120;

    // do it here
    // for content in message split by chunksize:
    //     send content

    for(int i = 0; i < messageLengthWhateverThatIs; i += splitLength ){
        int endPoint = i + splitLength;
        if (endPoint > messageLengthWhateverThatIs) {
            endPoint = messageLengthWhateverThatIs;
        }

        int len = endPoint - i + 1;
        char whatYouNeeded[len];
        for (int j = 0; j < len; j++)
            whatYouNeeded[i] = *(message + i + j);
        whatYouNeeded[len] = 0;

        //char[splitLength] whatYouNeeded = substr(message , i, endPoint - i);
    }

    return 0;
}