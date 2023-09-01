#include <stdio.h>      // Include standard I/O for printf and other functions
#include <stdlib.h>     // Include standard library for exit and other utility functions
#include <string.h>     // Include string.h for string manipulation functions like strlen
#include <unistd.h>     // Include unistd.h for close function
#include <arpa/inet.h>  // Include arpa/inet.h for inet_pton and other network functions
// Include sys/types and sys/socket for socket programming types and functions
#include <sys/types.h>
#include <sys/socket.h>

#define MAX_INPUT_SIZE 100 

static const char HOST[10] = "127.0.0.1";

struct sockaddr_in setupServer() {
    struct sockaddr_in server_address;
    server_address.sin_family = AF_INET; 
    server_address.sin_port = htons(9001); // Garante que a porta será Big Endian (network byte order)
    return server_address;
}

int setupClient() {
    int client_socket; // Armazena o descriptor do cliente
    client_socket = socket(AF_INET, SOCK_STREAM, 0); // IPv4, stream oriented e TCP

    if (client_socket == -1) {
        perror("Não foi possível criar socket do cliente!");
        return EXIT_FAILURE;
    }
    return client_socket;
}

int main() {
    char client_buffer[MAX_INPUT_SIZE]; // buffer p/ guardar o que o cliente digita 
    char server_buffer[1024] = {0}; // inicializa buffer p/ guardar msgs recebidas para todo de zeros 
    ssize_t bytes_received; // independente de plataforma: define tamanho maximo que permite I/O

    struct sockaddr_in server_address = setupServer();

    printf("Seja bem vindo(a) ao servidor de piadas!\nVocê tem 1 minuto entre conversas para não explodir :)");
    printf("\n");

    // Converte HOST para binário e guarda no struct
    if (inet_pton(AF_INET, HOST, &server_address.sin_addr) <= 0) {
        perror("Endereço de IP inválido! Verifique se o entrou corretamente.");
        exit(EXIT_FAILURE);
    }

    while (1) {
        printf("> ");

        if (fgets(client_buffer, sizeof(client_buffer), stdin) == NULL) {
            break;
        }

        client_buffer[strcspn(client_buffer, "\n")] = '\0';

        if (strcmp(client_buffer, "exit") == 0) {
            break;
        }

    }

    printf("Agradecemos a preferência!Volte logo!\n");
    return 0;
}
