#include <stdio.h>      
#include <stdlib.h>     
#include <string.h>     
#include <unistd.h>     
#include <arpa/inet.h>  
#include <sys/types.h>
#include <sys/socket.h>

#define MAX_INPUT_SIZE 1024 

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

void displayServerBuffer(int client_socket, char* server_buffer, size_t buffer_size) {
    ssize_t bytes_received; // independente de plataforma: define tamanho maximo que permite I/O

    // if (bytes_received <= 0) {
    //     perror("Erro na conexão: o servidor não mandou dados ou a conexão foi interrompida");
    //     return;
    // }

    bytes_received = recv(client_socket, server_buffer, buffer_size - 1, 0);
    server_buffer[bytes_received] = '\0'; 
    printf("%s\n", server_buffer);
}

void getClientBuffer(char* client_buffer, size_t buffer_size, int client_socket) {
    if (fgets(client_buffer, buffer_size, stdin) == NULL) {
            return;
        }
        client_buffer[strcspn(client_buffer, "\n")] = '\0';
        if (strcmp(client_buffer, "exit") == 0) {
            return;
        }
        if (send(client_socket, client_buffer, strlen(client_buffer), 0) < 0) {
            perror("Não foi possível mandar mensagem para o servidor!");
            exit(EXIT_FAILURE);
        }
}

int main() {
    char client_buffer[MAX_INPUT_SIZE]; // buffer p/ guardar o que o cliente digita 
    size_t client_buffer_size = sizeof(client_buffer);
    int client_socket = setupClient();

    char server_buffer[1024] = {0}; // inicializa buffer p/ guardar msgs recebidas para todo de zeros 
    size_t server_buffer_size = sizeof(server_buffer);
    struct sockaddr_in server_address = setupServer();

    printf("Seja bem vindo(a) ao servidor de piadas!\nVocê tem 1 minuto entre conversas para não explodir :)\n");
    
    if (inet_pton(AF_INET, HOST, &server_address.sin_addr) <= 0) {
        perror("Endereço de IP inválido! Verifique se o entrou corretamente.");
        exit(EXIT_FAILURE);
    }

    if (connect(client_socket, (struct sockaddr*)&server_address, sizeof(server_address)) < 0) {
        perror("Não foi possível conectar ao servidor!");
        exit(EXIT_FAILURE);
    }

    while (1) {
        displayServerBuffer(client_socket, server_buffer, server_buffer_size);

        printf("> ");
        getClientBuffer(client_buffer, client_buffer_size, client_socket);

        displayServerBuffer(client_socket, server_buffer, server_buffer_size);

        printf("> ");
        getClientBuffer(client_buffer, client_buffer_size, client_socket);

        displayServerBuffer(client_socket, server_buffer, server_buffer_size);

        break;
    }

    close(client_socket);
    printf("\nAgradecemos a preferência! Volte logo!\n");
    return 0;
}
