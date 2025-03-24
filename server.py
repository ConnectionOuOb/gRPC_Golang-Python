import grpc
from concurrent import futures
import time

from proto import calc_pb2
from proto import calc_pb2_grpc


class CalculatorServicer(calc_pb2_grpc.CalculatorServicer):
    def Calculate(self, request, context):
        # 建立一個 List 存放所有 Outer
        outers_list = []
        # 對於每個輸入的字串，產生兩個 Outer
        for input_str in request.inputs:
            # 第一個 Outer，包含兩個 Inner
            outer1 = calc_pb2.Outer()
            inner1 = calc_pb2.Inner()
            inner2 = calc_pb2.Inner()
            # inner1 依字元計算（例如：ascii 碼除以 100）
            for ch in input_str:
                inner1.values.append(float(ord(ch)) / 100.0)
            # inner2 依字元計算（例如：ascii 碼除以 50）
            for ch in input_str:
                inner2.values.append(float(ord(ch)) / 50.0)
            outer1.inners.extend([inner1, inner2])

            # 第二個 Outer，包含兩個 Inner 以固定值示範
            outer2 = calc_pb2.Outer()
            inner3 = calc_pb2.Inner()
            inner4 = calc_pb2.Inner()
            inner3.values.extend([1.1, 2.2, 3.3])
            inner4.values.extend([4.4, 5.5])
            outer2.inners.extend([inner3, inner4])

            outers_list.extend([outer1, outer2])

        response = calc_pb2.CalcResponse()
        response.outers.extend(outers_list)
        return response


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    calc_pb2_grpc.add_CalculatorServicer_to_server(CalculatorServicer(), server)
    server.add_insecure_port("[::]:50051")
    server.start()
    print("Python gRPC server started on port 50051")
    try:
        while True:
            time.sleep(86400)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == "__main__":
    serve()
