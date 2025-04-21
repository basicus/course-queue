import sys
from faker import Faker
import struct

from kafka import KafkaProducer


def usage():
    print(f"Usage: {sys.argv[0]} <bootstrap_servers> <topic>")


def main(argv):
    if len(argv) < 2:
        usage()
        sys.exit(2)

    bootstrap_servers = argv[0]
    topic = argv[1]

    producer = KafkaProducer(bootstrap_servers=bootstrap_servers.split(','),
                             value_serializer=lambda v: v.encode('utf-8'),
                             key_serializer=lambda k: str(k).encode('utf-8'),
                             )
    fake = Faker()
    try:
        for i in range(10000):
            w = fake.name()
            producer.send(topic, key=i, value=w)
            print(f"Sent: key={i}\tvalue={w}")

    except KeyboardInterrupt:
        print("Shutting down.")
    finally:
        producer.close()


if __name__ == "__main__":
    main(sys.argv[1:])
