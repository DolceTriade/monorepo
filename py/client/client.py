import logging
import random
from typing import Iterable, List
import grpc

from proto import test_pb2
from proto import test_pb2_grpc


def rpc(addr, req):
    with grpc.insecure_channel(addr) as channel:
        stub = test_pb2_grpc.TestServiceStub(channel)
        req = test_pb2.TestRequest()
        req.request = "hi"
        res = stub.Test(req)
        if res:
            return res.response
