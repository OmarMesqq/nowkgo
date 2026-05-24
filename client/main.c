#include <arpa/inet.h>
#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <unistd.h>

#define MAX_INPUT_SIZE 1024
#define HOST "127.0.0.1"
#define PORT 9001

static void getClientBuffer(char* client_buffer, int buffer_size, int client_socket);
static void displayServerBuffer(int client_socket, char* server_buffer, size_t buffer_size);
static int setupClient(void);
static struct sockaddr_in setupServer(void);

int main(void) {
  char client_buffer[MAX_INPUT_SIZE] = {0};  // stores client-side entered text
  int client_socket = setupClient();

  char server_buffer[MAX_INPUT_SIZE] = {0};  // stores received messages
  struct sockaddr_in server_address = setupServer();

  printf("\nWelcome to the joke theater! You have five minutes between interactions before exploding :)\n");

  if (inet_pton(AF_INET, HOST, &server_address.sin_addr) <= 0) {
    perror("[!] Invalid IP address. Check it again.");
    exit(EXIT_FAILURE);
  }

  if (connect(client_socket, (struct sockaddr*)&server_address, sizeof(server_address)) != 0) {
    fprintf(stderr, "[!] Couldn't connect to server. Reason: ");
    if (errno == ECONNREFUSED) {
      fprintf(stderr, "Connection refused. Theater probably isn't open.\n");
    } else {
      fprintf(stderr, "Unknown error: %d\n", errno);
    }
    exit(EXIT_FAILURE);
  }

  while (1) {
    displayServerBuffer(client_socket, server_buffer, MAX_INPUT_SIZE);

    printf("> ");
    getClientBuffer(client_buffer, MAX_INPUT_SIZE, client_socket);

    displayServerBuffer(client_socket, server_buffer, MAX_INPUT_SIZE);

    printf("> ");
    getClientBuffer(client_buffer, MAX_INPUT_SIZE, client_socket);

    displayServerBuffer(client_socket, server_buffer, MAX_INPUT_SIZE);

    break;
  }

  close(client_socket);
  printf("Thank you for your time!\n");
  return 0;
}

static struct sockaddr_in setupServer(void) {
  struct sockaddr_in server_address;
  server_address.sin_family = AF_INET;    // IPv4
  server_address.sin_port = htons(PORT);  // Ensures port number is Big Endian (network byte order)
  return server_address;
}

static int setupClient(void) {
  int client_socket;
  client_socket = socket(AF_INET, SOCK_STREAM, 0);  // IPv4, stream oriented (TCP)

  if (client_socket == -1) {
    perror("[!] Couldn't create client socket.");
    return EXIT_FAILURE;
  }
  return client_socket;
}

static void displayServerBuffer(int client_socket, char* server_buffer, size_t buffer_size) {
  // plataform independent: contains maximum allowed size for I/O
  ssize_t bytes_received;
  bytes_received = recv(client_socket, server_buffer, buffer_size - 1, 0);

  if (bytes_received <= 0) {
    perror("[!] Connection error: the server didn't send any data or the connection was closed.");
    return;
  }

  if (strcmp(server_buffer, "> KABOOM!") == 0) {
    printf("> KABOOM!");
    exit(0);
  }

  server_buffer[bytes_received] = '\0';
  printf("%s\n", server_buffer);
}

static void getClientBuffer(char* client_buffer, int buffer_size, int client_socket) {
  if (fgets(client_buffer, buffer_size, stdin) == NULL) {
    return;
  }

  size_t clientMsgLen = strcspn(client_buffer, "\n");

  if (client_buffer[clientMsgLen] != '\n') {
    fprintf(stderr, "Your message is just too long...\n");
    close(client_socket);
    exit(-1);
  }

  client_buffer[clientMsgLen] = '\0';
  if (strcmp(client_buffer, "exit") == 0) {
    return;
  }
  if (send(client_socket, client_buffer, strlen(client_buffer), 0) < 0) {
    perror("[!] Couldn't send message to server.");
    exit(EXIT_FAILURE);
  }
}
