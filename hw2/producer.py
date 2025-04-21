import random
import time
import sys
import json
import argparse
from kafka import KafkaProducer


def generate_payment_event():
    payment_id = random.randint(1, 1000)
    amount = round(random.uniform(10.00, 100.00), 2)
    timestamp = time.time()
    return {
        "payment_id": payment_id,
        "timestamp": timestamp,
        "amount": amount,
        "status": random.choice(['successful', 'failed']),
        "customer_id": random.randint(1, 9999)
    }


def produce_data():
    # Parse command line arguments
    parser = argparse.ArgumentParser(description='Generate synthetic stream workload data')
    parser.add_argument('--bootstrap_servers',
                        nargs='+',
                        type=str,
                        help='List of Kafka bootstrap servers (e.g., localhost:9092)')
    parser.add_argument('--sleep_time',
                        type=float,
                        default=1.0,
                        help='Time to sleep between sending events (default: 1.0 seconds)')
    parser.add_argument('--topic',
                        type=str,
                        default='payments',
                        help='Kafka topic for payments (default: payments)')

    # Parse arguments
    args = parser.parse_args()

    # Initialize producers for each topic
    payment_producer = KafkaProducer(
        bootstrap_servers=args.bootstrap_servers,
        value_serializer=lambda v: json.dumps(v).encode('utf-8'),
        key_serializer=lambda k: str(k).encode('utf-8'),
        acks=-1,
        batch_size=10,
    )

    try:
        while True:
            # Generate events
            payment_event = generate_payment_event()
            try:
                payment_producer.send(args.topic, value=payment_event, key=payment_event['customer_id'])
                print(f"Sent event: {str(payment_event)}")
            except Exception as e:
                print(f"Error sending payment event: {str(e)}")

            # Sleep to prevent overwhelming Kafka
            time.sleep(args.sleep_time)
    except KeyboardInterrupt:
        print("\nCtrl+D received. Closing producers and finalizing error counts...")
        sys.exit(1)
    finally:
        payment_producer.close()


if __name__ == "__main__":
    produce_data()
