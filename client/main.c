#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_INPUT_SIZE 100 

void echoInput(const char* input) {
    printf("%s\n", input);
}

int main() {
    char input_buffer[MAX_INPUT_SIZE];

    while (1) {
        printf("> ");

        // Joga entrada do stdin pro array de entrada até o tamanho especificado 
        // Se der NULL, quebra o ciclo
        if (fgets(input_buffer, sizeof(input_buffer), stdin) == NULL) {
            break;
        }

        // Substitui o caracter newline no input buffer 
        // pelo null terminating 
        // A função strcspn retorna a posição do segundo caracter 
        // na primeira string 
        input_buffer[strcspn(input_buffer, "\n")] = '\0';

        if (strcmp(input_buffer, "exit") == 0) {
            break;
        }

        echoInput(input_buffer);
    }

    printf("Exiting. Goodbye!\n");
    return 0;
}