package handlers

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *ask_cow(char phrase[]) {
    int phrase_len = strlen(phrase);
    char *buf = (char *)malloc(sizeof(char) * (160 + (phrase_len + 2) * 3));
    strcpy(buf, " ");

    for (unsigned int i = 0; i < phrase_len + 2; ++i) {
        strcat(buf, "_");
    }

    strcat(buf, "\n< ");
    strcat(buf, phrase);
    strcat(buf, " ");
    strcat(buf, ">\n ");

    for (unsigned int i = 0; i < phrase_len + 2; ++i) {
        strcat(buf, "-");
    }
    strcat(buf, "\n");
    strcat(buf, "        \\   ^__^\n");
    strcat(buf, "         \\  (oo)\\_______\n");
    strcat(buf, "            (__)\\       )\\/\\\n");
    strcat(buf, "                ||----w |\n");
    strcat(buf, "                ||     ||\n");
    return buf;
}
*/
import "C"
import (
	"unsafe"
)

func AskCow(phrase string) string {
	cPhrase := C.CString(phrase)
	defer C.free(unsafe.Pointer(cPhrase))

	cResult := C.ask_cow(cPhrase)
	defer C.free(unsafe.Pointer(cResult))

	return C.GoString(cResult)
}
