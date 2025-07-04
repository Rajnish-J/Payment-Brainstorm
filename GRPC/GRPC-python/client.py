# client.py
import grpc
import chat_pb2
import chat_pb2_grpc

def run():
    channel = grpc.insecure_channel('localhost:9000')
    stub = chat_pb2_grpc.ChatServiceStub(channel)

    request = chat_pb2.Message(body="Hello from Python gRPC client!")
    response = stub.SayHello(request)

    print("Response from server:", response.body)

if __name__ == '__main__':
    run()
