import argparse
from py.client import client


def main():
    print("Trying to RPC to the gRPC server")
    parser = argparse.ArgumentParser()
    parser.add_argument("address")
    args = parser.parse_args()

    print(client.rpc(args.address, "Hello"))


if __name__ == "__main__":
    main()
