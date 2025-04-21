import argparse
from kafka import KafkaConsumer
import struct


def main():
    # Parse command line arguments
    parser = argparse.ArgumentParser(description='Kafka Consumer App')
    parser.add_argument('--bootstrap_servers',
                        nargs='+',
                        type=str,
                        required=True,
                        help='List of Kafka bootstrap servers (e.g., localhost:9092)')
    parser.add_argument('--topic',
                        type=str,
                        default='payments',
                        help='Kafka topic for payments (default: payments)')
    parser.add_argument('--group',
                        type=str,
                        default='pgroup',
                        help='Kafka topic for payments (default: pgroup)')
    args = parser.parse_args()

    # Connect to Kafka cluster
    consumer = KafkaConsumer(
        args.topic,
        bootstrap_servers=args.bootstrap_servers,
        group_id=args.group,
        auto_offset_reset='earliest',
        enable_auto_commit=False,
    )

    try:
        for message in consumer:
            partition = message.partition
            offset = message.offset
            #key = struct.unpack('<H', message.key)[0]  # Assuming the key is an unsigned 32-bit integer
            print(f"Partition: {partition}, Offset: {offset}, client_id (key): {message.key.decode('utf-8')}, Value: {message.value.decode('utf-8')}")
            print(f"{message}Received message: {message.value.decode('utf-8')}")
            consumer.commit()
    except KeyboardInterrupt:
        print("Received keyboard interrupt, closing the consumer.")
    finally:
        consumer.close()


if __name__ == "__main__":
    main()
