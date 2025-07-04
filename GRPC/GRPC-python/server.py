# server.py
from concurrent import futures
import grpc
import chat_pb2
import chat_pb2_grpc

# Implementation of the ChatService defined in chat.proto
class ChatService(chat_pb2_grpc.ChatServiceServicer):
    def SayHello(self, request, context):
        print("Received from client:", request.body)
        return chat_pb2.Message(body="Hello from Python gRPC server!")

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    chat_pb2_grpc.add_ChatServiceServicer_to_server(ChatService(), server)
    server.add_insecure_port('[::]:9000')
    server.start()
    print("gRPC server started on port 9000")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
