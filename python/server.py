from concurrent import futures
import time
import rpc_pb2
import grpc
import rpc_pb2_grpc

class Add(rpc_pb2_grpc.AddServicer):
# 实现 proto 文件中定义的 rpc 调用
    def Adds(self, request, context):
        return rpc_pb2.Result({"num":3})

# 定义开启4个线程处理接收到的请求
server = grpc.server(futures.ThreadPoolExecutor(max_workers=4))
# 将编译出来的rpc_pb2_grpc的add_HelloServiceServicer_to_server函数添加到server中
rpc_pb2_grpc.add_AddServicer_to_server(Add, server)

# 定义服务端端口1234
server.add_insecure_port('127.0.0.1:1234')
server.start()

# 长期监听
try:
    while True:
        time.sleep(60 * 60 * 24)
except KeyboardInterrupt:
    server.stop(0)